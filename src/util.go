package tome

import (
	"fmt"
	"os"

	"github.com/spf13/viper"
)

// Check if error is not nil and panic in case it is.
func Check(e error) {
	if e != nil {
		if (viper.GetBool("debug")) {
			panic(e)
		} else {
			fmt.Printf("Encountered error: %v\n", e)
		}
		os.Exit(1)
	}
}
