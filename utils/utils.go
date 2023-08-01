package utils

import (
	"regexp"
)

func CheckLink(link string) ([]string, error) {
	r := regexp.MustCompile("(?i)^(vmess|trojan|vless|ss|ssr|http)://.+")
	matches := r.FindStringSubmatch(link)
	if len(matches) < 2 {
		return nil, NewError("Not Suported Link")
	}
	return matches, nil
}

func FindAllLink(link string) ([][]string, error) {
	r := regexp.MustCompile("(trojan|vmess|vless|ss|ssr)://[a-zA-Z0-9+/ =:#@_.?;%&-|]+")
	// matches := r.FindStringSubmatch(link)
	matches := r.FindAllStringSubmatch(link, -1)
	if len(matches) == 0 {
		return nil, NewError("Not Suported Link")
	}
	return matches, nil
}
