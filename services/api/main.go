package main

import (
	"flag"
	"log"
	"net/http"

	add "github.com/vems/pb/add"

	"google.golang.org/grpc"
)

type client struct {
	add.AddClient
}

func mustDial(addr *string) *grpc.ClientConn {
	conn, err := grpc.Dial(*addr, grpc.WithInsecure())

	if err != nil {
		log.Fatalf("failed to dial: %v", err)
		panic(err)
	}

	return conn
}

func main() {
	var (
		port    = flag.String("port", "8080", "The server port")
		addAddr = flag.String("add.addr", "add:8080", "The Add server address in the format of host:port")
	)
	flag.Parse()

	c := client{
		AddClient: add.NewAddClient(mustDial(addAddr)),
	}

	http.HandleFunc("/sum", http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		sumRequestHandler(c, w, r)
	}))

	log.Fatal(http.ListenAndServe(":"+*port, nil))
}
