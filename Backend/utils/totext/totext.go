package totext

import (
	"github.com/PuerkitoBio/goquery"
	"github.com/yuin/goldmark"
	"log"
	"strings"
)

func StripHTMLTags(htmlContent *string) string {
	doc, err := goquery.NewDocumentFromReader(strings.NewReader(*htmlContent))
	defer func() {
		if r := recover(); r != nil {
			log.Println("Recovered from:", r)
		}
	}()
	if err != nil {
		panic("解析html失败: " + err.Error())
	}
	return doc.Text()
}

func MarkdownToText(markdownContent *string) string {
	var buf strings.Builder
	err := goldmark.Convert([]byte(*markdownContent), &buf)
	defer func() {
		if r := recover(); r != nil {
			log.Println("Recovered from:", r)
		}
	}()

	if err != nil {
		panic("解析Markdown失败: " + err.Error())
	}
	htmlContent := buf.String()
	return StripHTMLTags(&htmlContent)
}
