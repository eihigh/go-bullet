package main

import (
	"flag"
	"log"
	"os"
	"os/exec"
	"path/filepath"
	"strings"
	"time"

	"github.com/fsnotify/fsnotify"
)

// example:
// $ go run ./watch -build watch/build.sh -run watch/run.sh

var (
	buildcmd = flag.String("build", "", "build command")
	runcmd   = flag.String("run", "", "run command")
)

type empty = struct{}

func main() {
	log.SetFlags(log.Ltime | log.Lmicroseconds)
	flag.Parse()
	if *buildcmd == "" {
		log.Fatal("build command (-build) is required")
	}
	if *runcmd == "" {
		log.Fatal("run command (-run) is required")
	}

	changed := make(chan empty, 1)
	built := make(chan empty, 1)
	go build(changed, built)
	go run(built)
	watch(changed)
}

func watch(changed chan<- empty) {
	w, err := fsnotify.NewWatcher()
	if err != nil {
		log.Fatal(err)
	}
	defer w.Close()

	// Walk the path to add all subdirectories to the watcher
	err = filepath.Walk(".", func(path string, info os.FileInfo, err error) error {
		if err != nil {
			return err
		}
		if info.IsDir() {
			return w.Add(path)
		}
		return nil
	})
	if err != nil {
		log.Fatal(err)
	}

	trySend(changed, empty{})
	log.Println("watch started")

	t := time.Now()
	for {
		select {
		case event, ok := <-w.Events:
			if !ok {
				log.Println("watch stopped")
				return
			}
			// if strings.HasPrefix(filepath.Base(event.Name), ".") {
			// 	break // skip hidden files
			// }
			if !strings.HasSuffix(event.Name, ".go") {
				break
			}
			if time.Now().Sub(t) > time.Second {
				log.Println("changed:", event.Name)
				t = time.Now()
				trySend(changed, empty{})
			}

		case err, ok := <-w.Errors:
			if !ok {
				log.Println("watch stopped")
				return
			}
			log.Println("watch failed:", err)
		}
	}
}

func build(changed <-chan empty, built chan<- empty) {
	for range changed {
		log.Println("build triggered")
		if err := execute(*buildcmd); err != nil {
			log.Println("build failed:", err)
		} else {
			log.Println("build succeeded")
		}
		trySend(built, empty{})
	}
}

func run(built <-chan empty) {
	<-built
	t := time.NewTimer(time.Second)
	for {
		log.Println("run triggered")
		t.Reset(time.Second)
		if err := execute(*runcmd); err != nil {
			log.Println("run failed:", err)
		}
		<-t.C
	}
}

func trySend[T any](ch chan<- T, v T) {
	select {
	case ch <- v:
	default:
	}
}

func execute(cmd string, args ...string) error {
	c := exec.Command(cmd, args...)
	c.Stdout = os.Stdout
	c.Stderr = os.Stderr
	return c.Run()
}
