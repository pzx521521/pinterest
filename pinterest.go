package pinterest

import (
	"encoding/json"
	"errors"
	"fmt"
	"io"
	"net/http"
	"net/url"
	"path/filepath"
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
	requestData := PinterestRequest{
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
	fullURL, err := requestData.ToURL(urlApiPrefix)
	if err != nil {
		return nil, err
	}
	req, err := http.NewRequest("GET", fullURL, nil)
	req.Header.Set("x-pinterest-appstate", "background")
	req.Header.Set("x-pinterest-pws-handler", "www/[username]/[slug].js")
	get, err := client.Do(req)
	defer get.Body.Close()
	data, err := io.ReadAll(get.Body)
	if err != nil {
		return nil, err
	}
	if get.StatusCode != 200 {
		return nil, errors.New(string(data))
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

func GetPinsUrl(cli *http.Client, userName string, boardName string) ([]string, error) {
	boardsResp, err := GetBoards(cli, userName)
	if err != nil {
		return nil, err
	}
	boards := boardsResp.InitialReduxState.Boards
	var imgs []string
	for _, board := range boards {
		if filepath.Base(board.Url) == boardName || boardName == "" {
			pins, err := GetPins(cli, &board, "")
			if err != nil {
				return nil, err
			}

			for _, imgUrl := range pins.ResourceResponse.ResourceData {
				img := imgUrl.GetOrigin()
				if img != nil {
					imgs = append(imgs, img.Url)
				}
			}
		}
	}
	return imgs, nil
}
