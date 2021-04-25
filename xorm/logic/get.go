package logic

import (
	"github.com/gunsluo/go-example/xorm/models"
	"xorm.io/xorm"
)

func GetInspectorById(db *xorm.Session, inspector *models.Inspector, id int) (bool, error) {
	return db.Where("id = ?", id).Get(inspector)
}
