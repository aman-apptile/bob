/*
Copyright © 2024 Mohammed Aman Khan <mohammed.aman@apptile.io>
*/
package main

import (
	"log"

	"github.com/aman-apptile/bob/cmd"
	"github.com/joho/godotenv"
)

func main() {
	if err := godotenv.Load(".config"); err != nil {
		log.Fatal("Error loading .config file")
	}
	cmd.Execute()
}
