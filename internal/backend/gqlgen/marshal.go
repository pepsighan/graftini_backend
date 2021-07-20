package gqlgen

import (
	"io"

	"github.com/99designs/gqlgen/graphql"
	"github.com/google/uuid"
	"github.com/pepsighan/graftini_backend/internal/pkg/logger"
)

func MarshalUUID(b uuid.UUID) graphql.Marshaler {
	return graphql.WriterFunc(func(w io.Writer) {
		w.Write([]byte("\""))
		w.Write([]byte(b.String()))
		w.Write([]byte("\""))
	})
}

func UnmarshalUUID(v interface{}) (uuid.UUID, error) {
	switch v := v.(type) {
	case string:
		return uuid.Parse(v)
	default:
		return uuid.UUID{}, logger.Errorf("%T is not a bool", v)
	}
}
