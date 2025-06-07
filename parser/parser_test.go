// Copyright 2024 Jiachen Shen.
//
// Licensed under the Apache License, Version 2.0 (the "License");
// you may not use this file except in compliance with the License.
// You may obtain a copy of the License at
//
//     http://www.apache.org/licenses/LICENSE-2.0
//
// Original Author: Jiachen Shen <malloc_realloc_free@outlook.com>
// Litex email: <litexlang@outlook.com>
// Litex website: https://litexlang.org
// Litex github repository: https://github.com/litexlang/golitex
// Litex discord server: https://discord.gg/uvrHM7eS

package litex_parser

import (
	"fmt"
	ast "golitex/ast"
	num "golitex/number"
	"strings"
	"testing"
)

func sourceCodeToFc(sourceCode ...string) ([]ast.Fc, error) {
	blocks, err := makeTokenBlocks(sourceCode, NewParserEnv())
	if err != nil {
		return nil, err
	}

	ret := []ast.Fc{}
	for _, block := range blocks {
		cur, err := block.header.rawFc()
		if err != nil {
			return nil, err
		}
		ret = append(ret, cur)
	}

	return ret, nil
}

func TestOrder(t *testing.T) {
	sourceCode := []string{
		"1+2*(4+ t(x)(x)) + 9 + 4*F(t) + (x-y)*(a+b) ",
		"x + x",
		"2*x",
	}
	fcSlice := []ast.Fc{}
	for _, code := range sourceCode {
		fc, err := sourceCodeToFc(code)
		if err != nil {
			t.Fatal(err)
		}
		fcSlice = append(fcSlice, fc...)
	}

	for _, fc := range fcSlice {
		bracketedStr := num.FcStringForParseAndExpandPolynomial(fc)
		fmt.Println(bracketedStr)
		ploy := num.ParseAndExpandPolynomial(bracketedStr)
		var parts []string
		for _, t := range ploy {
			parts = append(parts, t.String())
		}
		fmt.Println("Expanded:", strings.Join(parts, " + "))
	}
}
