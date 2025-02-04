package utils

func Pagination(hasNext bool, cursor string) interface{} {
	return map[string]interface{}{
		"has_next": hasNext,
		"cursor":   cursor,
	}
}