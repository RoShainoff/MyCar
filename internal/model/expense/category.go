package expense

type CategoryKind int

const (
	UnknownCategory CategoryKind = 0
	Fuel            CategoryKind = 1
	Repair          CategoryKind = 2
	Maintenance     CategoryKind = 3
	Insurance       CategoryKind = 4
	Taxes           CategoryKind = 5
	Other           CategoryKind = 6
)

type Category struct {
	Id   CategoryKind
	Name string
}

var Categories = []Category{
	{Id: UnknownCategory, Name: "Unknown"},
	{Id: Fuel, Name: "Fuel"},
	{Id: Repair, Name: "Repair"},
	{Id: Maintenance, Name: "Maintenance"},
	{Id: Insurance, Name: "Insurance"},
	{Id: Taxes, Name: "Taxes"},
	{Id: Other, Name: "Other"},
}

func (ck CategoryKind) String() string {
	for _, c := range Categories {
		if c.Id == ck {
			return c.Name
		}
	}
	return "Unknown"
}

func (ck CategoryKind) GetCategory() Category {
	for _, c := range Categories {
		if c.Id == ck {
			return c
		}
	}
	return Category{}
}

func isValidCategory(cat Category) bool {
	if cat == (Category{}) {
		return false
	}

	switch cat.Id {
	case Fuel, Repair, Maintenance, Insurance, Taxes, Other:
		return true
	default:
		return false
	}
}

func isValidCategoryKind(cat CategoryKind) bool {
	switch cat {
	case Fuel, Repair, Maintenance, Insurance, Taxes, Other:
		return true
	default:
		return false
	}
}
