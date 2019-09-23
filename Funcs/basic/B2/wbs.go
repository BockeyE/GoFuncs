package B2

import (
	"fmt"
	"log"
	"net/http"
)

func main() {
	http.HandleFunc("/hello", handleHello)
	fmt.Println("SERVING on 7777")
	log.Fatal(http.ListenAndServe("localhost:7777", nil))
}

func handleHello(writer http.ResponseWriter, request *http.Request) {
	log.Println("serving", request.URL)
	fmt.Fprintln(writer, "HELLO WORLD")
}

func handleSearch(w http.ResponseWriter, r *http.Request) {
	log.Println("serving", r.URL)
	query := r.FormValue("q")
	if query == "" {
		http.Error(w, "q missed", http.StatusBadRequest)
		return
	}

}
