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
// Litex Zulip community: https://litex.zulipchat.com/join/c4e7foogy6paz2sghjnbujov/

package litex_env

import (
	ast "golitex/ast"
	glob "golitex/glob"
)

func (e *Env) InsertFnInFnTT(fc ast.Fc, templateWhereFcIs *ast.FcFn, fnTNoName *ast.FnTStruct) error {
	memory := e.FnInFnTemplateFactsMem
	fnDefs, ok := memory[fc.String()]
	if !ok {
		memory[fc.String()] = []FnInFnTMemItem{
			{
				AsFcFn:      templateWhereFcIs,
				AsFnTStruct: fnTNoName,
			},
		}
		return nil
	} else {
		fnDefs = append(fnDefs, FnInFnTMemItem{
			AsFcFn:      templateWhereFcIs,
			AsFnTStruct: fnTNoName,
		})
		memory[fc.String()] = fnDefs
		return nil
	}
}

// 从后往前找，直到找到有个 fnHead 被已知在一个 fnInFnTInterface 中
// 比如 f(a)(b,c)(e,d,f) 我不知道 f(a)(b,c) 是哪个 fn_template 里的，但我发现 f(a) $in T 是知道的。那之后就是按T的返回值去套入b,c，然后再把e,d,f套入T的返回值的返回值
// 此时 indexWhereLatestFnTIsGot 就是 1, FnToFnItemWhereLatestFnTIsGot 就是 f(a) 的 fnInFnTMemItem
func (e *Env) FindRightMostResolvedFn_Return_ResolvedIndexAndFnTMemItem(fnHeadChain_AndItSelf []ast.Fc) (int, *FnInFnTMemItem) {
	indexWhereLatestFnTIsGot := 0
	var latestFnT *FnInFnTMemItem = nil
	for i := len(fnHeadChain_AndItSelf) - 2; i >= 0; i-- {
		fnHead := fnHeadChain_AndItSelf[i]
		if fnInFnTMemItem, ok := e.GetLatestFnT_GivenNameIsIn(fnHead.String()); ok {
			latestFnT = fnInFnTMemItem
			indexWhereLatestFnTIsGot = i
			break
		}
	}

	return indexWhereLatestFnTIsGot, latestFnT
}

func (e *Env) GetFnTStructOfFnInFnTMemItem(fnInFnTMemItem *FnInFnTMemItem) *ast.FnTStruct {
	if fnInFnTMemItem.AsFcFn != nil {
		if ok, paramSets, retSet := fnInFnTMemItem.AsFcFn.IsFnT_FcFn_Ret_ParamSets_And_RetSet(fnInFnTMemItem.AsFcFn); ok {
			excelNames := glob.GenerateNamesLikeExcelColumnNames(len(paramSets))
			return ast.NewFnTStruct(excelNames, paramSets, retSet, []ast.FactStmt{}, []ast.FactStmt{})
		}
	}

	return fnInFnTMemItem.AsFnTStruct
}
