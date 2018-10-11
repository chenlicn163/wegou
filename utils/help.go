package utils

import "reflect"

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
