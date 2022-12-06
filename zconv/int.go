package zconv

import "strconv"

func Int64ToString(i int64) string {
	return strconv.FormatInt(i, 10)
}

func IntToString(i int) string {
	return strconv.Itoa(i)
}
