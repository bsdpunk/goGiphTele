package main

import "github.com/godbus/dbus"
import (
	//	"context"
	"encoding/json"
	"fmt"

	"github.com/go-telegram-bot-api/telegram-bot-api"
	"log"
	"net/http"
	"os"
	//	"cloud.google.com/go/storage"
	//	"golang.org/x/oauth2/google"
	//	cloudkms "google.golang.org/api/cloudkms/v1"
	//	"google.golang.org/api/iterator"
	//	"google.golang.org/api/option"
	//	"strconv"
)

type Giph struct {
	Data struct {
		Type             string `json:"type"`
		ID               string `json:"id"`
		URL              string `json:"url"`
		Slug             string `json:"slug"`
		BitlyGifURL      string `json:"bitly_gif_url"`
		BitlyURL         string `json:"bitly_url"`
		EmbedURL         string `json:"embed_url"`
		Username         string `json:"username"`
		Source           string `json:"source"`
		Title            string `json:"title"`
		Rating           string `json:"rating"`
		ContentURL       string `json:"content_url"`
		SourceTld        string `json:"source_tld"`
		SourcePostURL    string `json:"source_post_url"`
		IsSticker        int    `json:"is_sticker"`
		ImportDatetime   string `json:"import_datetime"`
		TrendingDatetime string `json:"trending_datetime"`
		Images           struct {
			DownsizedLarge struct {
				Height string `json:"height"`
				Size   string `json:"size"`
				URL    string `json:"url"`
				Width  string `json:"width"`
			} `json:"downsized_large"`
			FixedHeightSmallStill struct {
				Height string `json:"height"`
				Size   string `json:"size"`
				URL    string `json:"url"`
				Width  string `json:"width"`
			} `json:"fixed_height_small_still"`
			Original struct {
				Frames   string `json:"frames"`
				Hash     string `json:"hash"`
				Height   string `json:"height"`
				Mp4      string `json:"mp4"`
				Mp4Size  string `json:"mp4_size"`
				Size     string `json:"size"`
				URL      string `json:"url"`
				Webp     string `json:"webp"`
				WebpSize string `json:"webp_size"`
				Width    string `json:"width"`
			} `json:"original"`
			FixedHeightDownsampled struct {
				Height   string `json:"height"`
				Size     string `json:"size"`
				URL      string `json:"url"`
				Webp     string `json:"webp"`
				WebpSize string `json:"webp_size"`
				Width    string `json:"width"`
			} `json:"fixed_height_downsampled"`
			DownsizedStill struct {
				Height string `json:"height"`
				Size   string `json:"size"`
				URL    string `json:"url"`
				Width  string `json:"width"`
			} `json:"downsized_still"`
			FixedHeightStill struct {
				Height string `json:"height"`
				Size   string `json:"size"`
				URL    string `json:"url"`
				Width  string `json:"width"`
			} `json:"fixed_height_still"`
			DownsizedMedium struct {
				Height string `json:"height"`
				Size   string `json:"size"`
				URL    string `json:"url"`
				Width  string `json:"width"`
			} `json:"downsized_medium"`
			Downsized struct {
				Height string `json:"height"`
				Size   string `json:"size"`
				URL    string `json:"url"`
				Width  string `json:"width"`
			} `json:"downsized"`
			PreviewWebp struct {
				Height string `json:"height"`
				Size   string `json:"size"`
				URL    string `json:"url"`
				Width  string `json:"width"`
			} `json:"preview_webp"`
			OriginalMp4 struct {
				Height  string `json:"height"`
				Mp4     string `json:"mp4"`
				Mp4Size string `json:"mp4_size"`
				Width   string `json:"width"`
			} `json:"original_mp4"`
			FixedHeightSmall struct {
				Height   string `json:"height"`
				Mp4      string `json:"mp4"`
				Mp4Size  string `json:"mp4_size"`
				Size     string `json:"size"`
				URL      string `json:"url"`
				Webp     string `json:"webp"`
				WebpSize string `json:"webp_size"`
				Width    string `json:"width"`
			} `json:"fixed_height_small"`
			FixedHeight struct {
				Height   string `json:"height"`
				Mp4      string `json:"mp4"`
				Mp4Size  string `json:"mp4_size"`
				Size     string `json:"size"`
				URL      string `json:"url"`
				Webp     string `json:"webp"`
				WebpSize string `json:"webp_size"`
				Width    string `json:"width"`
			} `json:"fixed_height"`
			DownsizedSmall struct {
				Height  string `json:"height"`
				Mp4     string `json:"mp4"`
				Mp4Size string `json:"mp4_size"`
				Width   string `json:"width"`
			} `json:"downsized_small"`
			Preview struct {
				Height  string `json:"height"`
				Mp4     string `json:"mp4"`
				Mp4Size string `json:"mp4_size"`
				Width   string `json:"width"`
			} `json:"preview"`
			FixedWidthDownsampled struct {
				Height   string `json:"height"`
				Size     string `json:"size"`
				URL      string `json:"url"`
				Webp     string `json:"webp"`
				WebpSize string `json:"webp_size"`
				Width    string `json:"width"`
			} `json:"fixed_width_downsampled"`
			FixedWidthSmallStill struct {
				Height string `json:"height"`
				Size   string `json:"size"`
				URL    string `json:"url"`
				Width  string `json:"width"`
			} `json:"fixed_width_small_still"`
			FixedWidthSmall struct {
				Height   string `json:"height"`
				Mp4      string `json:"mp4"`
				Mp4Size  string `json:"mp4_size"`
				Size     string `json:"size"`
				URL      string `json:"url"`
				Webp     string `json:"webp"`
				WebpSize string `json:"webp_size"`
				Width    string `json:"width"`
			} `json:"fixed_width_small"`
			OriginalStill struct {
				Height string `json:"height"`
				Size   string `json:"size"`
				URL    string `json:"url"`
				Width  string `json:"width"`
			} `json:"original_still"`
			FixedWidthStill struct {
				Height string `json:"height"`
				Size   string `json:"size"`
				URL    string `json:"url"`
				Width  string `json:"width"`
			} `json:"fixed_width_still"`
			Looping struct {
				Mp4     string `json:"mp4"`
				Mp4Size string `json:"mp4_size"`
			} `json:"looping"`
			FixedWidth struct {
				Height   string `json:"height"`
				Mp4      string `json:"mp4"`
				Mp4Size  string `json:"mp4_size"`
				Size     string `json:"size"`
				URL      string `json:"url"`
				Webp     string `json:"webp"`
				WebpSize string `json:"webp_size"`
				Width    string `json:"width"`
			} `json:"fixed_width"`
			PreviewGif struct {
				Height string `json:"height"`
				Size   string `json:"size"`
				URL    string `json:"url"`
				Width  string `json:"width"`
			} `json:"preview_gif"`
			Four80WStill struct {
				URL    string `json:"url"`
				Width  string `json:"width"`
				Height string `json:"height"`
			} `json:"480w_still"`
		} `json:"images"`
		ImageOriginalURL             string `json:"image_original_url"`
		ImageURL                     string `json:"image_url"`
		ImageMp4URL                  string `json:"image_mp4_url"`
		ImageFrames                  string `json:"image_frames"`
		ImageWidth                   string `json:"image_width"`
		ImageHeight                  string `json:"image_height"`
		FixedHeightDownsampledURL    string `json:"fixed_height_downsampled_url"`
		FixedHeightDownsampledWidth  string `json:"fixed_height_downsampled_width"`
		FixedHeightDownsampledHeight string `json:"fixed_height_downsampled_height"`
		FixedWidthDownsampledURL     string `json:"fixed_width_downsampled_url"`
		FixedWidthDownsampledWidth   string `json:"fixed_width_downsampled_width"`
		FixedWidthDownsampledHeight  string `json:"fixed_width_downsampled_height"`
		FixedHeightSmallURL          string `json:"fixed_height_small_url"`
		FixedHeightSmallStillURL     string `json:"fixed_height_small_still_url"`
		FixedHeightSmallWidth        string `json:"fixed_height_small_width"`
		FixedHeightSmallHeight       string `json:"fixed_height_small_height"`
		FixedWidthSmallURL           string `json:"fixed_width_small_url"`
		FixedWidthSmallStillURL      string `json:"fixed_width_small_still_url"`
		FixedWidthSmallWidth         string `json:"fixed_width_small_width"`
		FixedWidthSmallHeight        string `json:"fixed_width_small_height"`
		Caption                      string `json:"caption"`
	} `json:"data"`
	Meta struct {
		Status     int    `json:"status"`
		Msg        string `json:"msg"`
		ResponseID string `json:"response_id"`
	} `json:"meta"`
}

