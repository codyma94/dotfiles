package tapestry

import (
	"fmt"
	"net/rpc"
	"reflect"
	"sync"
)

/** Tapestries can be manually registered / unregistered to use caching rather than RPCs */
var doCaching = true
var cachedTapestries = make(map[string]*Tapestry)
var tapestryMapMutex = &sync.RWMutex{}

func getCachedTapestry(address string) *Tapestry {
	if doCaching {
		tapestryMapMutex.RLock()
		result := cachedTapestries[address]
		tapestryMapMutex.RUnlock()
		return result
	} else {
		return nil
	}
}

/*
	The methods defined in this file parallel the methods defined in tapestry-local.
	These methods take an additional argument, the node on which the method should be invoked.
	Calling any of these methods will invoke the corresponding method on the specified remote node.
*/

// Remote API: ping an address to get tapestry node info
func Hello_RPC(localNode Node, address string) (rsp Node, err error) {
	cached := getCachedTapestry(address)
	if cached != nil {
		cached.server.Hello(localNode, &rsp)
	} else {
		err = makeRemoteCall(address, "TapestryRPCServer", "Hello", localNode, &rsp)
	}
	return
}

// Remote API: makes a remote call to the GetNextHop function
func GetNextHop_RPC(remote Node, id ID) (bool, Node, error) {
	var rsp NextHopResponse
	err := makeRemoteNodeCall(remote, "GetNextHop", NextHopRequest{remote, id}, &rsp)
	return rsp.HasNext, rsp.Next, err
}

// Remote API: makes a remote call to the Register function
func Register_RPC(remote Node, replica Node, key string) (bool, error) {
	var rsp RegisterResponse
	err := makeRemoteNodeCall(remote, "Register", RegisterRequest{remote, replica, key}, &rsp)
	return rsp.IsRoot, err
}

// Remote API: makes a remote call to the Fetch function
func Fetch_RPC(remote Node, key string) (bool, []Node, error) {
	var rsp FetchResponse
	err := makeRemoteNodeCall(remote, "Fetch", FetchRequest{remote, key}, &rsp)
	return rsp.IsRoot, rsp.Values, err
}

// Remote API: makes a remote call to the RemoveBadNodes function
func RemoveBadNodes_RPC(remote Node, toremove []Node) error {
	return makeRemoteNodeCall(remote, "RemoveBadNodes", RemoveBadNodesRequest{remote, toremove}, &Node{})
}

// Remote API: makes a remote call to the AddNode function
func AddNode_RPC(remote Node, newnode Node) (neighbors []Node, err error) {
	err = makeRemoteNodeCall(remote, "AddNode", NodeRequest{remote, newnode}, &neighbors)
	return
}

// Remote API: makes a remote call to the AddNodeMulticast function
func AddNodeMulticast_RPC(remote Node, newnode Node, level int) (neighbors []Node, err error) {
	err = makeRemoteNodeCall(remote, "AddNodeMulticast", AddNodeMulticastRequest{remote, newnode, level}, &neighbors)
	return
}

// Remote API: makes a remote call to the Transfer function
func Transfer_RPC(remote Node, from Node, data map[string][]Node) error {
	return makeRemoteNodeCall(remote, "Transfer", TransferRequest{remote, from, data}, &Node{})
}

// Remote API: makes a remote call to the AddBackpointer function
func AddBackpointer_RPC(remote Node, toAdd Node) error {
	return makeRemoteNodeCall(remote, "AddBackpointer", NodeRequest{remote, toAdd}, &Node{})
}

// Remote API: makes a remote call to the RemoveBackpointer function
func RemoveBackpointer_RPC(remote Node, toRemove Node) error {
	return makeRemoteNodeCall(remote, "RemoveBackpointer", NodeRequest{remote, toRemove}, &Node{})
}

// Remote API: makes a remote call to the GetBackpointers function
func GetBackpointers_RPC(remote Node, from Node, level int) (neighbors []Node, err error) {
	err = makeRemoteNodeCall(remote, "GetBackpointers", GetBackpointersRequest{remote, from, level}, &neighbors)
	return
}

// Remote API: makes a remote call to the NotifyLeave function
func NotifyLeave_RPC(remote Node, from Node, replacement *Node) (err error) {
	return makeRemoteNodeCall(remote, "NotifyLeave", NotifyLeaveRequest{remote, from, replacement}, &Node{})
}

// Helper function to makes a remote call
func makeRemoteNodeCall(remote Node, method string, req interface{}, rsp interface{}) error {
	Debug.Printf("%v(%v)\n", method, req)
	return makeRemoteCall(remote.Address, "TapestryRPCServer", method, req, rsp)
}

// Helper function to makes a remote call to a cached tapestry
func makeCachedCall(address string, structtype string, method string, req interface{}, rsp interface{}) (bool, error) {
	t := getCachedTapestry(address)
	if t == nil {
		return false, nil
	}

	//fmt.Printf("Invoking %v %v %v\n", method, req, rsp)

	inputs := []reflect.Value{reflect.ValueOf(req), reflect.ValueOf(rsp)}
	interf := reflect.ValueOf(t.server).MethodByName(method).Call(inputs)[0].Interface()
	//fmt.Printf("%v\n", interf)
	if interf == nil {
		return true, nil
	} else {
		return true, interf.(error)
	}
}

// Helper function to makes a remote call
func makeRemoteCall(address string, structtype string, method string, req interface{}, rsp interface{}) error {
	// See if we can make a cached call
	cached, err := makeCachedCall(address, structtype, method, req, rsp)
	if cached {
		return err
	}

	// Dial the server
	client, err := rpc.Dial("tcp", address)
	if err != nil {
		return err
	}

	// Make the request
	fqm := fmt.Sprintf("%v.%v", structtype, method)
	err = client.Call(fqm, req, rsp)

	client.Close()
	if err != nil {
		return err
	}

	return nil
}
