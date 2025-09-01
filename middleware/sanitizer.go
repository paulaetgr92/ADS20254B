package middleware

import (
	"fmt"
	"regexp"
	"strings"
)

func FilterSanitizer(filter string) string {

	splited := strings.Split(filter, ` `)

	for _, i := range splited {
		switch {
		case strings.Contains(i, "eq"):
			filter = strings.Replace(filter, ` eq `, ` = `, -1)
		case strings.Contains(i, "ne"):
			filter = strings.Replace(filter, ` ne `, `!= `, -1)
		case strings.Contains(i, "gt"):
			filter = strings.Replace(filter, ` gt `, ` > `, -1)
		case strings.Contains(i, "ge"):
			filter = strings.Replace(filter, ` ge `, ` >=`, -1)
		case strings.Contains(i, "lt"):
			filter = strings.Replace(filter, ` lt `, ` < `, -1)
		case strings.Contains(i, "le"):
			filter = strings.Replace(filter, ` le `, ` <= `, -1)
		case strings.Contains(i, "not"):
			filter = strings.Replace(filter, ` not`, `!`, -1)
		case strings.Contains(i, "isent"):
			filter = strings.Replace(filter, ` isent `, ` is not `, -1)
		case strings.Contains(i, "is"):
			filter = strings.Replace(filter, ` is `, ` is `, -1)
		}

	}

	reLike := regexp.MustCompile(`(startswith|contains|endswith)\((?:(tolower|toupper)\()?(\w+)\)?,\s*'(.+?)'\)`)

	_ = reLike.ReplaceAllStringFunc(filter, func(match string) string {

		parts := reLike.FindStringSubmatch(match)
		function, caseFunc, field, value := parts[1], parts[2], parts[3], parts[4]

		sqlField := field
		if caseFunc == "tolower" {
			sqlField = fmt.Sprintf("LOWER(%s)", field)
		} else if caseFunc == "toupper" {
			sqlField = fmt.Sprintf("UPPER(%s)", field)
		}

		switch function {
		case "startswith":
			filter = strings.Replace(filter, parts[0], fmt.Sprintf("%s LIKE '%s%%'", sqlField, value), -1)
			return filter
		case "contains":
			filter = strings.Replace(filter, parts[0], fmt.Sprintf("%s LIKE '%%%s%%'", sqlField, value), -1)
			return filter
		case "endswith":
			filter = strings.Replace(filter, parts[0], fmt.Sprintf("%s LIKE '%%%s'", sqlField, value), -1)
			return filter
		default:
			return filter
		}
	})

	re := regexp.MustCompile(`(tolower|toupper)\((\w+)\)`)

	_ = re.ReplaceAllStringFunc(filter, func(match string) string {

		parts := re.FindStringSubmatch(match)
		function, field := parts[1], parts[2]

		if function == "tolower" {
			filter = strings.Replace(filter, parts[0], fmt.Sprintf("LOWER(%s)", field), -1)
		} else if function == "toupper" {
			filter = strings.Replace(filter, parts[0], fmt.Sprintf("UPPER(%s)", field), -1)
		}
		return filter
	})

	return filter
}
