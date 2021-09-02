package core

import (
	"io/ioutil"
	"net/http"
	"strings"
)

type IpData struct {
	Ip   string `json:"ip"`
	Date int    `json:date`
}

func GetIP() (ip string) {
	res, err := http.Get("http://members.3322.org/dyndns/getip")
	if err != nil {
		panic(err)
	}
	body, err := ioutil.ReadAll(res.Body)
	if err != nil {
		panic(err)
	}
	return strings.TrimSpace(string(body))
}
