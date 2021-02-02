package main

import (
	"golang.org/x/net/webdav"
	"log"
	"net/http"
)

func main() {
	//fs := webdav.NewMemFS()
	fs := fs{
		root: "/var/tmp/vfs",
	}
	ls := webdav.NewMemLS()
	h := webdav.Handler{
		FileSystem: fs,
		LockSystem: ls,
	}

	if err := http.ListenAndServe(":8082", &h); err != nil {
		log.Fatal(err)
	}
}
