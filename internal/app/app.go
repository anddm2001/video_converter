package app

import (
	"errors"
	"fmt"
	"log"
	"os"
	"runtime"
	"video_filter/internal/config"
	"video_filter/pkg/logger"

	"go.uber.org/zap"
)

type App struct {
	logger *zap.Logger
	config *config.Config
}

func New() *App {
	app := &App{}

	app.logger = logger.New()
	app.config = config.Setup(app.logger)

	return app
}

func (a *App) Run(arg string) error {
	entries, err := os.ReadDir(a.config.ImportPath)
	if err != nil {
		log.Fatal(err)
	}

	os := runtime.GOOS

	switch os {
	case "windows":
		fmt.Println("Выполнение рантайма в ос windows")
		a.logger.Error("Выполнение рантайма в ос windows")
	case "linux":
		fmt.Println("Выполнение рантайма в ос linux")
		a.logger.Error("Выполнение рантайма в ос linux")
	case "darwin":
		fmt.Println("Выполнение рантайма в mac os")
		a.logger.Error("Выполнение рантайма в mac os")

		return errors.New("Нет платформозависимого кода для mac os")
	default:
		fmt.Println("Неопознанная операционная система- не могу продолжить выполнение")
		a.logger.Error("Неопознанная операционная система- не могу продолжить выполнение")

		return errors.New("Неопознанная операционная система- не могу продолжить выполнение")
	}

	for _, e := range entries {
		fmt.Println("Начинаем перекодировку видеофайла: " + e.Name())
		a.logger.Info("Начинаем перекодировку видеофайла: " + e.Name())

		if len(arg) == 0 {
			if os == "windows" {
				a.convertFPSWin(e)
			}

			if os == "linux" {
				a.convertFPSLinux(e)
			}
		} else {
			if arg == "chcodec" {
				var crf string
				if a.config.VideoCodec == "libaom-av1" || a.config.VideoCodec == "libvpx-vp9" {
					crf = "30"
				}

				if a.config.VideoCodec == "libx265" {
					crf = "28"
				}
				if os == "windows" {
					a.changeCodecWindows(e, a.config.VideoCodec, crf)
				}
				if os == "linux" {

					a.changeCodecLinux(e, a.config.VideoCodec, crf)
				}
			}

			if arg == "mobile" {
				if os == "windows" {
					a.changeResolutionMobileWindows(e, a.config.VideoCodec)
				}
				if os == "linux" {

					a.changeResolutionMobileLinux(e, a.config.VideoCodec)
				}
			}

			if arg == "hd" {
				if os == "windows" {
					a.changeResolutionHDWindows(e, a.config.VideoCodec)
				}
				if os == "linux" {

					a.changeResolutionHDLinux(e, a.config.VideoCodec)
				}
			}
		}

		fmt.Println("Закончили перекодировку видеофайла: " + e.Name())
		a.logger.Info("Закончили перекодировку видеофайла: " + e.Name())
	}

	return nil
}
