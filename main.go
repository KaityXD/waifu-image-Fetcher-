package main

import (
    "bytes"
    "encoding/json"
    "flag"
    "fmt"
    "io"
    "log"
    "os"
    "path/filepath"
    "sync"

    "github.com/valyala/fasthttp"
)

type WaifuResponse struct {
    URL string `json:"url"`
}

func downloadImage(imageURL, fileName string, wg *sync.WaitGroup) {
    defer wg.Done()

    statusCode, body, err := fasthttp.Get(nil, imageURL)
    if err != nil {
        log.Printf("Error fetching image: %v", err)
        return
    }
    if statusCode != 200 {
        log.Printf("Unexpected status code %d while fetching image", statusCode)
        return
    }

    file, err := os.Create(fileName)
    if err != nil {
        log.Printf("Error creating file: %v", err)
        return
    }
    defer file.Close()

    if _, err := io.Copy(file, bytes.NewReader(body)); err != nil {
        log.Printf("Error writing image data to file: %v", err)
        return
    }

    fmt.Printf("Downloaded image: %s\n", fileName)
}

func fetchWaifu(nsfw bool, savePath string, count int, experimental bool) {
    category := "sfw"
    if nsfw {
        category = "nsfw"
    }
    url := fmt.Sprintf("https://api.waifu.pics/%s/waifu", category)

    if err := os.MkdirAll(savePath, os.ModePerm); err != nil {
        log.Fatalf("Error creating save path: %v", err)
    }

    var wg sync.WaitGroup

    for i := 0; i < count; i++ {
        statusCode, respBody, err := fasthttp.Get(nil, url)
        if err != nil {
            log.Printf("Error fetching waifu URL: %v", err)
            continue
        }
        if statusCode != 200 {
            log.Printf("Unexpected status code %d while fetching waifu URL", statusCode)
            continue
        }

        var waifuResponse WaifuResponse
        if err := json.Unmarshal(respBody, &waifuResponse); err != nil {
            log.Printf("Error parsing response JSON: %v", err)
            continue
        }

        imageURL := waifuResponse.URL
        if imageURL == "" {
            log.Printf("Oops! Couldn't find a waifu for you on attempt %d. Please try again later.", i+1)
            continue
        }

        fileName := filepath.Join(savePath, fmt.Sprintf("waifu_%s_%d.jpg", category, i+1))

        if experimental {
            wg.Add(1)
            go downloadImage(imageURL, fileName, &wg)
        } else {
            downloadImage(imageURL, fileName, nil)
        }
    }

    if experimental {
        wg.Wait()
    }
}

func main() {
    nsfw := flag.Bool("nsfw", false, "Fetch NSFW images.")
    count := flag.Int("count", 1, "Number of waifu images to fetch (default is 1).")
    savePath := flag.String("save-path", "./waifus", "Directory to save images (default is './waifus').")
    experimental := flag.Bool("experimental", false, "Enable experimental multithreaded downloads.")
    flag.Parse()

    fetchWaifu(*nsfw, *savePath, *count, *experimental)
}

