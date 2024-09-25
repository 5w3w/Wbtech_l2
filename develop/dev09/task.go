package main

import (
	"fmt"
	"io"
	"net/http"
	"net/url"
	"os"
	"path"

	"golang.org/x/net/html"
)

// Загружает содержимое страницы по указанному URL
func downloadPage(urlStr, outputDir string) error {
	resp, err := http.Get(urlStr)
	if err != nil {
		return err
	}
	defer resp.Body.Close()

	// Создаем папку для сохранения сайта
	if err := os.MkdirAll(outputDir, 0755); err != nil {
		return err
	}

	// Определяем имя файла для сохранения страницы
	pageFile := path.Join(outputDir, "index.html")
	file, err := os.Create(pageFile)
	if err != nil {
		return err
	}
	defer file.Close()

	// Копируем содержимое страницы в файл
	_, err = io.Copy(file, resp.Body)
	if err != nil {
		return err
	}

	// Парсим страницу, чтобы найти все ресурсы для загрузки
	doc, err := html.Parse(resp.Body)
	if err != nil {
		return err
	}

	resources := findResources(doc, urlStr)

	// Загружаем все ресурсы
	for _, resource := range resources {
		if err := downloadResource(resource, outputDir); err != nil {
			fmt.Printf("Ошибка загрузки ресурса %s: %v\n", resource, err)
		}
	}

	return nil
}

// Ищет ресурсы на странице (CSS, JS, изображения)
func findResources(n *html.Node, baseURL string) []string {
	var resources []string
	var f func(*html.Node)
	f = func(n *html.Node) {
		if n.Type == html.ElementNode {
			for _, attr := range n.Attr {
				if attr.Key == "src" || attr.Key == "href" {
					resources = append(resources, attr.Val)
				}
			}
		}
		for c := n.FirstChild; c != nil; c = c.NextSibling {
			f(c)
		}
	}
	f(n)
	return resolveRelativeUrls(resources, baseURL)
}

// Преобразует относительные ссылки в абсолютные
func resolveRelativeUrls(resources []string, baseURL string) []string {
	var resolved []string
	base, err := url.Parse(baseURL)
	if err != nil {
		return resolved
	}
	for _, res := range resources {
		u, err := url.Parse(res)
		if err != nil {
			continue
		}
		resolved = append(resolved, base.ResolveReference(u).String())
	}
	return resolved
}

// Загружает ресурс (например, изображение или CSS-файл)
func downloadResource(resourceURL, outputDir string) error {
	resp, err := http.Get(resourceURL)
	if err != nil {
		return err
	}
	defer resp.Body.Close()

	// Определяем имя файла для сохранения ресурса
	u, err := url.Parse(resourceURL)
	if err != nil {
		return err
	}

	filePath := path.Join(outputDir, path.Base(u.Path))
	file, err := os.Create(filePath)
	if err != nil {
		return err
	}
	defer file.Close()

	_, err = io.Copy(file, resp.Body)
	return err
}

func main() {
	if len(os.Args) < 2 {
		fmt.Println("Использование: wget <url>")
		return
	}

	urlStr := os.Args[1]
	outputDir := "./website"

	if err := downloadPage(urlStr, outputDir); err != nil {
		fmt.Printf("Ошибка загрузки: %v\n", err)
	} else {
		fmt.Println("Загрузка завершена.")
	}
}
