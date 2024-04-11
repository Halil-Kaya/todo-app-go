package main

import "todo/config"

func main() {
	appConfig := config.New()
	appConfig.Print()
}
