package moment

import (
	"strings"
	"time"
)

//FromSeconds Extract the time duration from unix seconds
//Usage:
//MomentFromSeconds(time.Now().Unix())
func FromSeconds(sec int64) time.Duration {
	e := time.Unix(sec, 0)
	d := time.Now().Sub(e)
	return d
}

//FromNanoSeconds Extract the time duration from unix nanoseconds
//Usage:
//MomentFromNanoSeconds(time.Now().UnixNano())
func FromNanoSeconds(sec int64) time.Duration {
	e := time.Unix(sec, 0)
	d := time.Now().Sub(e)
	return d
}

//FromSecondsPoints Extract the time duration seconds of two different time points
func FromSecondsPoints(pA int64, pB int64) time.Duration {
	//var d time.Duration
	if pA > pB {
		g := time.Unix(pA, 0)
		l := time.Unix(pB, 0)
		i := g.Sub(l)
		return i

	}
	g := time.Unix(pB, 0)
	l := time.Unix(pA, 0)
	i := g.Sub(l)
	return i
}

//FromNanoSecondsPoints Extract the time duration nanoseconds of two different time points
func FromNanoSecondsPoints(pA int64, pB int64) time.Duration {
	//var d time.Duration
	if pA > pB {
		g := time.Unix(pA, 0)
		l := time.Unix(pB, 0)
		i := g.Sub(l)
		return i

	}
	g := time.Unix(pB, 0)
	l := time.Unix(pA, 0)
	i := g.Sub(l)
	return i
}

//HumanizeDuration Represents time duration as a map for fast readabilty and usage
//with keys of "hours","minutes","seconds"
func HumanizeDuration(d time.Duration) map[string]int64 {
	var fm map[string]int64
	ds := d.String()
	//contains := strings.ContainsAny(ds, "h &m & s")
	if strings.Contains(ds, "h") == true && strings.Contains(ds, "m") == true && strings.Contains(ds, "s") == true {

		fm = splitterFull(ds)
	}
	if strings.Contains(ds, "h") == false && strings.Contains(ds, "m") == true && strings.Contains(ds, "s") == true {

		fm = splitterWithoutHour(ds)
	}
	if strings.Contains(ds, "h") == false && strings.Contains(ds, "m") == false && strings.Contains(ds, "s") == true {

		fm = splitterWithoutMinutes(ds)
	}
	return fm
}
