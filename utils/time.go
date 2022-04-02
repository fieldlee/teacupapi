package utils

import (
	"errors"
	"fmt"
	"strconv"
	"strings"
	"time"
)

const (
	TimeBarFormat          = "2006-01-02 15:04:05"
	TimeBarFormatPM        = "2006-01-02 15:04:05 PM"
	TimeFormatHMS          = "20060102150405"
	TimeUnderlineYearMonth = "2006_01"
	TimeBarYYMMDD          = "2006-01-02"
	TimeHHMMSS             = "15:04:05"
	TimeYYMMDD             = "20060102"
	TimeYMDHM              = "200601021504"
	TimeBEIJINGFormat      = "2006-01-02 15:04:05 +08:00"
	TimeGDFormat           = "01/02/2006 15:04:05"
	TimeTFormat            = "2006-01-02T15:04:05"
	TimeTBjFormat          = "2006-01-02T15:04:05+08:00"

	Minute   = 60
	HourVal  = Minute * 60
	DayVal   = HourVal * 24
	MonthVal = DayVal * 30
	YearVal  = MonthVal * 365

	BeiJinAreaTime = "Asia/Shanghai"
)

func GetBjTimeLoc() *time.Location {
	// 获取北京时间, 在 windows系统上 time.LoadLocation 会加载失败, 最好的办法是用 time.FixedZone
	var bjLoc *time.Location
	var err error
	bjLoc, err = time.LoadLocation(BeiJinAreaTime)
	if err != nil {
		bjLoc = time.FixedZone("CST", 8*3600)
	}

	return bjLoc
}

func GetBjNowTime() time.Time {
	// 获取北京时间, 在 windows系统上 time.LoadLocation 会加载失败, 最好的办法是用 time.FixedZone
	var bjLoc *time.Location
	var err error
	bjLoc, err = time.LoadLocation(BeiJinAreaTime)
	if err != nil {
		bjLoc = time.FixedZone("CST", 8*3600)
	}

	return time.Now().In(bjLoc)
}

// 将北京时间 2006-01-02 15:04:05 类型的时间转换为时间
func BjTBarFmtTime(timeStr string) (time.Time, error) {
	if timeStr == "" {
		return time.Time{}, errors.New("time is empty")
	}

	bjTimeLoc := GetBjTimeLoc()
	return time.ParseInLocation(TimeBarFormat, timeStr, bjTimeLoc)
}

// 将时间戳转换为北京时间
func FmtUnixToBjTime(timestamp int64) time.Time {
	bjTimeLoc := GetBjTimeLoc()

	utcTime := time.Unix(timestamp, 0)
	return utcTime.In(bjTimeLoc)
}

// 将 2019-08-15T16:00:00+08:00 类型的时间数据转化为多少小时或分钟前
func GetTimeInterval(timeStr string) string {
	if timeStr == "" {
		return ""
	}

	bjTime, err := time.ParseInLocation(TimeTBjFormat, timeStr, GetBjTimeLoc())
	if err != nil {
		return "30分钟前"
	}
	//fmt.Println("bjTime: ", bjTime.Format(TimeBarFormat))

	interval := GetBjNowTime().Unix() - bjTime.Unix()
	if interval < 60 {
		return "刚刚"
	}

	if interval/Minute > 0 && interval/Minute < Minute {
		return fmt.Sprintf("%v分钟前", interval/(Minute))
	} else if interval/HourVal > 0 && interval/HourVal < 24 {
		return fmt.Sprintf("%v小时前", interval/HourVal)
	} else if interval/DayVal > 0 && interval/DayVal < 30 {
		return fmt.Sprintf("%v天前", interval/DayVal)
	} else if interval/MonthVal > 0 && interval/MonthVal < 12 {
		return fmt.Sprintf("%v月前", interval/MonthVal)
	} else if interval/YearVal > 0 {
		return fmt.Sprintf("%v年前", interval/YearVal)
	}

	return "刚刚"
}

// 秒
func GetTwoTimesInterval(timeStart, timeEnd string) int64 {
	if timeStart == "" || timeEnd == "" {
		return 0
	}

	bjStartTime, err := time.ParseInLocation(TimeTBjFormat, timeStart, GetBjTimeLoc())
	if err != nil {
		return 0
	}

	bjEndTime, err := time.ParseInLocation(TimeTBjFormat, timeEnd, GetBjTimeLoc())
	if err != nil {
		return 0
	}

	return bjEndTime.Unix() - bjStartTime.Unix()

}

