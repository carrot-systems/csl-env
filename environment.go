package csl_env

import (
	"errors"
	"github.com/joho/godotenv"
	"log"
	"os"
	"strconv"
	"strings"
)

func LoadEnv() {
	err := godotenv.Load()

	if err != nil {
		log.Fatalln("Unable to load env file.")
	}
}

func GetEnvString(name string) (string, error) {
	envVariable, exists := os.LookupEnv(name)

	if exists == false {
		log.Println("Env variable: ", name, " is not found")
		return "", errors.New("env variable not found")
	}

	if strings.Contains(name, "SECRET") || strings.Contains(name, "TOKEN") || strings.Contains(name, "PASSWORD") {
		log.Println("[", name, "] = ", "****")
	} else {
		log.Println("[", name, "] = ", envVariable)
	}

	return envVariable, nil
}

func GetEnvInt(name string) (int, error) {
	value, err := GetEnvString(name)

	if err != nil {
		return 0, err
	}

	envVariable, err := strconv.Atoi(value)

	if err != nil {
		log.Println(name, " value is invalid, should be integer")
		return 0, errors.New("value is not int")
	}

	return envVariable, nil
}

func GetEnvBool(name string) (bool, error) {
	value, err := GetEnvString(name)

	if err != nil {
		return false, err
	}

	return value == "true" || value == "TRUE" || value == "1" || value == "yes" || value == "YES", nil
}

func RequireEnvString(name string) string {
	value, err := GetEnvString(name)

	if err != nil {
		log.Fatalln(err.Error())
	}

	return value
}

func RequireEnvInt(name string) int {
	envVariable, err := GetEnvInt(name)

	if err != nil {
		log.Fatalln(name, " value is invalid")
	}

	return envVariable
}

func RequireEnvBool(name string) bool {
	value, err := GetEnvBool(name)

	if err != nil {
		log.Fatalln(err.Error())
	}

	return value
}
