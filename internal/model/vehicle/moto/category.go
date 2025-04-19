package moto

type CategoryKind int

const (
	CategoryKindNone CategoryKind = 0
	Offroad          CategoryKind = 1
	Road             CategoryKind = 2
	Cruisers         CategoryKind = 3
	Sports           CategoryKind = 4
	Supermoto        CategoryKind = 5
	Trikes           CategoryKind = 6
	Tourism          CategoryKind = 7
	Allround         CategoryKind = 8
	Enduro           CategoryKind = 9
	DualPurpose      CategoryKind = 10
	SportEnduro      CategoryKind = 11
	TouringEnduro    CategoryKind = 12
	Naked            CategoryKind = 13
	RoadBike         CategoryKind = 14
	Classic          CategoryKind = 15
	Custom           CategoryKind = 16
	Cruiser          CategoryKind = 17
	Chopper          CategoryKind = 18
	Cross            CategoryKind = 19
	Speedway         CategoryKind = 20
	Kids             CategoryKind = 21
	Minibike         CategoryKind = 22
	Pitbike          CategoryKind = 23
	Trial            CategoryKind = 24
	Sportbike        CategoryKind = 25
	SportTouring     CategoryKind = 26
	Supersport       CategoryKind = 27
	Trike            CategoryKind = 28
	ThreeWheeler     CategoryKind = 29
)

type Category struct {
	id               CategoryKind
	name             string
	parentCategoryId CategoryKind
}

var ParentCategories = []Category{
	{id: Offroad, name: "Внедорожные"},
	{id: Road, name: "Дорожные"},
	{id: Cruisers, name: "Круизеры/Чопперы"},
	{id: Sports, name: "Спортивные"},
	{id: Supermoto, name: "Супермото"},
	{id: Trikes, name: "Трайки"},
	{id: Tourism, name: "Туристические"},
}

var ChildCategories = []Category{
	// Подкатегории внедорожных
	{id: Allround, name: "Allround", parentCategoryId: Offroad},
	{id: Enduro, name: "Внедорожный Эндуро", parentCategoryId: Offroad},
	{id: DualPurpose, name: "Мотоцикл повышенной проходимости", parentCategoryId: Offroad},
	{id: SportEnduro, name: "Спортивный Эндуро", parentCategoryId: Offroad},
	{id: TouringEnduro, name: "Туристический Эндуро", parentCategoryId: Offroad},

	// Подкатегории дорожных
	{id: Naked, name: "Naked bike", parentCategoryId: Road},
	{id: RoadBike, name: "Дорожный", parentCategoryId: Road},
	{id: Classic, name: "Классик", parentCategoryId: Road},

	// Подкатегории круизеров
	{id: Custom, name: "Кастом", parentCategoryId: Cruisers},
	{id: Cruiser, name: "Круизер", parentCategoryId: Cruisers},
	{id: Chopper, name: "Чоппер", parentCategoryId: Cruisers},

	// Подкатегории спортивных
	{id: Cross, name: "Кросс", parentCategoryId: Sports},
	{id: Speedway, name: "Speedway", parentCategoryId: Sports},
	{id: Kids, name: "Детский", parentCategoryId: Sports},
	{id: Minibike, name: "Минибайк", parentCategoryId: Sports},
	{id: Pitbike, name: "Питбайк", parentCategoryId: Sports},
	{id: Trial, name: "Триал", parentCategoryId: Sports},
	{id: Sportbike, name: "Спорт-байк", parentCategoryId: Sports},
	{id: SportTouring, name: "Спорт-туризм", parentCategoryId: Sports},
	{id: Supersport, name: "Супер-спорт", parentCategoryId: Sports},

	// Подкатегории трайков
	{id: Trike, name: "Трайк", parentCategoryId: Trikes},
	{id: ThreeWheeler, name: "Трицикл", parentCategoryId: Trikes},
}

func (id CategoryKind) GetSubcategories() []Category {
	var result []Category
	for _, c := range ChildCategories {
		if c.parentCategoryId == id {
			result = append(result, c)
		}
	}
	return result
}

func (id CategoryKind) GetParentCategory() Category {
	var result Category
	for _, c := range ParentCategories {
		if c.id == id {
			return c
		}
	}
	return result
}

func (id CategoryKind) GetCategory() Category {
	var result Category
	for _, c := range ChildCategories {
		if c.id == id {
			return c
		}
	}
	return result
}

func (cc Category) GetCategoryName() string {
	for _, c := range ChildCategories {
		if c.id == cc.id {
			return c.name
		}
	}
	return "Unknown"
}
