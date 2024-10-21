package main

import (
	"fmt"
	"github.com/gorilla/websocket"
	"log"
	"net/http"
	"os"
)

// Upgrader - глобальный объект для установки соединения WebSocket
var upgrader = websocket.Upgrader{
	ReadBufferSize:  1024,
	WriteBufferSize: 1024,
	CheckOrigin: func(r *http.Request) bool {
		return true // Разрешаем соединения с любого источника
	},
}

func homeHandler(w http.ResponseWriter, r *http.Request) {
	fmt.Println("bebra2")
	// Читаем HTML-файл
	htmlFile, err := os.ReadFile("./templates/index.html")
	if err != nil {
		log.Println(err)
		http.Error(w, "Internal server error", http.StatusInternalServerError)
		return
	}
	// Отправляем HTML-файл клиенту
	w.Header().Set("Content-Type", "text/html; charset=utf-8")
	w.Write(htmlFile)
}

func wsHandler(w http.ResponseWriter, r *http.Request) {
	// Устанавливаем соединение WebSocket
	fmt.Println("bebra")
	conn, err := upgrader.Upgrade(w, r, nil)
	if err != nil {
		log.Println(err)
		return
	}
	defer conn.Close()

	// Пример простого цикла обработки сообщений WebSocket
	for {
		_, message, err := conn.ReadMessage()
		if err != nil {
			log.Println("Ошибка чтения сообщения:", err)
			break
		}
		fmt.Printf("Получено сообщение: %s\n", message)

		// Пример ответа клиенту
		if err := conn.WriteMessage(websocket.TextMessage, []byte("Ответ от сервера")); err != nil {
			log.Println("Ошибка отправки сообщения:", err)
			break
		}
	}
}

func main() {
	// Назначаем обработчики
	http.HandleFunc("/", homeHandler)
	http.HandleFunc("/ws", wsHandler)

	// Запускаем сервер
	log.Fatal(http.ListenAndServe(":8080", nil))
}
