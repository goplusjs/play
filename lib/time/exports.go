package time

import (
	"reflect"
	"time"

	qspec "github.com/qiniu/goplus/exec.spec"
	"github.com/qiniu/goplus/gop"
)

// func time.After(d time.Duration) <-chan time.Time
func execAfter(_ int, p *gop.Context) {
	args := p.GetArgs(1)
	ret := time.After(time.Duration(args[0].(int64)))
	p.Ret(1, ret)
}

// func time.AfterFunc(d time.Duration, f func()) *time.Timer
func execAfterFunc(_ int, p *gop.Context) {
	args := p.GetArgs(2)
	ret := time.AfterFunc(time.Duration(args[0].(int64)), args[1].(func()))
	p.Ret(2, ret)
}

// func time.Date(year int, month time.Month, day int, hour int, min int, sec int, nsec int, loc *time.Location) time.Time
func execDate(_ int, p *gop.Context) {
	args := p.GetArgs(8)
	ret := time.Date(args[0].(int), time.Month(args[1].(int)), args[2].(int), args[3].(int), args[4].(int), args[5].(int), args[6].(int), args[7].(*time.Location))
	p.Ret(8, ret)
}

// func (time.Duration).Hours() float64
func execmDurationHours(_ int, p *gop.Context) {
	args := p.GetArgs(1)
	ret := args[0].(time.Duration).Hours()
	p.Ret(1, ret)
}

// func (time.Duration).Microseconds() int64
func execmDurationMicroseconds(_ int, p *gop.Context) {
	args := p.GetArgs(1)
	ret := args[0].(time.Duration).Microseconds()
	p.Ret(1, ret)
}

// func (time.Duration).Milliseconds() int64
func execmDurationMilliseconds(_ int, p *gop.Context) {
	args := p.GetArgs(1)
	ret := args[0].(time.Duration).Milliseconds()
	p.Ret(1, ret)
}

// func (time.Duration).Minutes() float64
func execmDurationMinutes(_ int, p *gop.Context) {
	args := p.GetArgs(1)
	ret := args[0].(time.Duration).Minutes()
	p.Ret(1, ret)
}

// func (time.Duration).Nanoseconds() int64
func execmDurationNanoseconds(_ int, p *gop.Context) {
	args := p.GetArgs(1)
	ret := args[0].(time.Duration).Nanoseconds()
	p.Ret(1, ret)
}

// func (time.Duration).Round(m time.Duration) time.Duration
func execmDurationRound(_ int, p *gop.Context) {
	args := p.GetArgs(2)
	ret := args[0].(time.Duration).Round(time.Duration(args[1].(int64)))
	p.Ret(2, ret)
}

// func (time.Duration).Seconds() float64
func execmDurationSeconds(_ int, p *gop.Context) {
	args := p.GetArgs(1)
	ret := args[0].(time.Duration).Seconds()
	p.Ret(1, ret)
}

// func (time.Duration).String() string
func execmDurationString(_ int, p *gop.Context) {
	args := p.GetArgs(1)
	ret := args[0].(time.Duration).String()
	p.Ret(1, ret)
}

// func (time.Duration).Truncate(m time.Duration) time.Duration
func execmDurationTruncate(_ int, p *gop.Context) {
	args := p.GetArgs(2)
	ret := args[0].(time.Duration).Truncate(time.Duration(args[1].(int64)))
	p.Ret(2, ret)
}

// func time.FixedZone(name string, offset int) *time.Location
func execFixedZone(_ int, p *gop.Context) {
	args := p.GetArgs(2)
	ret := time.FixedZone(args[0].(string), args[1].(int))
	p.Ret(2, ret)
}

// func time.LoadLocation(name string) (*time.Location, error)
func execLoadLocation(_ int, p *gop.Context) {
	args := p.GetArgs(1)
	ret, ret1 := time.LoadLocation(args[0].(string))
	p.Ret(1, ret, ret1)
}

