package clash

// package main

import (
	"github.com/go-yaml/yaml"
)

type Clash struct {
	MiedPort           int              `yaml:"mixed-port"`
	AllowLan           bool             `yaml:"allow-lan"`
	ExternalController string           `yaml:"external-controller"`
	Mode               string           `yaml:"mode"`
	Proxies            []map[string]any `yaml:"proxies"`
	// ProxyGroups        []ProxyGroups    `yaml:"proxy-groups"`
	RuleProviders any `yaml:"rule-providers"`
	Rules         any `yaml:"rules"`
}

// type ProxyGroups struct {
// 	Name    string   `yaml:"name"`
// 	Type    string   `yaml:"type"`
// 	Proxies []string `yaml:"proxies"`
// }

// 判断是否为clash
func IsClash(str string) bool {
	var c Clash

	yaml.Unmarshal([]byte(str), &c)
	return len(c.Proxies) > 0

}
