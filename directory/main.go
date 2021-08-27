package directory

import (
	"io"
	"io/ioutil"
	"log"
	"os"
)

func RemoveAll(path string) (ok bool, message string) {
	message = path + " successfully removed"
	ok = true

	err := os.RemoveAll(path)
	if err != nil {
		log.Println(err)
		message = err.Error()
		ok = false
	}

	return
}

func Create(path string) (created bool) {
	if _, err := os.Stat(path); os.IsNotExist(err) {
		err := os.MkdirAll(path, 0755)
		if err != nil {
			log.Println(err)
		} else {
			created = true
		}
	}

	return
}

func IsEmpty(path string, createPath bool) (empty bool) {

	if _, err := os.Stat(path); os.IsNotExist(err) && createPath {
		Create(path)
	}

	files, err := ioutil.ReadDir(path)
	if err != nil {
		log.Println(err)
	}

	return len(files) == 0
}

func Copy(src, dst string) (ok bool) {
	sourceFileStat, err := os.Stat(src)
	if err != nil {
		log.Println(err)
		return
	}

	if !sourceFileStat.Mode().IsRegular() {
		log.Printf("%s is not a regular file", src)
		return
	}

	source, err := os.Open(src)
	if err != nil {
		log.Println(err)
		return
	}
	defer source.Close()

	destination, err := os.Create(dst)
	if err != nil {
		log.Println(err)
		return
	}
	defer destination.Close()

	_, err = io.Copy(destination, source)
	if err != nil {
		log.Println(err)
		return
	}

	return true
}
