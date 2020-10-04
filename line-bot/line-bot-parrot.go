package main

import (
	"encoding/json"
	"fmt"
	"github.com/aws/aws-lambda-go/events"
	"github.com/aws/aws-lambda-go/lambda"
	"github.com/line/line-bot-sdk-go/linebot"
	"log"
	"os"
)

func main() {
	// Lambda関数のエントリポイント
	lambda.Start(HandleRequest)
}

func HandleRequest(request events.APIGatewayProxyRequest) (events.APIGatewayProxyResponse, error) {

	// LINEプラットフォームで生成されたWebhookオブジェクト
	fmt.Println("=== Body ===")
	fmt.Println(request.Body)

	// JSONをデコード
	fmt.Println("=== JSON decode ===")
	myLineRequest, err := UnmarshalLineRequest([]byte(request.Body))
	if err != nil {
		log.Fatal(err)
	}

	// ボットの定義
	fmt.Println("=== linebot new ===")
	bot, err := linebot.New(
		os.Getenv("LINE_CHANNEL_SECRET"),
		os.Getenv("LINE_CHANNEL_TOKEN"),
	)
	if err != nil {
		log.Fatal(err)
	}

	// 応答メッセージの生成
	fmt.Println("=== reply ===")
	replyMessage := myLineRequest.Events[0].Message.Text
	if _, err = bot.ReplyMessage(myLineRequest.Events[0].ReplyToken, linebot.NewTextMessage(replyMessage)).Do(); err != nil {
		log.Fatal(err)
	}

	// 終了
	fmt.Println("=== end ===")
	return events.APIGatewayProxyResponse{Body: request.Body, StatusCode: 200}, nil

	// events, err := bot.ParseRequest(request)
	// if err != nil {
	// 	status := 200
	// 	if err == bot.ErrInvalidSignature {
	// 		status = 400
	// 	} else {
	// 		status = 500
	// 	}
	// 	return events.APIGatewayProxyResponse{StatusCode: status}, errors.New("Bat Request")
	// }

	// for _, event := range events {
	// 	if event.Type == linebot.EventTypeMessage {
	// 		switch message := event.Message.(type) {
	// 		case *linebot.TextMessage:
	// 			if _, err = bot.ReplyMessage(event.ReplyToken, linebot.NewTextMessage(message.Text)).Do(); err != nil {
	// 				log.Print(err)
	// 			}
	// 		case *linebot.StickerMessage:
	// 			replyMessage := fmt.Sprintf(
	// 				"sticker id is %s, stickerResourceType is %s", message.StickerID, message.StickerResourceType)
	// 			if _, err = bot.ReplyMessage(event.ReplyToken, linebot.NewTextMessage(replyMessage)).Do(); err != nil {
	// 				log.Print(err)
	// 			}
	// 		}
	// 	}
	// }
}

// API Gatewayから受け取ったevents.APIGatewayProxyResponseのBody（JSON）をParseする
func UnmarshalLineRequest(data []byte) (LineRequest, error) {
	var r LineRequest
	err := json.Unmarshal(data, &r)
	return r, err
}

func (r *LineRequest) Marshal() ([]byte, error) {
	return json.Marshal(r)
}

type LineRequest struct {
	Destination string  `json:"destination"`
	Events      []Event `json:"events"`
}

type Event struct {
	ReplyToken string   `json:"replyToken"`
	Type       string   `json:"type"`
	Mode       string   `json:"mode"`
	Timestamp  int64    `json:"timestamp"`
	Source     Source   `json:"source"`
	Message    *Message `json:"message,omitempty"`
}

type Message struct {
	ID   string `json:"id"`
	Type string `json:"type"`
	Text string `json:"text"`
}

type Source struct {
	Type   string `json:"type"`
	UserID string `json:"userId"`
}
