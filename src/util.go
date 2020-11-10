package tome

import "fmt"

func Check(e error) {
	if e != nil {
		panic(fmt.Sprintf("Encountered fatal error: %v\n", e))
	}
}
