package dotenv

import (
	"github.com/gookit/ini/v2/dotenv"
	"log"
)

func init() {
	err := dotenv.Load("./", ".env")
	if err != nil {
		log.Println("Error loading .env file", err)
		return
	}
	log.Println("Loaded .env file successfully!")
}

func Get(name string) string {
	return dotenv.Get(name)
}
