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

package litex_global

func CartesianProduct[T any](input [][]T) [][]T {
	var res [][]T
	var dfs func(int, []T)

	dfs = func(depth int, path []T) {
		if depth == len(input) {
			combo := make([]T, len(path))
			copy(combo, path)
			res = append(res, combo)
			return
		}
		for _, val := range input[depth] {
			dfs(depth+1, append(path, val))
		}
	}

	if len(input) == 0 {
		return nil
	}

	dfs(0, []T{})
	return res
}
