package repository

import (
	"time"

	"github.com/MetalMatze/Krautreporter-API/domain/entity"
	"github.com/jinzhu/gorm"
)

type CrawlRepository struct {
	DB *gorm.DB
}

func (r CrawlRepository) FindOutdatedAuthors() ([]entity.Author, error) {
	var crawls []*entity.Crawl
	r.DB.Where("next < ?", time.Now()).Where("crawlable_type = ?", "authors").Order("next").Find(&crawls)

	var IDs []int
	for _, c := range crawls {
		IDs = append(IDs, c.CrawlableID)
	}

	var authors []entity.Author
	r.DB.LogMode(true)
	r.DB.Preload("Crawl").Where(IDs).Find(&authors)
	r.DB.LogMode(false)

	return authors, nil
}
