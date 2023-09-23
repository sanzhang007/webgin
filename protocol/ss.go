package protocol

import (
	"encoding/base64"
	"fmt"
	"net"
	"net/url"
	"regexp"
	"strings"

	"github.com/sanzhang007/webgin/base64decode"
	"github.com/xxf098/lite-proxy/utils"
)

type Ss struct {
	Id       int `gorm:"primaryKey"`
	Name     string
	Password string
	Port     string
	Server   string `gorm:"unique"`
	Cipher   string
	Type     string
	Link     string `gorm:"size:1024"`
}

func (ss *Ss) ParseTemp(urlstring string) {
	builder := new(strings.Builder)
	urlSpilt := strings.Split(urlstring, "://")
	builder.WriteString(urlSpilt[0])
	builder.WriteString("://")
	r := regexp.MustCompile("[a-zA-Z0-9+/=]+")
	match := r.FindAllString(urlSpilt[1], -1)[0]
	urlSpilt[1] = strings.Replace(urlSpilt[1], match, base64decode.Base64Decode(match), -1)
	builder.WriteString(base64decode.Base64Decode(urlSpilt[1]))
	urlstring = builder.String()
	u, err := url.Parse(urlstring)
	if err != nil {
		fmt.Printf("err: %v\n", err)
		return
	}
	userString := strings.Split(base64decode.Base64Decode(u.User.String()), ":")

	ss.Name = u.Fragment
	// ss.Name = "unkownName"
	ss.Server = u.Hostname()
	ss.Port = u.Port()
	ss.Type = "ss"
	ss.Link = urlstring

	if len(userString) == 1 {

		b, _ := base64.RawStdEncoding.DecodeString(userString[0])
		tmp := strings.Split(string(b), ":")
		if len(tmp) == 1 {
			ss = &Ss{}
			return
		}
		ss.Password = tmp[1]
		ss.Cipher = tmp[0]
	} else {
		ss.Cipher = userString[0]
		ss.Password = userString[1]
	}
}

func decodeB64SS(link string) (string, error) {
	if strings.Contains(link, "@") {
		return link, nil
	}
	regex := regexp.MustCompile(`^ss://([A-Za-z0-9+-=/_]+)`)
	res := regex.FindAllStringSubmatch(link, 1)
	b64 := ""
	if len(res) > 0 && len(res[0]) > 1 {
		b64 = res[0][1]
	}
	if b64 == "" {
		return link, nil
	}
	uri, err := utils.DecodeB64(b64)
	if err != nil {
		return "", err
	}
	return fmt.Sprintf("ss://%s", uri), nil
}

func (ss *Ss) Parse(link1 string) {
	link, err := decodeB64SS(link1)
	if err != nil {
		return
	}
	u, err := url.Parse(link)
	if err != nil {
		return
	}
	if u.Scheme != "ss" {
		return
	}
	pass := u.User.Username()
	hostport := u.Host
	host, port1, err := net.SplitHostPort(hostport)
	if err != nil {
		return
	}
	userinfo, err := utils.DecodeB64(pass)
	if err != nil || !strings.Contains(userinfo, ":") {
		pw, _ := u.User.Password()
		if pw == "" {
			return
		}
		userinfo = fmt.Sprintf("%s:%s", u.User.Username(), pw)
	}
	splits := strings.SplitN(userinfo, ":", 2)
	method := splits[0]
	pass = splits[1]
	remarks := u.Fragment
	if remarks == "" {
		if splits := strings.Split(link1, "#"); len(splits) > 1 {
			if rmk, err := url.QueryUnescape(splits[1]); err == nil {
				remarks = rmk
			}
		}
	}
	ss.Name = remarks
	ss.Server = host
	ss.Port = port1
	ss.Password = pass
	ss.Cipher = method
	ss.Type = "ss"
	ss.Link = link1
}
