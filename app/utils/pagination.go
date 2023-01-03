package utils

func TranslatePagination(page, perPage int) (int, int) {
	limit := perPage
	if limit < 1 {
		limit = 1
	}

	offset := 0
	if page > 1 {
		offset = (page * perPage) - perPage
	}
	return perPage, offset
}
