package fetcher

import (
	"bufio"
	"distributed-crawler/config"
	"fmt"
	"io/ioutil"
	"log"
	"net/http"
	"strings"
	"time"

	"golang.org/x/text/encoding/unicode"

	"golang.org/x/net/html/charset"
	"golang.org/x/text/encoding"
	"golang.org/x/text/transform"
)

var rateLimiter = time.Tick(333 * time.Millisecond)

func Fetch(url string) ([]byte, error) {
	<-rateLimiter
	client := &http.Client{}
	newUrl := strings.Replace(url, "http://", "https://", 1)
	request, err := http.NewRequest(http.MethodGet, newUrl, nil)
	if err != nil {
		return nil, err
	}

	request.Header.Set("User-Agent", config.Useragent)

	request.Header.Add("cookie", config.Cookies)

	resp, err := client.Do(request)
	if err != nil {
		return nil, err
	}
	defer resp.Body.Close()

	if resp.StatusCode != http.StatusOK {
		return nil, fmt.Errorf("error: status code %d", resp.StatusCode)
	}

	bodyReader := bufio.NewReader(resp.Body)
	e := determineEncoding(bodyReader)
	utf8Reader := transform.NewReader(bodyReader, e.NewDecoder())
	return ioutil.ReadAll(utf8Reader)
}

func determineEncoding(r *bufio.Reader) encoding.Encoding {
	bytes, err := r.Peek(1024)
	if err != nil {
		log.Printf("Fetcher error: %v", err)
		return unicode.UTF8
	}
	e, _, _ := charset.DetermineEncoding(bytes, "")
	return e
}
