package bind

import (
	"net/http"
	"reflect"
	"strconv"
)

func ParseQuery(r *http.Request, req interface{}) error {
	// Parse form body (application/x-www-form-urlencoded or multipart/form-data)
	if err := r.ParseForm(); err != nil {
		return err
	}

	// Get the struct's reflection value
	val := reflect.ValueOf(req).Elem()
	typ := val.Type()

	// Iterate over struct fields
	for i := 0; i < typ.NumField(); i++ {
		field := typ.Field(i)
		fieldVal := val.Field(i)

		// Get the form tag to determine the parameter name
		formTag := field.Tag.Get("form")
		if formTag == "" || formTag == "-" {
			continue
		}

		// Check form body first, then query parameters
		paramValue := r.Form.Get(formTag)
		if paramValue == "" {
			paramValue = r.URL.Query().Get(formTag)
		}

		if paramValue == "" {
			continue
		}

		// Set the field based on its type
		switch fieldVal.Kind() {
		case reflect.String:
			fieldVal.SetString(paramValue)
		case reflect.Int64:
			if num, err := strconv.ParseInt(paramValue, 10, 64); err == nil {
				fieldVal.SetInt(num)
			} else {
				return err
			}
		case reflect.Bool:
			fieldVal.SetBool(paramValue == "true")
		default:
			// Add support for other types as needed
			continue
		}
	}

	return nil
}
