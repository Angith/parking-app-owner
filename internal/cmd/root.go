package cmd

import (
	"fmt"
)

func execute() {
	db := prepareDB()
	if db != nil {
		fmt.Println("success!")
	}
}
