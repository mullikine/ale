// Code generated by "stringer -type=Opcode"; DO NOT EDIT.

package isa

import "strconv"

func _() {
	// An "invalid array index" compiler error signifies that the constant values have changed.
	// Re-run the stringer command to generate them again.
	var x [1]struct{}
	_ = x[Label-256]
	_ = x[Add-0]
	_ = x[Arg-1]
	_ = x[ArgLen-2]
	_ = x[Bind-3]
	_ = x[BindRef-4]
	_ = x[Call-5]
	_ = x[Call0-6]
	_ = x[Call1-7]
	_ = x[Closure-8]
	_ = x[CondJump-9]
	_ = x[Const-10]
	_ = x[Declare-11]
	_ = x[Deref-12]
	_ = x[Div-13]
	_ = x[Dup-14]
	_ = x[Eq-15]
	_ = x[False-16]
	_ = x[Gt-17]
	_ = x[Gte-18]
	_ = x[Jump-19]
	_ = x[Load-20]
	_ = x[Lt-21]
	_ = x[Lte-22]
	_ = x[MakeCall-23]
	_ = x[MakeTruthy-24]
	_ = x[Mod-25]
	_ = x[Mul-26]
	_ = x[Neg-27]
	_ = x[NegInf-28]
	_ = x[NegOne-29]
	_ = x[Neq-30]
	_ = x[NewRef-31]
	_ = x[NoOp-32]
	_ = x[Not-33]
	_ = x[Null-34]
	_ = x[One-35]
	_ = x[Panic-36]
	_ = x[Pop-37]
	_ = x[PosInf-38]
	_ = x[Resolve-39]
	_ = x[RestArg-40]
	_ = x[RetEmptyList-41]
	_ = x[RetFalse-42]
	_ = x[RetNull-43]
	_ = x[RetTrue-44]
	_ = x[Return-45]
	_ = x[Self-46]
	_ = x[Store-47]
	_ = x[Sub-48]
	_ = x[TailCall-49]
	_ = x[True-50]
	_ = x[Two-51]
	_ = x[Zero-52]
}

const (
	_Opcode_name_0 = "AddArgArgLenBindBindRefCallCall0Call1ClosureCondJumpConstDeclareDerefDivDupEqFalseGtGteJumpLoadLtLteMakeCallMakeTruthyModMulNegNegInfNegOneNeqNewRefNoOpNotNullOnePanicPopPosInfResolveRestArgRetEmptyListRetFalseRetNullRetTrueReturnSelfStoreSubTailCallTrueTwoZero"
	_Opcode_name_1 = "Label"
)

var (
	_Opcode_index_0 = [...]uint16{0, 3, 6, 12, 16, 23, 27, 32, 37, 44, 52, 57, 64, 69, 72, 75, 77, 82, 84, 87, 91, 95, 97, 100, 108, 118, 121, 124, 127, 133, 139, 142, 148, 152, 155, 159, 162, 167, 170, 176, 183, 190, 202, 210, 217, 224, 230, 234, 239, 242, 250, 254, 257, 261}
)

func (i Opcode) String() string {
	switch {
	case 0 <= i && i <= 52:
		return _Opcode_name_0[_Opcode_index_0[i]:_Opcode_index_0[i+1]]
	case i == 256:
		return _Opcode_name_1
	default:
		return "Opcode(" + strconv.FormatInt(int64(i), 10) + ")"
	}
}
