package vm

import (
	"monkey-lang/ast"
	"monkey-lang/compiler"
	"monkey-lang/lexer"
	"monkey-lang/object"
	"monkey-lang/parser"
	"testing"
)

func TestIntegerArithmetic(t *testing.T) {
	tests := []struct { input string; expected int64 }{
		{"1", 1},
		{"2", 2},
		{"1 + 2", 3},
		{"1 - 2", -1},
		{"1 * 2", 2},
		{"4 / 2", 2},
	}
	for _, tt := range tests {
		program := parse(tt.input)
		comp := compiler.New()
		if err := comp.Compile(program); err != nil { t.Fatalf("compiler error: %s", err) }
		machine := New(comp.Bytecode())
		if err := machine.Run(); err != nil { t.Fatalf("vm error: %s", err) }
		stackElem := machine.LastPoppedStackElem()
		integer, ok := stackElem.(*object.Integer)
		if !ok { t.Fatalf("object is not Integer. got=%T (%+v)", stackElem, stackElem) }
		if integer.Value != tt.expected { t.Fatalf("wrong result. want=%d, got=%d", tt.expected, integer.Value) }
	}
}

func parse(input string) *ast.Program { l := lexer.New(input); p := parser.New(l); return p.ParseProgram() }