// func time.LoadLocationFromTZData(name string, data []byte) (*time.Location, error)
func execLoadLocationFromTZData(_ int, p *gop.Context) {
	args := p.GetArgs(2)
	ret, ret1 := time.LoadLocationFromTZData(args[0].(string), args[1].([]byte))
	p.Ret(2, ret, ret1)
}

// func (*time.Location).String() string
func execmLocationString(_ int, p *gop.Context) {
	args := p.GetArgs(1)
	ret := args[0].(*time.Location).String()
	p.Ret(1, ret)
}

// func (time.Month).String() string
func execmMonthString(_ int, p *gop.Context) {
	args := p.GetArgs(1)
	ret := args[0].(time.Month).String()
	p.Ret(1, ret)
}

// func time.NewTicker(d time.Duration) *time.Ticker
func execNewTicker(_ int, p *gop.Context) {
	args := p.GetArgs(1)
	ret := time.NewTicker(time.Duration(args[0].(int64)))
	p.Ret(1, ret)
}

// func time.NewTimer(d time.Duration) *time.Timer
func execNewTimer(_ int, p *gop.Context) {
	args := p.GetArgs(1)
	ret := time.NewTimer(time.Duration(args[0].(int64)))
	p.Ret(1, ret)
}

// func time.Now() time.Time
func execNow(_ int, p *gop.Context) {
	ret := time.Now()
	p.Ret(0, ret)
}

// func time.Parse(layout string, value string) (time.Time, error)
func execParse(_ int, p *gop.Context) {
	args := p.GetArgs(2)
	ret, ret1 := time.Parse(args[0].(string), args[1].(string))
	p.Ret(2, ret, ret1)
}

// func time.ParseDuration(s string) (time.Duration, error)
func execParseDuration(_ int, p *gop.Context) {
	args := p.GetArgs(1)
	ret, ret1 := time.ParseDuration(args[0].(string))
	p.Ret(1, ret, ret1)
}

// func (*time.ParseError).Error() string
func execmParseErrorError(_ int, p *gop.Context) {
	args := p.GetArgs(1)
	ret := args[0].(*time.ParseError).Error()
	p.Ret(1, ret)
}

// func time.ParseInLocation(layout string, value string, loc *time.Location) (time.Time, error)
func execParseInLocation(_ int, p *gop.Context) {
	args := p.GetArgs(3)
	ret, ret1 := time.ParseInLocation(args[0].(string), args[1].(string), args[2].(*time.Location))
	p.Ret(3, ret, ret1)
}

// func time.Since(t time.Time) time.Duration
func execSince(_ int, p *gop.Context) {
	args := p.GetArgs(1)
	ret := time.Since(args[0].(time.Time))
	p.Ret(1, ret)
}

// func time.Sleep(d time.Duration)
func execSleep(_ int, p *gop.Context) {
	args := p.GetArgs(1)
	time.Sleep(time.Duration(args[0].(int64)))
}

// func time.Tick(d time.Duration) <-chan time.Time
func execTick(_ int, p *gop.Context) {
	args := p.GetArgs(1)
	ret := time.Tick(time.Duration(args[0].(int64)))
	p.Ret(1, ret)
}

// func (*time.Ticker).Stop()
func execmTickerStop(_ int, p *gop.Context) {
	args := p.GetArgs(1)
	args[0].(*time.Ticker).Stop()
}

// func (time.Time).Add(d time.Duration) time.Time
func execmTimeAdd(_ int, p *gop.Context) {
	args := p.GetArgs(2)
	ret := args[0].(time.Time).Add(time.Duration(args[1].(int64)))
	p.Ret(2, ret)
}

// func (time.Time).AddDate(years int, months int, days int) time.Time
func execmTimeAddDate(_ int, p *gop.Context) {
	args := p.GetArgs(4)
	ret := args[0].(time.Time).AddDate(args[1].(int), args[2].(int), args[3].(int))
	p.Ret(4, ret)
}

// func (time.Time).After(u time.Time) bool
func execmTimeAfter(_ int, p *gop.Context) {
	args := p.GetArgs(2)
	ret := args[0].(time.Time).After(args[1].(time.Time))
	p.Ret(2, ret)
}

