package main

import (
	"encoding/json"
	"fmt"
	"io"
	"log"
	"net/http"
	"os"
	"regexp"
	"strconv"
	"strings"
	"time"

	"github.com/samber/lo"
	"golang.org/x/net/html"
	"synchronizer/gocn/model"
)

var (
	newsReg = regexp.MustCompile(`(\d\.)(.+?)(https?:\/\/\S+)\s?(https?:\/\/\S+)?`)
	textReg = regexp.MustCompile(`(\d\.)(.+)`)
)

func main() {
	syncTime := time.Now().AddDate(0, -1, 0)
	dir := syncTime.Format("200601")
	syncGoCNNews(dir)
	readme, err := os.OpenFile("./gocn/README.md", os.O_APPEND|os.O_WRONLY, os.ModeAppend)
	if err != nil {
		panic(err)
	}
	defer readme.Close()

	_, err = readme.WriteString("- [" + dir + "](news/" + dir + ".md)\n")
	if err != nil {
		panic(err)
	}
}

func getFileList(dir string) []string {
	url := fmt.Sprintf("https://github.com/gocn/news/tree/master/%s", dir)
	req, err := http.NewRequest(http.MethodGet, url, nil)
	if err != nil {
		panic(err)
	}
	req.Header.Set("Accept", "application/json")

	resp, err := http.DefaultClient.Do(req)
	if err != nil {
		panic(err)
	}
	defer resp.Body.Close()

	body, err := io.ReadAll(resp.Body)
	if err != nil {
		panic(err)
	}

	var getGoCNDirResp model.GetGoCNDirResp
	err = json.Unmarshal(body, &getGoCNDirResp)
	if err != nil {
		panic(err)
	}

	return lo.Map(getGoCNDirResp.Payload.Tree.Items, func(i model.Item, _ int) string {
		return i.Name
	})
}

func syncGoCNNews(dir string) {
	list := getFileList(dir)
	if len(list) == 0 {
		log.Fatalln("File list is empty.")
	}

	newsFile, err := os.Create("./gocn/news/" + dir + ".md")
	if err != nil {
		panic(err)
	}
	defer newsFile.Close()

	for _, f := range list {
		process(dir, f, newsFile)
	}
}

func process(dir, mdFile string, newsFile *os.File) {
	link := fmt.Sprintf("https://raw.githubusercontent.com/gocn/news/master/%s/%s", dir, mdFile)

	resp, err := http.Get(link)
	if err != nil {
		panic(err)
	}
	defer resp.Body.Close()

	doc, err := html.Parse(resp.Body)
	if err != nil {
		panic(err)
	}

	var text string
	var f func(*html.Node)
	f = func(n *html.Node) {
		if n.Type == html.TextNode {
			text += n.Data
		}
		for c := n.FirstChild; c != nil; c = c.NextSibling {
			f(c)
		}
	}
	f(doc)

	_, err = newsFile.WriteString("## " + mdFile[:len(mdFile)-3] + "\n")
	if err != nil {
		panic(err)
	}
	strs := textReg.FindAllString(strings.TrimSpace(text), -1)
	count := 1
	for _, str := range strs {
		submatch := newsReg.FindStringSubmatch(str)
		if len(submatch) == 0 {
			continue
		}
		title := strings.TrimSpace(submatch[2])
		url := strings.TrimSpace(submatch[3])
		_, err := newsFile.WriteString(strconv.Itoa(count) + ". [" + title + "](" + url + ")\n")
		if err != nil {
			panic(err)
		}
		count++
	}
}
