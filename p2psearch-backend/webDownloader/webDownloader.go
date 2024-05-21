// Alex Kefer // January 2023 // Package to download webpages to be able to run them locally
// Will include options for all pages or just the page itself and where to save it
// Helper functions will assist in translating the html css and js files

package webDownloader

import (
	"errors"
	"github.com/alexkefer/p2psearch-backend/fileData"
	"github.com/alexkefer/p2psearch-backend/fileTypes"
	"github.com/alexkefer/p2psearch-backend/log"
	"github.com/alexkefer/p2psearch-backend/utils"
	"io"
	"net/http"
	"strings"
)

func CacheResource(url string, fileDataStore *fileData.FileDataStore) error {
	content, contentType, err := downloadResource(url)
	if err != nil {
		log.Warn("error downloading page: %s", err)
		return err
	}

	filename := utils.UrlToFilename(url)

	if strings.Contains(contentType, "text/html") {
		modifiedHtml := DownloadAllAssets(url, string(content), fileDataStore)
		SaveFile([]byte(modifiedHtml), filename, fileTypes.Html, fileDataStore)
		log.Info("cached webpage at %s", url)
	} else {
		SaveFile(content, filename, contentType, fileDataStore)
		log.Info("cached resource at %s", url)
	}

	return nil
}

func downloadResource(url string) ([]byte, string, error) {
	resp, err := http.Get(url)
	if err != nil {
		return nil, "", err
	}

	if resp.StatusCode != 200 {
		log.Warn("couldn't download web page at %s", url)
		return nil, "", errors.New("error getting url: " + url)
	}

	contentType := resp.Header.Get("Content-Type")
	data, err := io.ReadAll(resp.Body)

	defer resp.Body.Close()

	if err != nil {
		return nil, contentType, errors.New("error getting url: " + url)
	}
	// Print the content
	// fmt.Println(content)
	return data, contentType, nil
}
