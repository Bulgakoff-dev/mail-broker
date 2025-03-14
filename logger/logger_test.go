package logger

import (
	"io/ioutil"
	"strings"
	"testing"
	"time"
)

func TestInitLogger(t *testing.T) {
	logger := InitLogger("app.log")
	logger.Info().Msg("Test message for file and console")

	time.Sleep(1 * time.Second)

	// Читаем содержимое файла логов.
	data, err := ioutil.ReadFile("app.log")
	if err != nil {
		t.Fatalf("Ошибка чтения файла логов: %v", err)
	}

	if !strings.Contains(string(data), "Test message for file and console") {
		t.Errorf("В файле логов не найдено ожидаемое сообщение")
	}
}

func TestGetLogger(t *testing.T) {
	InitLogger("app.log")
	logger := GetLogger()

	logger.Info().Msg("Test message for file and console")
	time.Sleep(1 * time.Second)
	data, err := ioutil.ReadFile("app.log")
	if err != nil {
		t.Fatalf("Error with log read: %v", err)
	}

	if !strings.Contains(string(data), "Test message for file and console") {
		t.Errorf("Log file not consists message")
	}
}
