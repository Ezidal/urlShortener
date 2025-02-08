package main

import (
	"fmt"
	parseconfig "urlShortener/internal/parseConfig"
)

func main() {
	config, err := parseconfig.LoadConfig()
	if err != nil {
		panic(err)
	}
	fmt.Println(config)
}
