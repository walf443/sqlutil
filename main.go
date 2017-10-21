package sqlutil

import (
	"strings"
	"strconv"
)

func ReplaceIn(sql string, args []string) string {
	placeholder := strings.Join(args, ",")
	return strings.Replace(sql, "%s", "(" + placeholder + ")", 1)
}

func ReplaceInWithInt(sql string, args []int) string {
	sargs := make([]string, 0, len(args))
	for _, v := range(args) {
		sargs = append(sargs, strconv.Itoa(v))
	}
	return ReplaceIn(sql, sargs)
}
