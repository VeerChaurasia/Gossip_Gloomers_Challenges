package main

import (
    "encoding/json"
    "log"
	maelstrom "github.com/jepsen-io/maelstrom/demo/go"
)
func main(){
	n:=maelstrom.NewNode()
	var arr[] int
	n.Handle("broadcast",func(msg maelstrom.Message)error{
		var body struct{
			Type string `json:"type"`
			Message int `json:"message"`
		}
		if err:=json.Unmarshal(msg.Body, &body);err!=nil{
			return err

		}
		message:=body.Message
		arr = append(arr, message)
		return n.Reply(msg,struct{
			Type string `json:"type"`
		}{
			Type:"broadcast_ok",
		})

	})
	n.Handle("read",func(msg maelstrom.Message)error{
		var body struct{
			Type string `json:"type"`
		}
		if err:=json.Unmarshal(msg.Body, &body);err!=nil{
			return err

		}
		return n.Reply(msg,struct{
			Type string `json:"type"`
			Messages []int	`json:"messages"`
		}{
			Type:"read_ok",
			Messages:arr,
		})

	})
	n.Handle("topology",func(msg maelstrom.Message)error{
		var body struct{
			Type string `json:"type"`
		}
		if err:=json.Unmarshal(msg.Body, &body);err!=nil{
			return err

		}
		return n.Reply(msg,struct{
			Type string `json:"type"`
		}{
			Type:"topology_ok",
		})
		

	})
	if err := n.Run(); err != nil {
		log.Fatal(err)
	}

	
}