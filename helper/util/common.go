package util

import (
	"strings"

	"github.com/online_marketplace/helper/define"
)

func SortOrder(s string) string {
	s = strings.ToLower(s)
	if s == "" || (s != string(define.OrderAsc) && s != string(define.OrderDesc)) {
		return string(define.OrderAsc)
	}
	return s
}
