package config

import (
	"encoding/json"
	"fmt"
	"os"
)

type Config struct {
	DbUrl           string `json:"db_url"`
	CurrentUserName string `json:"current_user_name"`
}

// Read from ~/.gatorconfig.json
// return Config struct from json
func ReadConfig() (Config, error) {
	path, err := getUserHomePath()
	if err != nil {
		fmt.Println("Error getting user home path:", err)
		return Config{}, err
	}
	path = path + "/.gatorconfig.json"

	jsonData, err := os.ReadFile(path)
	if err != nil {
		fmt.Println("Error reading config file:", err)
		return Config{}, err
	}
	parsedData := Config{}
	err = json.Unmarshal(jsonData, &parsedData)
	if err != nil {
		fmt.Println("Error parsing config file", err)
		return Config{}, err
	}
	return parsedData, nil
}

func getUserHomePath() (string, error) {
	usr, err := os.UserHomeDir()
	if err != nil {
		return "", err
	}
	return usr, nil
}
