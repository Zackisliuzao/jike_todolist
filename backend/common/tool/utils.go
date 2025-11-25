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

	parsedDate, err := time.Parse("2006-01-02", dateStr)
	if err != nil {
		return sql.NullTime{Valid: false}, err
	}

	return sql.NullTime{Time: parsedDate, Valid: true}, nil
}
