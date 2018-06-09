package main

import (
	"context"
	"fmt"
	"net"
	"net/http"

	"github.com/guregu/kami"
)

func main() {
	kami.Get("/", func(ctx context.Context, w http.ResponseWriter, r *http.Request) {
		fmt.Fprintf(w, "Hello World!")
	})

	l, _ := net.Listen("tcp", ":8088")
	kami.ServeListener(l)
}
