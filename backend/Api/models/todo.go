package models

import "github.com/jinzhu/gorm"

type Todo struct {
	ID        uint64 `gorm:"primary_key;auto_increment" json:"id"`
	Item      string `gorm:"text;not null;" json:"item"`
	Completed bool   `gorm:"boolean;" json:"content"`
}

// SaveToDo is a method that saves a task to the DB
// It takes a pointer to a gorm.DB as an argument and returns a pointer to the saved task.
func (t *Todo) SaveToDo(db *gorm.DB) (*Todo, error) {
	if err := db.Debug().Model(&Todo{}).Create(&t).Error; err != nil || t.ID != 0 {
		return &Todo{}, err
	}
	return t, nil
}

// UpdateAToDo is a method that updates a task in the DB.
// It takes a pointer to a gorm.DB as an argument and returns a pointer to the updated task.
func (t *Todo) UpdateAToDo(db *gorm.DB) (*Todo, error) {

	if err := db.Debug().Model(&Todo{}).Where("id = ?", t.ID).Updates(Todo{Completed: t.Completed, Item: t.Item}).Error; err != nil || t.ID != 0 {
		return &Todo{}, err
	}
	return t, nil
}

// DeleteAtTodo deletes a task from the DB.
// It takes a pointer to a gorm.DB as argument and returns the number of affected rows.
func (t *Todo) DeleteAToDo(db *gorm.DB) (int64, error) {
	deletedDB := db.Debug().Model(&Todo{}).Where("id = ?", t.ID).Take(&Todo{}).Delete(&Todo{})
	if deletedDB.Error != nil {
		return 0, db.Error
	}
	return deletedDB.RowsAffected, nil
}
