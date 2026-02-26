package main

import (
	"net/http"
	"github.com/gin-gonic/gin"
)

// Novel 定义小说结构
type Novel struct {
	ID          string
	Title       string
	Description string
	Content     string
}

func main() {
	r := gin.Default()

	
	r.LoadHTMLGlob("templates/*")

	mockNovels := map[string]Novel{
		"the-lost-city": {
			ID:          "1",
			Title:       "The Lost City of Gold",
			Description: "An epic adventure of a traveler seeking the mythical city.",
			Content:     "Once upon a time, in a land far away...",
		},
	}


	r.GET("/", func(c *gin.Context) {
		c.HTML(http.StatusOK, "index.html", gin.H{
			"Title":  "Best English Web Novels - FullStoryHub",
			"Novels": mockNovels,
		})
	})


	r.GET("/novel/:slug", func(c *gin.Context) {
		slug := c.Param("slug")
		novel, exists := mockNovels[slug]
		if !exists {
			c.String(http.StatusNotFound, "Novel not found")
			return
		}

		
		c.HTML(http.StatusOK, "novel.html", gin.H{
			"Title":       novel.Title + " - Read Online Free | FullStoryHub",
			"Description": novel.Description,
			"Novel":       novel,
		})
	})


	r.GET("/sitemap.xml", func(c *gin.Context) {
		c.Header("Content-Type", "application/xml")
	
		c.String(http.StatusOK, `<?xml version="1.0" encoding="UTF-8"?>
<urlset xmlns="http://www.sitemaps.org/schemas/sitemap/0.9">
  <url><loc>https://www.fullstoryhub.com/</loc></url>
</urlset>`)
	})

	r.Run(":8080")
}
