package filter

import (
	"fmt"
	"go/ast"
	"go/parser"
	"go/token"
	"strings"
)

// Parse parses a filter expression string:
//
//  `Not(Combine(Extensions(".go"), Extensions(".html")))`
//
func Parse(filter string) (Func, error, ast.Node) {
	expr, err := parser.ParseExpr(filter)
	if err != nil {
		return nil, err, nil
	}
	return parseExpr(expr)
}

// parseExpr parses a single call expression.
func parseExpr(e ast.Expr) (Func, error, ast.Node) {
	// Validate node is a call expression with an identity.
	v, ok := e.(*ast.CallExpr)
	if !ok {
		return nil, fmt.Errorf("expected call expression"), e
	}
	ident, ok := v.Fun.(*ast.Ident)
	if !ok {
		return nil, fmt.Errorf("expected ident"), ident
	}

	// Parse each filter.
	switch ident.Name {
	case "Combine":
		return parseCombine(v, ident)
	case "Extensions":
		return parseExtensions(v, ident)
	case "Not":
		return parseNot(v, ident)
	default:
		return nil, fmt.Errorf("Unknown filter %q", ident.Name), ident
	}
}

// parseCombine parses a Combine filter.
func parseCombine(v *ast.CallExpr, ident *ast.Ident) (Func, error, ast.Node) {
	// Require at least one argument.
	if len(v.Args) == 0 {
		return nil, fmt.Errorf("%q expects at least one argument", ident.Name), ident
	}

	// Parse each filter argument / call expression and accumulate.
	var filters []Func
	for _, arg := range v.Args {
		filter, err, node := parseExpr(arg)
		if err != nil {
			return nil, err, node
		}
		filters = append(filters, filter)
	}
	return Combine(filters...), nil, nil
}

// parseExtensions parses a Extensions filter.
func parseExtensions(v *ast.CallExpr, ident *ast.Ident) (Func, error, ast.Node) {
	// Require at least one argument.
	if len(v.Args) == 0 {
		return nil, fmt.Errorf("%q expects at least one argument", ident.Name), ident
	}

	// Parse each file extension string literal.
	var exts []string
	for _, arg := range v.Args {
		// Must be a basic literal.
		s, ok := arg.(*ast.BasicLit)
		if !ok {
			return nil, fmt.Errorf("expected string literal"), arg
		}

		// Must be a string.
		if s.Kind != token.STRING {
			return nil, fmt.Errorf("expected string literal"), arg
		}

		// Must be a non-empty string.
		if s.Value == "" {
			return nil, fmt.Errorf("expected filepath extension with leading dot"), arg
		}

		// Trim the quote prefix/suffix from the value.
		s.Value = s.Value[1 : len(s.Value)-1]

		// Ensure the extension starts with a dot (.go, .html, etc).
		if !strings.HasPrefix(s.Value, ".") {
			fmt.Printf("%q\n", s.Value)
			return nil, fmt.Errorf("expected filepath extension with leading dot"), arg
		}
		exts = append(exts, s.Value)
	}
	return Extensions(exts...), nil, nil
}

// parseNot parses a Not filter.
func parseNot(v *ast.CallExpr, ident *ast.Ident) (Func, error, ast.Node) {
	// Require exactly one argument.
	if len(v.Args) != 1 {
		return nil, fmt.Errorf("%q expects one argument", ident.Name), ident
	}

	// Parse filter argument / call expression.
	filter, err, node := parseExpr(v.Args[0])
	if err != nil {
		return nil, err, node
	}
	return Not(filter), nil, nil
}
