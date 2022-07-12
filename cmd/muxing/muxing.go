package main

import (
	"fmt"
	"io"
	"log"
	"net/http"
	"os"
	"strconv"

	"github.com/gorilla/mux"
)

/**
Please note Start functions is a placeholder for you to start your own solution.
Feel free to drop gorilla.mux if you want and use any other solution available.

main function reads host/port from env just for an example, flavor it following your taste
*/

// Start /** Starts the web server listener on given host and port.
func Start(host string, port int) {
	router := mux.NewRouter()

	router.HandleFunc("/name/{param}", GetName).Methods("GET")
	router.HandleFunc("/bad", GetBadRequest).Methods("GET")
	router.HandleFunc("/data", PostData).Methods("POST")
	router.HandleFunc("/headers", PostHeader).Methods("POST")

	log.Println(fmt.Printf("Starting API server on %s:%d\n", host, port))
	if err := http.ListenAndServe(fmt.Sprintf("%s:%d", host, port), router); err != nil {
		log.Fatal(err)
	}
}

func GetName(w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)
	resp := "Hello, " + vars["param"] + "!"
	w.Write([]byte(resp))

}

func PostData(w http.ResponseWriter, r *http.Request) {
	data, err := io.ReadAll(r.Body)
	if err == nil {
		resp := "I got message:\n" + string(data)
		w.Write([]byte(resp))
	}
}

func GetBadRequest(w http.ResponseWriter, r *http.Request) {
	w.WriteHeader(http.StatusInternalServerError)
}

func PostHeader(w http.ResponseWriter, r *http.Request) {
	head := r.Header

	sum := 0
	for _, key := range []string{"A", "B"} {
		if val, ok := head[key]; ok {
			val_int, err := strconv.Atoi(val[0])
			if err != nil {
				w.WriteHeader(http.StatusBadRequest)
				return
			}
			sum += val_int
		}
		w.Header().Set("a+b", strconv.Itoa(sum))
	}
}

//main /** starts program, gets HOST:PORT param and calls Start func.
func main() {
	host := os.Getenv("HOST")
	port, err := strconv.Atoi(os.Getenv("PORT"))
	if err != nil {
		port = 8081
	}
	Start(host, port)
}
