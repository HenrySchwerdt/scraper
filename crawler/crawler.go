package crawler

import (
	"bufio"
	"bytes"
	"io"
	"log"
	"os"
	"sort"
	"strconv"
	"strings"
	"unicode"

	"github.com/PuerkitoBio/goquery"
	"github.com/gocolly/colly"
	prose "github.com/jdkato/prose/tokenize"
	"github.com/schollz/progressbar/v3"
)

type Crawler struct {
	InputFile  string
	OutputFile string
}

func NewCrawler(inputFile, outputFile string) *Crawler {
	return &Crawler{
		InputFile:  inputFile,
		OutputFile: outputFile,
	}
}

func (c *Crawler) lineCounter() (int, error) {
	buf := make([]byte, 32*1024)
	count := 0
	lineSep := []byte{'\n'}
	file, err := os.Open(c.InputFile)
	if err != nil {
		log.Fatal(err)
	}
	defer file.Close()
	for {
		c, err := file.Read(buf)
		count += bytes.Count(buf[:c], lineSep)

		switch {
		case err == io.EOF:
			return count, nil

		case err != nil:
			return count, err
		}
	}
}

func (c *Crawler) Crawl() {
	file, err := os.Create(c.OutputFile)
	if err != nil {
		log.Fatal(err)
	}
	file.Close()
	file, err = os.Open(c.InputFile)
	if err != nil {
		log.Fatal(err)
	}
	defer file.Close()

	scanner := bufio.NewScanner(file)
	numberOfWebsites, err := c.lineCounter()
	if err != nil {
		log.Fatal(err)
	}
	bar := progressbar.Default(int64(numberOfWebsites))
	for scanner.Scan() {
		c.crawlURL(scanner.Text())
		bar.Add(1)
	}

	if err := scanner.Err(); err != nil {
		log.Fatal(err)
	}

	bar.Finish()
}

func (c *Crawler) crawlURL(url string) {

	collector := colly.NewCollector()
	collector.OnResponse(func(r *colly.Response) {
		doc, err := goquery.NewDocumentFromReader(bytes.NewReader(r.Body))
		if err != nil {
			log.Fatal(err)
		}

		wordCount := make(map[string]int)
		tokenizer := prose.NewTreebankWordTokenizer()

		doc.Find("*").Each(func(i int, s *goquery.Selection) {
			// Skip if the element is script, style, img, svg, iframe, or has children
			if s.Is("style, img, svg, iframe, script") || s.Children().Length() > 0 {
				return
			}

			TxtContent := strings.TrimSpace(s.Text())
			for _, word := range tokenizer.Tokenize(TxtContent) {
				if len(word) != 0 && isAlphaOrHyphenOrDot(word) {
					wordCount[word]++
				}
			}

		})

		c.appendCsvLine(url, wordCount)
	})

	collector.Visit(url)
}

func (c *Crawler) createCsvLine(url string, wordCount map[string]int) string {
	var line string

	line += url + ","
	type kv struct {
		Key   string
		Value int
	}
	var ss []kv
	for k, v := range wordCount {
		ss = append(ss, kv{k, v})
	}
	sort.Slice(ss, func(i, j int) bool {
		return ss[i].Value > ss[j].Value
	})

	for i, kv := range ss {
		line += strings.ToLower(kv.Key) + ";" + strconv.Itoa(kv.Value)
		if i < len(ss)-1 {
			line += ","
		}
	}
	return line + "\n"
}

func (c *Crawler) appendCsvLine(url string, wordCount map[string]int) {

	file, err := os.OpenFile(c.OutputFile, os.O_APPEND|os.O_WRONLY, os.ModeAppend)
	if err != nil {
		log.Fatal(err)
	}
	defer file.Close()

	writer := bufio.NewWriter(file)

	writer.Write([]byte(c.createCsvLine(url, wordCount)))
	writer.Flush()
}

func isAlphaOrHyphenOrDot(s string) bool {
	runes := []rune(s)
	first, last := runes[0], runes[len(runes)-1]
	if s == "https" || s == "http" || s == "www" || s == "iframe" {
		return false
	}
	if !unicode.IsLetter(first) || !unicode.IsLetter(last) {
		return false
	}
	for _, r := range s {
		if !unicode.IsLetter(r) && r != '-' && r != '.' {
			return false
		}
	}
	return true
}
