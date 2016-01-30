/*                                                                           */
/*  Brown University, CS138, Spring 2016                                     */
/*                                                                           */
/*  Purpose: Chord struct and related functions to create new nodes, etc.    */
/*                                                                           */
/*                                                                           */
/*                                                                           */
/*  Acknowledgments: Eli Rosenthal (ezr) and Joshua Liebow-Feeser (jliebowf) */
/*  	whose code for syncing goroutines we adapted from their 2015 	     */
/*	Chord handin. 							     */
/*                                                                           */

package chord

import (
	"fmt"
	"log"
	"net"
	"net/rpc"
	"sync"
	"time"

	"../../cs138"
)

// Number of bits (i.e. m value). Assumes <= 128 and divisible by 8
const KEY_LENGTH = 8

// Timeout for Accept() blocking on the TCPListener for this node.
const LISTEN_TIMEOUT = time.Duration(100) * time.Millisecond

// Non-local node representation
type RemoteNode struct {
	Id   []byte
	Addr string
}

// Local node representation
type Node struct {
	Id         []byte           /* Unique Node ID */
	Listener   *net.TCPListener /* Node listener socket */
	Addr       string           /* String of listener address */
	RemoteSelf *RemoteNode      /* Remote node of our self */

	Successor   *RemoteNode  /* This Node's successor */
	Predecessor *RemoteNode  /* This Node's predecessor */
	sLock       sync.RWMutex /* RWLock for successor */
	pLock       sync.RWMutex /* RWLock for predecessor */

	IsShutdown bool         /* Is node in process of shutting down? */
	sdLock     sync.RWMutex /* RWLock for shutdown flag */

	FingerTable []FingerEntry /* Finger table entries */
	FtLock      sync.RWMutex  /* RWLock for finger table */

	dataStore map[string]string /* Local datastore for this node */
	DsLock    sync.RWMutex      /* RWLock for datastore */

	wg sync.WaitGroup /* WaitGroup of concurrent goroutines to sync before exiting */
}

// Creates a Chord node with a pre-defined ID (useful for testing).
func CreateDefinedNode(parent *RemoteNode, definedId []byte) (*Node, error) {
	node := new(Node)
	err := node.init(parent, definedId)
	if err != nil {
		return nil, err
	}
	return node, err
}

// Create Chord node with random ID based on listener address.
func CreateNode(parent *RemoteNode) (*Node, error) {
	node := new(Node)
	err := node.init(parent, nil)
	if err != nil {
		return nil, err
	}
	return node, err
}

// Initailize a Chord node, start listener, RPC server, and go routines.
func (node *Node) init(parent *RemoteNode, definedId []byte) error {
	if KEY_LENGTH > 128 || KEY_LENGTH%8 != 0 {
		log.Fatal(fmt.Sprintf("KEY_LENGTH of %v is not supported! Must be <= 128 and divisible by 8", KEY_LENGTH))
	}

	listener, _, err := cs138.OpenTCPListener()
	if err != nil {
		return err
	}

	node.Id = HashKey(listener.Addr().String())
	if definedId != nil {
		node.Id = definedId
	}

	node.Listener = listener
	node.Addr = listener.Addr().String()
	node.IsShutdown = false
	node.dataStore = make(map[string]string)

	// Populate RemoteNode that points to self
	node.RemoteSelf = new(RemoteNode)
	node.RemoteSelf.Id = node.Id
	node.RemoteSelf.Addr = node.Addr

	// Join this node to the same chord ring as parent
	err = node.join(parent)
	if err != nil {
		return err
	}

	// Populate finger table
	node.initFingerTable()

	// Thread 1: start RPC server on this connection

	rpc.RegisterName(node.Addr, node)
	node.spawn(func() { node.startRpcServer() })

	// Thread 2: kick off timer to stabilize periodically

	ticker1 := time.NewTicker(time.Millisecond * 100) //freq
	node.spawn(func() { node.stabilize(ticker1) })

	// Thread 3: kick off timer to fix finger table periodically

	ticker2 := time.NewTicker(time.Millisecond * 90) //freq
	node.spawn(func() { node.fixNextFinger(ticker2) })

	return err
}

// Go routine to accept and process RPC requests.
func (node *Node) startRpcServer() {
	for {
		node.sdLock.RLock()
		sd := node.IsShutdown
		node.sdLock.RUnlock()
		if sd {
			Debug.Printf("[%v] Shutting down RPC server\n", HashStr(node.Id))
			return
		}

		node.Listener.SetDeadline(time.Now().Add(LISTEN_TIMEOUT))
		if conn, err := node.Listener.Accept(); err != nil {
			type timeout interface {
				Timeout() bool
			}

			// Ignore timeout errors since we set an Accept() deadline.
			if t, ok := err.(timeout); ok && t.Timeout() {
				continue
			}

			log.Fatal("accept error: " + err.Error())
		} else {
			go func() {
				rpc.ServeConn(conn)
				conn.Close()
			}()
		}
	}
}

// Gracefully shutdown a specified Chord node.
func ShutdownNode(node *Node) {
	node.sdLock.Lock()
	node.IsShutdown = true
	node.sdLock.Unlock()
	//TODO students should modify this method to gracefully shutdown a node

	// Wait for all go routines to exit.
	node.wg.Wait()
	node.Listener.Close()
}

// Adds a new goroutine to the WaitGroup, spawns the go routine,
// and removes the goroutine from the WaitGroup on exit.
func (node *Node) spawn(fun func()) {
	go func() {
		node.wg.Add(1)
		fun()
		node.wg.Done()
	}()
}
