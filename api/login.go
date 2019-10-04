package api

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"net/http"
	"net/http/cookiejar"
	"net/url"
	"strings"
)

type LoginResult struct {
	Error     *ErrorDef `json:"error"` // nil on success
	Anonymous bool      `json:"anonymouse"`
	IsAdmin   bool      `json:"isAdmin"`
	LoginTime string    `json:"logintime"` // localtime like "2019-10-03 15:43:15"
	Username  string    `json:"username"`
}

func Login(qnapUrl string, username string, password string) (*LoginResult, []*http.Cookie, error) {
	jar, err := cookiejar.New(nil)
	if err != nil {
		return nil, nil, err
	}
	client := &http.Client{Jar: jar}
	values := url.Values{}
	values.Set("username", username)
	values.Set("password", password)
	loginUrl := fmt.Sprintf("%s%s%s", qnapUrl, ApiV1, LoginPath)
	res, err := client.Post(
		loginUrl,
		"application/x-www-form-urlencoded",
		strings.NewReader(values.Encode()))
	if err != nil {
		return nil, nil, err
	}

	defer res.Body.Close()

	body, err := ioutil.ReadAll(res.Body)
	if err != nil {
		return nil, nil, err
	}

	var loginResult LoginResult
	if err := json.Unmarshal(body, &loginResult); err != nil {
		return nil, nil, err
	}
	urlObj, err := url.Parse(qnapUrl)
	if err != nil {
		return nil, nil, err
	}
	cookies := jar.Cookies(urlObj)

	return &loginResult, cookies, nil
}
