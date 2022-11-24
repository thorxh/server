package phrase

import (
	"bytes"
	"github.com/PuerkitoBio/goquery"
	"golang.org/x/text/encoding/simplifiedchinese"
	"golang.org/x/text/transform"
	"io"
	"log"
	"net/http"
	"net/url"
	"strings"
)

func QueryPhrase(queryPhrase string) string {
	gbkStr, _ := utf8ToGbk(queryPhrase)
	data := url.Values{"txtname": {gbkStr}}
	resp, err := http.PostForm("http://ciyu.kaishicha.com/socy.asp", data)
	if err != nil {
		return "抱歉，查询失败"
	}

	if resp.StatusCode == 404 {
		return "抱歉，查不到该词语"
	}

	if resp.StatusCode != 200 {
		return "抱歉，查询失败"
	}

	defer func(Body io.ReadCloser) {
		err := Body.Close()
		if err != nil {

		}
	}(resp.Body)

	doc, err := goquery.NewDocumentFromReader(resp.Body)
	if err != nil {
		log.Fatal(err)
	}

	// 获取到页面主体
	contentDiv := doc.Find("div.zhongnr2")
	// 获取解释
	detailDiv := contentDiv.Find("div.tab-content")
	ulList := detailDiv.Find("ul")
	// 词语解释
	var explain string
	// 词语语法
	var usage string
	ulList.Find("li").Each(func(i int, node *goquery.Selection) {
		text, _ := gbkToUtf8(node.Text())
		if strings.Contains(text, "解释") {
			explain = strings.Replace(text, "【解释】：", "", 1)
		}
		if strings.Contains(text, "语法") {
			splitStr := strings.Split(text, "；")
			for _, val := range splitStr {
				if strings.Contains(val, "用于") {
					usage = val
					break
				}
				if strings.Contains(val, "含") {
					usage = val
					break
				}
			}
		}
	})

	if len(explain) == 0 {
		contentDiv.Find("p").Each(func(i int, node *goquery.Selection) {
			if i == 1 {
				text, _ := gbkToUtf8(node.Text())
				explain = strings.TrimSpace(text)
			}
		})
	}

	if len(usage) > 0 {
		return explain + usage + "。"
	} else {
		return explain
	}
}

func gbkToUtf8(src string) (string, error) {
	reader := transform.NewReader(bytes.NewReader([]byte(src)), simplifiedchinese.GBK.NewDecoder())
	d, e := io.ReadAll(reader)
	if e != nil {
		return "", e
	}
	return string(d), nil
}

func utf8ToGbk(src string) (string, error) {
	reader := transform.NewReader(bytes.NewReader([]byte(src)), simplifiedchinese.GBK.NewEncoder())
	d, e := io.ReadAll(reader)
	if e != nil {
		return "", e
	}
	return string(d), nil
}
