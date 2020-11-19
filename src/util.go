package tome

import (
	"os"

	"github.com/spf13/viper"
)

// Min calculates the minimum of two ints.
func Min(left int64, right int64) int64 {
	if left < right {
		return left
	} else {
		return right
	}
}

// Max calculates the maximum of two ints.
func Max(left int64, right int64) int64 {
	if left > right {
		return left
	} else {
		return right
	}
}

// Check if error is not nil and panic in case it is.
func Check(e error) {
	if e != nil {
		if viper.GetBool("debug") {
			panic(e)
		} else {
			Logger.Printf("Encountered error: %v\n", e)
		}
		os.Exit(1)
	}
}

func GetUserName() (string, error) {
	userName, err := GetGitConfigSetting("user.name")
	if err != nil {
		return "", err
	}
	return userName, nil
}
