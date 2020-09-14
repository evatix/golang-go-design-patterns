package stringsextend

import "strings"

func IsNull(s *string) bool {
	return s == nil
}

func IsNullOrEmpty(s *string) bool {
	return IsNull(s) || len(*s) == 0 || len(strings.Trim(*s, " ")) == 0
}

func IsNullOrWhitespace(s *string) bool {
	return IsNullOrEmpty(s) || len(strings.Trim(*s, " ")) == 0
}

func IsEmpty(s *string) bool {
	return len(*s) == 0
}

func IsBlank(s *string) bool {
	return IsNullOrWhitespace(s)
}

func HasCharacter(s *string) bool {
	return !IsNullOrWhitespace(s)
}

func IsDefined(s *string) bool {
	return !IsNullOrWhitespace(s)
}

func IndexOf(s, findingString string, stringCase StringCase) int {
	if stringCase == Sensitive {
		return strings.Index(s, findingString)
	}

	lowerCase := strings.ToLower(s)
	findingLower := strings.ToLower(findingString)

	return strings.Index(lowerCase, findingLower)
}

func IsExists(s, findingString string, stringCase StringCase) bool {
	return IndexOf(s, findingString, stringCase) > -1
}

func DoesntExist(s, findingString string, stringCase StringCase) bool {
	return !IsExists(s, findingString, stringCase)
}
