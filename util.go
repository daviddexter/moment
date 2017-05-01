package moment

import "time"
import "errors"

const (
	secInMin     = 60
	secInHour    = secInMin * 60
	secInDay     = secInHour * 24
	ADDOPERATION = "add"
	SUBOPERATION = "sub"
	HOURLEAP     = "hour"
	DAYLEAP      = "day"
	WEEKLEAP     = "week"
	YEARLEAP     = "year"
)

/*Strcuture for utiliy actions*/
type UtilBuilder struct {
	Operation string
	Leap      string
	Step      int64
}

/*Add operation to add hour,day,week or year. Result depends on the leap*/
func (v *UtilBuilder) Add(t time.Time) (res time.Time, err error) {
	if v.Operation != ADDOPERATION {
		err = errors.New("Invalid Operation")
		res = t
	} else {
		switch v.Leap {
		case HOURLEAP:
			res = addHour(t, v.Step)
		case DAYLEAP:
			res = addDay(t, v.Step)
		case WEEKLEAP:
			res = addWeek(t, v.Step)
		case YEARLEAP:
			res = addYear(t, v.Step)
		default:
			err = errors.New("Undefined Operation")
			res = t
		}
	}
	return
}

/*Subtract operation to add hour,day,week or year. Result depends on the leap*/
func (v *UtilBuilder) Sub(t time.Time) (res time.Time, err error) {
	if v.Operation != SUBOPERATION {
		err = errors.New("Invalid Operation")
		res = t
	} else {
		switch v.Leap {
		case HOURLEAP:
			res = subHour(t, v.Step)
		case DAYLEAP:
			res = subDay(t, v.Step)
		case WEEKLEAP:
			res = subWeek(t, v.Step)
		case YEARLEAP:
			res = subYear(t, v.Step)
		default:
			err = errors.New("Undefined Operation")
			res = t
		}
	}
	return
}
