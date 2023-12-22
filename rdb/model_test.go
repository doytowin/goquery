package rdb

import . "github.com/doytowin/go-query/core"

type AccountOr struct {
	Username *string
	Email    *string
	Mobile   *string
}

type TestEntity struct {
	Id       *int
	Username *string
	Email    *string
	Mobile   *string
}

func (e TestEntity) GetTableName() string {
	return "t_user"
}

type TestQuery struct {
	PageQuery
	AccountOr *AccountOr
	Account   *string `condition:"(username = ? OR email = ?)"`
	Deleted   *bool
}
