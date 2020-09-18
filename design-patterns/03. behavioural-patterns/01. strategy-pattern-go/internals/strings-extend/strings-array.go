package stringsextend

import "strings"


func IndexOf(stringArray []*string, findingString string, stringCase StringCase) int {
	if(stringArray == nil || len(stringArray) == 0){
		return -1
	}

	return
}

func IsExists(s, findingString string, stringCase StringCase) bool {
	return IndexOf(s, findingString, stringCase) > -1
}

func DoesntExist(s, findingString string, stringCase StringCase) bool {
	return !IsExists(s, findingString, stringCase)
}
