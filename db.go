package model

import (
	"github.com/jinzhu/gorm"
	_ "github.com/jinzhu/gorm/dialects/mysql"
)

type Idlibatch struct {
	Sambar int
	Vadas  int
	Hot    string
}

type DBV interface {
	Create(value interface{})
	Save(value interface{})
	Find(dest interface{}, conds interface{})
}

type Mysql struct {
	db *gorm.DB
}

func Createdb() *Mysql {

	d, err := gorm.Open("mysql", "root:Balaram@123@tcp(127.0.0.1:3306)/sys?charset=utf8&parseTime=True")

	if err != nil {

	}

	return &Mysql{db: d}
}

func (p *Mysql) Create(value interface{}) {

	p.db.Create(value)

}

func (p *Mysql) Save(value interface{}) {

	p.db.Save(value)
}

func (p *Mysql) Find(dest interface{}, conds interface{}) {

	p.db.Find(dest, conds)

}
  