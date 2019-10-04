package api

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"net/http"
	"net/http/cookiejar"
	"net/url"
)

type LogoutResult struct {
	Error    *ErrorDef `json:"error"` // nil on success
	Username string    `json:"username"`
}

func Logout(qnapUrl string, cookies []*http.Cookie) (*LogoutResult, error) {
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
	logoutUrl := fmt.Sprintf("%s%s%s", qnapUrl, ApiV1, LogoutPath)
	req, err := http.NewRequest("PUT", logoutUrl, nil)
	if err != nil {
		return nil, err
	}
	res, err := client.Do(req)
	if err != nil {
		return nil, err
	}

	defer res.Body.Close()

	body, err := ioutil.ReadAll(res.Body)
	if err != nil {
		return nil, err
	}

	var logoutResult LogoutResult
	if err := json.Unmarshal(body, &logoutResult); err != nil {
		return nil, err
	}

	return &logoutResult, nil
}
