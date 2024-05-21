/* This is a helper utility built to regex through the html and modify the locations to where they are downloaded rather than their links */

package webDownloader

import (
	"github.com/alexkefer/p2psearch-backend/fileData"
	"github.com/alexkefer/p2psearch-backend/log"
	"github.com/alexkefer/p2psearch-backend/utils"
	"golang.org/x/net/html"
	"io"
	"net/http"
	"strings"
)

func retrieveAsset(url string) ([]byte, string) {
	// takes in the url and returns the asset
	req, _ := http.NewRequest("GET", url, nil)
	req.Header.Set("User-Agent", "P2PWebCache")

	client := &http.Client{}
	resp, err := client.Do(req)
	if err != nil {
		log.Error("error downloading asset:", url, err)
		return nil, ""
	}

	contentType := resp.Header.Get("Content-Type")

	defer resp.Body.Close()
	body, err := io.ReadAll(resp.Body)
	if err != nil {
		log.Error("error reading asset content:", err)
		return nil, contentType
	}
	return body, contentType
}

func DownloadAllAssets(url, htmlContent string, fileStore *fileData.FileDataStore) string {
	tokenizer := html.NewTokenizer(strings.NewReader(htmlContent))
	modifiedHtml := ""

	for {
		tokenType := tokenizer.Next()
		switch tokenType {
		case html.ErrorToken:
			return modifiedHtml
		case html.StartTagToken, html.SelfClosingTagToken:
			token := tokenizer.Token()
			switch token.Data {
			case "link": // Download CSS
				for i, attr := range token.Attr {
					if attr.Key == "href" {
						rel, ok := getAttributeValue(token, "rel")
						if ok && (rel == "stylesheet") {
							link := url + attr.Val
							log.Debug("retrieving stylesheet asset: " + link)
							content, contentType := retrieveAsset(link)
							if content != nil {
								filename := utils.UrlToFilename(link)
								SaveFile(content, filename, contentType, fileStore)
								attr.Val = "/retrieve?path=" + filename
								token.Attr[i] = attr
							}
						}
					}
				}
			case "script": // Download JS
				for i, attr := range token.Attr {
					if attr.Key == "src" {
						link := url + attr.Val
						log.Debug("retrieving Asset: " + link)
						content, contentType := retrieveAsset(link)
						if content != nil {
							filename := utils.UrlToFilename(link)
							SaveFile(content, filename, contentType, fileStore)
							attr.Val = "/retrieve?path=" + filename
							token.Attr[i] = attr
						}
					}
				}
			case "img": // Download Images
				for i, attr := range token.Attr {
					if attr.Key == "src" {
						link := url + attr.Val
						log.Debug("retrieving Asset: " + link)
						content, contentType := retrieveAsset(link)
						if content != nil {
							filename := utils.UrlToFilename(link)
							SaveFile(content, filename, contentType, fileStore)
							attr.Val = "/retrieve?path=" + filename
							token.Attr[i] = attr
						}
					}
				}
			}

			modifiedHtml += html.UnescapeString(token.String())

		default:
			token := tokenizer.Token()
			modifiedHtml += html.UnescapeString(token.String())
		}
	}
}

/* Helper Functions */

// using the url from the token, it will determine the asset type (css, php, js, img, etc)
func detectAssetType(url string) string {
	if strings.Contains(url, ".css") {
		return "css"
	} else if strings.Contains(url, ".js") {
		return "js"
	} else if strings.Contains(url, ".php") {
		return "php"
	} else if strings.Contains(url, ".jpg") || strings.Contains(url, ".jpeg") || strings.Contains(url, ".png") || strings.Contains(url, ".gif") || strings.Contains(url, ".svg") || strings.Contains(url, ".bmp") || strings.Contains(url, ".webp") || strings.Contains(url, ".ico") {
		return "img"
	} else {
		return "unknown"
	}
}

func getAttributeValue(token html.Token, key string) (string, bool) {
	for _, attr := range token.Attr {
		if attr.Key == key {
			return attr.Val, true
		}
	}
	return "", false
}

func trimLongURL(url string) string {
	// takes in url and returns the trimmed url
	if len(url) > 50 {
		return url[:50]
	}
	return url
}

func buildPageUrl(url string, assetUrl string) string {
	// takes in the source url and the asset url and returns the full url
	if strings.HasPrefix(assetUrl, "http://") || strings.HasPrefix(assetUrl, "https://") {
		return assetUrl
	}
	if strings.HasPrefix(assetUrl, "//") {
		assetUrl = "https:" + assetUrl
	} else if assetUrl[0] == '/' {
		assetUrl = parsePageSource(parsePageLocation(url)) + assetUrl
	} else {
		for i := 0; i < len(assetUrl); i++ {
			if assetUrl[i] == '.' {
				assetUrl = "https://" + assetUrl
				break
			}
			if assetUrl[i] == '/' {
				assetUrl = parsePageSource(parsePageLocation(url)) + assetUrl
				break
			}
		}
	}
	return assetUrl
}

func parsePageSource(url string) string {
	// takes in the url and returns only the https://www.{url}
	for i := 0; i < len(url); i++ {
		if url[i] == '/' {
			return "https://" + url[:i]
		}
	}
	return "https://" + url
}
