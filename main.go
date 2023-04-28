package main

import (
	"flag"
	"fmt"
	"log"
	"net/http"
	"os"
	"path"
)

const (
	DEFAULT_PORT       = 3000
	DEFAULT_ADDR       = ""
	DEAFULT_STATIC_DIR = ""
	USAGE              = `
Usage: %s
Run Simple HTTP file server. Serves contents of a directory, defaults to index.html if path not found.

Options:
`
)

func FileServerWithDefaultFile(fs http.FileSystem) http.Handler {
	fsh := http.FileServer(fs)
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		_, err := fs.Open(path.Clean(r.URL.Path))
		if os.IsNotExist(err) {
			r.URL.Path = "/"
		}
		fsh.ServeHTTP(w, r)
	})
}

func main() {

	portPtr := flag.Int("p", DEFAULT_PORT, "tcp binding port")
	addrPtr := flag.String("a", DEFAULT_ADDR, "tcp binding addr")
	staticDirPtr := flag.String("d", DEAFULT_STATIC_DIR, "path to directory to serve statically")

	flag.Usage = func() {
		fmt.Fprintf(flag.CommandLine.Output(), USAGE, os.Args[0])
		flag.PrintDefaults()
	}
	flag.Parse()

	listenOn := fmt.Sprintf("%s:%d", *addrPtr, *portPtr)
	log.Printf("Serving contents of %s on %s", *staticDirPtr, listenOn)

	log.Fatal(http.ListenAndServe(listenOn, FileServerWithDefaultFile(http.Dir(*staticDirPtr))))
}
