package helpers

import (
	"fmt"
	"log"
	"net/http"
)

func Serve(path string, port string) {
	fmt.Println(path)
	fmt.Println("Starting serving at: ", port)
	http.Handle("/", http.FileServer(http.Dir(path)))
	err := http.ListenAndServe(fmt.Sprintf(":%v", port), nil)
	if err != nil {
		log.Fatal(err)
	}
}
