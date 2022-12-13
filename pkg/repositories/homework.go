package repositories

import (
	"errors"
	"golang-united-homework/pkg/database"
	"log"
	"time"
)

type Homework struct {
	Id          string `gorm:"type:uuid;primarykey;default:gen_random_uuid()"`
	LectureId   string `gorm:"index"`
	Title       string
	Description string
	CreatedAt   time.Time
	CreatedBy   string
	UpdatedAt   time.Time
	UpdatedBy   string
	DeletedAt   time.Time
	DeletedBy   string
}

func CreateHomework(homework *Homework) error {

	err := database.DB.Create(homework).Error
	if err != nil {
		log.Printf("Error on create homework: %s", err.Error())
		return errors.New("DB error")
	}

	return nil

}

func UpdateHomework(homework *Homework) error {

	err := database.DB.Updates(homework).Error
	if err != nil {
		log.Printf("Error on update homework: %s", err.Error())
		return errors.New("DB error")
	}

	return nil

}

func DeleteHomework(homework *Homework) error {

	err := database.DB.Delete(homework).Error
	if err != nil {
		log.Printf("Error on delete homework: %s", err.Error())
		return errors.New("DB error")
	}

	return nil

}

func GetHomework(id string) (*Homework, error) {

	homework := Homework{}

	err := database.DB.First(&homework, "id = ?", id).Error
	if err != nil {
		log.Printf("Error on select homework Id = %s: %s", id, err.Error())
		return nil, errors.New("DB error")
	}

	return &homework, nil

}

func GetHomeworkList(courseId string, showDeleted bool, limit uint32, offset uint32) (*[]Homework, error) {

	homework := []Homework{}

	query := database.DB.Model(&Homework{})

	if limit > 0 {
		query.Limit(int(limit))
	}

	if offset > 0 {
		query.Offset(int(offset))
	}

	if courseId != "" {
		query.Where("course_id = ?", courseId)
	}

	if !showDeleted {
		query.Where("deleted_at = '0001-01-01 00:00:00+00'")
	}

	query.Order("created_at asc")

	err := query.Find(&homework).Error
	if err != nil {
		log.Printf("Error on get list of homework: %s", err.Error())
		return nil, errors.New("DB error")
	}

	return &homework, nil

}
