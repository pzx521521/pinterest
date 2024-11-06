package pinterest

import (
	"golang.org/x/sync/errgroup"
	"io"
	"log"
	"net/http"
	"os"
	"path/filepath"
)

func DownloadImageMuti(client *http.Client, imageURLs []string, dirName string, poolSize int) error {
	var eg errgroup.Group
	semaphore := make(chan struct{}, poolSize)
	for index, imageUrl := range imageURLs {
		eg.Go(func() error {
			semaphore <- struct{}{}
			log.Printf("%d开始下载:%s ", index, imageUrl)
			defer func() {
				<-semaphore
			}()
			return downloadImageMuti(client, imageUrl, dirName)
		})
	}
	err := eg.Wait()
	if err != nil {
		return err
	}
	return nil
}
func downloadImageMuti(client *http.Client, imgUrl, dirPath string) error {
	filePath := filepath.Join(dirPath, filepath.Base(imgUrl))
	_, err := os.Stat(filePath)
	if !os.IsNotExist(err) {
		return nil
	}

	// 发送 HTTP 请求获取图片内容
	resp, err := client.Get(imgUrl)
	if err != nil {
		return err
	}
	defer resp.Body.Close()
	data, _ := io.ReadAll(resp.Body)
	// 创建文件
	os.WriteFile(filePath, data, 0644)
	return nil
}
