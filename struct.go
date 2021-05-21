package main

// User represents information determined about the caller by speedtest.net
type User struct {
	IP  string `xml:"ip,attr"`
	Lat string `xml:"lat,attr"`
	Lon string `xml:"lon,attr"`
	Isp string `xml:"isp,attr"`
}

// Users for decode xml
type Users struct {
	Users []User `xml:"client"`
}
