package reqdump

import (
	"fmt"
	"log"
	"net/http"
)

func Handle(w http.ResponseWriter, r *http.Request) {
	fr := fmt.Sprintf("%+v", r)
	log.Println(fr)
	w.Write(append([]byte("Congratulations, you are being spied on:"), []byte(fr)...))
}