// func (time.Time).AppendFormat(b []byte, layout string) []byte
func execmTimeAppendFormat(_ int, p *gop.Context) {
	args := p.GetArgs(3)
	ret := args[0].(time.Time).AppendFormat(args[1].([]byte), args[2].(string))
	p.Ret(3, ret)
}

// func (time.Time).Before(u time.Time) bool
func execmTimeBefore(_ int, p *gop.Context) {
	args := p.GetArgs(2)
	ret := args[0].(time.Time).Before(args[1].(time.Time))
	p.Ret(2, ret)
}

// func (time.Time).Clock() (hour int, min int, sec int)
func execmTimeClock(_ int, p *gop.Context) {
	args := p.GetArgs(1)
	ret, ret1, ret2 := args[0].(time.Time).Clock()
	p.Ret(1, ret, ret1, ret2)
}

// func (time.Time).Date() (year int, month time.Month, day int)
func execmTimeDate(_ int, p *gop.Context) {
	args := p.GetArgs(1)
	ret, ret1, ret2 := args[0].(time.Time).Date()
	p.Ret(1, ret, ret1, ret2)
}

// func (time.Time).Day() int
func execmTimeDay(_ int, p *gop.Context) {
	args := p.GetArgs(1)
	ret := args[0].(time.Time).Day()
	p.Ret(1, ret)
}

// func (time.Time).Equal(u time.Time) bool
func execmTimeEqual(_ int, p *gop.Context) {
	args := p.GetArgs(2)
	ret := args[0].(time.Time).Equal(args[1].(time.Time))
	p.Ret(2, ret)
}

// func (time.Time).Format(layout string) string
func execmTimeFormat(_ int, p *gop.Context) {
	args := p.GetArgs(2)
	ret := args[0].(time.Time).Format(args[1].(string))
	p.Ret(2, ret)
}

// func (*time.Time).GobDecode(data []byte) error
func execmTimeGobDecode(_ int, p *gop.Context) {
	args := p.GetArgs(2)
	ret := args[0].(*time.Time).GobDecode(args[1].([]byte))
	p.Ret(2, ret)
}

// func (time.Time).GobEncode() ([]byte, error)
func execmTimeGobEncode(_ int, p *gop.Context) {
	args := p.GetArgs(1)
	ret, ret1 := args[0].(time.Time).GobEncode()
	p.Ret(1, ret, ret1)
}

// func (time.Time).Hour() int
func execmTimeHour(_ int, p *gop.Context) {
	args := p.GetArgs(1)
	ret := args[0].(time.Time).Hour()
	p.Ret(1, ret)
}

// func (time.Time).ISOWeek() (year int, week int)
func execmTimeISOWeek(_ int, p *gop.Context) {
	args := p.GetArgs(1)
	ret, ret1 := args[0].(time.Time).ISOWeek()
	p.Ret(1, ret, ret1)
}

// func (time.Time).In(loc *time.Location) time.Time
func execmTimeIn(_ int, p *gop.Context) {
	args := p.GetArgs(2)
	ret := args[0].(time.Time).In(args[1].(*time.Location))
	p.Ret(2, ret)
}

// func (time.Time).IsZero() bool
func execmTimeIsZero(_ int, p *gop.Context) {
	args := p.GetArgs(1)
	ret := args[0].(time.Time).IsZero()
	p.Ret(1, ret)
}

// func (time.Time).Local() time.Time
func execmTimeLocal(_ int, p *gop.Context) {
	args := p.GetArgs(1)
	ret := args[0].(time.Time).Local()
	p.Ret(1, ret)
}

// func (time.Time).Location() *time.Location
func execmTimeLocation(_ int, p *gop.Context) {
	args := p.GetArgs(1)
	ret := args[0].(time.Time).Location()
	p.Ret(1, ret)
}

// func (time.Time).MarshalBinary() ([]byte, error)
func execmTimeMarshalBinary(_ int, p *gop.Context) {
	args := p.GetArgs(1)
	ret, ret1 := args[0].(time.Time).MarshalBinary()
	p.Ret(1, ret, ret1)
}

