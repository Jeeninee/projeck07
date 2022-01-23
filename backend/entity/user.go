package entity
import (
	//  "time"
	  "gorm.io/gorm"
	)
	 
	//-------------------ระบบบันทึกผลการเรียน--------------------------//

	type Grades struct {
		gorm.Model
		Grade 	string

		IncreaseGrades	[]IncreaseGrades `gorm:"foreignKey:GradesID"`
	}

	type Course struct{
		gorm.Model
		Coursename string
		Coursenumber int32
	
		IncreaseGrades	[]IncreaseGrades `gorm:"foreignKey:GradesID"`
	}

	type Registrar struct{
		gorm.Model
		ID_registrar string `gorm:"uniqueIndex"`
		Name        string
		Email       string //`gorm:"uniqueIndex"`
		Password    string

		IncreaseGrades	[]IncreaseGrades `gorm:"foreignKey:GradesID"`
	}

	type Student struct {
		gorm.Model
		ID_student string `gorm:"uniqueIndex"`
		Prefix     string
		Name       string
		Major      string
		Email      string // `gorm:"uniqueIndex"`
		Year       uint
		Password   string

		IncreaseGrades	[]IncreaseGrades `gorm:"foreignKey:GradesID"`
		
	}

	type IncreaseGrades struct {
		gorm.Model
		date	time.Time

		StudentID 	*uint
		Student		Student `gorm:"references:id"`

		GradesID	*uint
		Grades		Grades `gorm:"references:id"`

		CourseID 	*uint
		Course		Course `gorm:"references:id"`

		RegistrarID		*uint
		Registrar		Registrar `gorm:"references:id"`
	}