package api

import (
	"fmt"
	"io"
	"io/ioutil"
	"log"
	"net/http"
	"strings"
)

func webRequest(method string, command string, params map[string]string) (result []byte, err error) {

	var reader io.Reader
	var hc http.Client

	conf, err := getConfig()

	if err != nil {
		return nil, err
	}

	uri := fmt.Sprintf("%s://%s:%d/Api/v1/%s?key=%s&format=JSON",
		getProtocol(conf.Https), conf.Host, conf.Port, command, conf.Key)

	if strings.ToLower(method) == "get" {
		if params != nil {
			uri += joinStringMapForUri(params)
		}
	} else {
		if params != nil {
			readerData := joinStringMapForUri(params)
			reader = strings.NewReader(readerData)
		}
	}

	log.Println(uri)

	req, err := http.NewRequest(method, uri, reader)

	if err != nil {
		return nil, err
	}

	resp, err := hc.Do(req)

	if err != nil {
		return nil, err
	}

	body, err := ioutil.ReadAll(resp.Body)

	if err != nil {
		return nil, err
	}

	defer resp.Body.Close()

	if err != nil {
		return nil, err
	}

	return body, err
}

func getProtocol(isHttps bool) string {
	if isHttps {
		return "https"
	} else {
		return "http"
	}
}

func joinStringMapForUri(params map[string]string) string {
	var concat string

	for key, value := range params {
		concat += fmt.Sprintf("&%s=%s", key, value)
	}

	return concat
}
