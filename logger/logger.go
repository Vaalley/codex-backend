package logger

import (
	"log"
	"os"
)

func init() {
	// Set up logging with timestamp and file info
	log.SetFlags(log.Ldate | log.Ltime | log.Lshortfile)
	log.SetOutput(os.Stdout)
}
