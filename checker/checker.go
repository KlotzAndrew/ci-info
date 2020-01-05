package checker

import (
	"os"
)

func IsCI() bool {
	for _, vendor := range Vendors {
		if checkEnv(vendor) {
			return true
		}
	}
	return false
}

func IsPR() bool {
	for _, vendor := range Vendors {
		pr := vendor["pr"]
		str, isString := pr.(string)
		if isString {
			_, found := os.LookupEnv(str)
			if found {
				return true
			}
		}

		obj, isObj := pr.(map[string]string)
		if isObj {
			if obj["env"] != "" {
				val, found := os.LookupEnv(obj["env"])
				if found && val != obj["ne"] {
					return true
				}
			} else {
				for k := range obj {
					_, found := os.LookupEnv(k)
					if found {
						return true
					}
				}
			}
		}

		any, isAny := pr.(map[string][]string)
		if isAny {
			if len(any["any"]) > 0 {
				for _, val := range any["any"] {
					if _, found := os.LookupEnv(val); found {
						return true
					}
				}
			}
		}
	}
	return false
}

func CIName() string {
	for _, vendor := range Vendors {
		if checkEnv(vendor) {
			name, ok := vendor["name"].(string)
			if ok {
				return name
			}
		}
	}
	return ""
}

func checkEnv(vendor map[string]interface{}) bool {
	for _, val := range vendor {
		v, ok := val.(string)
		if !ok {
			continue
		}
		_, found := os.LookupEnv(v)
		if found {
			return true
		}
	}
	return false
}
