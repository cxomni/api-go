package gobilla

import (
	"testing"
	"time"

	"github.com/usabilla/gobilla/test"
)

var (
	date = time.Date(2015, time.February, 10, 23, 0, 0, 0, time.UTC)
)

func Test_GetRFC1123GMT(t *testing.T) {
	spec := test.Spec(t)
	rfcDate := getRFC1123GMT(date)
	expected := "Tue, 10 Feb 2015 23:00:00 GMT"
	spec.Expect(rfcDate).ToEqual(expected)
}

func Test_GetShortDate(t *testing.T) {
	spec := test.Spec(t)
	shortDate := getShortDate(date)
	expected := "20150210"
	spec.Expect(shortDate).ToEqual(expected)
}

func Test_GetShortDateTime(t *testing.T) {
	spec := test.Spec(t)
	shortDateTime := getShortDateTime(date)
	expected := "20150210T230000Z"
	spec.Expect(shortDateTime).ToEqual(expected)
}