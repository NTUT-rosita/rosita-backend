package sliceConverter

import "database/sql"

func NullStringToString(nullStrings []sql.NullString) []string {
	result := make([]string, 0)
	for _, nullString := range nullStrings {
		if nullString.Valid {
			result = append(result, nullString.String)
		}
	}

	return result
}
