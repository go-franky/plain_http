package api

import (
	"time"

	"github.com/99designs/gqlgen/graphql"
)

// MarshalTimestamp redefines the scalar timestamp to be formatted properly
func MarshalTimestamp(t time.Time) graphql.Marshaler {
	return graphql.MarshalTime(t)
	// return graphql.WriterFunc(func(w io.Writer) {
	// io.WriteString(w, strconv.Quote(t.UTC().Format(time.RFC3339)))
	// })
}

// UnmarshalTimestamp does the same thing for the unmarshaler
func UnmarshalTimestamp(v interface{}) (time.Time, error) {
	// if tmpStr, ok := v.(string); ok {
	// return time.Parse(time.RFC3339, tmpStr)
	// }
	// return time.Time{}, errors.New("time should be RFC3339 formatted string")
	return graphql.UnmarshalTime(v)
}
