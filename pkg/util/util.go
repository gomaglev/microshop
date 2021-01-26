package util

import (
	"github.com/jinzhu/copier"
)

// StructMapToStruct
func StructMapToStruct(s, ts interface{}) error {
	return copier.Copy(ts, s)
}
