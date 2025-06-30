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
	Id               CategoryKind
	Name             string
	ParentCategoryId CategoryKind
}

var ParentCategories = []Category{
	{Id: Offroad, Name: "Внедорожные"},
	{Id: Road, Name: "Дорожные"},
	{Id: Cruisers, Name: "Круизеры/Чопперы"},
	{Id: Sports, Name: "Спортивные"},
	{Id: Supermoto, Name: "Супермото"},
	{Id: Trikes, Name: "Трайки"},
	{Id: Tourism, Name: "Туристические"},
}

var ChildCategories = []Category{
	// Подкатегории внедорожных
	{Id: Allround, Name: "Allround", ParentCategoryId: Offroad},
	{Id: Enduro, Name: "Внедорожный Эндуро", ParentCategoryId: Offroad},
	{Id: DualPurpose, Name: "Мотоцикл повышенной проходимости", ParentCategoryId: Offroad},
	{Id: SportEnduro, Name: "Спортивный Эндуро", ParentCategoryId: Offroad},
	{Id: TouringEnduro, Name: "Туристический Эндуро", ParentCategoryId: Offroad},

	// Подкатегории дорожных
	{Id: Naked, Name: "Naked bike", ParentCategoryId: Road},
	{Id: RoadBike, Name: "Дорожный", ParentCategoryId: Road},
	{Id: Classic, Name: "Классик", ParentCategoryId: Road},

	// Подкатегории круизеров
	{Id: Custom, Name: "Кастом", ParentCategoryId: Cruisers},
	{Id: Cruiser, Name: "Круизер", ParentCategoryId: Cruisers},
	{Id: Chopper, Name: "Чоппер", ParentCategoryId: Cruisers},

	// Подкатегории спортивных
	{Id: Cross, Name: "Кросс", ParentCategoryId: Sports},
	{Id: Speedway, Name: "Speedway", ParentCategoryId: Sports},
	{Id: Kids, Name: "Детский", ParentCategoryId: Sports},
	{Id: Minibike, Name: "Минибайк", ParentCategoryId: Sports},
	{Id: Pitbike, Name: "Питбайк", ParentCategoryId: Sports},
	{Id: Trial, Name: "Триал", ParentCategoryId: Sports},
	{Id: Sportbike, Name: "Спорт-байк", ParentCategoryId: Sports},
	{Id: SportTouring, Name: "Спорт-туризм", ParentCategoryId: Sports},
	{Id: Supersport, Name: "Супер-спорт", ParentCategoryId: Sports},

	// Подкатегории трайков
	{Id: Trike, Name: "Трайк", ParentCategoryId: Trikes},
	{Id: ThreeWheeler, Name: "Трицикл", ParentCategoryId: Trikes},
}

func (id CategoryKind) GetSubcategories() []Category {
	var result []Category
	for _, c := range ChildCategories {
		if c.ParentCategoryId == id {
			result = append(result, c)
		}
	}
	return result
}

func (id CategoryKind) GetParentCategory() Category {
	var result Category
	for _, c := range ParentCategories {
		if c.Id == id {
			return c
		}
	}
	return result
}

func (id CategoryKind) GetCategory() Category {
	var result Category
	for _, c := range ChildCategories {
		if c.Id == id {
			return c
		}
	}
	return result
}

func (cc Category) GetCategoryName() string {
	for _, c := range ChildCategories {
		if c.Id == cc.Id {
			return c.Name
		}
	}
	return "Unknown"
}
