package query

import (
	"monolith/internal/module/customer/repository/internal/model"
	"monolith/pkg/goqu"
)

func GetInsert(model model.Customer) (string, []any, error) {
	record := make(goqu.Record)

	if model.Username.Valid {
		record["username"] = model.Username
	}

	record["email"] = model.Email
	record["password_hash"] = model.Password

	return goqu.Dialect.
		Insert(model.TableName()).
		Rows(record).
		Returning("*").
		Prepared(true).
		ToSQL()
}

func GetSelect(model model.Customer) (string, []any, error) {
	whereExpressions := make(goqu.Ex)

	if model.Id != 0 {
		whereExpressions["id"] = model.Id
	} else {
		whereExpressions["email"] = model.Email
		whereExpressions["password_hash"] = model.Password
	}

	return goqu.Dialect.
		Select("*").
		From(model.TableName()).
		Where(whereExpressions).
		Prepared(true).
		ToSQL()
}

func GetDelete(model model.Customer) (string, []any, error) {
	return goqu.Dialect.
		Delete(model.TableName()).
		Where(goqu.C("id").Eq(model.Id)).
		Prepared(true).
		ToSQL()
}

func GetUpdate(oldModel, model model.Customer) (string, []any, error) {
	record := make(goqu.Record)

	model.Id = oldModel.Id

	if oldModel.Username.Valid || model.Username.Valid {

	}

	if oldModel.Email != model.Email {
		record["email"] = model.Email
	}
	if oldModel.Password != model.Password {
		record["password"] = model.Password
	}

	return goqu.Dialect.
		Update(model.TableName()).
		Set(record).
		Where(goqu.C("id").Eq(model.Id)).
		Prepared(true).
		ToSQL()
}
