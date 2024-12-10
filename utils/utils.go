package utils

import "sync"

const (
	maxID = 9999
)

var id int = 0
var m sync.Mutex

func GenerateID() int {
	m.Lock()
	defer m.Unlock()
	id += 1
	if id > maxID {
		id = 1
	}
	return id
}

func CurID() int {
	return id
}

func SetID(i int) {
	id = i
}
