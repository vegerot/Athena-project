// Code generated by go tool dist; DO NOT EDIT.
// This is a bootstrap copy of /tmp/tmp.nzFSOQ2SAm/Go-1.12/go/src/cmd/internal/objabi/symkind_string.go

//line /tmp/tmp.nzFSOQ2SAm/Go-1.12/go/src/cmd/internal/objabi/symkind_string.go:1
// Code generated by "stringer -type=SymKind"; DO NOT EDIT.

package objabi

import "strconv"

const _SymKind_name = "SxxxSTEXTSRODATASNOPTRDATASDATASBSSSNOPTRBSSSTLSBSSSDWARFINFOSDWARFRANGESDWARFLOCSDWARFMISCSABIALIAS"

var _SymKind_index = [...]uint8{0, 4, 9, 16, 26, 31, 35, 44, 51, 61, 72, 81, 91, 100}

func (i SymKind) String() string {
	if i >= SymKind(len(_SymKind_index)-1) {
		return "SymKind(" + strconv.FormatInt(int64(i), 10) + ")"
	}
	return _SymKind_name[_SymKind_index[i]:_SymKind_index[i+1]]
}
