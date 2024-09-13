package server

import (
	"fmt"
	"net/http"
	"os"
)

func RunServer() {
	
	serviceRoutes := HandlersMap{
		"/":               handleHome,
		"/viewRSS/":       handleViewRSS,
		"/addRSS":         handleAddRSS,
		"/addCategorie":   handleAddCategorie,
		"/getAllRSSFeeds": handlegetAllRSSFeeds,
		"/deleteRSSFeed":  handleDeleteRSSFeed,
		"/login": handleLogin,
		"/register": handleRegister,
	}

	dir, _ := os.Getwd()
	// fmt.Println("Open on", port)
	fs := http.FileServer(http.Dir(dir)) // setup the directory of files
	fmt.Println("test")

	serviceRoutes.Root()
	
	http.Handle("/rsc/", fs)
	Init_db()
	
	fmt.Println("Server running on: http://localhost:" + PORT)
	
	http.ListenAndServe(":"+PORT, nil)
	fmt.Println("test2")
}