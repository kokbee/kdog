package main

import (
	"bytes"
	"encoding/xml"
	"fmt"
	"html/template"
	"io/ioutil"
	"log"
	"net/http"
)

func clientUser() (*User, error) {
	resp, err := http.Get("http://speedtest.net/speedtest-config.php")
	if err != nil {
		log.Fatal(err)
	}
	body, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		log.Fatal(err)
	}
	defer resp.Body.Close()

	// Decode xml
	decoder := xml.NewDecoder(bytes.NewReader(body))
	users := Users{}
	for {
		t, _ := decoder.Token()
		if t == nil {
			break
		}
		switch se := t.(type) {
		case xml.StartElement:
			decoder.DecodeElement(&users, &se)
		}
	}
	if users.Users == nil {
		log.Fatal(err)
	}

	return &users.Users[0], nil
}

// String representation of User
func (u *User) String() string {
	return fmt.Sprintf("%s, (%s) [%s, %s]", u.IP, u.Isp, u.Lat, u.Lon)
}

func handleMain(w http.ResponseWriter, r *http.Request) {
	i, err := template.ParseFiles(
		"assets/html/header.html",
		"assets/html/index.html",
		"assets/html/footer.html",
	)
	if err != nil {
		log.Fatal(err)
	}

	user, err := clientUser()
	if err != nil {
		log.Fatal(err)
	}

	userData := new(User)
	userData.IP = user.IP
	userData.Isp = user.Isp
	userData.Lat = user.Lat
	userData.Lon = user.Lon

	fmt.Println(userData)

	err = i.ExecuteTemplate(w, "index", userData)
	if err != nil {
		log.Fatal(err)
	}
}

func webserver() {
	http.HandleFunc("/", handleMain)

	fs := http.FileServer(http.Dir("assets/"))
	http.Handle("/assets/", http.StripPrefix("/assets/", fs))

	err := http.ListenAndServe(*flagHTTPPort, nil)
	if err != nil {
		log.Fatal(err)
	}
}
