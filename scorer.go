// Copyright 2015 Mathieu MAST. All rights reserved.
// Use of this source code is governed by a MIT style
// license that can be found in the LICENSE file.
package goloc

import (
	"strings"
)

type Scorer func(*Result) int

func DefaultScorer(result *Result) int {
	s := Score(result.Search, result.Name)
	if strings.HasPrefix(result.Type, "street") {
		s -= 3
	}
	if strings.HasPrefix(result.Type, "poi") {
		s -= 2
	}
	if strings.HasPrefix(result.Type, "zone") {
		s -= 1
	}
	return s
}