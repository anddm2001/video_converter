package app

import (
	"fmt"
	"io/fs"
	"os"
	"os/exec"
	"strings"
)

// Сжатие видео посредством упаковки другим кодеком- строго говоря необязательно "сжатие", скорее просто перекодирование с другим кодеком
// Платформозависимый метод для ОС Windows
func (a *App) changeCodecWindows(e fs.DirEntry, codec string, crf string) {

	if len(codec) == 0 {
		codec = "libx265"
	}

	var command string

	if codec == "libvpx-vp9" {
		filename := strings.Replace(e.Name(), ".MP4", ".webm", -1)
		command = fmt.Sprintf("%s -hide_banner -i %s/%s -vcodec %s -b:v 0 -crf %s %s/%s", a.config.FFMPEGPath, a.config.ImportPath, e.Name(), codec, crf, a.config.ExportPath, filename)
	} else {
		command = fmt.Sprintf("%s -hide_banner -i %s/%s -vcodec %s -crf %s %s/%s", a.config.FFMPEGPath, a.config.ImportPath, e.Name(), codec, crf, a.config.ExportPath, e.Name())
	}

	cmd := exec.Command("cmd", "/C", command)

	cmd.Stderr = os.Stderr
	cmd.Stdout = os.Stdout

	err := cmd.Run()
	if err != nil {
		fmt.Println("Ошибка при перекодировке видеофайла: " + e.Name())
		a.logger.Info("Ошибка при перекодировке видеофайла: " + e.Name())
	}
}

// Сжатие видео посредством упаковки другим кодеком- строго говоря необязательно "сжатие", скорее просто перекодирование с другим кодеком
// Платформозависимый метод для ОС linux
func (a *App) changeCodecLinux(e fs.DirEntry, codec string, crf string) {

	if len(codec) == 0 {
		codec = "libx265"
	}

	var cmd *exec.Cmd

	if codec == "libvpx-vp9" {
		filename := strings.Replace(e.Name(), ".MP4", ".webm", -1)
		cmd = exec.Command(a.config.FFMPEGPath, "-hide_banner", "-i", a.config.ImportPath+"/"+e.Name(), "-vcodec", codec, "-crf", crf, "-b:v", "0", a.config.ExportPath+"/"+filename)
	} else {
		cmd = exec.Command(a.config.FFMPEGPath, "-hide_banner", "-i", a.config.ImportPath+"/"+e.Name(), "-vcodec", codec, "-crf", crf, a.config.ExportPath+"/"+e.Name())
	}

	cmd.Stderr = os.Stderr
	cmd.Stdout = os.Stdout

	err := cmd.Run()
	if err != nil {
		fmt.Println("Ошибка при перекодировке видеофайла: " + e.Name())
		a.logger.Info("Ошибка при перекодировке видеофайла: " + e.Name())
	}
}
