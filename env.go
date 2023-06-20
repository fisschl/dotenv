package dotenv

import (
	"github.com/gookit/goutil/cliutil"
	"github.com/gookit/ini/v2/dotenv"
)

func init() {
	err := dotenv.Load("./", ".env")
	if err != nil {
		cliutil.Redln("Error loading .env file,", err)
		panic(err)
	}
	cliutil.Greenln("Loaded .env file successfully!")
}

var Get = dotenv.Get
