package main

import (
	"errors"
	"fmt"
	"os"
	"strconv"
)

type Student struct {
	id      int
	name    string
	address string
	job     string
	reason  string
}

func searchStudent(key int, students []Student) (Student, error) {
	for _, student := range students {
		if student.id == key {
			return student, nil
		}
	}
	return Student{}, errors.New("siswa tidak ditemukan")
}

func initStudent() []Student {
	studentsArr := [][]string{
		{"MUHAMMAD ZUNAN ALFIKRI", "Jakarta", "Student", "Ekspektasi di course ini aku dapat memahami bahasa go lebih dalam dan microservice lebih baik."},
		{"ARFAN JADULHAQ", "Bandung", "Student", "Agar dapat memahami lebih dalam bahasa Go terutama implementasi di bidang backend web."},
		{"TRIYONO", "Surakarta", "Student", "Berharap bisa membuat api service dengan golang,  buat suplay aplikasi flutter ku :grin:"},
		{"ADITYA RIZKI PRATAMA", "Semarang", "Student", "Ekspektasi saya dengan mengikuti program studi independent Golang dari hacktiv8, bisa upgrade skill backend, belajar micro service."},
		{"YULYANO THOMAS DJAYA", "Jakarta", "Student", "menambah pengetahuan dan pengalaman dibidang Backend dengan bahasa Go"},
		{"ARFINAL", "ACEH", "Student", "menambah pengetahuan dan pengalaman di bidang backend bahasa go dan juga bahasa golang ni baru saya pelajari dari bahasa pemrograman yang lain"},
		{"FELIX YANGSEN", "Makasar", "Student", " memperdalam skill dalam backend yang memang menjadi pekerjaan saya sekarang"},
		{"WAHYU DWI RAMADHAN", "Jepara", "Student", "Agar dapat memahami lebih dalam bahasa Go terutama implementasi di bidang backend web."},
		{"MUHAMMAD HANIF NAUFAL EKA W", "Bali", "Student", " bisa belajar lebih banyak mengenai bahasa Go dan Microservice"},
		{"THOBIB KHOIRUL ANNAS", "Serang", "Student", "Agar dapat memahami lebih dalam bahasa Go terutama implementasi di bidang backend web."},
	}

	var students []Student
	id := 1
	for _, item := range studentsArr {
		// insert student to struct
		student := Student{
			id:      id,
			name:    item[0],
			address: item[1],
			job:     item[2],
			reason:  item[3],
		}
		// insert student object to students slice
		students = append(students, student)
		id++
	}

	return students
}

func main() {
	// get first args
	numAbsent := os.Args[1]
	key, err := strconv.Atoi(numAbsent)

	if err != nil {
		fmt.Print(err)
		os.Exit(0)
	}

	students := initStudent()
	student, errSearch := searchStudent(key, students)

	if errSearch != nil {
		fmt.Println(errSearch)
		os.Exit(0)
	}

	fmt.Println("No Absent: ", student.id)
	fmt.Println("Nama \t :", student.name)
	fmt.Println("Alamat \t :", student.address)
	fmt.Println("Job \t :", student.job)
	fmt.Println("Reason\t :", student.reason)
}
