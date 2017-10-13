package main

import (
	"flag"
	"log"
)

var (
	listen  = flag.String("listen", ":9090", "Address to listen on")
	address = flag.String("address", "http://localhost:9090/greeter", "Address to listen on")
	mode    = flag.String("mode", "client", "Mode: server/client")
)

func main() {
	flag.Parse()
	switch *mode {
	case "client":
		if len(flag.Args()) != 1 {
			log.Fatal("must provide a name")
		}
		name := flag.Arg(0)
		client(*address, name)
	case "server":
		server(*listen)
	default:
		log.Fatalf("Unknow mode: %s", *mode)
	}
}
