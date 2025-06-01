package plot2d

import (
	"fmt"
	"os"
	"path/filepath"
	"runtime"

	"golang.org/x/image/font/opentype"
)

// BuscarFuentes busca archivos .ttf en directorios comunes del sistema operativo y verifica si son fuentes válidas
func BuscarFuentes() ([]string, error) {
	// Directorios comunes de fuentes por sistema operativo
	var directoriosFuentes []string
	switch runtime.GOOS {
	case "windows":
		directoriosFuentes = []string{
			filepath.Join(os.Getenv("WINDIR"), "Fonts"),
			filepath.Join(os.Getenv("LOCALAPPDATA"), "Microsoft", "Windows", "Fonts"),
		}
	case "darwin": // macOS
		directoriosFuentes = []string{
			"/System/Library/Fonts",
			"/Library/Fonts",
			filepath.Join(os.Getenv("HOME"), "Library/Fonts"),
		}
	case "linux":
		directoriosFuentes = []string{
			"/usr/share/fonts",
			"/usr/local/share/fonts",
			filepath.Join(os.Getenv("HOME"), ".fonts"),
			filepath.Join(os.Getenv("HOME"), ".local/share/fonts"),
		}
	default:
		return nil, fmt.Errorf("sistema operativo no soportado: %s", runtime.GOOS)
	}

	// Lista para almacenar rutas de fuentes válidas
	var fuentesValidas []string

	// Recorrer cada directorio y buscar archivos .ttf
	for _, dir := range directoriosFuentes {
		err := filepath.Walk(dir, func(path string, info os.FileInfo, err error) error {
			if err != nil {
				return nil // Ignorar errores de acceso a directorios
			}
			if !info.IsDir() && filepath.Ext(path) == ".ttf" {
				// Verificar si el archivo es una fuente TTF válida
				data, err := os.ReadFile(path)
				if err != nil {
					return nil // Ignorar archivos no legibles
				}
				_, err = opentype.Parse(data)
				if err == nil {
					fuentesValidas = append(fuentesValidas, path)
				}
			}
			return nil
		})
		if err != nil {
			fmt.Printf("Error al buscar en %s: %v\n", dir, err)
		}
	}

	if len(fuentesValidas) == 0 {
		return nil, fmt.Errorf("no se encontraron fuentes .ttf válidas")
	}

	return fuentesValidas, nil
}
