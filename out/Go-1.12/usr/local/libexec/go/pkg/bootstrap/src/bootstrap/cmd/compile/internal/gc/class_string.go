// Code generated by go tool dist; DO NOT EDIT.
// This is a bootstrap copy of /tmp/tmp.nzFSOQ2SAm/Go-1.12/go/src/cmd/compile/internal/gc/class_string.go

//line /tmp/tmp.nzFSOQ2SAm/Go-1.12/go/src/cmd/compile/internal/gc/class_string.go:1
// Code generated by "stringer -type=Class"; DO NOT EDIT.

package gc

import "strconv"

const _Class_name = "PxxxPEXTERNPAUTOPAUTOHEAPPPARAMPPARAMOUTPFUNCPDISCARD"

var _Class_index = [...]uint8{0, 4, 11, 16, 25, 31, 40, 45, 53}

func (i Class) String() string {
	if i >= Class(len(_Class_index)-1) {
		return "Class(" + strconv.FormatInt(int64(i), 10) + ")"
	}
	return _Class_name[_Class_index[i]:_Class_index[i+1]]
}
