package db

import (
	"database/sql"
	"fmt"
	"log"
	"reflect"
	"strings"
)

func buildInsertQuery(doc any) (string, []any, error) {
	val := reflect.ValueOf(doc)
	if val.Kind() == reflect.Ptr {
		val = val.Elem()
	}
	if val.Kind() != reflect.Struct {
		return "", nil, fmt.Errorf("doc must be a struct")
	}

	var fields []string
	var placeholders []string
	var args []any

	typ := val.Type()
	for i := 0; i < val.NumField(); i++ {
		field := typ.Field(i)
		dbTag := field.Tag.Get("db")
		if dbTag == "" || dbTag == "id" { // Skip ID or untagged fields
			continue
		}
		fields = append(fields, dbTag)
		placeholders = append(placeholders, "?")
		args = append(args, val.Field(i).Interface())
	}

	query := fmt.Sprintf("(%s) VALUES (%s)", strings.Join(fields, ", "), strings.Join(placeholders, ", "))
	return query, args, nil
}

func buildBatchInsertQuery(docs any) (string, []any, error) {
	sliceVal := reflect.ValueOf(docs)
	if sliceVal.Kind() != reflect.Slice {
		return "", nil, fmt.Errorf("docs must be a slice")
	}

	if sliceVal.Len() == 0 {
		return "", nil, fmt.Errorf("empty slice")
	}

	var fields []string
	var placeholders []string
	var args []any

	first := sliceVal.Index(0)
	if first.Kind() == reflect.Ptr {
		first = first.Elem()
	}
	typ := first.Type()
	for i := 0; i < first.NumField(); i++ {
		field := typ.Field(i)
		dbTag := field.Tag.Get("db")
		if dbTag == "" || dbTag == "id" {
			continue
		}
		fields = append(fields, dbTag)
	}

	for i := 0; i < sliceVal.Len(); i++ {
		item := sliceVal.Index(i)
		if item.Kind() == reflect.Ptr {
			item = item.Elem()
		}
		var rowPlaceholders []string
		for j := 0; j < item.NumField(); j++ {
			dbTag := typ.Field(j).Tag.Get("db")
			if dbTag == "" || dbTag == "id" {
				continue
			}
			rowPlaceholders = append(rowPlaceholders, "?")
			args = append(args, item.Field(j).Interface())
		}
		placeholders = append(placeholders, fmt.Sprintf("(%s)", strings.Join(rowPlaceholders, ", ")))
	}

	query := fmt.Sprintf("(%s) VALUES %s", strings.Join(fields, ", "), strings.Join(placeholders, ", "))
	return query, args, nil
}

func buildUpdateQuery(doc any) (string, []any, error) {
	val := reflect.ValueOf(doc)
	if val.Kind() == reflect.Ptr {
		val = val.Elem()
	}
	if val.Kind() != reflect.Struct {
		return "", nil, fmt.Errorf("doc must be a struct")
	}

	var sets []string
	var args []any
	var id any

	typ := val.Type()
	for i := 0; i < val.NumField(); i++ {
		field := typ.Field(i)
		dbTag := field.Tag.Get("db")
		if dbTag == "" {
			continue
		}
		if dbTag == "id" {
			id = val.Field(i).Interface()
			continue
		}
		sets = append(sets, fmt.Sprintf("%s = ?", dbTag))
		args = append(args, val.Field(i).Interface())
	}

	if id == nil {
		return "", nil, fmt.Errorf("id field is required for update")
	}

	args = append(args, id)
	query := fmt.Sprintf("SET %s WHERE id = ?", strings.Join(sets, ", "))
	return query, args, nil
}

func buildDeleteQuery(tableName string, value any, opts ...FindOption) (string, []any) {
	query := fmt.Sprintf("DELETE FROM %s", tableName)
	args := []any{}
	if len(opts) > 0 {
		opt := getOption(opts...)
		if len(opt.query) > 0 {
			var conditions []string
			for _, q := range opt.query {
				conditions = append(conditions, q.Query)
				args = append(args, q.Args...)
			}
			query += " WHERE " + strings.Join(conditions, " AND ")
		}
	}
	return query, args
}

func buildSelectQuery(tableName string, result any, opts ...FindOption) (string, []any) {
	query := fmt.Sprintf("SELECT * FROM %s", tableName)
	args := []any{}
	opt := getOption(opts...)

	if len(opt.preloads) > 0 {
		// Placeholder for JOINs if needed; not implemented for simplicity
	}

	if len(opt.query) > 0 {
		var conditions []string
		for _, q := range opt.query {
			conditions = append(conditions, q.Query)
			args = append(args, q.Args...)
		}
		query += " WHERE " + strings.Join(conditions, " AND ")
	}

	if opt.order != "" {
		query += fmt.Sprintf(" ORDER BY %v", opt.order)
	}

	if opt.offset != 0 {
		query += fmt.Sprintf(" OFFSET %d", opt.offset)
	}

	if opt.limit != 0 {
		query += fmt.Sprintf(" LIMIT %d", opt.limit)
	}

	log.Printf("buildSelectQuery: %s, Args: %v", query, args)
	return query, args
}

