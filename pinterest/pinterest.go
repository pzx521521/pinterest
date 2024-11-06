package pinterest

import (
	"encoding/json"
	"errors"
	"fmt"
	"io"
	"net/http"
	"net/url"
	"regexp"
)

func GetProxyHttpClient(proxyUrl string) (*http.Client, error) {
	// 创建代理 URL
	proxyURL, err := url.Parse(proxyUrl)
	if err != nil {
		return nil, err
	}
	// 创建一个带有代理的 Transport
	transport := &http.Transport{
		Proxy: http.ProxyURL(proxyURL),
	}
	// 创建一个带有自定义 Transport 的 Client
	client := &http.Client{
		Transport: transport,
	}
	return client, nil
}

func GetPins(client *http.Client, board *Board, Bookmark string) (*RespPins, error) {
	urlApiPrefix := "https://www.pinterest.com/resource/BoardFeedResource/get/"
	request := PinterestRequest{
		SourceUrl: board.Url,
		Data: PinterestData{
			Options: PinterestOptions{
				BoardId:   board.ID,
				BoardUrl:  board.Url,
				PageSize:  250,
				Bookmarks: []string{Bookmark},
			},
		},
	}
	fullURL, err := request.ToURL(urlApiPrefix)
	if err != nil {
		return nil, err
	}
	get, err := client.Get(fullURL)
	defer get.Body.Close()
	data, err := io.ReadAll(get.Body)
	if err != nil {
		return nil, err
	}
	var respPins RespPins
	err = json.Unmarshal(data, &respPins)
	if err != nil {
		return nil, err
	}
	return &respPins, nil
}

func GetBoards(client *http.Client, userName string) (*RespUser, error) {
	get, err := client.Get(fmt.Sprintf("https://www.pinterest.com/%s/", userName))
	if err != nil {
		return nil, err
	}
	defer get.Body.Close()
	body, err := io.ReadAll(get.Body)
	if err != nil {
		return nil, err
	}
	respUser, err := GetRespUser(body)
	return respUser, nil
}

// <div hidden id="S:3">
// <script id="__PWS_INITIAL_PROPS__" type="application/json">
func GetRespUser(htmlContent []byte) (*RespUser, error) {
	re := regexp.MustCompile(`<script id="__PWS_INITIAL_PROPS__" type="application/json">(.*?)</script>`)
	match := re.FindSubmatch(htmlContent)

	if len(match) < 1 {
		return nil, errors.New("未找到 __PWS_INITIAL_PROPS__ 中的 JSON 数据")

	}
	jsonData := match[1]
	var respUser RespUser
	err := json.Unmarshal(jsonData, &respUser)
	return &respUser, err
}
