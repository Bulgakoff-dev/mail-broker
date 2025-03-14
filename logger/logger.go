package logger

import (
	"fmt"
	"github.com/rs/zerolog"
	"io"
	"os"
	"strings"
	"sync"
	"time"
)

var (
	instance zerolog.Logger
	once     sync.Once
)

func InitLogger(logFile string) *zerolog.Logger {
	// Открываем (или создаём) файл для логов.
	file, err := os.OpenFile(logFile, os.O_CREATE|os.O_WRONLY|os.O_APPEND, 0666)
	if err != nil {
		panic(fmt.Sprintf("Не удалось открыть лог-файл: %v", err))
	}

	// Создаем ConsoleWriter для вывода в консоль с полным отображением уровня, времени и т.д.
	consoleWriter := zerolog.ConsoleWriter{
		Out:        os.Stdout,
		TimeFormat: time.RFC3339,
		FormatLevel: func(i interface{}) string {
			// Выводим уровень в верхнем регистре, например: | INFO |
			return strings.ToUpper(fmt.Sprintf("| %s |", i))
		},
	}

	// Объединяем вывод в файл и консоль.
	multi := io.MultiWriter(consoleWriter, file)

	// Устанавливаем формат времени глобально
	zerolog.TimeFieldFormat = time.RFC3339

	once.Do(func() {
		instance = zerolog.New(multi).With().Timestamp().Logger()
	})

	return &instance
}
func GetLogger() *zerolog.Logger {
	return &instance
}
