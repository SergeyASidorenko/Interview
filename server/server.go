// Copyright 2020 The Home. All rights not reserved.
// Пакет с реализацией тестового задание на простейший веб-сервер
// Сервер получает целочисленный параметр от пользователя только через метод POST, находить остаток от деления на 2
// и возвращает ответ в виде быстрого хеширования md5

// Использован пакет gin-gonic
// Сведения о лицензии отсутствуют

package main

import (
	"crypto/md5"
	"encoding/hex"
	"net/http"
	"strconv"
	"strings"

	"github.com/gin-gonic/gin"
)

type Data struct {
	Id   string `json:"id"`
	Text string `json:"text"`
}

func main() {
	r := gin.Default()
	r.POST("/md5", handlePostMD5)
	r.Run(":8020")
}
func handlePostMD5(cont *gin.Context) {
	var data Data
	if cont.BindJSON(&data) == nil {

		idi, _ := strconv.Atoi(data.Id)
		text := data.Text
		//Простейшая проверка параметров тела запроса
		if (idi < 0) || (len(strings.Trim(text, " ")) == 0) || (len([]rune(text)) > 400) {
			errorMessage := "Неверный запрос! Проверьте параметры!"
			//Ответ 400, если параметры неверны
			cont.String(http.StatusBadRequest, errorMessage)
			return
		}
		//Остаток от деления на 2
		remainderi := idi % 2
		ids := strconv.Itoa(idi)
		remainders := strconv.Itoa(remainderi)
		message := ids + text + remainders
		//Массив из 16 элементов, представляющий нашу свертку
		hash := md5.Sum([]byte(message))
		//Составление ответа в виде строки md5+хеш
		response := "md5" + hex.EncodeToString(hash[:])
		cont.String(http.StatusOK, response)
	}
}

//Вспомогательная функция для тестирования
func GetServer() *gin.Engine {
	r := gin.Default()
	r.POST("/md5", handlePostMD5)
	r.Run(":8020")
	return r
}
