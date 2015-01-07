package api

import (
	"testing"
)

func Test_GetConfig(t *testing.T) {
	conf, err := getConfig()

	if err != nil {
		t.Error(err)
	} else {
		t.Log("Host: ", conf.Host)
		t.Log("Host: ", conf.Key)

		if conf.Key == "" {
			t.Error("Api Key is null")
		}
	}
}
