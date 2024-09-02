package server

import (
	"net/http"
	"net/http/httptest"
	"testing"

	"github.com/gorilla/websocket"
)

func TestConnectionErrorHandling(t *testing.T) {
	//Arrange
	server := httptest.NewServer(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		http.Error(w, "not allowed", http.StatusForbidden)
	}))

	defer server.Close()

	// Act
	wsURL := "ws" + server.URL[len("http"):]
	_, _, err := websocket.DefaultDialer.Dial(wsURL, nil)

	// Assert
	if err == nil {
		t.Fatalf("Expected an error when connecting to WebSocket but none occurred")
	}

}

func TestSendMessage(t *testing.T) {
	// Arrange
	server := httptest.NewServer(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		upgrader := websocket.Upgrader{}
		conn, err := upgrader.Upgrade(w, r, nil)
		if err != nil {
			t.Fatalf("Error upgrading to WebSocket: %v", err)
		}

		_, message, err := conn.ReadMessage()
		if err != nil {
			t.Fatalf("Error reading message: %v", err)
		}

		// Assert
		expectedMessage := "Hello"
		if string(message) != expectedMessage {
			t.Fatalf("Expected %s but received %s", expectedMessage, message)
		}
	}))
	defer server.Close()

	// Act
	wsURL := "ws" + server.URL[len("http"):]
	ws, _, err := websocket.DefaultDialer.Dial(wsURL, nil)
	if err != nil {
		t.Fatalf("Error connecting to WebSocket: %v", err)
	}
	defer ws.Close()

	err = ws.WriteMessage(websocket.TextMessage, []byte("Hello"))
	if err != nil {
		t.Fatalf("Error sending message: %v", err)
	}
}

func TestWebSocketConnection(t *testing.T) {
	// Arrange
	server := httptest.NewServer(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		upgrader := websocket.Upgrader{}
		_, err := upgrader.Upgrade(w, r, nil)
		if err != nil {
			t.Fatalf("Error upgrading to WebSocket: %v", err)
		}
	}))
	defer server.Close()

	// Act
	wsURL := "ws" + server.URL[len("http"):]
	_, _, err := websocket.DefaultDialer.Dial(wsURL, nil)

	// Assert
	if err != nil {
		t.Fatalf("Error connecting to WebSocket: %v", err)
	}
}
