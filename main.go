package main

import (
	"bytes"
	"fmt"
	"net/http"
	"os/exec"
)

func main() {
	http.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) {
		var out0, out1, out2 bytes.Buffer
		cmd0 := exec.Command("hostname")
		cmd1 := exec.Command("free", "-m")
		cmd2 := exec.Command("uptime")
		cmd0.Stdout = &out0
		cmd1.Stdout = &out1
		cmd2.Stdout = &out2
		cmd0.Run()
		cmd1.Run()
		cmd2.Run()
		fmt.Fprintln(w, "<!DOCTYPE html><html><head><title>Runner Status</title></head><body><pre>", out0.String(), out1.String(), out2.String(), "</pre></body>")
	})

	http.ListenAndServe(":8765", nil)
	fmt.Println("Running and serving on :8765")
}
