// Code generated by "stringer -type=Region"; DO NOT EDIT

package models

import "fmt"

const _Region_name = "USACanadaOther"

var _Region_index = [...]uint8{0, 3, 9, 14}

func (i Region) String() string {
	i -= 1
	if i < 0 || i >= Region(len(_Region_index)-1) {
		return fmt.Sprintf("Region(%d)", i+1)
	}
	return _Region_name[_Region_index[i]:_Region_index[i+1]]
}
