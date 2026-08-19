package log

import "log"

var logger = log.New(
	log.Writer(),
	"chardata: ",
	log.Ldate|log.Ltime|log.Lshortfile,
)
