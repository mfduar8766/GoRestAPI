package config

import (
	"encoding/json"
	"fmt"
	"github.com/joho/godotenv"
	"github.com/mfduar8766/GoRestAPI/utils"
	"os"
)

// DbConfig - DB config details
type DbConfig struct {
	User     string `json:"DB_USERNAME"`
	Password string `json:"DB_PASSWORD"`
	Host     string `json:"DB_HOST"`
	Port     string `json:"DB_PORT"`
	DBName   string `json:"DB_NAME"`
}

func getEnvData(key string) string {
	fmt.Println("Config getEnvData()")
	err := godotenv.Load(".env")
	utils.MustNotError(err)
	env, exists := os.LookupEnv(key)
	if !exists {
		fmt.Println("ENV var does not exist")
		return ""
	}
	return env
}

// InitDbConfig - Used to return an instance of the DB config
func InitDbConfig() *DbConfig {
	fmt.Println("Config InitDbConfig()")
	dbConfig := new(DbConfig)
	envVars := map[string]interface{}{
		"DB_USERNAME": "",
		"DB_PASSWORD": "",
		"DB_HOST":     "",
		"DB_PORT":     "",
		"DB_NAME":     "",
	}
	for key := range envVars {
		value := getEnvData(key)
		if value != "" {
			envVars[key] = value
		}
	}
	bytesData, err := json.Marshal(envVars)
	utils.MustNotError(err)
	err = json.Unmarshal(bytesData, &dbConfig)
	utils.MustNotError(err)
	return dbConfig
}
