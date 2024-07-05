package main

import (
	"html"
	"net/http"
	"strings"

	"github.com/gin-contrib/cors"
	"github.com/gin-gonic/gin"
)

type CompareRequest struct {
	Text1 string `json:"text1"`
	Text2 string `json:"text2"`
}

type CompareResponse struct {
	Comparison  string `json:"comparison"`
	Text1HTML   string `json:"text1Html"`
	Text2HTML   string `json:"text2Html"`
	Differences int    `json:"differences"`
}

func main() {
	r := gin.Default()

	config := cors.DefaultConfig()
	config.AllowOrigins = []string{"http://localhost:3000"}
	r.Use(cors.New(config))

	r.POST("/compare", compareTexts)

	r.Run(":8080")
}

func compareTexts(c *gin.Context) {
	var req CompareRequest
	if err := c.BindJSON(&req); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	lines1 := strings.Split(req.Text1, "\n")
	lines2 := strings.Split(req.Text2, "\n")

	comparison := "The texts are "
	if len(lines1) != len(lines2) {
		comparison += "different in number of lines. "
	} else {
		comparison += "of the same number of lines. "
	}

	differences := 0
	text1HTML := ""
	text2HTML := ""

	maxLines := len(lines1)
	if len(lines2) > maxLines {
		maxLines = len(lines2)
	}

	for i := 0; i < maxLines; i++ {
		if i < len(lines1) && i < len(lines2) {
			words1 := strings.Fields(lines1[i])
			words2 := strings.Fields(lines2[i])

			lineHTML1 := ""
			lineHTML2 := ""

			maxWords := len(words1)
			if len(words2) > maxWords {
				maxWords = len(words2)
			}

			for j := 0; j < maxWords; j++ {
				if j < len(words1) && j < len(words2) {
					if words1[j] != words2[j] {
						differences++
						lineHTML1 += "<span style='background-color: #ffcccb'>" + html.EscapeString(words1[j]) + "</span> "
						lineHTML2 += "<span style='background-color: #ffcccb'>" + html.EscapeString(words2[j]) + "</span> "
					} else {
						lineHTML1 += html.EscapeString(words1[j]) + " "
						lineHTML2 += html.EscapeString(words2[j]) + " "
					}
				} else if j < len(words1) {
					differences++
					lineHTML1 += "<span style='background-color: #ffcccb'>" + html.EscapeString(words1[j]) + "</span> "
				} else if j < len(words2) {
					differences++
					lineHTML2 += "<span style='background-color: #ffcccb'>" + html.EscapeString(words2[j]) + "</span> "
				}
			}

			text1HTML += strings.TrimSpace(lineHTML1) + "<br>"
			text2HTML += strings.TrimSpace(lineHTML2) + "<br>"
		} else if i < len(lines1) {
			differences++
			text1HTML += "<span style='background-color: #ffcccb'>" + html.EscapeString(lines1[i]) + "</span><br>"
		} else if i < len(lines2) {
			differences++
			text2HTML += "<span style='background-color: #ffcccb'>" + html.EscapeString(lines2[i]) + "</span><br>"
		}
	}

	comparison += "There are " + string(differences) + " differences between the texts."

	response := CompareResponse{
		Comparison:  comparison,
		Text1HTML:   strings.TrimSpace(text1HTML),
		Text2HTML:   strings.TrimSpace(text2HTML),
		Differences: differences,
	}

	c.JSON(http.StatusOK, response)
}
