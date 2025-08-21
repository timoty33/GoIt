package watchexec

import (
	"fmt"
	"os"
	"os/exec"
	"runtime"
)

func WatchexecInstalled() bool {
	// Verifica se já está instalado
	if _, err := exec.LookPath("watchexec"); err == nil {
		fmt.Println("✅ Watchexec já está instalado")
		return true
	}

	fmt.Println("⚠️ Watchexec não encontrado. Instalando...")

	goos := runtime.GOOS
	goarch := runtime.GOARCH

	switch goos {
	case "windows":
		// Preferir Scoop (usuário precisa ter o Scoop)
		if _, err := exec.LookPath("scoop"); err == nil {
			cmd := exec.Command("scoop", "install", "watchexec")
			cmd.Stdout = os.Stdout
			cmd.Stderr = os.Stderr
			if err := cmd.Run(); err != nil {
				fmt.Println("❌ Falha ao instalar via Scoop:", err)
				return false
			}
			return true
		} else if _, err := exec.LookPath("cargo"); err == nil {
			// Instalar via Cargo (se Rust disponível)
			cmd := exec.Command("cargo", "install", "watchexec-cli")
			cmd.Stdout = os.Stdout
			cmd.Stderr = os.Stderr
			if err = cmd.Run(); err != nil {
				fmt.Println("❌ Erro ao instalar via Cargo: %w", err)
				return false
			}
			return true
		} else {
			fmt.Printf("❌ Não foi possível instalar automaticamente no Windows (%s/%s). Instale manualmente em: https://github.com/watchexec/watchexec/releases\n", goos, goarch)
		}

	case "linux":
		if _, err := exec.LookPath("apt"); err == nil {
			cmd := exec.Command("sudo", "apt", "install", "-y", "watchexec")
			cmd.Stdout = os.Stdout
			cmd.Stderr = os.Stderr
			if err = cmd.Run(); err != nil {
				fmt.Println("❌ Erro ao instalar via Apt: %w", err)
				return false
			}
			return true
		} else if _, err := exec.LookPath("pacman"); err == nil {
			cmd := exec.Command("sudo", "pacman", "-S", "--noconfirm", "watchexec")
			cmd.Stdout = os.Stdout
			cmd.Stderr = os.Stderr
			if err = cmd.Run(); err != nil {
				fmt.Println("❌ Erro ao instalar via PacMan: %w", err)
				return false
			}
			return true
		} else if _, err := exec.LookPath("cargo"); err == nil {
			cmd := exec.Command("cargo", "install", "watchexec-cli")
			cmd.Stdout = os.Stdout
			cmd.Stderr = os.Stderr
			if err = cmd.Run(); err != nil {
				fmt.Println("❌ Erro ao instalar via Cargo: %w", err)
				return false
			}
			return true
		} else {
			fmt.Printf("❌ Não foi possível instalar automaticamente no Linux (%s/%s). Instale manualmente em: https://github.com/watchexec/watchexec/releases\n", goos, goarch)
		}

	case "darwin": // macOS
		if _, err := exec.LookPath("brew"); err == nil {
			cmd := exec.Command("brew", "install", "watchexec")
			cmd.Stdout = os.Stdout
			cmd.Stderr = os.Stderr
			if err = cmd.Run(); err != nil {
				fmt.Println("❌ Erro ao instalar via Brew: %w", err)
				return false
			}
			return true
		} else if _, err := exec.LookPath("cargo"); err == nil {
			cmd := exec.Command("cargo", "install", "watchexec-cli")
			cmd.Stdout = os.Stdout
			cmd.Stderr = os.Stderr
			if err = cmd.Run(); err != nil {
				fmt.Println("❌ Erro ao instalar via Cargo: %w", err)
				return false
			}
			return true
		} else {
			fmt.Printf("❌ Não foi possível instalar automaticamente no macOS (%s/%s). Instale manualmente em: https://github.com/watchexec/watchexec/releases\n", goos, goarch)
		}

	default:
		fmt.Printf("❌ SO não suportado para instalação automática: %s/%s\n", goos, goarch)
		return false
	}

	return false
}
