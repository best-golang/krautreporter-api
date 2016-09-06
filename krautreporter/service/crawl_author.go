package service

import (
	"log"
	"regexp"
	"strings"
	"time"

	"github.com/PuerkitoBio/goquery"
	"github.com/metalmatze/krautreporter-api/krautreporter/entity"
)

var authorPageSrcsetRegex = regexp.MustCompile(`(.*) 170w, (.*) 340w`)

func CrawlAuthor(a *entity.Author) error {
	doc, err := goquery.NewDocument(mainURL + a.URL)
	if err != nil {
		log.Println("Failed to fetch %s", a.URL)
		return err
	}

	doc.Find("header.article__header").Each(func(i int, s *goquery.Selection) {
		a.Biography = strings.TrimSpace(s.Find(".author__bio").Text())

		html, err := s.Find("#author-page--media-links").Html()
		if err == nil {
			a.SocialMedia = strings.TrimSpace(html)
		}

		if srcset, exists := s.Find(".author__monogram").Attr("srcset"); exists {
			matches := authorPageSrcsetRegex.FindStringSubmatch(srcset)
			if len(matches) == 3 {
				a.AddImage(entity.Image{Width: 170, Src: matches[1]})
				a.AddImage(entity.Image{Width: 340, Src: matches[2]})
			}
		}
	})

	a.Crawl.Next = time.Now().Add(6 * time.Hour)

	return nil
}
