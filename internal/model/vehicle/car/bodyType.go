package car

type BodyTypeKind int

const (
	BodyStyleNone  BodyTypeKind = 0
	Sedan          BodyTypeKind = 1
	Hatchback      BodyTypeKind = 2
	Hatchback3Door BodyTypeKind = 3
	Hatchback5Door BodyTypeKind = 4
	Liftback       BodyTypeKind = 5
	SUV            BodyTypeKind = 6
	SUV3Door       BodyTypeKind = 7
	SUV5Door       BodyTypeKind = 8
	Wagon          BodyTypeKind = 9
	Coupe          BodyTypeKind = 10
	Minivan        BodyTypeKind = 11
	Pickup         BodyTypeKind = 12
	Limousine      BodyTypeKind = 13
	Van            BodyTypeKind = 14
	Cabriolet      BodyTypeKind = 15
)

type BodyType struct {
	Id            BodyTypeKind
	Name          string
	ParentStyleId BodyTypeKind
}

var bodyTypes = []BodyType{
	{Id: Sedan, Name: "Седан"},
	{Id: Hatchback, Name: "Хэтчбек"},
	{Id: Hatchback3Door, Name: "Хэтчбек 3 дв.", ParentStyleId: Hatchback},
	{Id: Hatchback5Door, Name: "Хэтчбек 5 дв.", ParentStyleId: Hatchback},
	{Id: Liftback, Name: "Лифтбек", ParentStyleId: Hatchback},
	{Id: SUV, Name: "Внедорожник"},
	{Id: SUV3Door, Name: "Внедорожник 3 дв.", ParentStyleId: SUV},
	{Id: SUV5Door, Name: "Внедорожник 5 дв.", ParentStyleId: SUV},
	{Id: Wagon, Name: "Универсал"},
	{Id: Coupe, Name: "Купе"},
	{Id: Minivan, Name: "Минивэн"},
	{Id: Pickup, Name: "Пикап"},
	{Id: Limousine, Name: "Лимузин"},
	{Id: Van, Name: "Фургон"},
	{Id: Cabriolet, Name: "Кабриолет"},
}

func (id BodyTypeKind) GetSubstyles() []BodyType {
	var result []BodyType
	for _, bs := range bodyTypes {
		if bs.ParentStyleId == id {
			result = append(result, bs)
		}
	}
	return result
}

func (id BodyTypeKind) GetBodyType() BodyType {
	var result BodyType
	for _, c := range bodyTypes {
		if c.Id == id {
			return c
		}
	}
	return result
}

var bodyTypeNames = map[BodyTypeKind]string{
	BodyStyleNone:  "Любой",
	Sedan:          "Седан",
	Hatchback:      "Хэтчбек",
	Hatchback3Door: "Хэтчбек 3 дв.",
	Hatchback5Door: "Хэтчбек 5 дв.",
	Liftback:       "Лифтбек",
	SUV:            "Внедорожник",
	SUV3Door:       "Внедорожник 3 дв.",
	SUV5Door:       "Внедорожник 5 дв.",
	Wagon:          "Универсал",
	Coupe:          "Купе",
	Minivan:        "Минивэн",
	Pickup:         "Пикап",
	Limousine:      "Лимузин",
	Van:            "Фургон",
	Cabriolet:      "Кабриолет",
}

func (b BodyType) String() string {
	if name, ok := bodyTypeNames[b.Id]; ok {
		return name
	}
	return "Неизвестно"
}
