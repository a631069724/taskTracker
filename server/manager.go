package server

import (
	"os"
	"taskTracker/manager"
)

const File = "tasks.json"

var Manager *manager.Manager

func init() {
	Manager = manager.NewManager(999)
	f, err := os.OpenFile(File, os.O_RDWR|os.O_CREATE|os.O_APPEND, 0666)
	if err != nil {
		panic(err)
	}
	defer f.Close()
	Manager.LoadTasks(f)
}

func SaveTasks() {
	f, err := os.OpenFile(File, os.O_RDWR|os.O_CREATE, 0666)
	if err != nil {
		panic(err)
	}
	defer f.Close()
	Manager.SaveTasks(f)
}
