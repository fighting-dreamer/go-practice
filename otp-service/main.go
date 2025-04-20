package main

import (
	"fmt"

	"nipun.io/otp-service/app"
)

func main() {
	config := app.LoadConfig()
	fmt.Printf("%+v\n", *config)
	fmt.Println(config.AppName)
	fmt.Println(config.Port)
}
