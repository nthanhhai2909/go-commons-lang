package sysutils

import (
	"github.com/nthanhhai2909/go-commons-lang/stringutils"
	"testing"
)

func TestGetUserInformation(t *testing.T) {
	usrInformation := GetUserInformation()
	if stringutils.IsEmpty(usrInformation.Name) ||
		stringutils.IsEmpty(usrInformation.HomeDir) ||
		stringutils.IsEmpty(usrInformation.Username) {
		t.Errorf("GetUserInformation failed")
	}
}
