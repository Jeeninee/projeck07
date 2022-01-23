package entity

import (
	"gorm.io/driver/sqlite"
	"gorm.io/gorm"
)

var db *gorm.DB

func DB() *gorm.DB {
	return db
}

func SetupDatabase() {
	database, err := gorm.Open(sqlite.Open("se-64.db"), &gorm.Config{})
	if err != nil {
		panic("failed to connect database")
	}

	// Migrate the schema
	database.AutoMigrate(
		&IncreaseGrades{},
		&Grades{},
		&Staff{},
		&Course{},
		&Student{}
	)

	db = database

	//Student Data
		student1 := Student{
			ID_student: "B6202385",
			Prefix: "Miss",
			Name: "Phatcha Sisuwo",
			Major: "CPE",
			Year: 3,
			Email: "phatcha@gmail.com",
			Password: "",
		}
		db.Model(&Student{}).Create(&student1)
		student2 := Student{
			ID_student: "B6202743",
			Prefix: "Miss",
			Name: "Narudee Arunno",
			Major: "CPE",
			Year: 3,
			Email: "narudee@gmail.com",
			Password: "",
		}
		db.Model(&Student{}).Create(&student2)
		student3 := Student{
			ID_student: "B6214449",
			Prefix: "Miss",
			Name: "Suwanan Thamsui",
			Major: "CPE",
			Year: 3,
			Email: "suwanan@gmail.com",
			Password: "",
		}
		db.Model(&Student{}).Create(&student3)
		student4 := Student{
			ID_student: "B6230760",
			Prefix: "Miss",
			Name: "Patnarin Aiewchoei",
			Major: "CPE",
			Year: 3,
			Email: "patnarin@gmail.com",
			Password: "",
		}
		db.Model(&Student{}).Create(&student4)
		student5 := Student{
			ID_student: "B5924615",
			Prefix: "Mr.",
			Name: "Patnarin Aiewchoei",
			Major: "CPE",
			Year: 3,
			Email: "Pawarit Praneetponkrang",
			Password: "",
		}
		db.Model(&Student{}).Create(&student4)
		student5 := Student{
			ID_student: "B5901258",
			Prefix: "Miss.",
			Name: "Jeeninee Khuptawuttinun",
			Major: "CPE",
			Year: 3,
			Email: "Jeeninee.sai@gmail.com",
			Password: "",
		}

		db.Model(&Student{}).Create(&student6)

	//Registrar Data
		registrar1 := Registrar{
			ID_registrar: "R5678912",
			Name: "Prayut",
			Email: "Prayut@gmail.com",
			Password: "",
		}
		db.Model(&Registrar{}).Create(&registrar1)
		registrar2 := Registrar{
			ID_registrar: "R2034567",
			Name: "Praweena",
			Email: "Praweena@gmail.com",
			Password: "",
		}
		db.Model(&Registrar{}).Create(&registrar2)
		registrar3 := Registrar{
			ID_registrar: "R7891234",
			Name: "Prawit",
			Email: "Prawit@gmail.com",
			Password: "",
		}
		db.Model(&Registrar{}).Create(&registrar3)
		registrar4 := Registrar{
			ID_registrar: "R3454321",
			Name: "Sutin",
			Email: "Sutin@gmail.com",
			Password: "",
		}
		db.Model(&Registrar{}).Create(&registrar4)

		//Course Data
		course1 := Course{
			Coursename: "SOFTWARE ENGINEERING",
			Coursenumber: 523332,
		}
		db.Model(&Course{}).Create(&course1)
		course2 := Course{
			Coursename: "COMPUTER AND COMMUNICATION",
			Coursenumber: 523352,
		}
		db.Model(&Course{}).Create(&course2)
		course3 := Course{
			Coursename: "OPERATING SYSTEMS",
			Coursenumber: 523354,
		}
		db.Model(&Course{}).Create(&course3)
		course4 := Course{
			Coursename: "System Analysis and Design",
			Coursenumber: 523331,
		}
		db.Model(&Course{}).Create(&course4)
		course5 := Course{
			Coursename: "DATABASE SYSTEMS",
			Coursenumber: 523211,
		}
		db.Model(&Course{}).Create(&course5)
		course6 := Course{
			Coursename: "COMPUTER STATISTICS",
			Coursenumber: 523301,
		}
		db.Model(&Course{}).Create(&course6)

		//Grades Data
		grade1 := Grades{
			grade: "A",
		}
		db.Model(&Grades{}).Create(&grade1)
		grade2 := Grades{
			grade: "B+",
		}
		db.Model(&Grades{}).Create(&grade2)
		grade3 := Grades{
			grade: "B",
		}
		db.Model(&Grades{}).Create(&grade3)	
		grade4 := Grades{
			grade: "C+",
		}
		db.Model(&Grades{}).Create(&grade4)	
		grade5 := Grades{
			grade: "C",
		}
		db.Model(&Grades{}).Create(&grade5)	
		grade6 := Grades{
			grade: "D+",
		}
		db.Model(&Grades{}).Create(&grade6)	
		grade7 := Grades{
			grade: "D",
		}
		db.Model(&Grades{}).Create(&grade7)
		grade8 := Grades{
			grade: "F",
		}
		db.Model(&Grades{}).Create(&grade8)
}