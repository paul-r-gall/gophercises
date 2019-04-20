package cmd

import (
	"time"

	"github.com/jinzhu/gorm"
	// add driver
	_ "github.com/jinzhu/gorm/dialects/sqlite"
)

// task extends gorm.Model and contains fields
// any task can have at most one label.
// not exported, only accessible from withing pkg cmd
type task struct {
	gorm.Model
	Name       string `gorm:"unique"`
	Imp        bool
	Label      string
	CompTime   time.Time
	CreateTime time.Time
}

func (tsk *task) Complete() {
	tsk.CompTime = time.Now()
}

func (tsk *task) Promote() {
	tsk.Imp = true
}

func (tsk *task) Demote() {
	tsk.Imp = false
}
