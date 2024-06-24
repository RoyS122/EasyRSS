package server

import (
	"fmt"
	"net/http"
	"time"

	"github.com/gofrs/uuid"
	"golang.org/x/crypto/bcrypt"
)


func getUserBySessionId(session_id string) (u User) {
	db := Open_DB()
	defer db.Close()
	if session_id == "NULL" {
		return User{}
	}
	var _exp_date string
	err := db.QueryRow("SELECT user_id, username, email, session_expiration FROM users WHERE users.session_id = ?", session_id).Scan(&u.Id, &u.Username, &u.Email, &_exp_date)
	if err != nil {
		fmt.Println(err)
	}
	_exp_date_time, err := time.Parse(TIME_LAYOUT, _exp_date)
	if err != nil {
		fmt.Println(err)
		return User{}
	}
	if _exp_date_time.Before(time.Now()) {
		db.Exec(`UPDATE users SET session_id = NULL, session_expiration = NULL WHERE user_id = ?`, u.Id)
		return User{}
	}

	return
}

func (u User) Register(password string) {
	// Commandes pour enregistrer dans la db

	db := Open_DB()
	defer db.Close()
	var st string = `INSERT INTO users(user_id, username, password, email, date) VALUES (?, ?, ?, ?, ?)`
	req, err := db.Prepare(st)
	if err != nil {
		fmt.Println(err)
		return
	}
	passwordHash, _ :=  bcrypt.GenerateFromPassword([]byte(password), 10)
	id, _ := uuid.NewV4() 
	req.Exec(id, u.Username, string(passwordHash), u.Email, time.Now().Format(TIME_LAYOUT))
}

func checkUserExists(u User) (res [2]bool) {
	db := Open_DB()
	defer db.Close()

	var username int
	db.QueryRow("SELECT COUNT(*) FROM users WHERE username=?", u.Username).Scan(&username)

	var email int
	db.QueryRow("SELECT COUNT(*) FROM users WHERE email=?", u.Email).Scan(&email)

	return [2]bool{username > 0, email > 0}
}


func GetUserInformationsByCookies(cookies []*http.Cookie) (u User) {
	for _, k := range cookies {
		if k.Name == "sessionID" {
			u = getUserBySessionId(k.Value)
		}
	}
	return u
}



func (u User) createConnexionCookies() (cookie http.Cookie) {
	_uuid, _ := uuid.NewV4()

	db := Open_DB()
	defer db.Close()

	var st string = `UPDATE users SET session_id = ?, session_expiration = ? WHERE user_id = ?`
	req, err := db.Prepare(st)
	if err != nil {
		fmt.Println(err)
		return
	}
	_expirationdate := time.Now().Add(time.Hour * 24)
	req.Exec(_uuid, _expirationdate.Format(TIME_LAYOUT), u.Id)
	cookie = http.Cookie{
		Name:    "sessionID",
		Value:   _uuid.String(),
		Expires: _expirationdate,
	}
	return
}

func (u User) logout() {
	db := Open_DB()
	db.Exec(`UPDATE users SET session_id = NULL, session_expiration = NULL WHERE user_id = ?`, u.Id)
	db.Close()
}


// func login(req ReqLogin) (err_ret string, session http.Cookie) {
// 	db := Open_DB()
// 	defer db.Close()

// 	var _pass string

// 	var u User
// 	getEmail := false

// 	for _, v := range req.EmailOrUsername {
// 		if v == '@' {
// 			getEmail = true
// 			break
// 		}
// 	}

// 	if getEmail {
// 		u = getUserByMail(req.EmailOrUsername)
// 	} else {
// 		u = getUserByUsername(req.EmailOrUsername)
// 	}

// 	err := db.QueryRow("SELECT password FROM users WHERE users.user_id = ?", u.Id).Scan(&_pass)
// 	if err != nil {
// 		return "", session
// 	}

// 	err = bcrypt.CompareHashAndPassword([]byte(_pass), []byte(req.Password))
// 	if err != nil {
// 		return "", session
// 	}
// 	session = u.createConnexionCookies()
// 	return err_ret, session
// }

// func getUserByMail(mail string) (u User) {
// 	db := Open_DB()
// 	defer db.Close()
// 	err := db.QueryRow("SELECT user_id, username FROM users WHERE users.email = ?", mail).Scan(&u.Id, &u.Username)
// 	if err != nil {
// 		fmt.Println(err)
// 	}
// 	return u
// }

// func getUserByUsername(username string) (u User) {
// 	db := Open_DB()
// 	defer db.Close()
// 	err := db.QueryRow("SELECT user_id, email FROM users WHERE users.username = ?", username).Scan(&u.Id, &u.Email)
// 	if err != nil {
// 		fmt.Println(err)
// 	}
// 	return u
// }
