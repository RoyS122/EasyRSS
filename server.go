package main

import (
	"fmt"
	"net/http"
	"os"
)




func main() {
    handlers_map := map[string]func(http.ResponseWriter, *http.Request){
		"/": handleHome, 	
        "/viewRSS/": handleViewRSS,
        "/addRSS": handleAddRSS,
	}

	dir, _ := os.Getwd()
	// fmt.Println("Open on", port)
	fs := http.FileServer(http.Dir(dir)) // setup the directory of files
	fmt.Println("test")

	for i, k := range handlers_map {
		fmt.Println(i)
		http.HandleFunc(i, k)
	}

	http.Handle("/rsc/", fs)

	
	fmt.Println("Server running on: http://localhost:" + PORT)
	http.ListenAndServe(":"+PORT, nil)
	fmt.Println("test2")
}