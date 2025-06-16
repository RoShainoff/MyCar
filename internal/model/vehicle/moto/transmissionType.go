package moto

type TransmissionTypeKind int

const (
	TransmissionTypeUnknown   TransmissionTypeKind = 0
	TransmissionTypeAuto      TransmissionTypeKind = 1
	TransmissionTypeAutomatic TransmissionTypeKind = 2
	TransmissionTypeRobot     TransmissionTypeKind = 3
	TransmissionTypeCVT       TransmissionTypeKind = 4
	TransmissionTypeManual    TransmissionTypeKind = 5
)

type TransmissionType struct {
	Id       TransmissionTypeKind
	Name     string
	ParentId TransmissionTypeKind
}

var TransmissionTypes = []TransmissionType{
	{Id: TransmissionTypeUnknown, Name: "Любая"},
	{Id: TransmissionTypeAuto, Name: "Автомат"},
	{Id: TransmissionTypeAutomatic, Name: "Автоматическая", ParentId: TransmissionTypeAuto},
	{Id: TransmissionTypeRobot, Name: "Робот", ParentId: TransmissionTypeAuto},
	{Id: TransmissionTypeCVT, Name: "Вариатор", ParentId: TransmissionTypeAuto},
	{Id: TransmissionTypeManual, Name: "Механическая"},
}

func (t TransmissionTypeKind) String() string {
	for _, tt := range TransmissionTypes {
		if tt.Id == t {
			return tt.Name
		}
	}
	return "Unknown"
}

func (id TransmissionTypeKind) GetTransmissionType() TransmissionType {
	var result TransmissionType
	for _, t := range TransmissionTypes {
		if t.Id == id {
			return t
		}
	}
	return result
}

func (t TransmissionTypeKind) GetSubtypes() []TransmissionType {
	var result []TransmissionType
	for _, tt := range TransmissionTypes {
		if tt.ParentId == t {
			result = append(result, tt)
		}
	}
	return result
}

func (t TransmissionTypeKind) GetParentType() TransmissionType {
	for _, tt := range TransmissionTypes {
		if tt.Id == t && tt.ParentId != 0 {
			return tt.ParentId.GetTransmissionType()
		}
	}
	return TransmissionTypeUnknown.GetTransmissionType()
}
