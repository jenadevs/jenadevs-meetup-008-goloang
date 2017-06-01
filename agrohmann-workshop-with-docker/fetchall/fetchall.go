package main

func main() {
	start := time.Now()
	ch := make(chan string)
	urls := []string {
		"http://www.jena.today",
		"http://www.andreas-grohmann.com",
	}

	for _, url := range urls {
		go fetch(url, ch)
	}

	for range urls {
		fmt.Println(<-ch)
	}
	fmt.Printf("%0.2f elapsed\n", time.Since(start))
}

func fetch(url string, ch chan<- string){
	start := time.Now()
	resp, err := http.Get(url)

	if err != nil {
		ch <- fmt.Sprintf("%s", err)
		return
	}

	n, err := io.Copy(ioutil.Discard, resp.Body)

	if err != nil {
		ch <- fmt.Sprintf("While reading %s: %v", url, err)
		return
	}

	secs := time.Since(start).Seconds()

	TODO ... ADD CODE FROM SLIDES HERE ...

}