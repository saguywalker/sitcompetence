package db

/*
func (db *Database) GetStudentByID(id string) (*model.Student, error) {
	var student model.Student

	if err := db.First(&student, id).Error; err != nil {
		if gorm.IsRecordNotFoundError(err) {
			return nil, nil
		}

		return nil, errors.Wrap(err, "unable to get todo")
	}

	return &student, nil
}

func (db *Database) GetStudents() ([]*model.Student, error) {
	var students []*model.Student
	err := errors.Wrap(db.Find(&students).Error, "unable to get students")
	return students, err
}

func (db *Database) CreateStudent(student *model.Student) error {
	return errors.Wrap(db.Create(student).Error, "unable to create a student")
}
*/
