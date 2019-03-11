// Code generated by go tool dist; DO NOT EDIT.
// This is a bootstrap copy of /tmp/tmp.nzFSOQ2SAm/Go-1.12/go/src/cmd/compile/internal/gc/builtin.go

//line /tmp/tmp.nzFSOQ2SAm/Go-1.12/go/src/cmd/compile/internal/gc/builtin.go:1
// Code generated by mkbuiltin.go. DO NOT EDIT.

package gc

import "bootstrap/cmd/compile/internal/types"

var runtimeDecls = [...]struct {
	name string
	tag  int
	typ  int
}{
	{"newobject", funcTag, 4},
	{"panicindex", funcTag, 5},
	{"panicslice", funcTag, 5},
	{"panicdivide", funcTag, 5},
	{"panicmakeslicelen", funcTag, 5},
	{"throwinit", funcTag, 5},
	{"panicwrap", funcTag, 5},
	{"gopanic", funcTag, 7},
	{"gorecover", funcTag, 10},
	{"goschedguarded", funcTag, 5},
	{"printbool", funcTag, 12},
	{"printfloat", funcTag, 14},
	{"printint", funcTag, 16},
	{"printhex", funcTag, 18},
	{"printuint", funcTag, 18},
	{"printcomplex", funcTag, 20},
	{"printstring", funcTag, 22},
	{"printpointer", funcTag, 23},
	{"printiface", funcTag, 23},
	{"printeface", funcTag, 23},
	{"printslice", funcTag, 23},
	{"printnl", funcTag, 5},
	{"printsp", funcTag, 5},
	{"printlock", funcTag, 5},
	{"printunlock", funcTag, 5},
	{"concatstring2", funcTag, 26},
	{"concatstring3", funcTag, 27},
	{"concatstring4", funcTag, 28},
	{"concatstring5", funcTag, 29},
	{"concatstrings", funcTag, 31},
	{"cmpstring", funcTag, 33},
	{"intstring", funcTag, 36},
	{"slicebytetostring", funcTag, 38},
	{"slicebytetostringtmp", funcTag, 39},
	{"slicerunetostring", funcTag, 42},
	{"stringtoslicebyte", funcTag, 43},
	{"stringtoslicerune", funcTag, 46},
	{"slicecopy", funcTag, 48},
	{"slicestringcopy", funcTag, 49},
	{"decoderune", funcTag, 50},
	{"countrunes", funcTag, 51},
	{"convI2I", funcTag, 52},
	{"convT16", funcTag, 54},
	{"convT32", funcTag, 54},
	{"convT64", funcTag, 54},
	{"convTstring", funcTag, 54},
	{"convTslice", funcTag, 54},
	{"convT2E", funcTag, 55},
	{"convT2Enoptr", funcTag, 55},
	{"convT2I", funcTag, 55},
	{"convT2Inoptr", funcTag, 55},
	{"assertE2I", funcTag, 52},
	{"assertE2I2", funcTag, 56},
	{"assertI2I", funcTag, 52},
	{"assertI2I2", funcTag, 56},
	{"panicdottypeE", funcTag, 57},
	{"panicdottypeI", funcTag, 57},
	{"panicnildottype", funcTag, 58},
	{"ifaceeq", funcTag, 60},
	{"efaceeq", funcTag, 60},
	{"fastrand", funcTag, 62},
	{"makemap64", funcTag, 64},
	{"makemap", funcTag, 65},
	{"makemap_small", funcTag, 66},
	{"mapaccess1", funcTag, 67},
	{"mapaccess1_fast32", funcTag, 68},
	{"mapaccess1_fast64", funcTag, 68},
	{"mapaccess1_faststr", funcTag, 68},
	{"mapaccess1_fat", funcTag, 69},
	{"mapaccess2", funcTag, 70},
	{"mapaccess2_fast32", funcTag, 71},
	{"mapaccess2_fast64", funcTag, 71},
	{"mapaccess2_faststr", funcTag, 71},
	{"mapaccess2_fat", funcTag, 72},
	{"mapassign", funcTag, 67},
	{"mapassign_fast32", funcTag, 68},
	{"mapassign_fast32ptr", funcTag, 68},
	{"mapassign_fast64", funcTag, 68},
	{"mapassign_fast64ptr", funcTag, 68},
	{"mapassign_faststr", funcTag, 68},
	{"mapiterinit", funcTag, 73},
	{"mapdelete", funcTag, 73},
	{"mapdelete_fast32", funcTag, 74},
	{"mapdelete_fast64", funcTag, 74},
	{"mapdelete_faststr", funcTag, 74},
	{"mapiternext", funcTag, 75},
	{"mapclear", funcTag, 76},
	{"makechan64", funcTag, 78},
	{"makechan", funcTag, 79},
	{"chanrecv1", funcTag, 81},
	{"chanrecv2", funcTag, 82},
	{"chansend1", funcTag, 84},
	{"closechan", funcTag, 23},
	{"writeBarrier", varTag, 86},
	{"typedmemmove", funcTag, 87},
	{"typedmemclr", funcTag, 88},
	{"typedslicecopy", funcTag, 89},
	{"selectnbsend", funcTag, 90},
	{"selectnbrecv", funcTag, 91},
	{"selectnbrecv2", funcTag, 93},
	{"selectsetpc", funcTag, 58},
	{"selectgo", funcTag, 94},
	{"block", funcTag, 5},
	{"makeslice", funcTag, 95},
	{"makeslice64", funcTag, 96},
	{"growslice", funcTag, 98},
	{"memmove", funcTag, 99},
	{"memclrNoHeapPointers", funcTag, 100},
	{"memclrHasPointers", funcTag, 100},
	{"memequal", funcTag, 101},
	{"memequal8", funcTag, 102},
	{"memequal16", funcTag, 102},
	{"memequal32", funcTag, 102},
	{"memequal64", funcTag, 102},
	{"memequal128", funcTag, 102},
	{"int64div", funcTag, 103},
	{"uint64div", funcTag, 104},
	{"int64mod", funcTag, 103},
	{"uint64mod", funcTag, 104},
	{"float64toint64", funcTag, 105},
	{"float64touint64", funcTag, 106},
	{"float64touint32", funcTag, 107},
	{"int64tofloat64", funcTag, 108},
	{"uint64tofloat64", funcTag, 109},
	{"uint32tofloat64", funcTag, 110},
	{"complex128div", funcTag, 111},
	{"racefuncenter", funcTag, 112},
	{"racefuncenterfp", funcTag, 5},
	{"racefuncexit", funcTag, 5},
	{"raceread", funcTag, 112},
	{"racewrite", funcTag, 112},
	{"racereadrange", funcTag, 113},
	{"racewriterange", funcTag, 113},
	{"msanread", funcTag, 113},
	{"msanwrite", funcTag, 113},
	{"x86HasPOPCNT", varTag, 11},
	{"x86HasSSE41", varTag, 11},
	{"arm64HasATOMICS", varTag, 11},
}

