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
	var found Dong
	g.db.
		Raw("SELECT * FROM dongs WHERE id IN (SELECT id FROM dongs ORDER BY RANDOM() LIMIT 1)").
		Scan(&found)

	return found.ToDomainDong(), nil
}

func (g GormRepository) RandomByCategory(cat string) (domain.Dong, error) {
	return domain.Dong{}, nil
}
