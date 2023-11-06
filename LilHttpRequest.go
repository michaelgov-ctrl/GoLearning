package main

import (
	"context"
	"io"
	"net/http"
	"os"
	"time"
)

func getRemoteContent(method string, url string) (int64, error) {
	client := &http.Client{}
	req, err := http.NewRequest(method, url, nil)
	if err != nil {
		return 0, err
	}

	ctx, cancel := context.WithTimeout(context.Background(), 3*time.Second)
	req = req.WithContext(ctx)
	resp, err := client.Do(req)
	cancel()
	if err != nil {
		return 0, err
	}

	stdOut, err := io.Copy(os.Stdout, resp.Body)
	if err != nil {
		return 0, err
	}

	return stdOut, nil
}

func main() {
	getRemoteContent("GET", "https://raw.githubusercontent.com/michaelgov-ctrl/AzureLearning/main/ARMTemplates/DebianVM/azuredeploy.json")
}
