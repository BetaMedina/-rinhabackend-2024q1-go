package pkg

import (
	"log"
	"os"
	"regexp"

	"github.com/joho/godotenv"
)

const projectDirName = "rinha"

func GetEnvs(key string) string {
	projectName := regexp.MustCompile(`^(.*` + projectDirName + `)`)
	currentWorkDirectory, _ := os.Getwd()
	rootPath := projectName.Find([]byte(currentWorkDirectory))

	err := godotenv.Load(string(rootPath) + `/.env`)

	if err != nil {
		log.Fatalf("Error loading .env file")
	}

	return os.Getenv(key)
}
