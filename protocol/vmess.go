package clash

import (
	"strconv"
	"strings"
	"webgin/base64decode"

	"github.com/go-yaml/yaml"
)

type Vmess struct {
	Name    string `yaml:"name"`
	Port    string `yaml:"port"`
	Server  string `yaml:"server"`
	AlterId string `yaml:"alterId"`
	Cipher  string `yaml:"cipher"`
	Network string `yaml:"network"`
	Tls     bool   `yaml:"tls"`
	Type    string `yaml:"type"`
	Uuid    string `yaml:"uuid"`
	WsOpts  WsOpts `yaml:"ws-opts"`
}

type WsOpts struct {
	Headers Headers `yaml:"headers"`
	Path    string  `yaml:"path"`
}
type Headers struct {
	Host string `yaml:"host"`
}

var m map[string]string

func (vmess *Vmess) Parse(url string) {
	url = base64decode.Base64Decode(strings.Split(url, "://")[1])
	yaml.Unmarshal([]byte(url), &m)
	vmess.Type = "vmess"
	vmess.Name = m["ps"]
	vmess.Server = m["add"]
	vmess.Port = m["port"]
	vmess.Uuid = m["id"]
	vmess.AlterId = m["aid"]
	vmess.Cipher = "auto"
	vmess.Network = m["net"]
	// vmess.Tls = m["tls"]
	boolValue, _ := strconv.ParseBool(m["tls"])
	vmess.Tls = boolValue
	vmess.WsOpts.Path = m["path"]
	vmess.WsOpts.Headers.Host = m["host"]
}
