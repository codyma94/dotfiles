/*                                                                           */
/*  Brown University, CS138, Spring 2016                                     */
/*                                                                           */
/*  Purpose: API and interal functions to interact with the Key-Value store  */
/*           that the Chord ring is providing.                               */
/*                                                                           */

package chord

import (
	"fmt"
)

/*                             */
/* External API Into Datastore */
/*                             */

// Get a value in the datastore, given an abitrary node in the ring.
func Get(node *Node, key string) (string, error) {

	//TODO students should implement this method
	return "", nil
}

// Put a key/value in the datastore, given an abitrary node in the ring.
func Put(node *Node, key string, value string) error {

	//TODO students should implement this method
	return nil
}

// Internal helper method to find the appropriate node in the ring based on a key.
func (node *Node) locate(key string) (*RemoteNode, error) {

	//TODO students should implement this method
	return nil, nil
}

/*                                                         */
/* RPCs to assist with interfacing with the datastore ring */
/*                                                         */

/* RPC */
func (node *Node) GetLocal(req *KeyValueReq, reply *KeyValueReply) error {
	if err := validateRpc(node, req.NodeId); err != nil {
		return err
	}

	//TODO students should implement this method
	return nil
}

/* RPC */
func (node *Node) PutLocal(req *KeyValueReq, reply *KeyValueReply) error {
	if err := validateRpc(node, req.NodeId); err != nil {
		return err
	}

	//TODO students should implement this method
	return nil
}

/* RPC */
// Find locally stored keys that are between (predId : fromId].
// Any of these nodes should be moved to fromId.
func (node *Node) TransferKeys(req *TransferReq, reply *RpcOkay) error {
	if err := validateRpc(node, req.NodeId); err != nil {
		return err
	}

	//TODO students should implement this method
	return nil
}

// Print the contents of a node's datastore.
func PrintDataStore(node *Node) {
	node.DsLock.RLock()
	defer node.DsLock.RUnlock()
	fmt.Printf("Node %v datastore: %v\n", HashStr(node.Id), node.dataStore)
}
