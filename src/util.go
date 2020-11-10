package tome

import "fmt"

// Check if error is not nil and panic in case it is.
func Check(e error) {
	if e != nil {
		panic(fmt.Sprintf("Encountered fatal error: %v\n", e))
	}
}
