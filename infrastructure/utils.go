package infrastructure

import "os"

// TODO: Refactor this in go-like way idk
func GetCacheDir() string {
	dir, err := os.UserCacheDir()
	if err != nil {
		panic(err)
	}

	path := dir + "/pomodoro"
	os.MkdirAll(path, os.ModePerm)
	return path
}
