package error

import errField "field-service/constants/error/field"
import errFieldSchedule "field-service/constants/error/fieldschedule"

func ErrMapping(err error) bool {
	allErrors := append(append(GeneralErrors[:], errField.FieldErrors[:]...), errFieldSchedule.FieldScheduleErrors[:]...)

	for _, item := range allErrors {
		if err.Error() == item.Error() {
			return true
		}
	}
	return false
}
