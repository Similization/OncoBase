package model

type PatientWithDisease struct {
	Id         string
	FirstName  string
	MiddleName string
	LastName   string
	BirthDate  string
	Sex        string
	SNILS      string
	Phone      string
}

type Doctor struct {
	Id         string
	FirstName  string
	MiddleName string
	LastName   string
	BirthDate  string
	Sex        string
	Phone      string
}

type Course struct {
	Id        int
	Period    string
	Frequency string
	Dose      string
	Drug      string
}

type PatientCourse struct {
	Id        int
	Patient   int
	Disease   string
	Course    string
	Doctor    int
	BeginDate string
	EndDate   string
	Diagnosis string
}

type BloodCount struct {
	Id               string // global blood count id
	Description      string
	MinNormalValue   string
	MaxNormalValue   string
	MinPossibleValue string
	MaxPossibleValue string
	MeasureCode      string
}

type BloodCountValue struct {
	Disease     string // global disease id
	BloodCount  string
	Coefficient float32
	Description string
}

type CourseProcedure struct {
	PatientCourse string // global disease id
	BloodCount    string
	BeginDate     float32
	Period        string
	Result        int
	MeasureCode   string
}
