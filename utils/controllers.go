package utils

import (
	"fmt"
	"net/http"
	"net/url"
	"strconv"
)

func ExtractPaginationParams(r *http.Request) (int, int) {
	limitStr := r.URL.Query().Get("limit")
	if limitStr == "" {
		limitStr = "5"
	}

	limit, err := strconv.Atoi(limitStr)
	if err != nil {
		limit = 5
	}

	offsetStr := r.URL.Query().Get("offset")
	if offsetStr == "" {
		offsetStr = "0"
	}

	offset, err := strconv.Atoi(offsetStr)
	if err != nil {
		offset = 0
	}

	return limit, offset
}

func GetPaginationUrl(limit, offset, totalRows int, r *http.Request) string {
	remainingRows := totalRows - offset

	if remainingRows <= limit {
		limit = remainingRows
	}

	offset = offset + limit

	var nextPaginationUrl string

	if offset != totalRows {
		parsedUrl, err := url.Parse(r.URL.String())
		if err != nil {
			return ""
		}

		queryParams := parsedUrl.Query()
		queryParams.Set("limit", fmt.Sprintf("%d", limit))
		queryParams.Set("offset", fmt.Sprintf("%d", offset))
		parsedUrl.RawQuery = queryParams.Encode()
		nextPaginationUrl = parsedUrl.RequestURI()
	}

	return nextPaginationUrl
}
