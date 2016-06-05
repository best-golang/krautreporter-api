package marshaller

import (
	"github.com/MetalMatze/Krautreporter-API/krautreporter/entity"
)

type imageMarshaller struct {
	ID    int    `json:"id"`
	Width int    `json:"width"`
	Src   string `json:"src"`
}

func marshallImage(i entity.Image) imageMarshaller {
	return imageMarshaller{
		ID:    i.ID,
		Width: i.Width,
		Src:   KrautreporterURL + i.Src,
	}
}

func Images(images []entity.Image) map[string][]imageMarshaller {
	var im []imageMarshaller

	for _, i := range images {
		im = append(im, marshallImage(i))
	}

	return map[string][]imageMarshaller{
		"data": im,
	}
}