package sysutils

import (
	"log"
	"os/user"
)

func GetUserInformation() *user.User {
	usr, err := user.Current()
	if err != nil {
		log.Fatalf(err.Error())
	}
	return usr
}
