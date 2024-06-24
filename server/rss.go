package server

import (
	"encoding/xml"
	"fmt"
	"io/ioutil"
	"net/http"
)

func fetchRDF(flux Flux) (*RDF, error) {
	// Effectuer la requête HTTP
	resp, err := http.Get(flux.Link)
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
	fmt.Println(string(body))
	// Analyser le corps XML dans la structure RSS
	var rdf RDF 
	
	err = xml.Unmarshal(body, &rdf)


	if err != nil {
		fmt.Println(err)
		return nil, err
	}
	// fmt.Println("3")

	return &rdf, nil
}

func fetchRSS(flux Flux) (*RSS, error) {
	// Effectuer la requête HTTP
	resp, err := http.Get(flux.Link)
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
	fmt.Println(string(body))
	// Analyser le corps XML dans la structure RSS
	var rss RSS = RSS{Name: flux.Name}
	
	err = xml.Unmarshal(body, &rss)


	if err != nil {
		return nil, err
	}
	// fmt.Println("3")

	return &rss, nil
}
