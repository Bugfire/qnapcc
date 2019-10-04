package api

import (
	"io/ioutil"
	"net/http"
	"net/http/cookiejar"
	"net/url"
)

func Get(qnapUrl string, cookies []*http.Cookie) ([]byte, error) {
	jar, err := cookiejar.New(nil)
	if err != nil {
		return nil, err
	}
	urlObj, err := url.Parse(qnapUrl)
	if err != nil {
		return nil, err
	}
	jar.SetCookies(urlObj, cookies)

	client := &http.Client{Jar: jar}
	res, err := client.Get(qnapUrl)
	if err != nil {
		return nil, err
	}

	defer res.Body.Close()

	body, err := ioutil.ReadAll(res.Body)
	if err != nil {
		return nil, err
	}

	return body, nil
}
