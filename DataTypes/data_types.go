package datatypes

import "os"

type SVCS map[string]string
type User struct {
	Name         string
	TrackedFiles []*os.File
}
