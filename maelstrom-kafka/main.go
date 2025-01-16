package main

import (
	"context"
	"encoding/json"
	"log"

	maelstrom "github.com/jepsen-io/maelstrom/demo/go"
)
func main(){
	n:=maelstrom.NewNode()
	kv := maelstrom.NewSeqKV(n)
}