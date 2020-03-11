/*
Serve is a very simple static file server in go
Usage:
	-p="8100": port to serve on
	-d=".":    the directory of static files to host
Navigating to http://localhost:8100 will display the index.html or directory
listing file.
*/
package main

import (
	"bytes"
	"errors"
	"fart/messages"
	"flag"
	"io/ioutil"
	"log"
	"net/http"
	"os"
	"os/exec"
	"sync"
	"time"
)

// set as LDFLAGS
var (
	Version     = "v0.0.1"
	BuildTarget = "from-source"
	BuildDate   = "cold-build"
)

func main() {
	port := flag.String("p", "8100", "port to serve on")
	file := flag.String("f", "", "the file to provide for download")

	shareWithWorld := flag.Bool("w", false, "share file through ngrok")
	version := flag.Bool("v", false, "show version ouptut")
	debug := flag.Bool("d", false, "show debug output")

	flag.Parse()

	// create debugger
	debugger := messages.Debugger{ON: *debug}

	if *version {
		messages.OutputVersion(Version, BuildTarget, BuildDate, true, 0)
	}

	if *file == "" {
		messages.PrintFileNotProvided(true, 1)
	}

	// read provided file
	data, err := ioutil.ReadFile(*file)
	if err != nil {
		debugger.UnwrapAndPrint(err)
	}

	wg := new(sync.WaitGroup)

	wg.Add(1)

	go func() {
		debugger.Printf("Serving %s on HTTP port: %s\n", *file, *port)

		log.Fatal(http.ListenAndServe(":"+*port, http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {

			w.Header().Set("Content-Disposition", "attachment; filename="+*file)
			w.Header().Set("Content-Type", r.Header.Get("Content-Type"))
			http.ServeContent(w, r, *file, time.Now(), bytes.NewReader(data))
		})))

		wg.Done()
	}()

	if *shareWithWorld {
		debugger.Println("sharing with public via ngrok")
		cmd := exec.Command("ngrok", "http", *port)
		cmd.Stdout = os.Stdout
		cmd.Stderr = os.Stderr

		if err := cmd.Run(); err != nil {

			if errors.Unwrap(err) == exec.ErrNotFound || err == exec.ErrNotFound {
				messages.PrintNgrokNotFound(true, 1)
			}

			log.Fatalf("failed to share serve in ngrok %v", err)
		}
	}

	wg.Wait()

}
