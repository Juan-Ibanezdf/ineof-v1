package tests

import (
	"database/sql"
	"time"
)

// Converte string para sql.NullString (evita NULL no banco)
func sqlNullString(s string) sql.NullString {
	return sql.NullString{String: s, Valid: s != ""}
}

// Converte time.Time para sql.NullTime (evita NULL no banco)
func sqlNullTime(t time.Time) sql.NullTime {
	return sql.NullTime{Time: t, Valid: !t.IsZero()}
}

// Converte float64 para sql.NullFloat64 (evita NULL no banco)
func sqlNullFloat64(f float64) sql.NullFloat64 {
	return sql.NullFloat64{Float64: f, Valid: f != 0}
}

// Converte int32 para sql.NullInt32 (evita NULL no banco)
func sqlNullInt32(i int32) sql.NullInt32 {
	return sql.NullInt32{Int32: i, Valid: i != 0}
}
