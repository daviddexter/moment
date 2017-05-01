package moment

import (
	"fmt"
	"strings"
	"time"
)

/* Moment package implements as inspired by momentjs*/

//Extract the time duration from unix seconds
//Usage:
//MomentFromSeconds(time.Now().Unix())
func MomentFromSeconds(sec int64) time.Duration {
	e := time.Unix(sec, 0)
	d := time.Now().Sub(e)
	return d
}

//Extract the time duration from unix nanoseconds
//Usage:
//MomentFromSeconds(time.Now().UnixNano())
func MomentFromNanoSeconds(sec int64) time.Duration {
	e := time.Unix(sec, 0)
	d := time.Now().Sub(e)
	return d
}

//Extract the time duration seconds of two different time points
func MomentFromSecondsPoints(pA int64, pB int64) time.Duration {
	//var d time.Duration
	if pA > pB {
		g := time.Unix(pA, 0)
		l := time.Unix(pB, 0)
		i := g.Sub(l)
		return i

	} else {
		g := time.Unix(pB, 0)
		l := time.Unix(pA, 0)
		i := g.Sub(l)
		return i
	}
}

//Extract the time duration nanoseconds of two different time points
func MomentFromNanoSecondsPoints(pA int64, pB int64) time.Duration {
	//var d time.Duration
	if pA > pB {
		g := time.Unix(pA, 0)
		l := time.Unix(pB, 0)
		i := g.Sub(l)
		return i

	} else {
		g := time.Unix(pB, 0)
		l := time.Unix(pA, 0)
		i := g.Sub(l)
		return i
	}
}

//Represents time duration as a map for fast readabilty and usage
//with keys of "hours","minutes","seconds"
func HumanizeDuration(d time.Duration) map[string]int64 {
	var fm map[string]int64
	ds := d.String()
	//contains := strings.ContainsAny(ds, "h &m & s")
	if strings.Contains(ds, "h") == true && strings.Contains(ds, "m") == true && strings.Contains(ds, "s") == true {
		fmt.Println("line 69")
		fm = splitterFull(ds)
	}
	if strings.Contains(ds, "h") == false && strings.Contains(ds, "m") == true && strings.Contains(ds, "s") == true {
		fmt.Println("line 73")
		fm = splitterWithoutHour(ds)
	}
	if strings.Contains(ds, "h") == false && strings.Contains(ds, "m") == false && strings.Contains(ds, "s") == true {
		fmt.Println("line 77")
		fm = splitterWithoutMinutes(ds)
	}
	return fm
}
