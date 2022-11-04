package main

import (
	"flag"
	"fuser/fuser"
	"log"
	"os"
	"os/signal"
	"syscall"

	"bazil.org/fuse"
	"bazil.org/fuse/fs"
)

var in fs.FS = fuser.ParseSource(flag.String("in", "./openapi.json", "Source for the filesystem"))

var out string = fuser.ParseDestination(flag.String("out", "./mnt", "Destination for filesystem"))

func main() {
	conn, err := fuse.Mount(out)
	if err != nil {
		log.Fatal(err)
	}

	sig := make(chan os.Signal)
	signal.Notify(sig, syscall.SIGTERM, syscall.SIGINT)

	server := fs.New(conn, nil)

	go server.Serve(in)

	<-sig

	conn.Close()
}
