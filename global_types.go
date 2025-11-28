package netsuite

import (
	"strings"

	"github.com/cydev/zero"
)

type Links []Link

func (l Links) IsEmpty() bool {
	return zero.IsZero(l)
}

type Link struct {
	Rel  string `json:"rel"`
	Href string `json:"href"`
}

func (l Link) IsZero() bool {
	return zero.IsZero(l)
}

type Fields []Field

func (f Fields) IsZero() bool {
	return zero.IsZero(f)
}

func (ff Fields) MarshalSchema() string {
	fields := []string{}
	for _, f := range ff {
		fields = append(fields, string(f))
	}
	return strings.Join(fields, ",")
}

type Field string
