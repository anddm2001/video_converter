package app

import (
	"fmt"
	"io/fs"
	"os"
	"os/exec"
)

// Замедление видео и отключение звуковой дорожки- платформозависимый метод для ОС Windows
func (a *App) convertFPSWin(e fs.DirEntry) {

	command := fmt.Sprintf("%s -hide_banner -i %s/%s -an -vf setpts=2*PTS %s/%s", a.config.FFMPEGPath, a.config.ImportPath, e.Name(), a.config.ExportPath, e.Name())

	cmd := exec.Command("cmd", "/C", command)

	cmd.Stderr = os.Stderr
	cmd.Stdout = os.Stdout

	err := cmd.Run()
	if err != nil {
		fmt.Println("Ошибка при перекодировке видеофайла: " + e.Name())
		a.logger.Info("Ошибка при перекодировке видеофайла: " + e.Name())
	}
}

// Замедление видео и отключение звуковой дорожки- платформозависимый метод для ОС Linux
func (a *App) convertFPSLinux(e fs.DirEntry) {

	cmd := exec.Command(a.config.FFMPEGPath, "-hide_banner", "-i", a.config.ImportPath+"/"+e.Name(), "-an", "-vf", "setpts=2*PTS", a.config.ExportPath+"/"+e.Name())

	cmd.Stderr = os.Stderr
	cmd.Stdout = os.Stdout

	err := cmd.Run()
	if err != nil {
		fmt.Println("Ошибка при перекодировке видеофайла: " + e.Name())
		a.logger.Info("Ошибка при перекодировке видеофайла: " + e.Name())
	}
}
