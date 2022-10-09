package main

import (
	"strconv"
	"strings"
)

func subdomainVisits(cpdomains []string) []string {
	m := make(map[string]int)
	res := []string{}
	for _, cpdomain := range cpdomains {
		repAndDomain := strings.Split(cpdomain, " ")
		times, _ := strconv.Atoi(repAndDomain[0])
		domain := repAndDomain[1]
		d := strings.Split(domain, ".")
		path := ""
		for i := len(d) - 1; i >= 0; i-- {
			path = d[i] + path
			m[path] += times
			path = "." + path
		}
	}
	for key, value := range m {
		res = append(res, strconv.Itoa(value)+" "+key)
	}
	return res
}