// func (time.Time).MarshalJSON() ([]byte, error)
func execmTimeMarshalJSON(_ int, p *gop.Context) {
	args := p.GetArgs(1)
	ret, ret1 := args[0].(time.Time).MarshalJSON()
	p.Ret(1, ret, ret1)
}

// func (time.Time).MarshalText() ([]byte, error)
func execmTimeMarshalText(_ int, p *gop.Context) {
	args := p.GetArgs(1)
	ret, ret1 := args[0].(time.Time).MarshalText()
	p.Ret(1, ret, ret1)
}

// func (time.Time).Minute() int
func execmTimeMinute(_ int, p *gop.Context) {
	args := p.GetArgs(1)
	ret := args[0].(time.Time).Minute()
	p.Ret(1, ret)
}

// func (time.Time).Month() time.Month
func execmTimeMonth(_ int, p *gop.Context) {
	args := p.GetArgs(1)
	ret := args[0].(time.Time).Month()
	p.Ret(1, ret)
}

// func (time.Time).Nanosecond() int
func execmTimeNanosecond(_ int, p *gop.Context) {
	args := p.GetArgs(1)
	ret := args[0].(time.Time).Nanosecond()
	p.Ret(1, ret)
}

// func (time.Time).Round(d time.Duration) time.Time
func execmTimeRound(_ int, p *gop.Context) {
	args := p.GetArgs(2)
	ret := args[0].(time.Time).Round(time.Duration(args[1].(int64)))
	p.Ret(2, ret)
}

// func (time.Time).Second() int
func execmTimeSecond(_ int, p *gop.Context) {
	args := p.GetArgs(1)
	ret := args[0].(time.Time).Second()
	p.Ret(1, ret)
}

// func (time.Time).String() string
func execmTimeString(_ int, p *gop.Context) {
	args := p.GetArgs(1)
	ret := args[0].(time.Time).String()
	p.Ret(1, ret)
}

// func (time.Time).Sub(u time.Time) time.Duration
func execmTimeSub(_ int, p *gop.Context) {
	args := p.GetArgs(2)
	ret := args[0].(time.Time).Sub(args[1].(time.Time))
	p.Ret(2, ret)
}

// func (time.Time).Truncate(d time.Duration) time.Time
func execmTimeTruncate(_ int, p *gop.Context) {
	args := p.GetArgs(2)
	ret := args[0].(time.Time).Truncate(time.Duration(args[1].(int64)))
	p.Ret(2, ret)
}

// func (time.Time).UTC() time.Time
func execmTimeUTC(_ int, p *gop.Context) {
	args := p.GetArgs(1)
	ret := args[0].(time.Time).UTC()
	p.Ret(1, ret)
}

// func (time.Time).Unix() int64
func execmTimeUnix(_ int, p *gop.Context) {
	args := p.GetArgs(1)
	ret := args[0].(time.Time).Unix()
	p.Ret(1, ret)
}

// func (time.Time).UnixNano() int64
func execmTimeUnixNano(_ int, p *gop.Context) {
	args := p.GetArgs(1)
	ret := args[0].(time.Time).UnixNano()
	p.Ret(1, ret)
}

// func (*time.Time).UnmarshalBinary(data []byte) error
func execmTimeUnmarshalBinary(_ int, p *gop.Context) {
	args := p.GetArgs(2)
	ret := args[0].(*time.Time).UnmarshalBinary(args[1].([]byte))
	p.Ret(2, ret)
}

// func (*time.Time).UnmarshalJSON(data []byte) error
func execmTimeUnmarshalJSON(_ int, p *gop.Context) {
	args := p.GetArgs(2)
	ret := args[0].(*time.Time).UnmarshalJSON(args[1].([]byte))
	p.Ret(2, ret)
}

// func (*time.Time).UnmarshalText(data []byte) error
func execmTimeUnmarshalText(_ int, p *gop.Context) {
	args := p.GetArgs(2)
	ret := args[0].(*time.Time).UnmarshalText(args[1].([]byte))
	p.Ret(2, ret)
}

