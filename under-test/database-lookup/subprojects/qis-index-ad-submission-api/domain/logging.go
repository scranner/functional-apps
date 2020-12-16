package domain

import (
	"github.com/comail/colog"
	"log"
	"strings"
)

func ConfigureLogging() {
	colog.SetDefaultLevel(colog.LInfo)
	colog.SetMinLevel(colog.LInfo)
	colog.SetFormatter(&colog.StdFormatter{
		Colors: true,
		Flag:   log.Ldate | log.Ltime | log.Lshortfile,
	})
	colog.Register()
	log.Print("Log level was set to ", strings.ToUpper(colog.LInfo.String()))
	colog.SetMinLevel(colog.LInfo)
}
