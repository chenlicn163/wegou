package utils

import (
	"os"
	"reflect"
)

func InArray(need interface{}, needArr interface{}) (exists bool, index int) {
	exists = false
	index = -1
	switch reflect.TypeOf(needArr).Kind() {
	case reflect.Slice:
		s := reflect.ValueOf(needArr)

		for i := 0; i < s.Len(); i++ {
			if reflect.DeepEqual(need, s.Index(i).Interface()) == true {
				index = i
				exists = true
				return
			}
		}
	}
	return
}

func PathExists(path string) (bool, error) {
	_, err := os.Stat(path)
	if err == nil {
		return true, nil
	}
	if os.IsNotExist(err) {
		return false, nil
	}
	return false, err
}
