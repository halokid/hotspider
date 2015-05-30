//

package main

import (
    // "fmt"
    // "io/ioutil"
    "github.com/PuerkitoBio/goquery"
    "github.com/hu17889/go_spider/core/common/page"
    "github.com/hu17889/go_spider/core/common/request"
    "github.com/hu17889/go_spider/core/spider"
    // "strings"
)

type MyPageProcesser struct {
}

func NewMyPageProcesser() *MyPageProcesser {
    return &MyPageProcesser{}
}

// Parse html dom here and record the parse result that we want to crawl.
// Package goquery (http://godoc.org/github.com/PuerkitoBio/goquery) is used to parse html.
func (this *MyPageProcesser) Process(p *page.Page) {
    if !p.IsSucc() {
        println(p.Errormsg())
        return
    }

    query := p.GetHtmlParser()

    // parse HTML dom
    // name := query.Find(".lemmaTitleH1").Text()
    // name = strings.Trim(name, " \t\n")

    // summary := query.Find(".card-summary-content .para").Text()
    // summary = strings.Trim(summary, " \t\n")

    // the entity we want to save by Pipeline
    // p.AddField("name", name)
    // p.AddField("summary", summary)
    
    
    //百度热榜搜索
    // hotKey := query.Find(".list").Text();
    // hotKey := query.Find(".list-title").Text();
    // hotKey := query.Find(".list-title").Html();
    hotKey := query.Find(".list-title").Each(func( i int, s *goquery.Selection ){
        println(s.Text())
    })
    
    
    // hotKey = strings.Trim(hotKey, " \t\n");
    
    println(hotKey, "--------------------")
    
    // p.AddField("hotKey", hotKey)
    
    
}

func main() {

    // println("test")
    // exit
    
    // spider input:
    //  PageProcesser ;
    //  task name used in Pipeline for record;
    sp := spider.NewSpider(NewMyPageProcesser(), "TaskName")
    // GetWithParams Params:
    //  1. Url.
    //  2. Responce type is "html" or "json" or "jsonp" or "text".
    //  3. The urltag is name for marking url and distinguish different urls in PageProcesser and Pipeline.
    //  4. The method is POST or GET.
    //  5. The postdata is body string sent to sever.
    //  6. The header is header for http request.
    //  7. Cookies
    // req := request.NewRequest("http://baike.baidu.com/view/1628025.htm?fromtitle=http&fromid=243074&type=syn", "html", "", "GET", "", nil, nil, nil, nil)
    req := request.NewRequest("http://top.baidu.com/?fr=super_card", "html", "", "GET", "", nil, nil, nil, nil)
    
    pageItems := sp.GetByRequest(req)
    
    // pageItems := sp.Get("http://baike.baidu.com/view/1628025.htm?fromtitle=http&fromid=243074&type=syn", "html")

    url := pageItems.GetRequest().GetUrl()
    println("-----------------------------------spider.Get---------------------------------")
    println("url\t:\t" + url)
    for name, value := range pageItems.GetAll() {
        println(name + ":" + value)
        // fileName := fmt.Sprintf("%s.txt", name)
        // content := readContent(item.url)
        // ioutil.WriteFile(fileName, []byte(value), 0644)
        
        // println("-------------------------")
    }


    // println("\n--------------------------------spider.GetAll---------------------------------")
    // urls := []string{
        // "http://baike.baidu.com/view/1628025.htm?fromtitle=http&fromid=243074&type=syn",
        // "http://baike.baidu.com/view/383720.htm?fromtitle=html&fromid=97049&type=syn",
    // }
    // var reqs []*request.Request
    // for _, url := range urls {
        // req := request.NewRequest(url, "html", "", "GET", "", nil, nil, nil, nil)
        // reqs = append(reqs, req)
    // }
    // pageItemsArr := sp.SetThreadnum(2).GetAllByRequest(reqs)
    // pageItemsArr := sp.SetThreadnum(2).GetAll(urls, "html")
    // for _, item := range pageItemsArr {
        // url = item.GetRequest().GetUrl()
        // println("url\t:\t" + url)
        // fmt.Printf("item\t:\t%s\n", item.GetAll())
    // }
    
}












