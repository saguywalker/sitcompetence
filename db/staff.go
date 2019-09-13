package db

/*
func (db *Database) GetStaffByID(id string) (*model.Staff, error) {
	var staff model.Staff

	if err := db.First(&staff, id).Error; err != nil {
		if gorm.IsRecordNotFoundError(err) {
			return nil, nil
		}

		return nil, errors.Wrap(err, "unable to get staff")
	}

	return &staff, nil
}

func (db *Database) GetStaffs() ([]*model.Staff, error) {
	var staffs []*model.Staff
	err := errors.Wrap(db.Find(&staffs).Error, "unable to get staffs")
	return staffs, err
}

func (db *Database) CreateStaff(staff *model.Staff) error {
	return errors.Wrap(db.Create(staff).Error, "unable to create a staff")
}
*/
