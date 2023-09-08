package chrono

type DateTime[TZ TimeZone] struct {
	datetime NaiveDateTime
	offset   Offset
}

func DateTimeFromNaiveUTCAndOffset[TZ TimeZone](datetime NaiveDateTime, offset Offset) DateTime[TZ] {
	return DateTime[TZ]{
		datetime: datetime,
		offset:   offset,
	}
}