func main() {

	tkey := os.Getenv("TKEY")
	gkey := os.Getenv("GKEY")
	bot, err := tgbotapi.NewBotAPI(tkey)
	if err != nil {
		log.Panic(err)
	}

	bot.Debug = true

	log.Printf("Authorized on account %s", bot.Self.UserName)

	u := tgbotapi.NewUpdate(0)
	u.Timeout = 60

	updates, err := bot.GetUpdatesChan(u)

	for update := range updates {
		if update.Message == nil { // ignore any non-Message Updates
			continue
		}

		log.Printf("[%s] %s", update.Message.From.UserName, update.Message.Text)

		msg := tgbotapi.NewMessage(update.Message.Chat.ID, update.Message.Text)
		msg.ReplyToMessageID = update.Message.MessageID
		tag := update.Message.Text

		resp, _ := http.Get("https://api.giphy.com/v1/gifs/random?api_key=" + gkey + "&tag=" + tag + "&rating=r")
		fmt.Println(resp)
		var in Giph

		json.NewDecoder(resp.Body).Decode(&in)
		fmt.Println("=================")
		fmt.Println(in.Data.Images.DownsizedLarge.URL)

		conn, err := dbus.ConnectSessionBus()
		if err != nil {
			panic(err)
		}
		defer conn.Close()

		obj := conn.Object("org.freedesktop.Notifications", "/org/freedesktop/Notifications")
		call := obj.Call("org.freedesktop.Notifications.Notify", 0, "", uint32(0),
			"", "Applause Gif", in.Data.Images.DownsizedLarge.URL, []string{},
			map[string]dbus.Variant{}, int32(5000))

		if call.Err != nil {
			panic(call.Err)
		}

		bot.Send(msg)
		bot.Send(tgbotapi.NewMessage(update.Message.Chat.ID, in.Data.Images.DownsizedLarge.URL))
	}

	//explicit("/home/dusty/gkey.json", "ardent-strength-215718")
}
