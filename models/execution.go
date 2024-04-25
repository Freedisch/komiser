package models

import (
	"context"
	"database/sql"
	"fmt"

	"github.com/uptrace/bun"
)

type QueryType string

const (
	RAW    QueryType = "RAW"
	SELECT QueryType = "SELECT"
	INSERT QueryType = "INSERT"
	DELETE QueryType = "DELETE"
	UPDATE QueryType = "UPDATE"
)

func HandleQuery(ctx context.Context, db *bun.DB, queryTitle string, schema interface{}, conditions [][3]string) (sql.Result, error) {
	var resp sql.Result
	var err error
	query := Queries[queryTitle]
	switch query.Type {
	case RAW:
		err = executeRaw(ctx, db, query.Query, schema, conditions)

	case SELECT:
		err = executeSelect(ctx, db, query.Query, schema, conditions)

	case INSERT:
		resp, err = executeInsert(ctx, db, schema, conditions)

	case DELETE:
		resp, err = executeDelete(ctx, db, schema, query.Query, conditions)

	case UPDATE:
		resp, err = executeUpdate(ctx, db, schema, query.Params, conditions)
	}
	return resp, err
}

func executeRaw(ctx context.Context, db *bun.DB, query string, schema interface{}, additionals [][3]string) error {
	if len(additionals) > 0 {
		query = fmt.Sprintf("%s where", query)
	}

	for _, triplet := range additionals {
		key, op, value := triplet[0], triplet[1], triplet[2]
		query = fmt.Sprintf("%s %s %s '%s' and", query, key, op, value)
	}

	if len(additionals) > 0 {
		query = query[:len(query)-4]
	}

	err := db.NewRaw(query).Scan(ctx, schema)
	if err != nil {
		return err
	}
	return nil
}

func executeSelect(ctx context.Context, db *bun.DB, query string, schema interface{}, conditions [][3]string) error {
	q := db.NewSelect().Model(schema)

	q = addWhereClause(q.QueryBuilder(), conditions).Unwrap().(*bun.SelectQuery)

	return q.Scan(ctx, schema)
}

func executeInsert(ctx context.Context, db *bun.DB, schema interface{}, conditions [][3]string) (sql.Result, error) {
	resp, err := db.NewInsert().Model(schema).Exec(ctx)
	if err != nil {
		return resp, err
	}
	return resp, nil
}

func executeDelete(ctx context.Context, db *bun.DB, schema interface{}, query string, conditions [][3]string) (sql.Result, error) {
	q := db.NewDelete().Model(schema)

	q = addWhereClause(q.QueryBuilder(), conditions).Unwrap().(*bun.DeleteQuery)

	resp, err := q.Exec(ctx)
	if err != nil {
		return resp, err
	}
	return resp, nil
}

func executeUpdate(ctx context.Context, db *bun.DB, schema interface{}, columns []string, conditions [][3]string) (sql.Result, error) {
	q := db.NewUpdate().Model(schema).Column(columns...)

	q = addWhereClause(q.QueryBuilder(), conditions).Unwrap().(*bun.UpdateQuery)

	q = q.Returning("*")

	resp, err := q.Exec(ctx)
	if err != nil {
		return resp, err
	}
	return resp, nil
}

func addWhereClause(query bun.QueryBuilder, conditions [][3]string) bun.QueryBuilder {
	for _, triplet := range conditions {
		key, op, value := triplet[0], triplet[1], triplet[2]
		query = query.Where(fmt.Sprintf("%s %s ?", key, op), value)
	}
	return query
}
