
func main() {
	client := &http.Client{}
	req, err := http.NewRequest("GET", "http://www.golang.org", nil)
	if err != nil {
		fmt.Println("error: ", err)
		return
	}

	ctx, cancel := context.WithTimeout(context.Background(), 3*time.Second)

	req = req.WithContext(ctx)

	resp, err := client.Do(req)
	cancel()
	if err != nil {
		fmt.Println("error: ", err)
		return
	}

	io.Copy(os.Stdout, resp.Body)
}
