package tool

import (
	"database/sql"
	"time"
)

// formatNullableTime 格式化可为空的时间字段
func FormatNullableTime(t sql.NullTime) string {
	if t.Valid {
		return t.Time.Format(time.RFC3339)
	}
	return ""
}

// formatNullableDate 格式化可为空的日期字段
func FormatNullableDate(t sql.NullTime) string {
	if t.Valid {
		return t.Time.Format("2006-01-02")
	}
	return ""
}

// parseNullableDate 解析日期字符串为可为空的时间字段
func ParseNullableDate(dateStr string) (sql.NullTime, error) {
	if dateStr == "" {
		return sql.NullTime{Valid: false}, nil
	}

	// 支持多种日期格式
	dateFormats := []string{
		"2006-01-02",                 // YYYY-MM-DD
		"2006-01-02 15:04:05",       // YYYY-MM-DD HH:MM:SS
		time.RFC3339,                // RFC3339 format (with timezone)
		"2006-01-02T15:04:05Z07:00", // ISO8601 format with timezone
		"2006-01-02T15:04:05",      // ISO8601 format without timezone
	}

	for _, format := range dateFormats {
		parsedDate, err := time.Parse(format, dateStr)
		if err == nil {
			return sql.NullTime{Time: parsedDate, Valid: true}, nil
		}
	}

	return sql.NullTime{Valid: false}, nil
}
