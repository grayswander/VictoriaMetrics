package resolution

import "time"

/* This package handles the resolution of time series.
 * It is used everywhere in the codebase, where we need to calculate time from serialized data.
 * It is also used in the query engine, where we need to calculate the resolution of the time series.
 * Currently supported resolutions: milliseconds (default), nanoseconds
 */

var systemResolution Resolution = MillisecondsResolution{}

func SetResolution(resolution Resolution) {
	systemResolution = resolution
}

func GetResolution() Resolution {
	return systemResolution
}

func SetResolutionToMilliseconds() {
	SetResolution(MillisecondsResolution{})
}

func SetResolutionToNanoseconds() {
	SetResolution(NanosecondsResolution{})
}

type Resolution interface {
	GetTrimDuration(trimDuration time.Duration) int64
}

type MillisecondsResolution struct {
}

func (r MillisecondsResolution) GetTrimDuration(trimDuration time.Duration) int64 {
	return trimDuration.Milliseconds()
}

type NanosecondsResolution struct {
}

func (r NanosecondsResolution) GetTrimDuration(trimDuration time.Duration) int64 {
	return trimDuration.Nanoseconds()
}
