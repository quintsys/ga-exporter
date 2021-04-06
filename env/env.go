package env

import "os"

// Lookup retrieves the value of the environment variable named
// by the key. If the variable is present in the environment the
// value (which may be empty) is returned.
// Otherwise the returned value will be the default value.
func Lookup(key, defaultValue string) string {
	val, ok := os.LookupEnv(key)
	if !ok {
		return defaultValue
	}
	return val
}
