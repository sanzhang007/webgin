package protocol

import (
	"log"
	"net/url"
)

type Trojan struct {
	Id             int `gorm:"primaryKey"`
	Name           string
	Password       string
	Port           string
	Server         string `gorm:"unique"`
	SkipCertVerify bool
	Type           string
	Udp            bool
	Link           string
}

func (trojan *Trojan) Parse(urlstring string) {
	u, err := url.Parse(urlstring)
	if err != nil {
		log.Fatal(err)
	}

	trojan.Name = u.Fragment
	trojan.Type = "trojan"
	trojan.Port = u.Port()
	trojan.Server = u.Hostname()
	trojan.Password = u.User.String()
	query := u.Query()
	if query.Get("allowInsecure") == "1" {
		trojan.SkipCertVerify = true
	}
	trojan.Link = urlstring
}
