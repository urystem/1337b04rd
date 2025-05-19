package main

import (
	"fmt"
	"log/slog"
	"os"
)

func main() {
	slog.Info("dd")

	l := slog.Default()
	l.Info("ssss")
	fmt.Println()

	log := slog.New(slog.NewTextHandler(os.Stdout, nil))
	log.Info("dd")
	fmt.Println()

	ret := slog.With("with my", "ddd")
	ret.Info("dsdfd")
	fmt.Println()

	lret := l.With("with my", "ddd")
	lret.Info("dsdfd")
	fmt.Println()

	withGroup := l.WithGroup("gggggggg")
	withGroup.Info("sdd", "sadsd", "asdsdsd")
	withGroup.Info("sdd", slog.String("field1", "value1"), slog.Int("field2", 42))
	fmt.Println()

	
}
