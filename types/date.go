package types

import "time"

type DateRange struct {
	Start time.Time `form:"start" time_format:"2006-01-02" binding:"required" time_utc:"1"`
	End   time.Time `form:"end" time_format:"2006-01-02" binding:"required" time_utc:"1"`
	Type  string    `form:"type" binding:"required"`
}
