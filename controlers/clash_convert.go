package controlers

import (
	"fmt"
	"io"
	"log"
	"net/http"
	"strings"
	"text/template"
	"time"

	clash "github.com/sanzhang007/webgin/protocol"

	"github.com/sanzhang007/webgin/utils"

	"github.com/gin-gonic/gin"
	"github.com/go-yaml/yaml"
)

type ClashT struct {
	Proxies []Proxy `yaml:"proxies"`
	Time    string  `yaml:"time"`
	Nums    Num     `yaml:"nums"`
}

type Num struct {
	All    int
	Trojan int
	Ssr    int
	Ss     int
	Vmess  int
}

type Proxy struct {
	Name           string `yaml:"name"`
	Server         string `yaml:"server"`
	Port           string `yaml:"port"`
	Type           string `yaml:"type"`
	Password       string `yaml:"password"`
	Cipher         string `yaml:"cipher"`
	Uuid           string `yaml:"uuid"`
	AlterId        string `yaml:"alterId"`
	Tls            string `yaml:"tls"`
	Tfo            bool   `yaml:"tfo"`
	SkipCertVerify string `yaml:"skip-cert-verify"`
	Sni            string `yaml:"sni"`
	Udp            bool   `yaml:"udp"`
	Network        string `yaml:"network"`
	WsOpts         WsOpt  `yaml:"ws-opts"`

	//ss
	Obfs          string `yaml:"obfs"`
	ObfsParam     string `yaml:"obfs-param"`
	Protocol      string `yaml:"protocol"`
	ProtocolParam string `yaml:"protocol-param"`

	Id int
}

type WsOpt struct {
	Path    string `yaml:"path"`
	Headers Header `yaml:"headers"`
}
type Header struct {
	Host string `yaml:"host"`
}

type Ss struct {
	Name     string
	Password string
	Port     string
	Server   string
	Cipher   string
	Type     string
}

func ClashConvert(ctx *gin.Context) {
	//修改为text/plain
	ctx.Header("Content-Type", "text/plain; charset=utf-8")
	url := ctx.Query("link")
	b, err := curl(url)
	if err != nil {
		ctx.String(200, "%s\n", b)
		ctx.String(200, "[error]: %s\n", err.Error())
	}
	clash := templateClash(ClashByte(string(b)))
	// fmt.Printf("clash: %v\n", clash)
	// ctx.HTML(http.StatusOK, "config.tmpl", clash)
	t, err2 := template.ParseFiles("./config.tmpl")
	if err2 != nil {
		fmt.Printf("err2: %v\n", err2)
		return
	}
	t.Execute(ctx.Writer, clash)
	// fmt.Fprintln(ctx.Writer, clash)

}

func curl(link string) ([]byte, error) {
	var client http.Client
	r, err := client.Get(link)
	if err != nil {
		return nil, nil
	}
	body := r.Body
	defer body.Close()
	return io.ReadAll(body)
}

func templateClash(ctxByte []byte) ClashT {
	var c ClashT
	c.Time = time.Now().Format("2006-01-02 15:04:05")
	// var c map[string]any
	s := strings.Replace(string((ctxByte)), "Host", "host", -1)
	s = strings.Replace(s, "skipcertverify", "skip-cert-verify", -1)
	err := yaml.Unmarshal([]byte(s), &c)
	if err != nil {
		fmt.Printf("err: %v\n", err)
		return c
	}
	c.Nums.All = len(c.Proxies)
	var tmp []Proxy
	id := 0
	for i := 0; i < c.Nums.All; i++ {
		c.Proxies[i].Name = strings.Replace(c.Proxies[i].Name, `\`, ``, -1)
		if strings.ToLower(c.Proxies[i].Type) == "vmess" {
			tmp = append(tmp, c.Proxies[i])
			tmp[id].Id = id
			id++
			c.Nums.Vmess++
		}
		if strings.ToLower(c.Proxies[i].Type) == "ssr" {
			tmp = append(tmp, c.Proxies[i])
			tmp[id].Id = id
			id++
			c.Nums.Ssr++
		}
		if strings.ToLower(c.Proxies[i].Type) == "ss" {
			tmp = append(tmp, c.Proxies[i])
			tmp[id].Id = id
			id++
			c.Nums.Ss++
		}
		if strings.ToLower(c.Proxies[i].Type) == "trojan" {
			tmp = append(tmp, c.Proxies[i])
			tmp[id].Id = id
			id++
			c.Nums.Trojan++
		}
	}
	c.Proxies = tmp
	c.Nums.All = len(c.Proxies)
	return c
}

func ClashByte(ctx string) []byte {
	//为clash自己保存
	ClashAll := new(clash.Clash)
	list, err := utils.FindAllLink(ctx)
	if err != nil {
		log.Println(err)
		return nil
	}

	for _, v := range list {
		cl, err := clash.ClashParse(&v[0])
		if err != nil {
			log.Println(err)
			continue
		}
		y, _ := yaml.Marshal(cl)
		var m map[string]any
		yaml.Unmarshal(y, &m)
		if cl != nil {
			ClashAll.Proxies = append(ClashAll.Proxies, m)
		}
	}

	clashctx, err := yaml.Marshal(ClashAll)
	if err != nil {
		log.Println(err)
	}
	return clashctx

}
