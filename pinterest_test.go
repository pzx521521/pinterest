package main

import (
	"log"
	"testing"
)

func TestGetBoards(t *testing.T) {
	client, _ := GetProxyHttpClient("http://localhost:7897")
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
			boards, err := GetPins(client, val, "")
			if err != nil {
				log.Printf("%v\n", boards)
			}
		}
	}

}
