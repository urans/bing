package dict

import (
	"io/ioutil"
	"net/http"
	"strings"
	"time"

	"github.com/antchfx/htmlquery"
	"github.com/briandowns/spinner"
	"github.com/fatih/color"
	"golang.org/x/net/html"
)

var hrow = func() { color.Black("") }

func text(doc *html.Node) string {
	if doc == nil {
		return ""
	}
	return htmlquery.InnerText(doc)
}

// Query return the meaning of the query string
func Query(q string, sentNum int) error {
	spin := spinner.New(spinner.CharSets[39], time.Microsecond*100)
	spin.Suffix = "Querying ..."
	spin.Color("cyan")

	url := "https://www.bing.com/dict/search"
	req, err := http.NewRequest(http.MethodGet, url, nil)
	if err != nil {
		return err
	}

	params := req.URL.Query()
	params.Add("q", q)
	req.URL.RawQuery = params.Encode()
	resp, err := http.DefaultClient.Do(req)
	if err != nil {
		return err
	}
	defer resp.Body.Close()

	body, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		return err
	}

	ParseWithXPath(string(body), sentNum)
	return nil
}

// ParseWithXPath parse html via html query
func ParseWithXPath(html string, sentNum int) {
	doc, err := htmlquery.Parse(strings.NewReader(html))
	if err != nil {
		color.Red("parse html failed: %v", err)
	}

	parseBasicMean(doc)
	parseCrossMean(doc)
	parseWebMean(doc)
	parseSentences(doc, sentNum)
	parseTranslation(doc)
}

func parseBasicMean(doc *html.Node) {
	head := htmlquery.FindOne(doc, `//div[@id="headword"]/h1/strong`)
	voiceAM := htmlquery.FindOne(doc, `//div[@class="hd_prUS b_primtxt"]`)
	voiceEN := htmlquery.FindOne(doc, `//div[@class="hd_pr b_primtxt"]`)

	if head == nil {
		parseSynonym(doc)
		return
	}

	color.Green("%v  %v %v", text(head), text(voiceAM), text(voiceEN))

	hrow()
}

func parseCrossMean(doc *html.Node) {
	trs := htmlquery.Find(doc, `//div[@id="crossid"]//tr`)
	for _, tr := range trs {
		pos := htmlquery.FindOne(tr, `//div[@class="pos pos1"]`)
		def := htmlquery.FindOne(tr, `//span[@class="p1-1 b_regtxt"]`)
		link := htmlquery.FindOne(tr, `//a[@class="p1-12 b_alink"]`)
		color.Green("%v %v %v", text(pos), text(def), text(link))
	}
}

func parseWebMean(doc *html.Node) {
	means := []string{}
	wrs := htmlquery.Find(doc, `//div[@id="webid"]//tr`)
	for _, tr := range wrs {
		def := htmlquery.FindOne(tr, `//div[@class="p1-1 b_regtxt"]`)
		means = append(means, text(def))
	}

	if len(means) > 0 {
		color.Green("[网络] %v", strings.Join(means, "; "))
		hrow()
	}
}

func parseSentences(doc *html.Node, num int) {
	hrow()
	sents := htmlquery.Find(doc, `//div[@class="li_exs"]`)
	if len(sents) == 0 || num == 0 {
		return
	}

	if 0 < num && num < len(sents) {
		sents = sents[:num]
	}

	for i, sent := range sents {
		enRow := htmlquery.FindOne(sent, `//div[@class="val_ex"]`)
		cnRow := htmlquery.FindOne(sent, `//div[@class="bil_ex"]`)
		color.Green("%v. %v", i+1, text(enRow))
		color.Magenta(" %v", text(cnRow))
		hrow()
	}
}

func parseSynonym(doc *html.Node) {
	tips := htmlquery.FindOne(doc, `//div[@class="p2-2"]`)
	if tips == nil {
		return
	}
	color.Green("%v...", text(tips))
	hrow()

	items := htmlquery.Find(doc, `//div[@class="df_wb_c"]`)
	for i, item := range items {
		word := htmlquery.FindOne(item, `//a[@class="p1-3-1_dymp"]`)
		mean := htmlquery.FindOne(item, `//div[@class="df_wb_text"]`)
		color.Green("[%v] %v: %v", i, text(word), text(mean))
	}
	hrow()
}

func parseTranslation(doc *html.Node) {
	tips := htmlquery.FindOne(doc, `//div[@class="lf_area"]//div[@class="smt_hw"]`)
	orig := htmlquery.FindOne(doc, `//div[@class="lf_area"]//div[@class="p1-10"]`)
	tran := htmlquery.FindOne(doc, `//div[@class="lf_area"]//div[@class="p1-11"]`)

	if tips != nil {
		color.Green("[%v]", text(tips))
		hrow()
		color.Green("%v", text(orig))
		hrow()
		color.Green("%v", text(tran))
		hrow()
	}
}
