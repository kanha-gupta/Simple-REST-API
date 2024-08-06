package main

import (
	"encoding/json"
	"fmt"
	"net/http"
	"os/exec"
)

type CmdReq struct {
	Command string `json:"command"`
}

type CmdRes struct {
	Output string `json:"output"`
	Error  string `json:"error,omitempty"`
}

func main() {
	http.HandleFunc("/api/cmd", handler)
	fmt.Print("Server started at :8080")
	http.ListenAndServe(":8080", nil)
}

func handler(w http.ResponseWriter, r *http.Request) {
	switch r.Method {
	case http.MethodPost:
		execCmd(w, r)
	default:
		http.Error(w, "Method not allowed", http.StatusMethodNotAllowed)
	}
}

func execCmd(w http.ResponseWriter, r *http.Request) {
	var cmdReq CmdReq

	if r.Header.Get("Content-Type") == "application/json" {
		if err := json.NewDecoder(r.Body).Decode(&cmdReq); err != nil {
			http.Error(w, "Invalid request body", http.StatusBadRequest)
			return
		}
	} else {
		cmdReq.Command = r.URL.Query().Get("command")
		if cmdReq.Command == "" {
			http.Error(w, "Command not provided", http.StatusBadRequest)
			return
		}
	}

	cmd := exec.Command("sh", "-c", cmdReq.Command)
	output, err := cmd.CombinedOutput()
	exitError, ok := err.(*exec.ExitError)

	if err != nil {
		if ok && exitError.ExitCode() == 127 {
			errMessage := fmt.Sprintf("Error: Command not found (status code: %d)", http.StatusNotFound)
			http.Error(w, errMessage, http.StatusNotFound)
		} else {
			errMessage := fmt.Sprintf("Error: %s (status code: %d)", string(output), http.StatusBadRequest)
			http.Error(w, errMessage, http.StatusBadRequest)
		}
		return
	}

	cmdRes := CmdRes{Output: string(output)}
	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(cmdRes)
}
