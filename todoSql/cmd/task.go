package cmd

import (
	"github.com/jinzhu/gorm"
	// add driver
	_ "github.com/jinzhu/gorm/dialects/sqlite"
)

// task extends gorm.Model and contains fields
// any task can have at most one label.
// not exported, only accessible from withing pkg cmd
type task struct {
	gorm.Model
	Name  string `gorm:"unique"`
	Imp   bool
	Label string
}

func (tsk *task) Promote() {
	tsk.Imp = true
}

func (tsk *task) Demote() {
	tsk.Imp = false
}
