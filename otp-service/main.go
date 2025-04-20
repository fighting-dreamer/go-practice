package main

import (
	"fmt"

	"nipun.io/otp-service/config"
)

func main() {
	config := config.LoadConfig();
	fmt.Printf("%+v\n", *config);
	fmt.Println(config.AppName);
	fmt.Println(config.Port);
}