// func (time.Time).Weekday() time.Weekday
func execmTimeWeekday(_ int, p *gop.Context) {
	args := p.GetArgs(1)
	ret := args[0].(time.Time).Weekday()
	p.Ret(1, ret)
}

// func (time.Time).Year() int
func execmTimeYear(_ int, p *gop.Context) {
	args := p.GetArgs(1)
	ret := args[0].(time.Time).Year()
	p.Ret(1, ret)
}

// func (time.Time).YearDay() int
func execmTimeYearDay(_ int, p *gop.Context) {
	args := p.GetArgs(1)
	ret := args[0].(time.Time).YearDay()
	p.Ret(1, ret)
}

// func (time.Time).Zone() (name string, offset int)
func execmTimeZone(_ int, p *gop.Context) {
	args := p.GetArgs(1)
	ret, ret1 := args[0].(time.Time).Zone()
	p.Ret(1, ret, ret1)
}

// func (*time.Timer).Reset(d time.Duration) bool
func execmTimerReset(_ int, p *gop.Context) {
	args := p.GetArgs(2)
	ret := args[0].(*time.Timer).Reset(time.Duration(args[1].(int64)))
	p.Ret(2, ret)
}

// func (*time.Timer).Stop() bool
func execmTimerStop(_ int, p *gop.Context) {
	args := p.GetArgs(1)
	ret := args[0].(*time.Timer).Stop()
	p.Ret(1, ret)
}

// func time.Unix(sec int64, nsec int64) time.Time
func execUnix(_ int, p *gop.Context) {
	args := p.GetArgs(2)
	ret := time.Unix(args[0].(int64), args[1].(int64))
	p.Ret(2, ret)
}

// func time.Until(t time.Time) time.Duration
func execUntil(_ int, p *gop.Context) {
	args := p.GetArgs(1)
	ret := time.Until(args[0].(time.Time))
	p.Ret(1, ret)
}

// func (time.Weekday).String() string
func execmWeekdayString(_ int, p *gop.Context) {
	args := p.GetArgs(1)
	ret := args[0].(time.Weekday).String()
	p.Ret(1, ret)
}

// I is a Go package instance.
var I = gop.NewGoPackage("time")

