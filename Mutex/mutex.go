package mutex

import (
	"sync"
)

var UserDataMutex sync.Mutex

func GetUserMutexData() *sync.Mutex {
	return &UserDataMutex
}
