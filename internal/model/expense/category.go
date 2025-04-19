package expense

type Category int

const (
	Fuel Category = iota
	Repair
	Maintenance
	Insurance
	Taxes
	Other
)

func isValidCategory(cat Category) bool {
	switch cat {
	case Fuel, Repair, Maintenance, Insurance, Taxes, Other:
		return true
	default:
		return false
	}
}
