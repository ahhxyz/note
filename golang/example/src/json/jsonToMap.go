package main

import (
	"encoding/json"
	"fmt"
)

type Config struct {
	MODULE []string
	DB     DataBase
}
type DataBase struct {
	Host     string
	Driver   string
	User     string
	Password string
	DbName   string
	SSLMode  string
	Prefix   string
}

func main() {
	configs := `{
				"App": {
					"MODULE": ["erp","mes"],
					"DB": {
						"Host":"127.0.0.1",
						"Driver":   "postgres",
						"User":     "postgres",
						"Password": "123456",
						"DbName":   "cims",
						"Prefix":   "cims_",
						"SSLMode":  "disable"
					}
				}
			}`

	configMap := make(map[string]Config)
	json.Unmarshal([]byte(configs), &configMap)
	fmt.Printf("%v\n", configMap)
}
