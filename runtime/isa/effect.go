package isa

import "fmt"

// Error messages
const (
	EffectNotDeclared = "effect not declared for opcode: %s"
)

// Effect captures how an instruction affects the stack and PC
type Effect struct {
	Size   int
	Pop    int
	Push   int
	DPop   int
	DPush  int
	Ignore bool
	Exit   bool
}

// Effects is a lookup table of instruction effects
var Effects = map[Opcode]*Effect{
	Add:          {Size: 1, Pop: 2, Push: 1},
	Arg:          {Size: 2, Push: 1},
	ArgLen:       {Size: 1, Push: 1},
	Bind:         {Size: 1, Pop: 2},
	BindRef:      {Size: 1, Pop: 2},
	Call:         {Size: 2, Pop: 1, Push: 1, DPop: 1},
	Call0:        {Size: 1, Pop: 1, Push: 1},
	Call1:        {Size: 1, Pop: 2, Push: 1},
	Closure:      {Size: 2, Push: 1},
	CondJump:     {Size: 2, Pop: 1},
	Const:        {Size: 2, Push: 1},
	Declare:      {Size: 1, Pop: 1},
	Deref:        {Size: 1, Pop: 1, Push: 1},
	Div:          {Size: 1, Pop: 2, Push: 1},
	Dup:          {Size: 1, Pop: 1, Push: 2},
	Eq:           {Size: 1, Pop: 2, Push: 1},
	False:        {Size: 1, Push: 1},
	Gt:           {Size: 1, Pop: 2, Push: 1},
	Gte:          {Size: 1, Pop: 2, Push: 1},
	Jump:         {Size: 2},
	Label:        {Size: 2, Ignore: true},
	Load:         {Size: 2, Push: 1},
	Lt:           {Size: 1, Pop: 2, Push: 1},
	Lte:          {Size: 1, Pop: 2, Push: 1},
	MakeCall:     {Size: 1, Pop: 1, Push: 1},
	MakeTruthy:   {Size: 1, Pop: 1, Push: 1},
	Mod:          {Size: 1, Pop: 2, Push: 1},
	Mul:          {Size: 1, Pop: 2, Push: 1},
	Neg:          {Size: 1, Pop: 1, Push: 1},
	NegInf:       {Size: 1, Push: 1},
	NegOne:       {Size: 1, Push: 1},
	Neq:          {Size: 1, Pop: 2, Push: 1},
	NewRef:       {Size: 1, Push: 1},
	NoOp:         {Size: 1, Ignore: true},
	Not:          {Size: 1, Pop: 1, Push: 1},
	Null:         {Size: 1, Push: 1},
	One:          {Size: 1, Push: 1},
	Panic:        {Size: 1, Pop: 1, Exit: true},
	Pop:          {Size: 1, Pop: 1},
	PosInf:       {Size: 1, Push: 1},
	Resolve:      {Size: 1, Pop: 1, Push: 1},
	RestArg:      {Size: 2, Push: 1},
	RetEmptyList: {Size: 1, Exit: true},
	RetFalse:     {Size: 1, Exit: true},
	RetNull:      {Size: 1, Exit: true},
	RetTrue:      {Size: 1, Exit: true},
	Return:       {Size: 1, Pop: 1, Exit: true},
	Self:         {Size: 1, Push: 1},
	Store:        {Size: 2, Pop: 1},
	Sub:          {Size: 1, Pop: 2, Push: 1},
	TailCall:     {Size: 2, Pop: 1, DPop: 1},
	True:         {Size: 1, Push: 1},
	Two:          {Size: 1, Push: 1},
	Zero:         {Size: 1, Push: 1},
}

// MustGetEffect gives you effect information or explodes violently
func MustGetEffect(oc Opcode) *Effect {
	if effect, ok := Effects[oc]; ok {
		return effect
	}
	panic(fmt.Sprintf(EffectNotDeclared, oc.String()))
}
