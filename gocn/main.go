package main

import (
	"fmt"
	"log"
	"net/http"
	"os"
	"os/exec"
	"regexp"
	"strconv"
	"strings"
	"time"

	"golang.org/x/net/html"
)

var (
	newsReg = regexp.MustCompile(`(\d\.)(.+)(https?:\/\/\S+)`)
	textReg = regexp.MustCompile(`(\d\.)(.+)`)
)

func main() {
	time := syncGoCNNews()
	if time == "" {
		return
	}
	exec.Command("sync.sh", time)
}

func getFileList(dir string) []string {
	url := fmt.Sprintf("https://github.com/gocn/news/tree/master/%s", dir)
	resp, err := http.Get(url)
	if err != nil {
		panic(err)
	}
	defer resp.Body.Close()

	doc, err := html.Parse(resp.Body)
	if err != nil {
		panic(err)
	}

	var files []string
	getFiles(doc, &files, dir)

	return files
}

func getFiles(n *html.Node, files *[]string, dir string) {
	if n.Type == html.ElementNode && n.Data == "a" {
		for _, attr := range n.Attr {
			if attr.Key == "href" {
				href := attr.Val
				if strings.HasPrefix(href, fmt.Sprintf("/gocn/news/blob/master/%s/", dir)) {
					*files = append(*files, strings.TrimPrefix(href, fmt.Sprintf("/gocn/news/blob/master/%s/", dir)))
				}
			}
		}
	}
	for c := n.FirstChild; c != nil; c = c.NextSibling {
		getFiles(c, files, dir)
	}
}

func syncGoCNNews() string {
	syncTime := time.Now().AddDate(0, -1, 0)
	dir := syncTime.Format("200601")

	list := getFileList(dir)
	if len(list) == 0 {
		log.Println("File list is empty.")
		return ""
	}

	newsFile, err := os.Create("./gocn/" + dir + ".md")
	if err != nil {
		panic(err)
	}

	for _, f := range list {
		process(dir, f, newsFile)
	}
	return dir
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
	for i, str := range strs {
		submatch := newsReg.FindStringSubmatch(str)
		title := strings.TrimSpace(submatch[2])
		url := strings.TrimSpace(submatch[3])
		_, err := newsFile.WriteString(strconv.Itoa(i+1) + ". [" + title + "](" + url + ")\n")
		if err != nil {
			panic(err)
		}
	}
}
