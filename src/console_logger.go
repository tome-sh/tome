package tome

import (
	"log"
	"os"
)

var Logger = log.New(os.Stderr, "", 0)
