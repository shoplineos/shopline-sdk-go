package client

import (
	"fmt"
	"net/url"
	"strings"
	"time"
)

// BuildTimestamp Build timestamp
func BuildTimestamp() string {
	timestamp := time.Now().UnixMilli()
	timestampStr := fmt.Sprintf("%d", timestamp)
	return timestampStr
}

// GetStoreFullName Return the full store name, including .myshopline.com
func GetStoreFullName(name string) string {
	name = strings.TrimSpace(name)
	name = strings.Trim(name, ".")
	if strings.Contains(name, "myshopline.com") {
		return name
	}
	return name + ".myshopline.com"
}

// GetStoreShortName Return the short store name, excluding .myshopline.com
func GetStoreShortName(name string) string {
	return strings.Replace(GetStoreFullName(name), ".myshopline.com", "", -1)
}

// GetStoreBaseUrl Return the Store's base url.
// eg: https://storeName.myshopline.com
func GetStoreBaseUrl(name string) string {
	name = GetStoreFullName(name)
	return fmt.Sprintf("https://%s", name)
}

type OnlyDate struct {
	time.Time
}

func (c *OnlyDate) UnmarshalJSON(b []byte) error {
	value := strings.Trim(string(b), `"`)
	if value == "" || value == "null" {
		*c = OnlyDate{time.Time{}}
		return nil
	}

	t, err := time.Parse("2006-01-02", value)
	if err != nil {
		return err
	}
	*c = OnlyDate{t}
	return nil
}

func (c *OnlyDate) MarshalJSON() ([]byte, error) {
	return []byte(c.String()), nil
}

func (c *OnlyDate) EncodeValues(key string, v *url.Values) error {
	v.Add(key, c.String())
	return nil
}

func (c *OnlyDate) String() string {
	return `"` + c.Format("2006-01-02") + `"`
}

func TimePtr(v time.Time) *time.Time {
	return &v
}
