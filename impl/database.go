package impl

import (
	"github.com/jinzhu/gorm"
	"github.com/miodzie/dong"
)

// TODO: DELETE ME AND REFACTOR SCRAPER!
var db *gorm.DB

func NewGormRepository(database *gorm.DB) *GormRepository {
	database.AutoMigrate(&Dong{})
	return &GormRepository{db: database}
}

type GormRepository struct {
	db *gorm.DB
}

func (g GormRepository) Random() (dong.Emoji, error) {
	var ding Dong
	g.db.
		Raw(`SELECT * FROM dongs WHERE id IN (SELECT id FROM dongs ORDER BY RANDOM() LIMIT 1)`).
		Scan(&ding)

	return ding.ToDomainDong(), nil
}

func (g GormRepository) RandomByCategory(cat string) (dong.Emoji, error) {
	var ding Dong
	g.db.Raw(`SELECT * FROM dongs WHERE id 
                   IN (SELECT id FROM dongs WHERE category IN (?) ORDER BY RANDOM() LIMIT 1)`, cat).
		Scan(&ding)

	return ding.ToDomainDong(), nil
}

func (g GormRepository) Count() int64 {
	var count int64
	g.db.Model(&Dong{}).Count(&count)
	return count
}

func (g GormRepository) Categories() ([]string, error) {
	var cats []string
	rows, err := g.db.Model(&Dong{}).Select("category").Group("category").Rows()
	if err != nil {
		return cats, err
	}
	for rows.Next() {
		var c string
		rows.Scan(&c)
		cats = append(cats, c)
	}

	return cats, nil
}
