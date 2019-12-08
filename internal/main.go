package main

import (
	"flag"
	"github.com/BurntSushi/toml"
	"github.com/gin-gonic/gin"
	"log"
	"math/rand"
	"net/http"
	"time"
	"url_short/internal/conf"
	"url_short/internal/service"
)

const (
	_contentTypeTextPlain = "text/plain"
)

var config conf.Config
var svc service.Service

func main() {
	argConf := flag.String("conf", "config.toml", "-conf config.toml")

	flag.Parse()

	log.Println("server starting...")
	rand.Seed(time.Now().UnixNano())
	if _, err := toml.DecodeFile(*argConf, &config); err != nil {
		panic(err)
	}

	svc = service.NewService(&config)

	server := gin.Default()

	server.Handle(http.MethodGet, "/:key", handleGet)
	server.Handle(http.MethodPost, "/new", handleNew)

	if err := server.Run(config.Options.HttpAddr...); err != nil {
		panic(err)
	}
	log.Println("server stop.")
}

func handleGet(c *gin.Context) {
	url, err := svc.GetUrl(c.Param("key"))
	if err != nil {
		if err == service.ErrorNotExists {
			c.Data(http.StatusNotFound, _contentTypeTextPlain, []byte("not found."))
		} else {
			c.Data(http.StatusServiceUnavailable, _contentTypeTextPlain, []byte(err.Error()))
		}
		c.Abort()
		return
	}
	c.Redirect(http.StatusFound, url)
	c.Abort()
	return
}
func handleNew(c *gin.Context) {
	var body struct {
		Token  string `json:"token"`
		Expire int64  `json:"expire"`
		Url    string `json:"url"`
		KeyLen int    `json:"key_len"`
	}
	err := c.BindJSON(&body)
	if err != nil {
		c.Data(http.StatusBadRequest, _contentTypeTextPlain, []byte(err.Error()))
		c.Abort()
		return
	}
	if body.Expire <= 0 {
		body.Expire = -1
	}
	if body.KeyLen <= 0 {
		body.KeyLen = config.Options.KeyLength
	}
	access := false
	for _, token := range config.Options.Security.Token {
		if token == body.Token {
			access = true
		}
	}
	if !access {
		c.Data(http.StatusUnauthorized, _contentTypeTextPlain, []byte("unauthorized"))
		c.Abort()
		return
	}

	for i := 0; i < config.Options.RetryCount; i++ {
		key, err := svc.SaveUrl(body.KeyLen, body.Url, time.Duration(body.Expire)*time.Second)
		if err != nil {
			if err == service.ErrorExists {
				continue
			}
			c.Data(http.StatusBadRequest, _contentTypeTextPlain, []byte(err.Error()))
			c.Abort()
			return
		}
		c.JSON(http.StatusOK, map[string]interface{}{
			"short_url": key,
		})
		c.Abort()
		return
	}
	c.Data(http.StatusConflict, _contentTypeTextPlain, []byte("key exists"))
	c.Abort()
	return
}
