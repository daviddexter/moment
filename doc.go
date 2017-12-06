/*Moment package implements as inspired by momentjs

Extract the time duration from unix seconds

Extract the time duration from unix nanoseconds

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

    d := moment.FromSeconds(time.Now().Unix())
    dd := moment.HumanizeDuration(d)


    //get duration from now and unix epoch in nanoseconds

    e := moment.FromNanoSeconds(time.Now().UnixNano())
    ee := moment.HumanizeDuration(e)

    //get duration between two points in seconds

    f := moment.FromSecondsPoints(pA,pB)
    ff := moment.HumanizeDuration(f)

    The same applies for MomentFromNanoSecondsPoints function

    //To add leaps
    l := moment.UtilBuilder{momemt.Operation,moment.Leap,moment.Step}


*/

package moment
