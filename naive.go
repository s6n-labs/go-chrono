package chrono

import (
	"time"

	"github.com/s6n-labs/go-chrono/internal"
	"github.com/s6n-labs/go-chrono/opt"
	"github.com/s6n-labs/go-chrono/tup"
)

type NaiveDate struct {
	ymdf internal.DateImpl
}

func (d NaiveDate) of() internal.OF {
	return internal.OFFromDateImpl(d.ymdf)
}

func (d NaiveDate) mdf() internal.MDF {
	return d.of().ToMDF()
}

func (d NaiveDate) Year() int32 {
	return int32(d.ymdf >> 13)
}

func (d NaiveDate) Month() uint32 {
	return d.mdf().Month()
}

func (d NaiveDate) Day() uint32 {
	return d.mdf().Day()
}

func fromMDF(year int32, mdf internal.MDF) opt.Option[NaiveDate] {
	if year < internal.MinYear || year > internal.MaxYear {
		return opt.None[NaiveDate]()
	}

	return opt.Map(mdf.ToOF(), func(of internal.OF) NaiveDate {
		return NaiveDate{
			ymdf: internal.DateImpl(uint32(year<<13) | uint32(of)),
		}
	})
}

func FromYMD(year int32, month, day uint32) opt.Option[NaiveDate] {
	flags := internal.YearFlagsFromYear(year)

	return opt.AndThen(internal.NewMDF(month, day, flags), func(mdf internal.MDF) opt.Option[NaiveDate] {
		return fromMDF(year, mdf)
	})
}

type NaiveTime struct {
	secs uint32
	frac uint32
}

func (t NaiveTime) HMS() tup.Tuple3[uint32, uint32, uint32] {
	sec := t.secs % 60
	mins := t.secs / 60
	minute := mins % 60
	hour := mins / 60

	return tup.NewTuple3(hour, minute, sec)
}

func (t NaiveTime) Hour() uint32 {
	return t.HMS().V0
}

func (t NaiveTime) Minute() uint32 {
	return t.HMS().V1
}

func (t NaiveTime) Second() uint32 {
	return t.HMS().V2
}

func (t NaiveTime) Nanosecond() uint32 {
	return t.frac
}

func FromHMSNano(hour, min, sec, nano uint32) opt.Option[NaiveTime] {
	if hour >= 24 || min >= 60 || sec >= 60 || nano >= 2_000_000_000 {
		return opt.None[NaiveTime]()
	}

	return opt.Some(NaiveTime{
		secs: hour*3600 + min*60 + sec,
		frac: nano,
	})
}

func FromHMSMicro(hour, min, sec, micro uint32) opt.Option[NaiveTime] {
	return FromHMSNano(hour, min, sec, micro*1_000)
}

func FromHMSMilli(hour, min, sec, milli uint32) opt.Option[NaiveTime] {
	return FromHMSNano(hour, min, sec, milli*1_000_000)
}

func FromHMS(hour, min, sec uint32) opt.Option[NaiveTime] {
	return FromHMSNano(hour, min, sec, 0)
}

type NaiveDateTime struct {
	NaiveDate
	NaiveTime
}

func (d NaiveDateTime) AndLocalTimeZone(timezone TimeZone) {

}

func NewNaiveDateTime(date NaiveDate, time NaiveTime) NaiveDateTime {
	return NaiveDateTime{
		NaiveDate: date,
		NaiveTime: time,
	}
}

func FromStdTime(tm time.Time) opt.Option[NaiveDateTime] {
	return opt.Map(
		opt.Zip(
			FromYMD(int32(tm.Year()), uint32(tm.Month()), uint32(tm.Day())),
			FromHMSNano(uint32(tm.Hour()), uint32(tm.Minute()), uint32(tm.Second()), uint32(tm.Nanosecond())),
		),
		func(tuple tup.Tuple[NaiveDate, NaiveTime]) NaiveDateTime {
			return NewNaiveDateTime(tuple.Explode())
		},
	)
}
