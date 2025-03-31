package main

import "github.com/paulnaber/rssgator-go/internal/config"
import "fmt"

func main() {

	config, err := config.ReadConfig()
	if err != nil {
		fmt.Println("Error reading config:", err)
	}
	fmt.Printf("Config: %v", config)

}
