package main

import (
	"encoding/json"
	"fmt"
	"net/http"
	"text/template"
)

func createResponse(w http.ResponseWriter, templatePath string, data PageData) {
	t, err := template.ParseFiles(templatePath)
	if err != nil {
		fmt.Println(err)
		return
	}
//	fmt.Println("test")
	t.Execute(w, data)
}

func handleHome(w http.ResponseWriter, r *http.Request) {
	switch(r.Method){
	case "GET":
		var listFlux []Flux = getFluxFromJson( "/rsc/json/fluxlist.json")
		//fmt.Println(dir + "rsc/json/fluxlist.json")
		var pD  PageData
		fmt.Println("test")
		pD.RSSFluxArrays = make(map[string][]Flux)
		pD.RSSFluxArrays["listFlux"] = listFlux
		fmt.Println("test")
		createResponse(w, "rsc/html/index.html", pD)
	
	default: 
		errorHandler(w, r, http.StatusBadRequest)
		return
	}
	
}

func handleViewRSS(w http.ResponseWriter, r *http.Request) {
	switch(r.Method) {
	case "GET":
		var listFlux = getFluxFromJson("/rsc/json/fluxlist.json")
		var sURL []string = Split(r.URL.String(), '/')
		var pD PageData
	//	fmt.Println("super mega test rss")
		var RSSId int= Atoi(sURL[2])
	//	fmt.Println("super mega test rss", listFlux[RSSId])
		pD.RSSData = make(map[string]*RSS)
		pD.RSSData["currentRSS"], _ = fetchRSS(listFlux[RSSId].Link, listFlux[RSSId].Name)
		fmt.Println(pD.RSSData["currentRSS"].Channel.Items[0])
		createResponse(w, "rsc/html/viewRSS.html", pD)
	default:
		errorHandler(w, r, http.StatusBadRequest)
	}
}

func handleAddRSS(w http.ResponseWriter, r *http.Request) {
	
	switch(r.Method) {
	case "POST":
	
		var nRSS Flux
		err := json.NewDecoder(r.Body).Decode(&nRSS)
		if err != nil {
			fmt.Println(err)
			return
		}
		var lFlux []Flux = getFluxFromJson("/rsc/json/fluxlist.json")
		lFlux = append(lFlux, nRSS)
		nString, _ := json.Marshal(lFlux)
		err = writeFile("/rsc/json/fluxlist.json", string(nString))
		if err != nil {
			fmt.Println(err)
			return
		}
		http.Redirect(w, r, "/", http.StatusOK)
	}
}