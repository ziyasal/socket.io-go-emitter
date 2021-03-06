//A Go implementation of socket.io-emitter
//
//socket.io provides a hook point to easily allow you to emit events to browsers from anywhere so
//socket.io-go-emitter communicates with socket.io servers through redis.
//
//		package main
//
//		import (
//			"fmt"
//
//			"github.com/ziyasal/socket.io-go-emitter/emitter"
//		)
//
//		func main() {
//			opts := emitter.Options{
//				Host: "127.0.0.1:6379",
//				Key:  "socket.io",
//			}
//
//			sio := emitter.New(opts)
//
//			fmt.Println("Emit :", sio.Emit("broadcast event", "Hello from socket.io-go-emitter"))
//
//			fmt.Println("In Emit :", sio.In("test").Emit("broadcast event", "Hello from socket.io-go-emitter"))
//
//			fmt.Println("To Emit : ", sio.To("test").Emit("broadcast event", "Hello from socket.io-go-emitter"))
//
//			fmt.Println("Of Emit :", sio.Of("/nsp").Emit("broadcast event", "Hello from socket.io-go-emitter"))
//		}
//
//

package emitter
