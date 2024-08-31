/*
 * The Clear BSD License
 *
 * Copyright (c) 2024, DoytoWin, Inc.
 * All rights reserved.
 *
 * This source code is licensed under the BSD-style license found in the
 * LICENSE file in the root directory of this source tree.
 */

package mongodb

import (
	. "go.mongodb.org/mongo-driver/bson/primitive"
	"regexp"
)

var sortRgx = regexp.MustCompile("(?i)(\\w+)(,(asC|dEsc))?;?")

func buildSort(sort string) D {
	submatch := sortRgx.FindAllStringSubmatch(sort, -1)
	result := make(D, len(submatch))
	for i, group := range submatch {
		if group[3] != "" {
			result[i] = E{group[1], 7 - len(group[3])*2}
		} else {
			result[i] = E{group[1], 1}
		}
	}
	return result
}
