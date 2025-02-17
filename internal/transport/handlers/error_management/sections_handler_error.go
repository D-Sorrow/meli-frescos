package error_management

const TemperatureInvalid = "Description must not be empty"
const Current_capacityinvalid = "Rates must not be negative"

type Error struct {
	Message string
	Code    int
}

func (e Error) Error() string {
	return e.Message
}
