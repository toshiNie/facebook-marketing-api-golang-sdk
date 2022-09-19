package main

import (
	"context"
	"fmt"
	"github.com/toshiNie/facebook-marketing-api-golang-sdk/fb"
	v14 "github.com/toshiNie/facebook-marketing-api-golang-sdk/marketing/v14"
)

const (
	_accessToken  = "{{access_token}}"
	_clientSecret = "{{client_secret}}"
	_accountID    = "{{account_id}}"
)

func main() {
	client, err := v14.New(fb.NewClient(_accessToken, _clientSecret))
	if err != nil {
		panic(err)
	}
	vc := make(chan v14.Video, 1024)
	client.Videos.ReadList(context.Background(), _accountID, vc)
	close(vc)
	for video := range vc {
		fmt.Println(video.ID)
	}
}
