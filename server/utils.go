package server

import (
	"fmt"
	"mime/multipart"
	"net/smtp"
)


func removeFluxFromList(list []Flux, id uint) (nList []Flux) {
	for i := range list {
		if uint(i) != id {
			nList = append(nList, list[i])
		}
	}
	return nList
}

func TestImports() {
	fmt.Println("imports fonctionnels")
}

func Split(s string, sep rune) (res []string) {
	var word string
	for _, k := range s {
		if k == sep {
			res = append(res, word)
			word = ""
		} else {
			word += string(k)
		}
	}
	if len(word) > 0 {
		res = append(res, word)
	}
	return res
}

func SendMail(from, password, msg string, to []string) {
	// smtp server configuration.
	smtpHost := "smtp.gmail.com"
	smtpPort := "587"
  
	// Message.
	message := []byte(msg)
	
	// Authentication.
	auth := smtp.PlainAuth("", from, password, smtpHost)
	
	// Sending email.
	err := smtp.SendMail(smtpHost+":"+smtpPort, auth, from, to, message)
	if err != nil {
	  fmt.Println(err)
	  return
	}
	fmt.Println("Email Sent Successfully!")
}

func Atoi(str string) (res int) {
	var neg int = 1
   
	if str[0] == '-' {
	  str, neg = str[1:], -1
	}
  
	for i, k := range RevertString(str) {
   
	  res += PowInt(10, i)* int(rune(k) - '0') 
	 
	}
	
	return res * neg
  } 
  
  func RevertString(in_str string) (res string) {

	for i := len(in_str); i > 0; i -- {
	  
	  res += string(in_str[i - 1])
	} 
	return res
  } // This function revert a string
  
  func PowInt(a, b int) (res int) {
	if b == 0 {
	  return 1
	}
	res = a
	for i := 1; i < b; i ++ {
	  res = res * a
	}  
	
	return res
  } // This function get a ** b


  func removeChar(c rune, s string) (r string) {
	for _,k := range s {
		if k != c {
			r += string(k)
		}
	}
	return r
  }

  func CheckIsEmpty(s string) (r bool) {
	fmt.Println(removeChar(' ', s))
	if !(len(removeChar(' ', s)) > 0) {
		r = true
	} 
	return r
  }

  func CheckFile(file multipart.File, infos *multipart.FileHeader, ftype string, max_size int) (r bool){
	fmt.Println(infos.Header)
	fmt.Println(Split(infos.Header.Get("Content-Type"),'/')[0] == ftype)
	fmt.Println(infos.Size)
	return (Split(infos.Header.Get("Content-Type"),'/')[0] == ftype ) && (infos.Size < int64(max_size))
  }