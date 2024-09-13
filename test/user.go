/*
 * The Clear BSD License
 *
 * Copyright (c) 2024, DoytoWin, Inc.
 * All rights reserved.
 *
 * This source code is licensed under the BSD-style license found in the
 * LICENSE file in the root directory of this source tree.
 */

package test

import . "github.com/doytowin/goooqo/core"

type UserEntity struct {
	Int64Id
	Score *int    `json:"score"`
	Memo  *string `json:"memo"`
}

func (u UserEntity) GetTableName() string {
	return "User"
}

type UserQuery struct {
	PageQuery
	IdGt     *int
	IdIn     *[]int
	ScoreLt  *int
	MemoNull *bool
	MemoLike *string
	Deleted  *bool

	ScoreLtAvg *UserQuery `subquery:"select avg(score) from User"`
	ScoreLtAny *UserQuery `subquery:"SELECT score FROM User"`
	ScoreLtAll *UserQuery `subquery:"select score from User"`
}
