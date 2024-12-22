package repository

import (
	"a21hc3NpZ25tZW50/model"
	"fmt"

	"gorm.io/gorm"
)

type StudentRepository interface {
	FetchAll() ([]model.Student, error)
	FetchByID(id int) (*model.Student, error)
	Store(s *model.Student) error
	Update(id int, s *model.Student) error
	Delete(id int) error
	FetchWithClass() (*[]model.StudentClass, error)
}

type studentRepoImpl struct {
	db *gorm.DB
}

func NewStudentRepo(db *gorm.DB) *studentRepoImpl {
	return &studentRepoImpl{db}
}

func (s *studentRepoImpl) FetchAll() ([]model.Student, error) {
	rows, err := s.db.Table("students").Select("*").Rows()
	if err != nil {
		return nil, err
	}
	
	var students []model.Student // TODO: replace this
	for rows.Next(){
		var student model.Student
		s.db.ScanRows(rows, &student)
		if err != nil{
			return nil, err
		}
		students = append(students, student)
	}
	return students, nil// TODO: replace this
}// TODO: replace this


func (s *studentRepoImpl) Store(student *model.Student) error {
	if err := s.db.Create(student).Error; err != nil{
		return err
	}
	return nil // TODO: replace this
}

func (s *studentRepoImpl) Update(id int, student *model.Student) error {
	if err := s.db.Model(&model.Student{}).Where("id = ?", id).Updates(student).Error; err != nil{
		return err
	}
	return nil// TODO: replace this
}

func (s *studentRepoImpl) Delete(id int) error {
	if err := s.db.Delete(&model.Student{},id).Error; err != nil{
	   return fmt.Errorf("not implement") 
	}
	return nil // TODO: replace this
}

func (s *studentRepoImpl) FetchByID(id int) (*model.Student, error) {
	var student model.Student
	if err := s.db.First(&student, id).Error; err != nil{
		return nil, err
	}
	return &student, nil // TODO: replace this
}

func (s *studentRepoImpl) FetchWithClass() (*[]model.StudentClass, error) {
	var sClass []model.StudentClass
	if err  := s.db.Table("students").
		Select("students.name, students.address,  classes.name as class_name, classes.professor, classes.room_number").
		Joins("left join classes on classes.id = students.class_id").
		Scan(&sClass).Error; err !=nil{
			return nil, err
		}
		if sClass == nil{
			sClass =[]model.StudentClass{}
		}
		return &sClass, nil// TODO: replace this
}
