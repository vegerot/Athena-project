// Code generated by go tool dist; DO NOT EDIT.
// This is a bootstrap copy of /tmp/tmp.nzFSOQ2SAm/Go-1.12/go/src/debug/dwarf/class_string.go

//line /tmp/tmp.nzFSOQ2SAm/Go-1.12/go/src/debug/dwarf/class_string.go:1
// Code generated by "stringer -type=Class"; DO NOT EDIT.

package dwarf

import "strconv"

const _Class_name = "ClassUnknownClassAddressClassBlockClassConstantClassExprLocClassFlagClassLinePtrClassLocListPtrClassMacPtrClassRangeListPtrClassReferenceClassReferenceSigClassStringClassReferenceAltClassStringAlt"

var _Class_index = [...]uint8{0, 12, 24, 34, 47, 59, 68, 80, 95, 106, 123, 137, 154, 165, 182, 196}

func (i Class) String() string {
	if i < 0 || i >= Class(len(_Class_index)-1) {
		return "Class(" + strconv.FormatInt(int64(i), 10) + ")"
	}
	return _Class_name[_Class_index[i]:_Class_index[i+1]]
}
