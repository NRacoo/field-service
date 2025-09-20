package error

import "errors"

var (
	ErrFieldSchedule      = errors.New("field schedule not found")
	ErrFieldScheduleExist = errors.New("field schedule already exist")
)

var FieldScheduleErrors = []error{
	ErrFieldSchedule,
	ErrFieldScheduleExist,
}