func init() {
	I.RegisterConsts(
		I.Const("ANSIC", qspec.ConstBoundString, time.ANSIC),
		I.Const("April", reflect.Int, time.April),
		I.Const("August", reflect.Int, time.August),
		I.Const("December", reflect.Int, time.December),
		I.Const("February", reflect.Int, time.February),
		I.Const("Friday", reflect.Int, time.Friday),
		I.Const("Hour", qspec.Uint64, uint64(time.Hour)),
		I.Const("January", reflect.Int, time.January),
		I.Const("July", reflect.Int, time.July),
		I.Const("June", reflect.Int, time.June),
		I.Const("Kitchen", qspec.ConstBoundString, time.Kitchen),
		I.Const("March", reflect.Int, time.March),
		I.Const("May", reflect.Int, time.May),
		I.Const("Microsecond", reflect.Int64, time.Microsecond),
		I.Const("Millisecond", reflect.Int64, time.Millisecond),
		I.Const("Minute", qspec.Uint64, uint64(time.Minute)),
		I.Const("Monday", reflect.Int, time.Monday),
		I.Const("Nanosecond", reflect.Int64, time.Nanosecond),
		I.Const("November", reflect.Int, time.November),
		I.Const("October", reflect.Int, time.October),
		I.Const("RFC1123", qspec.ConstBoundString, time.RFC1123),
		I.Const("RFC1123Z", qspec.ConstBoundString, time.RFC1123Z),
		I.Const("RFC3339", qspec.ConstBoundString, time.RFC3339),
		I.Const("RFC3339Nano", qspec.ConstBoundString, time.RFC3339Nano),
		I.Const("RFC822", qspec.ConstBoundString, time.RFC822),
		I.Const("RFC822Z", qspec.ConstBoundString, time.RFC822Z),
		I.Const("RFC850", qspec.ConstBoundString, time.RFC850),
		I.Const("RubyDate", qspec.ConstBoundString, time.RubyDate),
		I.Const("Saturday", reflect.Int, time.Saturday),
		I.Const("Second", reflect.Int64, time.Second),
		I.Const("September", reflect.Int, time.September),
		I.Const("Stamp", qspec.ConstBoundString, time.Stamp),
		I.Const("StampMicro", qspec.ConstBoundString, time.StampMicro),
		I.Const("StampMilli", qspec.ConstBoundString, time.StampMilli),
		I.Const("StampNano", qspec.ConstBoundString, time.StampNano),
		I.Const("Sunday", reflect.Int, time.Sunday),
		I.Const("Thursday", reflect.Int, time.Thursday),
		I.Const("Tuesday", reflect.Int, time.Tuesday),
		I.Const("UnixDate", qspec.ConstBoundString, time.UnixDate),
		I.Const("Wednesday", reflect.Int, time.Wednesday),
	)
	I.RegisterVars(
		I.Var("Local", &time.Local),
		I.Var("UTC", &time.UTC),
	)
	I.RegisterTypes(
		I.Type("Duration", qspec.TyInt64),
		I.Rtype(reflect.TypeOf((*time.Location)(nil))),
		I.Type("Month", qspec.TyInt),
		I.Rtype(reflect.TypeOf((*time.ParseError)(nil))),
		I.Rtype(reflect.TypeOf((*time.Ticker)(nil))),
		I.Rtype(reflect.TypeOf((*time.Time)(nil))),
		I.Rtype(reflect.TypeOf((*time.Timer)(nil))),
		I.Type("Weekday", qspec.TyInt),
	)
	I.RegisterFuncs(
		I.Func("After", time.After, execAfter),
		I.Func("AfterFunc", time.AfterFunc, execAfterFunc),
		I.Func("Date", time.Date, execDate),
		I.Func("(Duration).Hours", (time.Duration).Hours, execmDurationHours),
		I.Func("(Duration).Microseconds", (time.Duration).Microseconds, execmDurationMicroseconds),
		I.Func("(Duration).Milliseconds", (time.Duration).Milliseconds, execmDurationMilliseconds),
		I.Func("(Duration).Minutes", (time.Duration).Minutes, execmDurationMinutes),
		I.Func("(Duration).Nanoseconds", (time.Duration).Nanoseconds, execmDurationNanoseconds),
		I.Func("(Duration).Round", (time.Duration).Round, execmDurationRound),
		I.Func("(Duration).Seconds", (time.Duration).Seconds, execmDurationSeconds),
		I.Func("(Duration).String", (time.Duration).String, execmDurationString),
		I.Func("(Duration).Truncate", (time.Duration).Truncate, execmDurationTruncate),
		I.Func("FixedZone", time.FixedZone, execFixedZone),
		I.Func("LoadLocation", time.LoadLocation, execLoadLocation),
		I.Func("LoadLocationFromTZData", time.LoadLocationFromTZData, execLoadLocationFromTZData),
		I.Func("(*Location).String", (*time.Location).String, execmLocationString),
		I.Func("(Month).String", (time.Month).String, execmMonthString),
		I.Func("NewTicker", time.NewTicker, execNewTicker),
		I.Func("NewTimer", time.NewTimer, execNewTimer),
		I.Func("Now", time.Now, execNow),
		I.Func("Parse", time.Parse, execParse),
		I.Func("ParseDuration", time.ParseDuration, execParseDuration),
		I.Func("(*ParseError).Error", (*time.ParseError).Error, execmParseErrorError),
		I.Func("ParseInLocation", time.ParseInLocation, execParseInLocation),
		I.Func("Since", time.Since, execSince),
		I.Func("Sleep", time.Sleep, execSleep),
		I.Func("Tick", time.Tick, execTick),
		I.Func("(*Ticker).Stop", (*time.Ticker).Stop, execmTickerStop),
		I.Func("(Time).Add", (time.Time).Add, execmTimeAdd),
		I.Func("(Time).AddDate", (time.Time).AddDate, execmTimeAddDate),
		I.Func("(Time).After", (time.Time).After, execmTimeAfter),
		I.Func("(Time).AppendFormat", (time.Time).AppendFormat, execmTimeAppendFormat),
		I.Func("(Time).Before", (time.Time).Before, execmTimeBefore),
		I.Func("(Time).Clock", (time.Time).Clock, execmTimeClock),
		I.Func("(Time).Date", (time.Time).Date, execmTimeDate),
		I.Func("(Time).Day", (time.Time).Day, execmTimeDay),
		I.Func("(Time).Equal", (time.Time).Equal, execmTimeEqual),
		I.Func("(Time).Format", (time.Time).Format, execmTimeFormat),
		I.Func("(*Time).GobDecode", (*time.Time).GobDecode, execmTimeGobDecode),
		I.Func("(Time).GobEncode", (time.Time).GobEncode, execmTimeGobEncode),
		I.Func("(Time).Hour", (time.Time).Hour, execmTimeHour),
		I.Func("(Time).ISOWeek", (time.Time).ISOWeek, execmTimeISOWeek),
		I.Func("(Time).In", (time.Time).In, execmTimeIn),
		I.Func("(Time).IsZero", (time.Time).IsZero, execmTimeIsZero),
		I.Func("(Time).Local", (time.Time).Local, execmTimeLocal),
		I.Func("(Time).Location", (time.Time).Location, execmTimeLocation),
		I.Func("(Time).MarshalBinary", (time.Time).MarshalBinary, execmTimeMarshalBinary),
		I.Func("(Time).MarshalJSON", (time.Time).MarshalJSON, execmTimeMarshalJSON),
		I.Func("(Time).MarshalText", (time.Time).MarshalText, execmTimeMarshalText),
		I.Func("(Time).Minute", (time.Time).Minute, execmTimeMinute),
		I.Func("(Time).Month", (time.Time).Month, execmTimeMonth),
		I.Func("(Time).Nanosecond", (time.Time).Nanosecond, execmTimeNanosecond),
		I.Func("(Time).Round", (time.Time).Round, execmTimeRound),
		I.Func("(Time).Second", (time.Time).Second, execmTimeSecond),
		I.Func("(Time).String", (time.Time).String, execmTimeString),
		I.Func("(Time).Sub", (time.Time).Sub, execmTimeSub),
		I.Func("(Time).Truncate", (time.Time).Truncate, execmTimeTruncate),
		I.Func("(Time).UTC", (time.Time).UTC, execmTimeUTC),
		I.Func("(Time).Unix", (time.Time).Unix, execmTimeUnix),
		I.Func("(Time).UnixNano", (time.Time).UnixNano, execmTimeUnixNano),
		I.Func("(*Time).UnmarshalBinary", (*time.Time).UnmarshalBinary, execmTimeUnmarshalBinary),
		I.Func("(*Time).UnmarshalJSON", (*time.Time).UnmarshalJSON, execmTimeUnmarshalJSON),
		I.Func("(*Time).UnmarshalText", (*time.Time).UnmarshalText, execmTimeUnmarshalText),
		I.Func("(Time).Weekday", (time.Time).Weekday, execmTimeWeekday),
		I.Func("(Time).Year", (time.Time).Year, execmTimeYear),
		I.Func("(Time).YearDay", (time.Time).YearDay, execmTimeYearDay),
		I.Func("(Time).Zone", (time.Time).Zone, execmTimeZone),
		I.Func("(*Timer).Reset", (*time.Timer).Reset, execmTimerReset),
		I.Func("(*Timer).Stop", (*time.Timer).Stop, execmTimerStop),
		I.Func("Unix", time.Unix, execUnix),
		I.Func("Until", time.Until, execUntil),
		I.Func("(Weekday).String", (time.Weekday).String, execmWeekdayString),
	)
}
