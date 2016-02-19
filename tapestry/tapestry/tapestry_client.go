package tapestry

import (
	"fmt"
)

// Invoke tapestry.Store on a remote tapestry node
func TapestryStore(remote Node, key string, value []byte) error {
	Debug.Printf("Making remote TapestryStore call\n")
	return makeRemoteNodeCall(remote, "TapestryStore", StoreRequest{remote, key, value}, &StoreResponse{})
}

// Invoke tapestry.Lookup on a remote tapestry node
func TapestryLookup(remote Node, key string) (nodes []Node, err error) {
	Debug.Printf("Making remote TapestryLookup call\n")
	var rsp LookupResponse
	err = makeRemoteNodeCall(remote, "TapestryLookup", LookupRequest{remote, key}, &rsp)
	nodes = rsp.Nodes
	return
}

// Get data from a tapestry node.  Looks up key then fetches directly
func TapestryGet(remote Node, key string) ([]byte, error) {
	Debug.Printf("Making remote TapestryGet call\n")
	// Lookup the key
	replicas, err := TapestryLookup(remote, key)
	if err != nil {
		return nil, err
	}
	if len(replicas) == 0 {
		return nil, fmt.Errorf("No replicas returned for key %v", key)
	}

	// Contact replicas
	var errs []error
	for _, replica := range replicas {
		blob, err := FetchRemoteBlob(replica, key)
		if err != nil {
			errs = append(errs, err)
		}
		if blob != nil {
			return *blob, nil
		}
	}

	return nil, fmt.Errorf("Error contacting replicas, %v: %v", replicas, errs)
}

type StoreRequest struct {
	To    Node
	Key   string
	Value []byte
}
type StoreResponse struct {
}

type LookupRequest struct {
	To  Node
	Key string
}
type LookupResponse struct {
	Nodes []Node
}

// Server: extension method to open up Store via RPC
func (server *TapestryRPCServer) TapestryStore(req StoreRequest, rsp *StoreResponse) (err error) {
	Debug.Printf("Received remote invocation of Tapestry.Store\n")
	return server.tapestry.Store(req.Key, req.Value)
}

// Server: extension method to open up Lookup via RPC
func (server *TapestryRPCServer) TapestryLookup(req LookupRequest, rsp *LookupResponse) (err error) {
	Debug.Printf("Received remote invocation of Tapestry.Lookup\n")
	rsp.Nodes, err = server.tapestry.Lookup(req.Key)
	return
}
