package query

import (
	"monolith/internal/module/customer/repository/internal/model"

	"github.com/doug-martin/goqu/v9"
)

func GetCreate(model model.Customer) (string, []any, error) {
	record := make(goqu.Record)

	if model.Username.Valid {
		record["username"] = model.Username
	}

	record["email"] = model.Email
	record["password_hash"] = model.Password

	return goqu.Insert(model.TableName()).
		Rows(record).
		Returning("*").
		Prepared(true).
		ToSQL()
}