func runtimeTypes() []*types.Type {
	var typs [114]*types.Type
	typs[0] = types.Bytetype
	typs[1] = types.NewPtr(typs[0])
	typs[2] = types.Types[TANY]
	typs[3] = types.NewPtr(typs[2])
	typs[4] = functype(nil, []*Node{anonfield(typs[1])}, []*Node{anonfield(typs[3])})
	typs[5] = functype(nil, nil, nil)
	typs[6] = types.Types[TINTER]
	typs[7] = functype(nil, []*Node{anonfield(typs[6])}, nil)
	typs[8] = types.Types[TINT32]
	typs[9] = types.NewPtr(typs[8])
	typs[10] = functype(nil, []*Node{anonfield(typs[9])}, []*Node{anonfield(typs[6])})
	typs[11] = types.Types[TBOOL]
	typs[12] = functype(nil, []*Node{anonfield(typs[11])}, nil)
	typs[13] = types.Types[TFLOAT64]
	typs[14] = functype(nil, []*Node{anonfield(typs[13])}, nil)
	typs[15] = types.Types[TINT64]
	typs[16] = functype(nil, []*Node{anonfield(typs[15])}, nil)
	typs[17] = types.Types[TUINT64]
	typs[18] = functype(nil, []*Node{anonfield(typs[17])}, nil)
	typs[19] = types.Types[TCOMPLEX128]
	typs[20] = functype(nil, []*Node{anonfield(typs[19])}, nil)
	typs[21] = types.Types[TSTRING]
	typs[22] = functype(nil, []*Node{anonfield(typs[21])}, nil)
	typs[23] = functype(nil, []*Node{anonfield(typs[2])}, nil)
	typs[24] = types.NewArray(typs[0], 32)
	typs[25] = types.NewPtr(typs[24])
	typs[26] = functype(nil, []*Node{anonfield(typs[25]), anonfield(typs[21]), anonfield(typs[21])}, []*Node{anonfield(typs[21])})
	typs[27] = functype(nil, []*Node{anonfield(typs[25]), anonfield(typs[21]), anonfield(typs[21]), anonfield(typs[21])}, []*Node{anonfield(typs[21])})
	typs[28] = functype(nil, []*Node{anonfield(typs[25]), anonfield(typs[21]), anonfield(typs[21]), anonfield(typs[21]), anonfield(typs[21])}, []*Node{anonfield(typs[21])})
	typs[29] = functype(nil, []*Node{anonfield(typs[25]), anonfield(typs[21]), anonfield(typs[21]), anonfield(typs[21]), anonfield(typs[21]), anonfield(typs[21])}, []*Node{anonfield(typs[21])})
	typs[30] = types.NewSlice(typs[21])
	typs[31] = functype(nil, []*Node{anonfield(typs[25]), anonfield(typs[30])}, []*Node{anonfield(typs[21])})
	typs[32] = types.Types[TINT]
	typs[33] = functype(nil, []*Node{anonfield(typs[21]), anonfield(typs[21])}, []*Node{anonfield(typs[32])})
	typs[34] = types.NewArray(typs[0], 4)
	typs[35] = types.NewPtr(typs[34])
	typs[36] = functype(nil, []*Node{anonfield(typs[35]), anonfield(typs[15])}, []*Node{anonfield(typs[21])})
	typs[37] = types.NewSlice(typs[0])
	typs[38] = functype(nil, []*Node{anonfield(typs[25]), anonfield(typs[37])}, []*Node{anonfield(typs[21])})
	typs[39] = functype(nil, []*Node{anonfield(typs[37])}, []*Node{anonfield(typs[21])})
	typs[40] = types.Runetype
	typs[41] = types.NewSlice(typs[40])
	typs[42] = functype(nil, []*Node{anonfield(typs[25]), anonfield(typs[41])}, []*Node{anonfield(typs[21])})
	typs[43] = functype(nil, []*Node{anonfield(typs[25]), anonfield(typs[21])}, []*Node{anonfield(typs[37])})
	typs[44] = types.NewArray(typs[40], 32)
	typs[45] = types.NewPtr(typs[44])
	typs[46] = functype(nil, []*Node{anonfield(typs[45]), anonfield(typs[21])}, []*Node{anonfield(typs[41])})
	typs[47] = types.Types[TUINTPTR]
	typs[48] = functype(nil, []*Node{anonfield(typs[2]), anonfield(typs[2]), anonfield(typs[47])}, []*Node{anonfield(typs[32])})
	typs[49] = functype(nil, []*Node{anonfield(typs[2]), anonfield(typs[2])}, []*Node{anonfield(typs[32])})
	typs[50] = functype(nil, []*Node{anonfield(typs[21]), anonfield(typs[32])}, []*Node{anonfield(typs[40]), anonfield(typs[32])})
	typs[51] = functype(nil, []*Node{anonfield(typs[21])}, []*Node{anonfield(typs[32])})
	typs[52] = functype(nil, []*Node{anonfield(typs[1]), anonfield(typs[2])}, []*Node{anonfield(typs[2])})
	typs[53] = types.Types[TUNSAFEPTR]
	typs[54] = functype(nil, []*Node{anonfield(typs[2])}, []*Node{anonfield(typs[53])})
	typs[55] = functype(nil, []*Node{anonfield(typs[1]), anonfield(typs[3])}, []*Node{anonfield(typs[2])})
	typs[56] = functype(nil, []*Node{anonfield(typs[1]), anonfield(typs[2])}, []*Node{anonfield(typs[2]), anonfield(typs[11])})
	typs[57] = functype(nil, []*Node{anonfield(typs[1]), anonfield(typs[1]), anonfield(typs[1])}, nil)
	typs[58] = functype(nil, []*Node{anonfield(typs[1])}, nil)
	typs[59] = types.NewPtr(typs[47])
	typs[60] = functype(nil, []*Node{anonfield(typs[59]), anonfield(typs[53]), anonfield(typs[53])}, []*Node{anonfield(typs[11])})
	typs[61] = types.Types[TUINT32]
	typs[62] = functype(nil, nil, []*Node{anonfield(typs[61])})
	typs[63] = types.NewMap(typs[2], typs[2])
	typs[64] = functype(nil, []*Node{anonfield(typs[1]), anonfield(typs[15]), anonfield(typs[3])}, []*Node{anonfield(typs[63])})
	typs[65] = functype(nil, []*Node{anonfield(typs[1]), anonfield(typs[32]), anonfield(typs[3])}, []*Node{anonfield(typs[63])})
	typs[66] = functype(nil, nil, []*Node{anonfield(typs[63])})
	typs[67] = functype(nil, []*Node{anonfield(typs[1]), anonfield(typs[63]), anonfield(typs[3])}, []*Node{anonfield(typs[3])})
	typs[68] = functype(nil, []*Node{anonfield(typs[1]), anonfield(typs[63]), anonfield(typs[2])}, []*Node{anonfield(typs[3])})
	typs[69] = functype(nil, []*Node{anonfield(typs[1]), anonfield(typs[63]), anonfield(typs[3]), anonfield(typs[1])}, []*Node{anonfield(typs[3])})
	typs[70] = functype(nil, []*Node{anonfield(typs[1]), anonfield(typs[63]), anonfield(typs[3])}, []*Node{anonfield(typs[3]), anonfield(typs[11])})
	typs[71] = functype(nil, []*Node{anonfield(typs[1]), anonfield(typs[63]), anonfield(typs[2])}, []*Node{anonfield(typs[3]), anonfield(typs[11])})
	typs[72] = functype(nil, []*Node{anonfield(typs[1]), anonfield(typs[63]), anonfield(typs[3]), anonfield(typs[1])}, []*Node{anonfield(typs[3]), anonfield(typs[11])})
	typs[73] = functype(nil, []*Node{anonfield(typs[1]), anonfield(typs[63]), anonfield(typs[3])}, nil)
	typs[74] = functype(nil, []*Node{anonfield(typs[1]), anonfield(typs[63]), anonfield(typs[2])}, nil)
	typs[75] = functype(nil, []*Node{anonfield(typs[3])}, nil)
	typs[76] = functype(nil, []*Node{anonfield(typs[1]), anonfield(typs[63])}, nil)
	typs[77] = types.NewChan(typs[2], types.Cboth)
	typs[78] = functype(nil, []*Node{anonfield(typs[1]), anonfield(typs[15])}, []*Node{anonfield(typs[77])})
	typs[79] = functype(nil, []*Node{anonfield(typs[1]), anonfield(typs[32])}, []*Node{anonfield(typs[77])})
	typs[80] = types.NewChan(typs[2], types.Crecv)
	typs[81] = functype(nil, []*Node{anonfield(typs[80]), anonfield(typs[3])}, nil)
	typs[82] = functype(nil, []*Node{anonfield(typs[80]), anonfield(typs[3])}, []*Node{anonfield(typs[11])})
	typs[83] = types.NewChan(typs[2], types.Csend)
	typs[84] = functype(nil, []*Node{anonfield(typs[83]), anonfield(typs[3])}, nil)
	typs[85] = types.NewArray(typs[0], 3)
	typs[86] = tostruct([]*Node{namedfield("enabled", typs[11]), namedfield("pad", typs[85]), namedfield("needed", typs[11]), namedfield("cgo", typs[11]), namedfield("alignme", typs[17])})
	typs[87] = functype(nil, []*Node{anonfield(typs[1]), anonfield(typs[3]), anonfield(typs[3])}, nil)
	typs[88] = functype(nil, []*Node{anonfield(typs[1]), anonfield(typs[3])}, nil)
	typs[89] = functype(nil, []*Node{anonfield(typs[1]), anonfield(typs[2]), anonfield(typs[2])}, []*Node{anonfield(typs[32])})
	typs[90] = functype(nil, []*Node{anonfield(typs[83]), anonfield(typs[3])}, []*Node{anonfield(typs[11])})
	typs[91] = functype(nil, []*Node{anonfield(typs[3]), anonfield(typs[80])}, []*Node{anonfield(typs[11])})
	typs[92] = types.NewPtr(typs[11])
	typs[93] = functype(nil, []*Node{anonfield(typs[3]), anonfield(typs[92]), anonfield(typs[80])}, []*Node{anonfield(typs[11])})
	typs[94] = functype(nil, []*Node{anonfield(typs[1]), anonfield(typs[1]), anonfield(typs[32])}, []*Node{anonfield(typs[32]), anonfield(typs[11])})
	typs[95] = functype(nil, []*Node{anonfield(typs[1]), anonfield(typs[32]), anonfield(typs[32])}, []*Node{anonfield(typs[53])})
	typs[96] = functype(nil, []*Node{anonfield(typs[1]), anonfield(typs[15]), anonfield(typs[15])}, []*Node{anonfield(typs[53])})
	typs[97] = types.NewSlice(typs[2])
	typs[98] = functype(nil, []*Node{anonfield(typs[1]), anonfield(typs[97]), anonfield(typs[32])}, []*Node{anonfield(typs[97])})
	typs[99] = functype(nil, []*Node{anonfield(typs[3]), anonfield(typs[3]), anonfield(typs[47])}, nil)
	typs[100] = functype(nil, []*Node{anonfield(typs[53]), anonfield(typs[47])}, nil)
	typs[101] = functype(nil, []*Node{anonfield(typs[3]), anonfield(typs[3]), anonfield(typs[47])}, []*Node{anonfield(typs[11])})
	typs[102] = functype(nil, []*Node{anonfield(typs[3]), anonfield(typs[3])}, []*Node{anonfield(typs[11])})
	typs[103] = functype(nil, []*Node{anonfield(typs[15]), anonfield(typs[15])}, []*Node{anonfield(typs[15])})
	typs[104] = functype(nil, []*Node{anonfield(typs[17]), anonfield(typs[17])}, []*Node{anonfield(typs[17])})
	typs[105] = functype(nil, []*Node{anonfield(typs[13])}, []*Node{anonfield(typs[15])})
	typs[106] = functype(nil, []*Node{anonfield(typs[13])}, []*Node{anonfield(typs[17])})
	typs[107] = functype(nil, []*Node{anonfield(typs[13])}, []*Node{anonfield(typs[61])})
	typs[108] = functype(nil, []*Node{anonfield(typs[15])}, []*Node{anonfield(typs[13])})
	typs[109] = functype(nil, []*Node{anonfield(typs[17])}, []*Node{anonfield(typs[13])})
	typs[110] = functype(nil, []*Node{anonfield(typs[61])}, []*Node{anonfield(typs[13])})
	typs[111] = functype(nil, []*Node{anonfield(typs[19]), anonfield(typs[19])}, []*Node{anonfield(typs[19])})
	typs[112] = functype(nil, []*Node{anonfield(typs[47])}, nil)
	typs[113] = functype(nil, []*Node{anonfield(typs[47]), anonfield(typs[47])}, nil)
	return typs[:]
}
