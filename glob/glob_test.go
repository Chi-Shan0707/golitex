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

package litex_global

import (
	"fmt"
	"log"
	"runtime"
	"testing"
)

func TestRuntimeGetFuncName(t *testing.T) {
	pc, _, _, _ := runtime.Caller(0)
	fcName := runtime.FuncForPC(pc).Name()
	log.Println(NewErrLinkAtFunc(nil, fcName, "").Error())
}

func TestIsValidName(t *testing.T) {
	tests := []struct {
		name string
		want bool
	}{
		{"变量", true},
		{"αβγ", true},
		{"_name", true},
		{"name123", true},
		{"🍎", true},         // emoji
		{"東京", true},        // 日文
		{"user@name", true}, // 特殊符号（现在允许）
		{"123name", false},  // 数字开头
		{"__secret", false}, // 双下划线开头
		{"", false},         // 空字符串
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := IsValidName(tt.name); (got == nil) != tt.want {
				t.Errorf("IsValidName(%q) = %v, want %v", tt.name, got, tt.want)
			}
		})
	}
}

func TestFcEval(t *testing.T) {
	testCases := []struct {
		a, b, expected string
	}{
		{"123.45", "67.89", "191.34"},
		{"0.1", "0.2", "0.3"},
		{"99999999999999999999.99999999999999999999", "0.00000000000000000001", "100000000000000000000.00000000000000000000"},
		{"1.0000000000000000000000000000000000000001", "2.0000000000000000000000000000000000000002", "3.0000000000000000000000000000000000000003"},
	}

	for _, tc := range testCases {
		result, _, _ := addBigFloat(tc.a, tc.b)
		fmt.Printf("%s + %s = %s (期望: %s) ", tc.a, tc.b, result, tc.expected)
		if cmpBigFloat(result, tc.expected) == 0 {
			fmt.Println("✓")
		} else {
			fmt.Println("✗")
		}
	}

	fmt.Println(cmpBigFloat("1.23", "1.23000"))    // 0
	fmt.Println(cmpBigFloat("1.23", "1.24"))       // -1
	fmt.Println(cmpBigFloat("123.456", "123.456")) // 0
	fmt.Println(cmpBigFloat("123.456", "123.455")) // 1
	fmt.Println(cmpBigFloat("00001.000", "1"))     // 0
	fmt.Println(cmpBigFloat("10.00001", "10"))     // 1
	fmt.Println(cmpBigFloat("9.9999", "10"))       // -1

}
