package utils

import (
	"log"
	"os"
)

var Logger = log.New(os.Stdout, "[CogAI] ", log.LstdFlags)
