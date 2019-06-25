package model

import (
	"github.com/jinzhu/gorm"
	_ "github.com/jinzhu/gorm/dialects/sqlite"
	"github.com/jinzhu/gorm/dialects/postgres"
	"time"
)

//type Project struct {
//	gorm.Model
//	Title    string `gorm:"unique" json:"title"`
//	Archived bool   `json:"archived"`
//	Tasks    []Task `gorm:"ForeignKey:ProjectID" json:"tasks"`
//}
//
//func (p *Project) Archive() {
//	p.Archived = true
//}
//
//func (p *Project) Restore() {
//	p.Archived = false
//}

type Task struct {
	gorm.Model
	Title     string     `json:"title"`
	//	Priority  string     `gorm:"type:ENUM('0', '1', '2', '3');default:'0'" json:"priority"`
	Deadline  *time.Time `gorm:"default:null" json:"deadline"`
	Status     postgres.Jsonb   `gorm:"type:json" json:"created"`
	Details    postgres.Jsonb  `gorm:"type:json" json:"details"`
	Host       string 		`json:"host"`
	Hostlist []TaskHost `gorm:"ForeignKey:TaskID" json:"hostlist"`
}


type TaskHost struct {
	TaskId uint `json:"task_id"  gorm:"primary_key;"`
	Host string `json:"host"  gorm:"primary_key;"`
	Response   postgres.Jsonb   `gorm:"type:json" json:"response"`
}

// DBMigrate will create and migrate the tables, and then make the some relationships if necessary
func DBMigrate(db *gorm.DB) *gorm.DB {
	db.AutoMigrate(&Task{}, &TaskHost{})
	//	db.Model(&Task{}).AddForeignKey("project_id", "projects(id)", "CASCADE", "CASCADE")
	return db
}
