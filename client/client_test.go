package main

import (
	"fmt"
	"net/rpc"
	"testing"
)

func BenchmarkCalls(b *testing.B) {
	c, err := rpc.Dial("tcp", "127.0.0.1:9999")
	if err != nil {
		fmt.Println(err)
		return
	}
	b.ResetTimer()
	var result int64
	for i := 0; i < b.N; i++ {
		err = c.Call("Server.Negate", int64(999), &result)
		if err != nil {
			fmt.Println(err)
		} else {
			fmt.Println("Server.Negate(999) =", result)
		}
	}
}
