package main

import (
	"encoding/xml"
	"fmt"
	"io/ioutil"
	"net/http"
	"os"
)

func fetchRSS(url, name string) (*RSS, error) {
    // Effectuer la requête HTTP
    resp, err := http.Get(url)
    if err != nil {
        return nil, err
    }
    fmt.Println("1")
    defer resp.Body.Close()

    // Lire le corps de la réponse
    body, err := ioutil.ReadAll(resp.Body)
    if err != nil {
        return nil, err
    }
    fmt.Println("2")

    // Analyser le corps XML dans la structure RSS
    var rss RSS = RSS{Name: name}
    err = xml.Unmarshal(body, &rss)
    
    if err != nil {
        return nil, err
    }
    fmt.Println("3")

    return &rss, nil
}


func main() {
    handlers_map := map[string]func(http.ResponseWriter, *http.Request){
		"/": handleHome, 	
        "/viewRSS/": handleViewRSS,
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