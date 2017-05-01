Moment package implements as inspired by momentjs

Extract the time duration from unix seconds
Usage:
MomentFromSeconds(time.Now().Unix())


Extract the time duration from unix nanoseconds
Usage:
MomentFromSeconds(time.Now().UnixNano())

Extract the time duration seconds of two different time points

Extract the time duration nanoseconds of two different time points

Represents time duration as a map for fast readabilty and usage
with keys of "hours","minutes","seconds"

Add leaps to time. Leaps of Hour,Day,week and year