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

Add leaps operations to time. Leaps of Hour,Day,week and year

Installing

    go get github.com/daviddexter/moment

Usage

    import (
	    "github.com/daviddexter/moment"
    )

    //get duration from now and unix epoch in seconds

    d := moment.MomentFromSeconds(time.Now().Unix())

    //get duration from now and unix epoch in nanoseconds

    e := moment.MomentFromNanoSeconds(time.Now().UnixNano())

    //get duration between two points in seconds

    f := moment.MomentFromSecondsPoints(pA,pB)

    The same applies for MomentFromNanoSecondsPoints function

    //Humanize time lapse.
    //Instead of returning the difference in " ... ago " format, moment returns the elapse time as a map

    h := moment.HumanizeDuration(time.Duration)

    //To add leaps
    l := moment.UtilBuilder{momemt.Operation,moment.Leap,moment.Step}

