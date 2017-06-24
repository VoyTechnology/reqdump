package main

import (
	"flag"
	"fmt"
	"net/http"

	"github.com/voytechnology/reqdump"
)

var (
	portFlag = flag.Int("port", 8080, "port to listen on")
)

func main() {
	http.HandleFunc("/", reqdump.Handle)
	http.ListenAndServe(fmt.Sprintf(":%d", *portFlag), nil)
}
