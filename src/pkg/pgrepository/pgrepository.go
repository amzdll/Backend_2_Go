package pgrepository

import (
	"context"
	"fmt"
	"github.com/amzdll/backend_2_go/src/internal/model"
	"github.com/georgysavva/scany/v2/pgxscan"
	"github.com/google/uuid"
	"github.com/jackc/pgx/v5/pgxpool"
	"reflect"
	"strings"
)

func Create(table string, pool *pgxpool.Pool, ctx context.Context, data any) error {
	if err := checkData(data); err != nil {
		return err
	}
	td := parseCreateData(data)
	query := fmt.Sprintf(`insert into %s (%s) values (%s)`,
		table,
		strings.Join(td.columns, ", "),
		strings.Join(td.placeholders, ", "),
	)
	fmt.Println(query, td.values)
	_, err := pool.Exec(ctx, query, td.values...)
	return err
}

func GetWithFilters[T any](
	table string, pool *pgxpool.Pool, ctx context.Context, filters []string, data interface{},
) ([]T, error) {
	err := checkData(data)
	if err != nil {
		return nil, err
	}
	dataType := reflect.TypeOf(data)
	dataValue := reflect.ValueOf(data)
	var values []interface{}
	var filter strings.Builder
	for _, filterName := range filters {
		field, found := dataType.FieldByName(filterName)
		if !found {
			return nil, fmt.Errorf("field %s not found in struct", filterName)
		}
		fieldValue := dataValue.FieldByName(filterName)
		values = append(values, fieldValue.Interface())
		filter.WriteString(field.Tag.Get("db") + " = ? AND ")
	}

	filterStr := filter.String()
	filterStr = filterStr[:len(filterStr)-5]

	var list []T

	query := fmt.Sprintf("SELECT * FROM %s WHERE %s", table, filterStr)
	err = pgxscan.Select(ctx, pool, &list, query, values...)
	if err != nil {
		return nil, err
	}

	return list, nil
}

func GetAll[T any](table string, pool *pgxpool.Pool, ctx context.Context, pagination model.Pagination) ([]T, error) {
	query := fmt.Sprintf(`select * from %s limit $1 offset $2`, table)
	var list []T
	err := pgxscan.Select(ctx, pool, &list, query, pagination.Limit, pagination.Offset)
	fmt.Println(query, pagination.Limit, pagination.Offset, reflect.TypeOf(list))
	if err != nil {
		return nil, err
	}
	return list, nil
}

func Update(
	table string, pool *pgxpool.Pool, ctx context.Context, filters []string, data interface{},
) error {

	return nil
}

func Delete(table string, pool *pgxpool.Pool, ctx context.Context, id uuid.UUID) error {
	const query = `delete from $1 where id = $2`
	result, err := pool.Exec(ctx, query, table, id)
	if err == nil && result.RowsAffected() == 0 {
		err = fmt.Errorf("not found")
	}
	return err
}
