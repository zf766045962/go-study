package main

import (
	"fmt"
)

type site interface {
	fetch()
}

type siteModel struct {
	URL string
}
type site1 struct {
	siteModel
}

func (s site1) fetch() {
	fmt.Println("site1 fetch data", s.URL)
}

func factory(s string) site {
	if s == "site" {
		var si site

		sitevalue := site1{
			siteModel{URL: "http://www.xxxx.com"},
		}

		si = sitevalue
		return si
	}
	return nil
}

func main() {
	s := factory("site")
	fmt.Println(s)
	s.fetch()
}
