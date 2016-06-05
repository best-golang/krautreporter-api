package entity

import "time"

type Author struct {
	ID          int       `json:"id"`
	Ordering    int       `json:"order"`
	Name        string    `json:"name"`
	Title       string    `json:"title"`
	URL         string    `json:"url"`
	Biography   string    `json:"biography"`
	SocialMedia string    `json:"social-media"`
	CreatedAt   time.Time `json:"created_at"`
	UpdatedAt   time.Time `json:"updated_at"`

	Images []Image `gorm:"polymorphic:Imageable;" json:"images"`
	Crawl  Crawl   `gorm:"polymorphic:Crawlable;" json:"-"`
}
