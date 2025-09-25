package client

import (
	"github.com/stretchr/testify/assert"
	"testing"
)

func TestParsePagination(t *testing.T) {

	// case 1
	linkHeader := "linkHeader"
	_, err := parsePagination(linkHeader)
	a := assert.New(t)
	a.NotNil(err)

	// case 2
	linkHeader = ""
	_, err = parsePagination(linkHeader)
	a.Nil(err)

	linkHeader = "<https://fafafa.myshopline.com/admin/openapi/v33322/products/products.json?limit=1&page_info=eyJzaW5jZUlkIjoiMTYwNTc1OTAxNTM4OTA4Mjk1MjExMTI3ODgiLCJkaXJlY3Rpb24iOiJuZXh0IiwibGltaXQiOjF9>; rel=\"next\",<https://raoruouor.myshopline.com/admin/openapi/fajlfja/products/products.json?limit=1&page_info=eyJzaW5jZUlkIjoiMTYwNTc2NjAxNzI1MjczOTI4MDEwOTI3ODgiLCJkaXJlY3Rpb24iOiJwcmV2IiwibGltaXQiOjF9>; rel=\"previous\""
	pagination, err := parsePagination(linkHeader)
	a.Nil(err)
	a.NotNil(pagination)

	a.NotNil(pagination.Previous)
	a.Equal(pagination.Previous.PageInfo, "eyJzaW5jZUlkIjoiMTYwNTc2NjAxNzI1MjczOTI4MDEwOTI3ODgiLCJkaXJlY3Rpb24iOiJwcmV2IiwibGltaXQiOjF9")
	a.Equal(pagination.Previous.Limit, 1)

	a.NotNil(pagination.Next)
	a.NotEmpty(pagination.Next.PageInfo)
	a.Equal(pagination.Next.PageInfo, "eyJzaW5jZUlkIjoiMTYwNTc1OTAxNTM4OTA4Mjk1MjExMTI3ODgiLCJkaXJlY3Rpb24iOiJuZXh0IiwibGltaXQiOjF9")
	a.Equal(pagination.Next.Limit, 1)

	linkHeader = "<https://fafafa.myshopline.com/admin/openapi/v33322/products/products.json?limit=1&page_info=eyJzaW5jZUlkIjoiMTYwNTc1OTAxNTM4OTA4Mjk1MjExMTI3ODgiLCJkaXJlY3Rpb24iOiJuZXh0IiwibGltaXQiOjF9>; rel=\"next\""
	pagination, err = parsePagination(linkHeader)
	a.Nil(err)
	a.NotNil(pagination)
	a.NotNil(pagination.Next)
	a.NotEmpty(pagination.Next.PageInfo)
	a.Equal(pagination.Next.Limit, 1)

}
