package main

import (
	"log"
	"os/exec"

	"github.com/fsnotify/fsnotify"
)

func main() {
	log.Println("start : file watcher")
	watcher, err := fsnotify.NewWatcher()
	if err != nil {
		log.Fatal(err)
	}
	defer watcher.Close()

	done := make(chan bool)
	go func() {
		for {
			select {
			case event := <-watcher.Events:
				log.Println("event:", event)
				if event.Op&fsnotify.Write == fsnotify.Write {
					log.Println("modified file:", event.Name)
					cmd := exec.Command("go", "run", "src/main.go")
					err := cmd.Start()
					if err != nil {
						log.Fatal(err)
					}
					log.Printf("restarted %v", cmd.Process.Pid)
				}
			case err := <-watcher.Errors:
				log.Println("error:", err)
			}
		}
	}()

	err = watcher.Add("./src")
	if err != nil {
		log.Fatal(err)
	}
	<-done
}