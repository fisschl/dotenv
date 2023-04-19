package dotenv

import (
	"github.com/gookit/goutil/cliutil"
	"github.com/gookit/ini/v2/dotenv"
)

func init() {
	err := dotenv.Load("./", ".env")
	if err != nil {
		cliutil.Redln("Error loading .env file,", err)
		return
	}
	cliutil.Greenln("Loaded .env file successfully!")
}

func Get(name string) string {
	return dotenv.Get(name)
}
