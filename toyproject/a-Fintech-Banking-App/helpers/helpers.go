package helpers

import (
	_ "golang.org/x/crypto"
)

func HandleErr(err error) {
	if err != nil {
		panic(err.Error())
	}
}
