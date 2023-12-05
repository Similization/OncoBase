package model

type Diagnosis struct {
	Id          string // diag id
	Description string
}

type Drug struct {
	Id          string // drug global id
	Indications string
	Description string
}

type Patient struct {
	Id                   string
	FirstName            string
	MiddleName           string
	LastName             string
	BirthDate            string
	Sex                  string
	SNILS                string
	Phone                string
	PreliminaryDiagnosis string
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
	Patient   string
	Disease   string
	Course    string
	BeginDate string
	EndDate   string
	result    string
}

type UnitMeasure struct {
	Id        string // global unit measure id
	Shorthand string
	FullText  string
	Global    string
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

type Disease struct {
	Id          string // global disease id
	Description string
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
