package timestamp

import (
	"fmt"
	"time"
)

const TimeLayoutFull = "2006-01-02 15:04:05"

// StampBeforeDay 获取前n天的开始时间和结束时间
func StampBeforeDay(n int, zone *time.Location) (int, int, bool) {
	var startTime, endTime int
	now := time.Now()
	beForeTime := now.Local().AddDate(0, 0, n)
	strStart, err := time.ParseInLocation(TimeLayoutFull, fmt.Sprintf("%v-%v-%v 00:00:00", beForeTime.Year(), beForeTime.String()[5:7], beForeTime.Day()), zone)
	if err != nil {
		return startTime, endTime, false
	}
	startTime = int(strStart.Unix())
	strEnd, err := time.ParseInLocation(TimeLayoutFull, fmt.Sprintf("%v-%v-%v 23:59:59", beForeTime.Year(), beForeTime.String()[5:7], beForeTime.Day()), zone)
	if err != nil {
		return startTime, endTime, false
	}
	endTime = int(strEnd.Unix())
	return startTime, endTime, true
}
