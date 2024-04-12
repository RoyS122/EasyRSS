package main

import (
	"encoding/xml"
	"io/ioutil"
	"net/http"
)

func fetchRSS(url, name string) (*RSS, error) {
	// Effectuer la requête HTTP
	resp, err := http.Get(url)
	if err != nil {
		return nil, err
	}
	// fmt.Println("1")
	defer resp.Body.Close()

	// Lire le corps de la réponse
	body, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		return nil, err
	}
	// fmt.Println("2")

	// Analyser le corps XML dans la structure RSS
	var rss RSS = RSS{Name: name}
	err = xml.Unmarshal(body, &rss)

	if err != nil {
		return nil, err
	}
	// fmt.Println("3")

	return &rss, nil
}
