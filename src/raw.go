package main

import (
	"fmt"
	"log"

	"github.com/valyala/fasthttp"
)

func main() {
	if err := fasthttp.ListenAndServe(":8088", fastHTTPRawHandler); err != nil {
		log.Fatalf("Error in ListenAndServe: %s", err)
	}
}

func fastHTTPRawHandler(ctx *fasthttp.RequestCtx) {
	if string(ctx.Method()) == "GET" {
		switch string(ctx.Path()) {
		case "/":
			fmt.Fprintf(ctx, "Hello, world!\n\n")
		default:
			ctx.Error("Unsupported path", fasthttp.StatusNotFound)
		}
		return
	}
	ctx.Error("Unsupported method", fasthttp.StatusMethodNotAllowed)
}
