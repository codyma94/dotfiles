package main

import (
	"bufio"
	"flag"
	"fmt"
	"log"
	"os"
	"strconv"
	"strings"

	"./chord"
)

func printHelp() {
	fmt.Println("Commands:")
	fmt.Println(" - help                    Prints this help message")
	fmt.Println(" - node                    Prints the id, predecessor and successor of for node(s)")
	fmt.Println(" - addr                    Display node listener address(es)")
	fmt.Println(" - data                    Display datastore(s) for node(s)")
	fmt.Println(" - table                   Display finger table information for node(s)")
	fmt.Println(" - put <key> <value>       Stores the provided key-value in the Chord ring")
	fmt.Println(" - get <key>               Looks up the value associated with key in the Chord ring")
	fmt.Println(" - debug on|off            Toggles debug printing statements on or off (on by default)")
	fmt.Println(" - quit                    Shutdown node(s), then quit this CLI")
}

func main() {
	countPtr := flag.Int("count", 1, "Total number of Chord nodes to start up in this process")
	addrPtr := flag.String("addr", "", "Address of a node in the Chord ring you wish to join")
	idPtr := flag.String("id", "", "ID of a node in the Chord ring you wish to join")
	flag.Parse()

	var parent *chord.RemoteNode
	if *addrPtr == "" {
		parent = nil
	} else {
		parent = new(chord.RemoteNode)
		val, _ := strconv.Atoi(*idPtr)
		parent.Id = []byte{byte(val)}
		parent.Addr = *addrPtr
		fmt.Printf("Attach this node to id:%v, addr:%v\n", parent.Id, parent.Addr)
	}

	var err error
	nodes := make([]*chord.Node, *countPtr)
	for i, _ := range nodes {
		nodes[i], err = chord.CreateNode(parent)
		if err != nil {
			fmt.Println("Unable to create new node!")
			log.Fatal(err)
		}
		if parent == nil {
			parent = nodes[i].RemoteSelf
		}
		fmt.Printf("Created -id %v -addr %v\n", chord.HashStr(nodes[i].Id), nodes[i].Addr)
	}

	chord.SetDebug(true)
	printHelp()
	for {
		fmt.Printf("> ")
		reader := bufio.NewReader(os.Stdin)
		line, _ := reader.ReadString('\n')
		line = strings.TrimSpace(line)
		args := strings.SplitN(line, " ", 3)

		cmd := args[0]
		switch cmd {
		case "node":
			for _, node := range nodes {
				fmt.Println(NodeStr(node))
			}
		case "table":
			for _, node := range nodes {
				chord.PrintFingerTable(node)
			}
		case "addr":
			for _, node := range nodes {
				fmt.Printf("Node %v: %v\n", node.Id, node.Addr)
			}
		case "data":
			for _, node := range nodes {
				chord.PrintDataStore(node)
			}
		case "get":
			if len(args) > 1 {
				val, err := chord.Get(nodes[0], args[1])
				if err != nil {
					fmt.Println(err)
				} else {
					fmt.Println(val)
				}
			} else {
				fmt.Printf("USAGE: %v <key>\n", cmd)
			}
		case "put":
			if len(args) > 2 {
				err := chord.Put(nodes[0], args[1], args[2])
				if err != nil {
					fmt.Println(err)
				}
			} else {
				fmt.Printf("USAGE: %v <key> <value>\n", cmd)
			}
		case "debug":
			if len(args) == 2 {
				val := strings.ToLower(args[1])
				if val == "on" || val == "true" {
					fmt.Println("Debug statements turned on")
					chord.SetDebug(true)
				} else if val == "off" || val == "false" {
					fmt.Println("Debug statements turned off")
					chord.SetDebug(false)
				} else {
					fmt.Printf("USAGE: %v on|off\n", cmd)
				}
			} else {
				fmt.Printf("USAGE: %v on|off\n", cmd)
			}

		case "quit":
			fmt.Println("goodbye")
			for _, node := range nodes {
				chord.ShutdownNode(node)
			}
			return
		case "help":
			printHelp()
		case "":
			continue
		default:
			fmt.Println("invalid command")
			continue
		}
	}
}

func NodeStr(node *chord.Node) string {
	var succ []byte
	var pred []byte
	if node.Successor != nil {
		succ = node.Successor.Id
	}
	if node.Predecessor != nil {
		pred = node.Predecessor.Id
	}

	return fmt.Sprintf("Node %v: {succ:%v, pred:%v}", node.Id, succ, pred)

}
