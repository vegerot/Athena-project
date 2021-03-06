// Code generated by go tool dist; DO NOT EDIT.
// This is a bootstrap copy of /tmp/tmp.nzFSOQ2SAm/Go-1.12/go/src/cmd/compile/internal/arm/galign.go

//line /tmp/tmp.nzFSOQ2SAm/Go-1.12/go/src/cmd/compile/internal/arm/galign.go:1
// Copyright 2009 The Go Authors. All rights reserved.
// Use of this source code is governed by a BSD-style
// license that can be found in the LICENSE file.

package arm

import (
	"bootstrap/cmd/compile/internal/gc"
	"bootstrap/cmd/compile/internal/ssa"
	"bootstrap/cmd/internal/obj/arm"
	"bootstrap/cmd/internal/objabi"
)

func Init(arch *gc.Arch) {
	arch.LinkArch = &arm.Linkarm
	arch.REGSP = arm.REGSP
	arch.MAXWIDTH = (1 << 32) - 1
	arch.SoftFloat = objabi.GOARM == 5
	arch.ZeroRange = zerorange
	arch.ZeroAuto = zeroAuto
	arch.Ginsnop = ginsnop
	arch.Ginsnopdefer = ginsnop

	arch.SSAMarkMoves = func(s *gc.SSAGenState, b *ssa.Block) {}
	arch.SSAGenValue = ssaGenValue
	arch.SSAGenBlock = ssaGenBlock
}
