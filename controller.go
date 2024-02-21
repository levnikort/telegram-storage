package main

import (
	"encoding/json"
	"fmt"
	"io"
	"net/http"

	"github.com/bluele/gcache"
	"github.com/go-martini/martini"
	tgbotapi "github.com/go-telegram-bot-api/telegram-bot-api/v5"
	"github.com/levnikort/telegram-storage/bot"
	"github.com/levnikort/telegram-storage/config"
)

type ResponseFileID struct {
	Ok     bool `json:"ok"`
	Result struct {
		FileSize int    `json:"file_size"`
		FilePath string `json:"file_path"`
	} `json:"result"`
}

type Controller struct {
	bot   *bot.Bot
	cache gcache.Cache
}

func (c *Controller) Download(params martini.Params, res http.ResponseWriter) any {
	fileId := params["file_id"]
	r := ResponseFileID{}

	if filePath, err := c.cache.Get(fileId); err == nil {
		io.Copy(res, c.getFile(filePath.(string), res))
		return nil
	}

	resp, err := http.Get("https://api.telegram.org/bot" + config.Config.TelegramBotToken + "/getFile?file_id=" + fileId)

	if err != nil {
		res.WriteHeader(http.StatusNotFound)
		return nil
	}

	b, _ := io.ReadAll(resp.Body)

	if err := json.Unmarshal([]byte(b), &r); err != nil {
		res.WriteHeader(http.StatusInternalServerError)
		fmt.Println(err)
		return "failed to convert json"
	}

	if !r.Ok {
		res.WriteHeader(http.StatusBadRequest)
		return nil
	}

	c.cache.Set(fileId, r.Result.FilePath)
	io.Copy(res, c.getFile(r.Result.FilePath, res))

	return nil
}

func (c *Controller) Upload(req *http.Request, res http.ResponseWriter, params martini.Params) any {
	var inputMedia interface{}
	fileType := params["file_type"]
	fileName := req.Header.Get("file-Name")
	fileReader := tgbotapi.FileReader{Name: fileName, Reader: req.Body}

	if fileType == "photo" {
		inputMedia = tgbotapi.NewInputMediaPhoto(fileReader)
	} else {
		inputMedia = tgbotapi.NewInputMediaDocument(fileReader)
	}

	metaData, err := c.bot.API.SendMediaGroup(
		tgbotapi.NewMediaGroup(
			int64(config.Config.TelegramChatID),
			[]interface{}{inputMedia},
		),
	)

	if err != nil {
		fmt.Println(err)
		res.WriteHeader(http.StatusBadRequest)
		return nil
	}

	defer req.Body.Close()

	var b []byte
	var errConvert error

	if fileType == "photo" {
		b, errConvert = json.Marshal(metaData[0].Photo)
	} else {
		b, errConvert = json.Marshal(metaData[0].Document)
	}

	if errConvert != nil {
		res.WriteHeader(500)
		return "json convert error"
	}

	return string(b)
}

func (c *Controller) getFile(filePath string, res http.ResponseWriter) io.ReadCloser {
	resp, err := http.Get("https://api.telegram.org/file/bot" + config.Config.TelegramBotToken + "/" + filePath)

	if err != nil {
		res.WriteHeader(http.StatusNotFound)
		return nil
	}

	return resp.Body
}
