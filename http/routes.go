package http

import (
	"net/http"

	"github.com/MetalMatze/Krautreporter-API/domain"
	"github.com/MetalMatze/Krautreporter-API/http/controller"
	"github.com/MetalMatze/gollection"
	"github.com/MetalMatze/gollection/router"
)

func Routes(g *gollection.Gollection, kr *domain.Krautreporter) func(router.Router) {
	return func(r router.Router) {
		r.GET("/", func(req router.Request, res router.Response) error {
			return res.String(http.StatusOK, "hi")
		})

		authorsController := controller.AuthorsController{AuthorInteractor: kr.AuthorInteractor, Log: g.Log}
		r.GET("/authors", authorsController.GetAuthors)
		r.GET("/authors/:id", authorsController.GetAuthor)

		articlesController := controller.ArticlesController{ArticleInteractor: kr.ArticleInteractor, Log: g.Log}
		r.GET("/articles", articlesController.GetArticles)
		r.GET("/articles/:id", articlesController.GetArticle)
	}
}
