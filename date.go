package chrono

type Date[TZ TimeZone] struct {
	date   NaiveDate
	offset Offset
}

func DateFromUTC[TZ TimeZone](date NaiveDate, offset Offset) Date[TZ] {
	return Date[TZ]{
		date:   date,
		offset: offset,
	}
}
