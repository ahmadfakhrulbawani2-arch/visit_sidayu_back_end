package helpers

import (
	"fmt"
	"reflect"
	"time"

	"github.com/gosimple/slug"
)

func GenerateSlugWithTimestamp(title string, item interface{}) {
	baseSlug := slug.Make(title)
	timeCode := time.Now().Format("15-04-05-02-01-2006")
	uniqueSlug := fmt.Sprintf("%s-%s", baseSlug, timeCode)

	v := reflect.ValueOf(item)

	if v.Kind() == reflect.Pointer {
		v = v.Elem()
	}
	field := v.FieldByName("Slug")
	if field.IsValid() && field.CanSet() && field.Kind() == reflect.String {
		field.SetString(uniqueSlug)
	}
}
