package dev

import (
	"fmt"
	"goit/utils"
	"log"
	"net"
	"os"
	"os/exec"
	"path/filepath"
	"sync"
	"time"

	"github.com/fsnotify/fsnotify"
	"github.com/google/shlex"
)

func RunDevBackend(configProject utils.ConfigProject) error {
	fmt.Println("Iniciando BackEnd em modo de dev!")

	command, err := shlex.Split(configProject.Run.Dev.InitCommandBackend)
	if err != nil {
		log.Panic(err)
	}

	var backendCmd *exec.Cmd
	var mu sync.Mutex

	startBackend := func() {
		mu.Lock()
		defer mu.Unlock()

		// Mata processo antigo se existir
		if backendCmd != nil && backendCmd.Process != nil {
			backendCmd.Process.Kill()
			backendCmd.Process.Wait() // espera realmente encerrar
			// Espera até a porta realmente liberar
			for i := 0; i < 10; i++ {
				conn, _ := net.DialTimeout("tcp", "localhost:8080", 100*time.Millisecond)
				if conn == nil {
					break // porta livre
				}
				time.Sleep(100 * time.Millisecond)
			}
		}

		backendCmd = exec.Command(command[0], command[1:]...)
		backendCmd.Stdout = os.Stdout
		backendCmd.Stderr = os.Stderr

		if err := backendCmd.Start(); err != nil {
			log.Fatalf("Erro ao iniciar backend: %v", err)
		}
	}

	startBackend()

	if configProject.Run.Dev.HotReloadBackend.Active {
		watcher, err := fsnotify.NewWatcher()
		if err != nil {
			log.Fatal(err)
		}
		defer watcher.Close()

		listenPath := configProject.Run.Dev.HotReloadBackend.ListenPath
		listenPath, err = filepath.Abs(listenPath)
		if err != nil {
			log.Fatalf("Erro ao obter caminho absoluto: %v", err)
		}

		filepath.WalkDir(listenPath, func(path string, d os.DirEntry, err error) error {
			if err != nil {
				return err
			}
			if d.IsDir() {
				return watcher.Add(path)
			}
			return nil
		})

		for event := range watcher.Events {
			if event.Op&(fsnotify.Create|fsnotify.Write) != 0 {
				if !shouldIgnore(event.Name, configProject.Run.Dev.Ignore) {
					fmt.Println("Reiniciando...")
					time.Sleep(400 * time.Millisecond)
					startBackend()
				}
			}
		}
	} else {
		if err := backendCmd.Wait(); err != nil {
			return err
		}
	}

	return nil
}
