package main

import (
	"log/slog"
	"os"
)

func main() {
	// Создаём свой логгер с текстовым обработчиком и кастомными опциями
	slog.Info("Это сообщение с старом дефолтным логгером")
	myLogger := slog.New(slog.NewTextHandler(os.Stdout, &slog.HandlerOptions{
		Level: slog.LevelDebug,
	}))

	// Устанавливаем его как дефолтный логгер
	slog.SetDefault(myLogger)

	// Теперь slog.Default() вернёт myLogger
	slog.Info("Это сообщение с новым дефолтным логгером")
}