// GetESTimeFomat return 2019-01-14T19:00:33+08:00
func GetESTimeFomat(timestr string) string {
	return fmt.Sprintf("%s+08:00", strings.Replace(strings.TrimSpace(timestr), " ", "T", -1))
}

func StrToTime(value string) time.Time {
	if value == "" {
		return time.Time{}
	}
	layouts := []string{
		"2006-01-02 15:04:05 -0700 MST",
		"2006-01-02 15:04:05 -0700",
		"2006-01-02 15:04:05",
		"2006/01/02 15:04:05 -0700 MST",
		"2006/01/02 15:04:05 -0700",
		"2006/01/02 15:04:05",
		"2006-01-02 -0700 MST",
		"2006-01-02 -0700",
		"2006-01-02",
		"2006/01/02 -0700 MST",
		"2006/01/02 -0700",
		"2006/01/02",
		"2006-01-02 15:04:05 -0700 -0700",
		"2006/01/02 15:04:05 -0700 -0700",
		"2006-01-02 -0700 -0700",
		"2006/01/02 -0700 -0700",
		time.ANSIC,
		time.UnixDate,
		time.RubyDate,
		time.RFC822,
		time.RFC822Z,
		time.RFC850,
		time.RFC1123,
		time.RFC1123Z,
		time.RFC3339,
		time.RFC3339Nano,
		time.Kitchen,
		time.Stamp,
		time.StampMilli,
		time.StampMicro,
		time.StampNano,
	}

	var t time.Time
	var err error
	for _, layout := range layouts {
		t, err = time.Parse(layout, value)
		if err == nil {
			return t
		}
	}

	return t
}

// 获取北京时间区的地理位置
func GetLoctionBJ() *time.Location {
	var beiJinLocation *time.Location
	var err error
	beiJinLocation, err = time.LoadLocation("Asia/Shanghai")
	if err != nil {
		beiJinLocation = time.FixedZone("CST", 8*3600)
	}
	return beiJinLocation
}

// BJNowTime 北京当前时间
func BJNowTime() time.Time {
	// 获取北京时间, 在 windows系统上 time.LoadLocation 会加载失败, 最好的办法是用 time.FixedZone, es 中的时间为: "2019-03-01T21:33:18+08:00"
	var beiJinLocation *time.Location
	var err error

	beiJinLocation, err = time.LoadLocation("Asia/Shanghai")
	if err != nil {
		beiJinLocation = time.FixedZone("CST", 8*3600)
	}

	nowTime := time.Now().In(beiJinLocation)

	return nowTime
}

func CurTimeToStr() string {
	curTime := BJNowTime()
	return curTime.Format(TimeBarFormat)
}

func GetFirstDateOfWeek() (weekMonday string) {
	now := time.Now()
	offset := int(time.Monday - now.Weekday())
	if offset > 0 {
		offset = -6
	}
	weekStartDate := time.Date(now.Year(), now.Month(), now.Day(), 0, 0, 0, 0, time.Local).AddDate(0, 0, offset)
	weekMonday = weekStartDate.Format(TimeBarFormat)
	return
}

func GetMinNumberForAward(timeCur time.Time) int {
	curmins := timeCur.Hour()*60 + timeCur.Minute()
	if curmins == 0 {
		return curmins
	}
	return curmins - 1
}

func GetMinNumber(timeCur time.Time) int {
	curmins := timeCur.Hour()*60 + timeCur.Minute()
	return curmins
}
func GetNextPeriod(period string) (string, error) {
	if len(period) != 12 {
		return "", errors.New("period 格式不对")
	}
	strDate := period[:8]
	strmins := period[8:]

	curDate, err := time.Parse("20060102", strDate)
	if err != nil {
		return "", err
	}
	mins, err := strconv.Atoi(strings.TrimLeft(strmins, "0"))
	if err != nil {
		return "", err
	}
	curTime := curDate.Add(time.Duration(mins+1) * time.Minute)
	periodDate := curTime.Format(TimeYYMMDD)
	periodMins := GetMinNumber(curTime)
	return fmt.Sprintf("%s%04d", periodDate, periodMins), nil
}
