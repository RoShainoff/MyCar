package expense

type Category string

const (
	Fuel        Category = "Fuel"
	Repair      Category = "Repair"
	Maintenance Category = "Maintenance"
	Insurance   Category = "Insurance"
	Taxes       Category = "Taxes"
	Other       Category = "Other"
)

func isValidCategory(cat Category) bool {
	switch cat {
	case Fuel, Repair, Maintenance, Insurance, Taxes, Other:
		return true
	default:
		return false
	}
}
