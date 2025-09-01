package middleware

import (
	"fmt"
	"reflect"
	"strconv"
	"strings"
)

func FieldMapper[T any](data T, queryField string) T {

	dataType := reflect.TypeOf(data)
	dataValue := reflect.ValueOf(&data).Elem()

	fields := strings.Split(queryField, ",")

	for _, v := range fields {
		fieldName := getFieldTag(dataType, v)
		if fieldName != "" {
			fieldValue := dataValue.FieldByName(fieldName)
			if fieldValue.IsValid() && fieldValue.CanSet() {
				if fieldValue.Type().Kind() == reflect.String {
					fieldValue.SetString(v)
				}
				if fieldValue.Type().Kind() == reflect.Int {
					r, _ := strconv.Atoi(v)
					fieldValue.SetInt(int64(r))
				}

			}
		}
	}

	return data
}

func getFieldTag(t reflect.Type, value string) string {
	for i := 0; i < t.NumField(); i++ {
		field := t.Field(i)
		tag := field.Tag.Get("json")
		tags := strings.Split(tag, ",")
		if tags[0] == value {
			return field.Name
		}
	}
	return ""
}

func MapSelectFields(selectStr string, mapping map[string]string) ([]string, error) {
	if selectStr == "" {
		return nil, nil // ou return todos os campos padrão
	}

	fields := strings.Split(selectStr, ",")
	var mapped []string
	for _, f := range fields {
		f = strings.TrimSpace(f)
		realField, ok := mapping[f]
		if !ok {
			return nil, fmt.Errorf("field %q not allowed", f)
		}
		mapped = append(mapped, realField)
	}
	return mapped, nil
}

func quoteIdentifier(identifier string) string {
	return `"` + strings.ReplaceAll(identifier, `"`, `""`) + `"`
}

func MapFormToStruct(form map[string][]string, dst interface{}) error {
	dstVal := reflect.ValueOf(dst).Elem()
	dstType := dstVal.Type()

	for i := 0; i < dstType.NumField(); i++ {
		field := dstType.Field(i)
		formTag := field.Tag.Get("form")
		if values, ok := form[formTag]; ok && len(values) > 0 {
			fieldVal := dstVal.FieldByName(field.Name)
			if fieldVal.CanSet() {
				fieldVal.SetString(values[0])
			}
		}
	}
	return nil
}

func ExtractNameFromContentDisposition(header string) string {
	parts := strings.Split(header, ";")
	for _, part := range parts {
		part = strings.TrimSpace(part)
		if strings.HasPrefix(part, "name=") {
			return strings.Trim(part[5:], `"`)
		}
	}
	return ""
}
