package main

import (
	"encoding/json"
	"fmt"
	"net/url"
	"path/filepath"
)

// https://www.pinterest.com/parapeng/
const Domain = "https://www.pinterest.com/"

type RespUser struct {
	InitialReduxState InitialReduxState `json:"initialReduxState"`
}
type InitialReduxState struct {
	Boards map[string]Board `json:"boards"`
}

func (bs *InitialReduxState) GetBoards(name string) []*Board {
	var ret []*Board
	for _, board := range bs.Boards {
		if name == filepath.Base(board.Url) || name == "" {
			ret = append(ret, &board)
		}
	}
	return ret
}

type Board struct {
	ID       string `json:"id"`
	PinCount int    `json:"pin_count"`
	Url      string `json:"url"`
}

type RespPins struct {
	ResourceResponse ResourceResponse `json:"resource_response"`
}
type ResourceResponse struct {
	ResourceData []ResourceData `json:"data"`
}
type ResourceData struct {
	Images map[string]*Image `json:"images"`
}

func (r *ResourceData) GetOrigin() *Image {
	if img, ok := r.Images["orig"]; ok {
		return img
	}
	for _, img := range r.Images {
		return img
	}
	return nil
}

type Image struct {
	Width  int    `json:"width"`
	Height int    `json:"height"`
	Url    string `json:"url"`
}

type PinterestRequest struct {
	SourceUrl string        `json:"-"`
	Data      PinterestData `json:"data"`
}

type PinterestData struct {
	Options PinterestOptions `json:"options"`
}

type PinterestOptions struct {
	BoardId   string   `json:"board_id"`
	BoardUrl  string   `json:"board_url"`
	PageSize  int      `json:"page_size"`
	Bookmarks []string `json:"bookmarks,omitempty"`
}

func (r *PinterestRequest) ToURL(baseURL string) (string, error) {
	// 将结构体转换为 JSON
	jsonBytes, err := json.Marshal(r.Data)
	if err != nil {
		return "", err
	}
	// 构建 URL
	fullURL := fmt.Sprintf("%s?source_url=%s&data=%s", baseURL, url.QueryEscape(r.SourceUrl), url.QueryEscape(string(jsonBytes)))
	return fullURL, nil
}
