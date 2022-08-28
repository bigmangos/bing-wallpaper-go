package html

import (
	"os"
	"path"
)

var (
	bingHtmlRoot          = "docs/"
	bingHtmlIndexTemplate = "docs/bing-template.html"
)

func readIndexTemplateFile() (string, error) {
	content, err := os.ReadFile(bingHtmlIndexTemplate)
	if err != nil {
		return "", err
	}
	return string(content), nil
}

func writeIndexHtml(html string) error {
	indexPath := path.Join(bingHtmlRoot, "index.html")
	if err := os.WriteFile(indexPath, []byte(html), 0644); err != nil {
		return err
	}
	return nil
}

func writeMonthHtml(month, html string) error {
	indexPath := path.Join(bingHtmlRoot, month+".html")
	if err := os.WriteFile(indexPath, []byte(html), 0644); err != nil {
		return err
	}
	return nil
}
