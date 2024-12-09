package env

import (
	"fmt"
	"os"
	"strings"

	"github.com/jolt9dev/go-platform"
)

const (
	// The current process environment variable.
	X_PROCESS = 0
	// The machine environment variable.
	X_MACHINE = 1
	// The user environment variable.
	X_USER = 2
)

// Gets the value of the environment variable named by the key.
//
// Parameters
//   - key: the name of the environment variable
func Get(key string) string {
	return os.Getenv(key)
}

// Sets the value of the environment variable named by the key.
//
// Parameters
//   - key: the name of the environment variable
//   - value: the value to set the environment variable to
func Set(key, value string) error {
	return os.Setenv(key, value)
}

// Deletes the environment variable named by the key.
//
// Parameters
//   - key: the name of the environment variable
func Delete(key string) error {
	return os.Unsetenv(key)
}

// Has returns true if the environment variable named by the key exists.
//
// Parameters
//   - key: the name of the environment variable
func Has(key string) bool {
	_, ok := os.LookupEnv(key)
	return ok
}

// All returns a map containing all the environment variables
// where the keys are the variable names and the values are the
// corresponding values. Only environment variables with non-empty
// values are included in the map.
func All() map[string]string {
	kv := make(map[string]string)
	for _, e := range os.Environ() {
		pair := strings.Split(e, "=")
		if len(pair) == 2 && len(pair[1]) > 0 {
			kv[pair[0]] = pair[1]
		}
	}

	return kv
}

// Print prints all the environment variables to the standard output.
func Print() {
	for k, v := range All() {
		fmt.Printf("%s=%s%s", k, v, platform.EOL)
	}
}

// GetPath returns the value of the PATH environment variable.
func GetPath() string {
	return os.Getenv(PATH)
}

// SetPath sets the value of the PATH environment variable.
//
// Parameters
//   - path: the value to set the PATH environment variable to
func SetPath(path string) error {
	return os.Setenv(PATH, path)
}

// HasPath returns true if the path exists in the PATH environment variable.
//
// Parameters
//   - path: the path to check for in the PATH environment variable
func HasPath(path string) bool {
	return hasPath(path, SplitPath())
}

// AppendPath appends the path to the PATH environment variable.
//
// Parameters
//   - path: the path to append to the PATH environment variable
func AppendPath(path string) error {
	paths := SplitPath()
	if hasPath(path, paths) {
		return nil
	}
	paths = append(paths, path)
	return SetPath(JoinPath(paths...))
}

// PrependPath prepends the path to the PATH environment variable.
//
// Parameters
//   - path: the path to prepend to the PATH environment variable
func PrependPath(path string) error {
	paths := SplitPath()
	if hasPath(path, paths) {
		return nil
	}
	paths = append([]string{path}, paths...)
	return SetPath(JoinPath(paths...))
}

// SplitPath splits the PATH environment variable into a slice of paths.
func SplitPath() []string {
	return strings.Split(GetPath(), string(os.PathListSeparator))
}

// JoinPath joins the paths into a single string using the os.PathListSeparator.
func JoinPath(paths ...string) string {
	return strings.Join(paths, string(os.PathListSeparator))
}
