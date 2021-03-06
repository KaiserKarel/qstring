package qstring

import (
	"strings"
)

// MapBuilders are used to construct queries which map against an objects fields.
//
// Prefer not using builders directly, instead using the constructors.
type MapBuilder struct {
	builder *strings.Builder
}

func (kv *MapBuilder) Has(key, value string) *Builder {
	kv.builder.WriteString(" has {key='")
	kv.builder.WriteString(escape(key))
	kv.builder.WriteString("' and value='")
	kv.builder.WriteString(escape(value))
	kv.builder.WriteString("'}")
	return &Builder{builder: kv.builder}
}
