package dbsql

import (
	"database/sql"
	"time"
)

// HandleNullString:return NULL if string == ""
func HandleNullString(str string) sql.NullString {
	if str != "" {
		return sql.NullString{String: str, Valid: true}
	}

	return sql.NullString{Valid: false}
}

// HandleNullSmallInt: return NULL if int16 == 0
func HandleNullSmallInt(i int16) sql.NullInt16 {
	if i != 0 {
		return sql.NullInt16{Int16: i, Valid: true}
	}

	return sql.NullInt16{Valid: false}
}

// HandleNullInteger: return NULL if int32 == 0
func HandleNullInteger(i int32) sql.NullInt32 {
	if i != 0 {
		return sql.NullInt32{Int32: i, Valid: true}
	}

	return sql.NullInt32{Valid: false}
}

// HandleNullBigInt: return NULL if int64 = 0
func HandleNullBigInt(i int64) sql.NullInt64 {
	if i != 0 {
		return sql.NullInt64{Int64: i, Valid: true}
	}

	return sql.NullInt64{Valid: false}
}

// HandleNullTimestamp: return NULL if time IsZero()
func HandleNullTimestamp(t time.Time) sql.NullTime {
	if !t.IsZero() {
		return sql.NullTime{Time: t, Valid: true}
	}

	return sql.NullTime{Valid: false}
}

// HandleNullBool: return NULL if boolean == false
func HandleNullBool(boolean bool) sql.NullBool {
	if boolean {
		return sql.NullBool{Bool: boolean, Valid: true}
	}

	return sql.NullBool{Valid: false}
}
