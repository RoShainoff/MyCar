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
	id            BodyTypeKind
	name          string
	parentStyleId BodyTypeKind
}

var bodyTypes = []BodyType{
	{id: Sedan, name: "Седан"},
	{id: Hatchback, name: "Хэтчбек"},
	{id: Hatchback3Door, name: "Хэтчбек 3 дв.", parentStyleId: Hatchback},
	{id: Hatchback5Door, name: "Хэтчбек 5 дв.", parentStyleId: Hatchback},
	{id: Liftback, name: "Лифтбек", parentStyleId: Hatchback},
	{id: SUV, name: "Внедорожник"},
	{id: SUV3Door, name: "Внедорожник 3 дв.", parentStyleId: SUV},
	{id: SUV5Door, name: "Внедорожник 5 дв.", parentStyleId: SUV},
	{id: Wagon, name: "Универсал"},
	{id: Coupe, name: "Купе"},
	{id: Minivan, name: "Минивэн"},
	{id: Pickup, name: "Пикап"},
	{id: Limousine, name: "Лимузин"},
	{id: Van, name: "Фургон"},
	{id: Cabriolet, name: "Кабриолет"},
}

func (id BodyTypeKind) GetSubstyles() []BodyType {
	var result []BodyType
	for _, bs := range bodyTypes {
		if bs.parentStyleId == id {
			result = append(result, bs)
		}
	}
	return result
}

func (id BodyTypeKind) GetBodyType() BodyType {
	var result BodyType
	for _, c := range bodyTypes {
		if c.id == id {
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
	if name, ok := bodyTypeNames[b.id]; ok {
		return name
	}
	return "Неизвестно"
}
