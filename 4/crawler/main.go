package main

import (
	"io/ioutil"
	"net/http"
	"regexp"
)

func Crawl(host string) []string {
	var URLs = []string{}
	var depth = map[string]int{}

	loop(URLs, host, depth)

	var count = len(depth)

	var res = make([]string, count)

	for url, i := range depth {
		res[i] = url
	}

	return res
}

func get_links(html string) []string {
	bodyStr := html

	r, _ := regexp.Compile(`(?s)href\s*=\s*[\'\"]([^\"\']+)[\'\"]`)

	sub := r.FindAllStringSubmatch(bodyStr, -1)

	var links = []string{}

	for i := range sub {
		href := sub[i][1]

		m1, _ := regexp.MatchString(`(?s)^\s*http`, href)
		m2, _ := regexp.MatchString(`(?s)^\s*/?index.html?`, href)

		if href == "/" || m1 || m2 {
			continue
		}

		if href[0] != '/' {
			href = "/" + href
		}

		links = append(links, href)
	}

	return links
}

func get_html(url string) (string, bool) {
	resp, err := http.Get(url)

	if err != nil {
		panic(err)
	}

	defer resp.Body.Close()

	html, err := ioutil.ReadAll(resp.Body)

	if err != nil {
		panic(err)
	}

	if resp.StatusCode == http.StatusOK {
		return string(html), true
	}

	return "", false
}

func loop(HREFs []string, host string, depth map[string]int) {
	for i := range HREFs {
		href := HREFs[i]

		html, ok := get_html(host + href)

		if ok {
			depth[href] = len(depth)

			links := get_links(html)

			var next_hrefs = []string{}

			for j := range links {
				href := links[j]

				if _, ok := depth[href]; !ok {
					next_hrefs = append(next_hrefs, href)

					loop(next_hrefs, host, depth)
				}
			}
		}
	}

	if len(HREFs) == 0 {
		resp, _ := http.Get(host)

		url := resp.Request.URL.RequestURI()
		HREFs = append(HREFs, url)

		loop(HREFs, host, depth)
	}

	return
}
