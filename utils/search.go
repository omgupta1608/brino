package utils

import "fmt"

func FindOne(obj interface{}, key string) (interface{}, bool) {

	//if the argument is not a map, ignore it
	mobj, ok := obj.(map[string]interface{})
	if !ok {
		fmt.Println("not map")
		return nil, false
	}

	for k, v := range mobj {
		// key match, return value
		if k == key {
			return v, true
		}

		// if the value is a map, search recursively
		if m, ok := v.(map[string]interface{}); ok {
			if res, ok := FindOne(m, key); ok {
				return res, true
			}
		}
		// if the value is an array, search recursively
		// from each element
		if va, ok := v.([]interface{}); ok {
			for _, a := range va {
				if res, ok := FindOne(a, key); ok {
					return res, true
				}
			}
		}
	}

	// element not found
	fmt.Println("Not found")
	return nil, false
}

// func FindAll(obj interface, key string) (interface, bool) {

// }
