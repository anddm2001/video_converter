package app

import (
	"fmt"
	"io/fs"
	"os"
	"os/exec"
)

func (a *App) changeResolutionMobileLinux(e fs.DirEntry, codec string) {

	if len(codec) == 0 {
		codec = "libx265"
	}

	cmd := exec.Command(a.config.FFMPEGPath, "-hide_banner", "-i", a.config.ImportPath+"/"+e.Name(), "-vf", "scale=1280:720", "-c:v", codec, a.config.ExportPath+"/"+e.Name())

	cmd.Stderr = os.Stderr
	cmd.Stdout = os.Stdout

	err := cmd.Run()
	if err != nil {
		fmt.Println("Ошибка при перекодировке видеофайла: " + e.Name())
		a.logger.Info("Ошибка при перекодировке видеофайла: " + e.Name())
	}
}

func (a *App) changeResolutionMobileWindows(e fs.DirEntry, codec string) {

	if len(codec) == 0 {
		codec = "libx265"
	}

	command := fmt.Sprintf("%s -hide_banner -i %s/%s -vf scale=1280:720 -c:v %s %s/%s", a.config.FFMPEGPath, a.config.ImportPath, e.Name(), codec, a.config.ExportPath, e.Name())

	cmd := exec.Command("cmd", "/C", command)

	cmd.Stderr = os.Stderr
	cmd.Stdout = os.Stdout

	err := cmd.Run()
	if err != nil {
		fmt.Println("Ошибка при перекодировке видеофайла: " + e.Name())
		a.logger.Info("Ошибка при перекодировке видеофайла: " + e.Name())
	}
}

func (a *App) changeResolutionHDLinux(e fs.DirEntry, codec string) {

	if len(codec) == 0 {
		codec = "libx265"
	}

	cmd := exec.Command(a.config.FFMPEGPath, "-hide_banner", "-i", a.config.ImportPath+"/"+e.Name(), "-vf", "scale=1920:1080", "-c:v", codec, a.config.ExportPath+"/"+e.Name())

	cmd.Stderr = os.Stderr
	cmd.Stdout = os.Stdout

	err := cmd.Run()
	if err != nil {
		fmt.Println("Ошибка при перекодировке видеофайла: " + e.Name())
		a.logger.Info("Ошибка при перекодировке видеофайла: " + e.Name())
	}
}

func (a *App) changeResolutionHDWindows(e fs.DirEntry, codec string) {

	if len(codec) == 0 {
		codec = "libx265"
	}

	command := fmt.Sprintf("%s -hide_banner -i %s/%s -vf scale=1920:1080 -c:v %s %s/%s", a.config.FFMPEGPath, a.config.ImportPath, e.Name(), codec, a.config.ExportPath, e.Name())

	cmd := exec.Command("cmd", "/C", command)

	cmd.Stderr = os.Stderr
	cmd.Stdout = os.Stdout

	err := cmd.Run()
	if err != nil {
		fmt.Println("Ошибка при перекодировке видеофайла: " + e.Name())
		a.logger.Info("Ошибка при перекодировке видеофайла: " + e.Name())
	}
}
