// Package osutil provides platform-independent interfaces
// to operating system functionality.

package osutil

import (
	"io/ioutil"
	"os"
	"time"
)

// EmptyAll will delete all contents of a directory, leaving
// the provided directory. This is different from os.Remove
// which also removes the directory provided.
func EmptyAll(name string) error {
	aEntries, err := ioutil.ReadDir(name)
	if err != nil {
		return err
	}
	for _, f := range aEntries {
		if f.Name() == "." || f.Name() == ".." {
			continue
		}
		err = os.Remove(name + "/" + f.Name())
		if err != nil {
			return err
		}
	}
	return nil
}

// Exists checks whether the named filepath exists or not for
// a file or directory.
func Exists(name string) (bool, error) {
	_, err := os.Stat(name)
	if err == nil {
		return true, nil
	}
	if os.IsNotExist(err) {
		return false, nil
	}
	return true, err
}

// FileModAge returns a time.Duration representing the age
// of the named file from FileInfo.ModTime().
func FileModAge(name string) (time.Duration, error) {
	stat, err := os.Stat(name)
	if err != nil {
		dur0, _ := time.ParseDuration("0s")
		return dur0, err
	}
	return time.Now().Sub(stat.ModTime()), nil
}
