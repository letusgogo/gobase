package util

import (
	"strconv"
	"strings"
)

// CompareVersions compares two version strings and returns the difference
// at the first different position. If the versions are the same, returns 0.
func CompareVersions(v1, v2 string) int {
	v1Parts := strings.Split(v1, ".")
	v2Parts := strings.Split(v2, ".")

	maxLength := len(v1Parts)
	if len(v2Parts) > maxLength {
		maxLength = len(v2Parts)
	}

	for i := 0; i < maxLength; i++ {
		var num1, num2 int

		if i < len(v1Parts) {
			num1, _ = strconv.Atoi(v1Parts[i])
		}

		if i < len(v2Parts) {
			num2, _ = strconv.Atoi(v2Parts[i])
		}

		if num1 != num2 {
			return num1 - num2
		}
	}

	return 0
}
