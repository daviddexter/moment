package moment

import "time"

func addHour(t time.Time, step int64) time.Time {
	lp := secInHour * step
	nd := t.Unix() + lp
	return time.Unix(nd, 0)
}

func addDay(t time.Time, step int64) time.Time {
	lp := secInDay * step
	nd := t.Unix() + lp
	return time.Unix(nd, 0)
}

func addWeek(t time.Time, step int64) time.Time {
	lp := (secInDay * 7) * step
	nd := t.Unix() + lp
	return time.Unix(nd, 0)
}

func addYear(t time.Time, step int64) time.Time {
	lp := (secInDay * 365) * step
	nd := t.Unix() + lp
	return time.Unix(nd, 0)
}

func subHour(t time.Time, step int64) time.Time {
	lp := secInHour * step
	nd := t.Unix() - lp
	return time.Unix(nd, 0)
}

func subDay(t time.Time, step int64) time.Time {
	lp := secInDay * step
	nd := t.Unix() - lp
	return time.Unix(nd, 0)
}

func subWeek(t time.Time, step int64) time.Time {
	lp := (secInDay * 7) * step
	nd := t.Unix() - lp
	return time.Unix(nd, 0)
}

func subYear(t time.Time, step int64) time.Time {
	lp := (secInDay * 365) * step
	nd := t.Unix() - lp
	return time.Unix(nd, 0)
}
