package client

import (
	"net/url"
	"regexp"
	"strconv"
	"strings"
)

// Pagination of results
// For more details, see：
// 中文: https://developer.shopline.com/zh-hans-cn/docs/apps/api-instructions-for-use/paging-mechanism?version=v20251201
// en: https://developer.shopline.com/docs/apps/api-instructions-for-use/paging-mechanism?version=v20251201
type Pagination struct {
	Next     *ListOptions
	Previous *ListOptions
}

// ListOptions
// General list options that can be used for most collections of entities.
type ListOptions struct {
	// PageInfo is used with new pagination search.
	PageInfo string `url:"page_info,omitempty"`

	SinceId *uint64 `url:"since_id,omitempty"`
	Limit   int     `url:"limit,omitempty"`
	Fields  string  `url:"fields,omitempty"`
}

// linkRegex is used to parse the pagination link from shopline API search results.
var linkRegex = regexp.MustCompile(`^ *<([^>]+)>; rel="(previous|next)" *$`)

// parsePagination
// linkHeader eg: <https://fafafa.myshopline.com/admin/openapi/v33322/products/products.json?limit=1&page_info=eyJzaW5jZUlkIjoiMTYwNTc1OTAxNTM4OTA4Mjk1MjExMTI3ODgiLCJkaXJlY3Rpb24iOiJuZXh0IiwibGltaXQiOjF9>; rel="next"
// 中文: https://developer.shopline.com/zh-hans-cn/docs/apps/api-instructions-for-use/paging-mechanism?version=v20251201
// en: https://developer.shopline.com/docs/apps/api-instructions-for-use/paging-mechanism?version=v20251201
func parsePagination(linkHeader string) (*Pagination, error) {
	if linkHeader == "" {
		return nil, nil
	}

	pagination := new(Pagination)
	for _, link := range strings.Split(linkHeader, ",") {
		match := linkRegex.FindStringSubmatch(link)
		// Make sure the link is not empty or invalid
		if len(match) != 3 {
			// We expect 3 values:
			// match[0] = full match
			// match[1] is the URL and match[2] is either 'previous' or 'next'
			err := ResponseDecodingError{
				Message: "could not extract pagination link header",
			}
			return pagination, err
		}

		queryURL, err := url.Parse(match[1])
		if err != nil {
			err = ResponseDecodingError{
				Message: "pagination does not contain a valid URL",
			}
			return pagination, err
		}

		params, err := url.ParseQuery(queryURL.RawQuery)
		if err != nil {
			return pagination, err
		}

		paginationListOptions := ListOptions{}

		paginationListOptions.PageInfo = params.Get("page_info")
		if paginationListOptions.PageInfo == "" {
			err = ResponseDecodingError{
				Message: "The page_info is missing",
			}
			return pagination, err
		}

		limit := params.Get("limit")
		if limit != "" {
			paginationListOptions.Limit, err = strconv.Atoi(params.Get("limit"))
			if err != nil {
				return pagination, err
			}
		}

		fields := params.Get("fields")
		if fields != "" {
			paginationListOptions.Fields = params.Get("fields")
		}

		// 'rel' is either next or previous
		if match[2] == "next" {
			pagination.Next = &paginationListOptions
		} else {
			pagination.Previous = &paginationListOptions
		}
	}

	return pagination, nil
}
