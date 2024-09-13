package server

import (
	"encoding/json"
	"fmt"
	"net/http"
	"text/template"

	"github.com/gofrs/uuid"
	"golang.org/x/crypto/bcrypt"
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


		//var listFlux []Flux = getFluxFromJson( "/rsc/json/fluxlist.json")
		//fmt.Println(dir + "rsc/json/fluxlist.json")
		var pD  PageData
		var cUser User = GetUserInformationsByCookies(r.Cookies());

		var listFlux []Flux = cUser.GetFlux();
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
		var u User = GetUserInformationsByCookies(r.Cookies()) 
		var listFlux = u.GetFlux()

		var sURL []string = Split(r.URL.String(), '/')
		var pD PageData
	//	fmt.Println("super mega test rss")
		var fID int = Atoi(sURL[2])
	//	fmt.Println("super mega test rss", listFlux[RSSId])
		switch(listFlux[fID].Version) {
		case "1.0":
			pD.RDFData = make(map[string]*RDF)
			pD.RDFData["currentRDF"], _ = fetchRDF(listFlux[fID])
		case "2.0": 
			pD.RSSData = make(map[string]*RSS)
			pD.RSSData["currentRSS"], _ = fetchRSS(listFlux[fID])
		default:
			pD.RSSData = make(map[string]*RSS)
			pD.RSSData["currentRSS"], _ = fetchRSS(listFlux[fID])
		}
		

		fmt.Println("su^per mega test qui tue	")

		
		fmt.Println(pD.RSSData["currentRSS"])
		fmt.Println(pD.RDFData["currentRDF"])
		fmt.Println("su^per mega test qui tue	2")
		createResponse(w, "rsc/html/viewRSS.html", pD)
	default:
		errorHandler(w, r, http.StatusBadRequest)
	}
}

func handleAddRSS(w http.ResponseWriter, r *http.Request) {
	
	switch(r.Method) {
	case "POST":
		var u User = GetUserInformationsByCookies(r.Cookies()) 

		if u.Username == "" {
			w.WriteHeader(http.StatusBadRequest)
			return;
		}
		var nRSS Flux
		err := json.NewDecoder(r.Body).Decode(&nRSS)
		if err != nil {
			fmt.Println(err)
			return
		}
		u.AddFlux(nRSS)
		w.WriteHeader(http.StatusOK)
	}
}

func handleAddCategorie(w http.ResponseWriter, r *http.Request) {
	switch(r.Method) {
	case "POST":
	
		var nCate Categorie
		err := json.NewDecoder(r.Body).Decode(&nCate)
		if err != nil {
			fmt.Println(err)
			return
		}
		var lCate []Categorie = getCategoriesFromJson("/rsc/json/categories.json")
		lCate = append(lCate, nCate)
		nString, _ := json.Marshal(lCate)
		err = writeFile("/rsc/json/categories.json", string(nString))
		if err != nil {
			fmt.Println(err)
			return
		}
		http.Redirect(w, r, "/", http.StatusSeeOther)
	}
}

func handlegetAllRSSFeeds(w http.ResponseWriter, r *http.Request) {
	switch(r.Method) {
	case "GET":
		var u User = GetUserInformationsByCookies(r.Cookies()) 
	
		fList := u.GetFlux()
		
		nText, _ := json.Marshal(fList)
		w.Header().Set("Content-Type", "application/json")
		w.Write(nText)
		fmt.Println(string(nText))
		
	default:
		return
	}
}

func handleDeleteRSSFeed(w http.ResponseWriter, r *http.Request) {
	switch(r.Method) {
	case "POST":
		var data map[string]interface{}
		err := json.NewDecoder(r.Body).Decode(&data)
		fmt.Println(data)
		if err != nil {
			fmt.Println(err)
		}
		//fmt.Println(r)
		id, ok := data["Id"].(float64) // JSON numbers are floats
		fmt.Println(id)
		nId := uint(id)
		if !ok {
			http.Error(w, "Invalid ID", http.StatusBadRequest)
			return
		}
		fmt.Println(nId)
		flux := getFluxFromJson("/rsc/json/fluxlist.json")
		flux = removeFluxFromList(flux, nId)
		nText, _ := json.Marshal(flux)
		fmt.Println(flux)
		writeFile("/rsc/json/fluxlist.json", string(nText))
		
		w.WriteHeader(http.StatusOK)
		
	default: 
		errorHandler(w, r, http.StatusBadRequest)
	}
}

func handleRegister(w http.ResponseWriter, r *http.Request) {
	var pD PageData
	switch (r.Method) {
	case "GET":
		createResponse(w, "rsc/html/login_register.html", pD)
	case "POST":
		//	fmt.Println("test2")
		var (
			_username              string = r.FormValue("username")
			_email                 string = r.FormValue("email")
			_password              string = r.FormValue("password")
			_password_confirmation string = r.FormValue("password_confirmation")
		)

		for _, v := range r.Form {
			for _, val := range v {
				if CheckIsEmpty(val) {
					pD.Errors = append(pD.Errors, "Empty value")
				}
			}
		}
		// fmt.Println(_username, _email, _password, _password_confirmation)
		if _password != _password_confirmation {
			pD.Errors = append(pD.Errors, "Not same password")
		}

		
		

		var _user User = User{
			Username: _username,
			Email:    _email,
		}
		var check_user [2]bool = checkUserExists(_user)
		if check_user[0] {
			pD.Errors = append(pD.Errors, "username already taken")
		}
		if check_user[1] {
			pD.Errors = append(pD.Errors, "email already taken")
		}

		if len(pD.Errors) != 0  {
			createResponse(w, "rsc/html/login_register.html", pD)
			return;
		}

		id, _ := uuid.NewV4()
		
		_user.Id = id.String();

		pswd, _ := bcrypt.GenerateFromPassword([]byte(_password), 6)
		_user.Register(string(pswd))
		var conn_cookie http.Cookie = _user.createConnexionCookies()
		http.SetCookie(w, &conn_cookie)
		http.Redirect(w, r, "/", http.StatusSeeOther)
		//handleIndex(w, r)
		return
	default:
		// Display error
	} 
}


func handleLogin(w http.ResponseWriter, r *http.Request) {
	var pD PageData
	switch (r.Method) {
	case "GET":
		createResponse(w, "rsc/html/login_register.html", pD)
	case "POST":
	//	fmt.Println("y test")
		var (
			_email    string = r.FormValue("email")
			_password string = r.FormValue("password")
		)
		var _user User = getUserByMail(_email)
		err, sess := _user.login(_password)
		
		if err != "" {
			pD.Errors = append(pD.Errors, err)
		}
			
		if len(pD.Errors) != 0 {
			createResponse(w, "rsc/html/login_register.html", pD)
			return
		}
		
		http.SetCookie(w, &sess)
		http.Redirect(w, r, "/", http.StatusSeeOther)
		//handleIndex(w, r)
		return
	default:
		// Display error
	} 
}