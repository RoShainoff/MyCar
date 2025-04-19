package car

type DriveTypeKind int

const (
	UnknownDriveType DriveTypeKind = 0
	FWD              DriveTypeKind = 1
	RWD              DriveTypeKind = 2
	AWD              DriveTypeKind = 3 // Новый родительский тип
	PermanentAWD     DriveTypeKind = 4
	SwitchableAWD    DriveTypeKind = 5
)

type DriveType struct {
	Id       DriveTypeKind
	Name     string
	ParentId DriveTypeKind
}

var driveTypes = []DriveType{
	{Id: FWD, Name: "Передний", ParentId: UnknownDriveType},
	{Id: RWD, Name: "Задний", ParentId: UnknownDriveType},
	{Id: AWD, Name: "Полный", ParentId: UnknownDriveType},
	{Id: PermanentAWD, Name: "Постоянный полный", ParentId: AWD},
	{Id: SwitchableAWD, Name: "Подключаемый полный", ParentId: AWD},
}

func (d DriveTypeKind) String() string {
	for _, t := range driveTypes {
		if t.Id == d {
			return t.Name
		}
	}
	return "Неизвестно"
}

func (id DriveTypeKind) GetDriveType() DriveType {
	for _, t := range driveTypes {
		if t.Id == id {
			return t
		}
	}
	return DriveType{}
}

func (id DriveTypeKind) GetSubtypes() []DriveType {
	var subtypes []DriveType
	for _, t := range driveTypes {
		if t.ParentId == id {
			subtypes = append(subtypes, t)
		}
	}
	return subtypes
}

func (id DriveTypeKind) GetParentType() *DriveType {
	for _, t := range driveTypes {
		if t.Id == id && t.ParentId != UnknownDriveType {
			for _, parent := range driveTypes {
				if parent.Id == t.ParentId {
					return &parent
				}
			}
		}
	}
	return nil
}