func buildCountQuery(tableName string, model any, opts ...FindOption) (string, []any) {
	query := fmt.Sprintf("SELECT COUNT(*) FROM %s", tableName)
	args := []any{}
	if len(opts) > 0 {
		opt := getOption(opts...)
		if len(opt.query) > 0 {
			var conditions []string
			for _, q := range opt.query {
				conditions = append(conditions, q.Query)
				args = append(args, q.Args...)
			}
			query += " WHERE " + strings.Join(conditions, " AND ")
		}
	}
	return query, args
}

// scanRow scans a single row into a struct.
func scanRow(row *sql.Row, result any) error {
	val := reflect.ValueOf(result)
	if val.Kind() != reflect.Ptr || val.Elem().Kind() != reflect.Struct {
		return fmt.Errorf("result must be a pointer to a struct")
	}

	columns, err := getColumns(val.Elem().Type())
	if err != nil {
		return err
	}

	values := make([]any, len(columns))
	for i := range values {
		values[i] = new(any)
	}

	err = row.Scan(values...)
	if err != nil {
		return err
	}

	return setStructFields(val.Elem(), columns, values)
}

// scanRows scans multiple rows into a slice of structs.
func scanRows(rows *sql.Rows, result any) error {
	sliceVal := reflect.ValueOf(result)
	if sliceVal.Kind() != reflect.Ptr || sliceVal.Elem().Kind() != reflect.Slice {
		return fmt.Errorf("result must be a pointer to a slice")
	}

	// Get the slice element type (e.g., *User)
	elemType := sliceVal.Elem().Type().Elem()
	// For getColumns, use the struct type (User)
	structType := elemType
	if elemType.Kind() == reflect.Ptr {
		structType = elemType.Elem()
	}
	if structType.Kind() != reflect.Struct {
		return fmt.Errorf("slice elements must be structs or pointers to structs, got %v", structType.Kind())
	}

	columns, err := getColumns(structType)
	if err != nil {
		return err
	}

	var results []reflect.Value
	for rows.Next() {
		values := make([]any, len(columns))
		for i := range values {
			values[i] = new(any)
		}

		err := rows.Scan(values...)
		if err != nil {
			return err
		}

		// Create a new *User (keep the pointer)
		newItem := reflect.New(structType)
		err = setStructFields(newItem.Elem(), columns, values)
		if err != nil {
			return err
		}
		results = append(results, newItem) // Store *User, not User
	}

	slice := reflect.MakeSlice(sliceVal.Elem().Type(), len(results), len(results))
	for i, item := range results {
		slice.Index(i).Set(item) // item is *User, matches slice type
	}
	sliceVal.Elem().Set(slice)

	return nil
}

// getColumns gets the database column names from struct tags.
func getColumns(t reflect.Type) ([]string, error) {
	var columns []string
	for i := 0; i < t.NumField(); i++ {
		dbTag := t.Field(i).Tag.Get("db")
		if dbTag == "" {
			continue
		}
		columns = append(columns, dbTag)
	}
	if len(columns) == 0 {
		return nil, fmt.Errorf("no db tags found in struct")
	}
	return columns, nil
}

// setStructFields, scanRow, scanRows, getColumns remain unchanged
func setStructFields(val reflect.Value, columns []string, values []any) error {
	for i, col := range columns {
		for j := 0; j < val.NumField(); j++ {
			if val.Type().Field(j).Tag.Get("db") == col {
				field := val.Field(j)
				if !field.CanSet() {
					continue
				}
				v := reflect.ValueOf(values[i]).Elem().Interface()
				switch field.Type().Kind() {
				case reflect.String:
					if bytes, ok := v.([]uint8); ok {
						field.SetString(string(bytes))
					} else {
						field.Set(reflect.ValueOf(v))
					}
				case reflect.Ptr:
					switch field.Type().Elem().Kind() {
					case reflect.String:
						if bytes, ok := v.([]uint8); ok {
							str := string(bytes)
							field.Set(reflect.ValueOf(&str))
						} else if str, ok := v.(string); ok {
							field.Set(reflect.ValueOf(&str))
						} else if v == nil {
							field.Set(reflect.Zero(field.Type()))
						}
					case reflect.Int:
						if intVal, ok := v.(int64); ok {
							int32Val := int(intVal)
							field.Set(reflect.ValueOf(&int32Val))
						} else if v == nil {
							field.Set(reflect.Zero(field.Type()))
						}
					default:
						field.Set(reflect.ValueOf(v))
					}
				default:
					field.Set(reflect.ValueOf(v))
				}
				break
			}
		}
	}
	return nil
}
