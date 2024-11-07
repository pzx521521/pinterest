package main

import (
	"flag"
	"github.com/pzx521521/pinterest"
	"log"
	"net/http"
	"os"
	"path/filepath"
)

func main() {
	p := flag.String("p", "", "proxy")
	o := flag.String("o", "./download", "proxy")
	ps := flag.Int("ps", 5, "poolsize")
	flag.Parse()

	args := flag.Args()
	userName := ""
	broadName := ""
	for i, arg := range args {
		if i == 0 {
			userName = arg
		} else {
			broadName = arg
		}
	}
	if userName == "" {
		log.Printf("请输入用户名\n")
		return
	}
	cli := http.DefaultClient
	if *p != "" {
		var err error
		cli, err = pinterest.GetProxyHttpClient(*p)
		if err != nil {
			log.Fatal("%v\n", err)
		}
	}
	boardsResp, err := pinterest.GetBoards(cli, userName)
	if err != nil {
		log.Fatal("%v\n", err)
	}
	boards := boardsResp.InitialReduxState.Boards
	for _, board := range boards {
		log.Printf("board:%s, count:%d\n", board.Url, board.PinCount)
	}
	boardsFilter := boardsResp.InitialReduxState.GetBoards(broadName)
	if len(boardsFilter) == 0 {
		log.Printf("can't find board:%v\n", broadName)
		return
	}
	for _, board := range boardsFilter {
		pins, err := pinterest.GetPins(cli, board, "")
		if err != nil {
			log.Printf("can't find pins in board:%v\n", board.Url)
			continue
		}
		var imgs []string
		for _, imgUrl := range pins.ResourceResponse.ResourceData {
			img := imgUrl.GetOrigin()
			if img != nil {
				imgs = append(imgs, img.Url)
			}
		}
		log.Printf("will download broad:%s, count:%d\n", board.Url, len(imgs))
		savePath, _ := filepath.Abs(*o)
		savePath = filepath.Join(savePath, filepath.Base(board.Url))
		os.MkdirAll(savePath, os.ModePerm)
		pinterest.DownloadImageMuti(cli, imgs, savePath, *ps)
	}
}
