package chrono

import (
	"time"

	"github.com/s6n-labs/go-chrono/opt"
)

type FixedOffset struct {
	localMinusUTC int32
}

func (o FixedOffset) Fix() FixedOffset {
	return o
}

func (o FixedOffset) LocalMinusUTC() int32 {
	return o.localMinusUTC
}

func (o FixedOffset) UTCMinusLocal() int32 {
	return -o.localMinusUTC
}

func FixedOffsetEast(secs int32) opt.Option[FixedOffset] {
	if secs <= -86_400 || secs >= 86_400 {
		return opt.None[FixedOffset]()
	}

	return opt.Some(FixedOffset{
		localMinusUTC: secs,
	})
}

func FixedOffsetWest(secs int32) opt.Option[FixedOffset] {
	if secs <= -86_400 || secs >= 86_400 {
		return opt.None[FixedOffset]()
	}

	return opt.Some(FixedOffset{
		localMinusUTC: -secs,
	})
}

type Offset interface {
	Fix() FixedOffset
}

type LocalResult[T Offset] interface {
	Single() opt.Option[T]
	Earliest() opt.Option[T]
	Latest() opt.Option[T]
}

type None[T Offset] struct{}

func (n None[T]) Single() opt.Option[T] {
	return opt.None[T]()
}

func (n None[T]) Earliest() opt.Option[T] {
	return opt.None[T]()
}

func (n None[T]) Latest() opt.Option[T] {
	return opt.None[T]()
}

type Single[T Offset] struct {
	offset T
}

func (s Single[T]) Single() opt.Option[T] {
	return opt.Some(s.offset)
}

func (s Single[T]) Earliest() opt.Option[T] {
	return opt.Some(s.offset)
}

func (s Single[T]) Latest() opt.Option[T] {
	return opt.Some(s.offset)
}

type Ambiguous[T Offset] struct {
	min T
	max T
}

func (a Ambiguous[T]) Single() opt.Option[T] {
	return opt.None[T]()
}

func (a Ambiguous[T]) Earliest() opt.Option[T] {
	return opt.Some(a.min)
}

func (a Ambiguous[T]) Latest() opt.Option[T] {
	return opt.Some(a.max)
}

type TimeZoneAbstract interface {
	Offset
}

type TimeZone interface {
	TimeZoneAbstract
}

type TimeZoneImpl[T TimeZoneAbstract] struct{}

type UTC struct {
	TimeZoneImpl[UTC]
}

func (u UTC) Fix() FixedOffset {
	return FixedOffsetEast(0).Unwrap()
}

func (u UTC) String() string {
	return "UTC"
}

type Local struct {
	TimeZoneImpl[FixedOffset]
}

func (l Local) Fix() FixedOffset {
	_, offset := time.Now().In(time.Local).Zone()

	return FixedOffsetEast(int32(offset)).Unwrap()
}

func (l Local) String() string {
	return time.Local.String()
}
