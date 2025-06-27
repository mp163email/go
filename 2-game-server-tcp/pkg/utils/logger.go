package utils

import (
	"log"
	"os"
)

var logger = log.New(os.Stdout, "[GameServer]", log.LstdFlags|log.Lmicroseconds|log.Lshortfile)
