package tplauncher

import "os"

// LaunchMethod is an interface defining all types of launch methods.
type LaunchMethod interface {
	GetCommandPartial() []string
	DoPreLaunch() error
}

//replace env vars in a string slice.
func escapableEnvVarReplaceSlice(s []string) []string {
	for key, val := range s {
		s[key] = escapableEnvVarReplace(val)
	}
	return s
}

//escapableEnvVarReplace wraps os.Getenv to allow for escaping with $$.
func escapableEnvVarReplace(s string) string {
	return os.Expand(s, func(s string) string {
		if s == "$" {
			return "$"
		}
		realEnvVal := os.Getenv(s)

		return realEnvVal
	})
}
