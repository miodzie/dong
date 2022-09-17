package impl

import (
	"github.com/jinzhu/gorm"
	"github.com/miodzie/dong/domain"
)

func NewGormRepository(database *gorm.DB) *GormRepository {
	database.AutoMigrate(&Dong{})
	return &GormRepository{db: database}
}

type GormRepository struct {
	db *gorm.DB
}

func (g GormRepository) Random() (domain.Dong, error) {
	var ding Dong
	g.db.
		Raw(`SELECT * FROM dongs WHERE id IN (SELECT id FROM dongs ORDER BY RANDOM() LIMIT 1)`).
		Scan(&ding)

	return ding.ToDomainDong(), nil
}

func (g GormRepository) RandomByCategory(cat string) (domain.Dong, error) {
	var ding Dong
	g.db.Raw(`SELECT * FROM dongs WHERE id 
                   IN (SELECT id FROM dongs WHERE category IN (?) ORDER BY RANDOM() LIMIT 1)`, cat).
		Scan(&ding)

	return ding.ToDomainDong(), nil
}
