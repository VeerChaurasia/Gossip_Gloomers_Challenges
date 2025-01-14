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

	
	n.Handle("add",func(msg maelstrom.Message)error{
		var body map[string] any
		if err:=json.Unmarshal(msg.Body,&body);err!=nil{
			return err
		}
		delta := body["delta"].(float64)
		for{
			v, err1 := kv.ReadInt(context.Background(), "counter")
			if err1 != nil {
				if maelstrom.ErrorCode(err1) == maelstrom.KeyDoesNotExist {
					v = 0
				} else {
					return err1
				}
			}
			value := v + int(delta)
			err2 := kv.CompareAndSwap(context.Background(), "counter", v, value, true)
			if err2 == nil {
				break
			}
		}
		body["type"] = "add_ok"
		delete(body,"delta")
		return n.Reply(msg,body)

	})
	n.Handle("read",func(msg maelstrom.Message)error{
		var body map[string] any
		if err:=json.Unmarshal(msg.Body,&body);err!=nil{
			return err
		}
		counter, err := kv.ReadInt(context.Background(), "counter")
		if err != nil {
		if maelstrom.ErrorCode(err) == maelstrom.KeyDoesNotExist {
			counter = 0
		} else {
			return err
		}
		}
		body["type"]="read_ok"
		body["value"]=counter
		return n.Reply(msg,body)

	})
	if err := n.Run(); err != nil {
		log.Fatal(err)
	}
}