package api

import (
	"testing"
)

func Test_GetWhoami(t *testing.T) {

	_user := User{}
	iam, err := _user.GetWhoami()

	if err != nil {
		t.Error(err)
	} else {
		t.Logf("Error Code: %d", iam.ErrorCode)
		t.Logf("Username: %s", iam.Details.Username)
	}
}
