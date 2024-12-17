package pinterest

import (
	"log"
	"os"
	"testing"
)

func TestGetBoards(t *testing.T) {
	//client, _ := GetProxyHttpClient("http://localhost:7897")
	client, _ := GetProxyHttpClient("http://localhost:8888")
	boards, err := GetBoards(client, "parapeng")
	if err != nil {
		log.Printf("%v\n", boards)
	}
}
func TestGetPins(t *testing.T) {
	client, _ := GetProxyHttpClient("http://localhost:7897")
	boards, err := GetBoards(client, "parapeng")
	if err != nil {
		log.Printf("%v\n", boards)
	}

	for key, val := range boards.InitialReduxState.Boards {
		if key == "946107902908880006" {
			pins, err := GetPins(client, &val, "")
			if err != nil {
				log.Printf("%v\n", pins)
			}
		}
	}
}
func TestGetPinsUrl(t *testing.T) {
	client, _ := GetProxyHttpClient("http://localhost:7897")
	imgs, err := GetPinsUrl(client, "parapeng", "wallpaper2")
	if err != nil {
		log.Printf("%v\n", err)
		return
	}
	log.Printf("%v\n", imgs)
	dir := "/Users/parapeng/Downloads/wait"
	os.MkdirAll(dir, os.ModePerm)
	DownloadImageMuti(client, imgs, dir, 10)
}
