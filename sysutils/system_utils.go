package sysutils

import (
	"log"
	"os/user"
)

const (
//PathSeparator = os.PathSeparator
)

func GetUserInformation() *user.User {
	usr, err := user.Current()
	if err != nil {
		log.Fatalf(err.Error())
	}
	return usr
}
