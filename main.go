package main

import (
	"fmt"
	"os/exec"
)

func main() {
	// Definir el comando y sus argumentos
	cmd := exec.Command("xrandr", "--output", "HDMI-1", "--mode", "1920x1080", "--rate", "144")

	// Ejecutar el comando y manejar errores
	if err := cmd.Run(); err != nil {
		fmt.Println("Error al ejecutar el comando:", err)
		return
	}

}
