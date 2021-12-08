package dbsql

import (
	"database/sql"
	"time"
)

// ParseNullString: return str if str.Valid else ""
func ParseNullString(str sql.NullString) string {
	if str.Valid {
		return str.String
	}

	return ""
}

// ParseNullSmallInt: return i if i.Valid else 0
func ParseNullSmallInt(i sql.NullInt16) int16 {
	if i.Valid {
		return i.Int16
	}

	return 0
}

// ParseNullInteger: return i if i.Valid else 0
func ParseNullInteger(i sql.NullInt32) int32 {
	if i.Valid {
		return i.Int32
	}

	return 0
}

// ParseNullBigInt: return i if i.Valid else 0
func ParseNullBigInt(i sql.NullInt64) int64 {
	if i.Valid {
		return i.Int64
	}

	return 0
}

// ParseNullTimestamp: return t if t.Valid else time.Time{}
func ParseNullTimestamp(t sql.NullTime) time.Time {
	if t.Valid {
		return t.Time
	}

	return time.Time{}
}

// ParseNullBool: return i if i.Valid else false
func ParseNullBool(b sql.NullBool) bool {
	if b.Valid {
		return b.Bool
	}

	return false
}
