package crawler

import (
	"fmt"
	"strings"

	"github.com/PuerkitoBio/goquery"
	"github.com/jimweng/crawler/crawler/plugins/inputs"
	"github.com/jimweng/crawler/crawler/utils"
)

type QueryUrl struct {
	Url string
}

func (q *QueryUrl) Gather() (interface{}, error) {
	doc, err := goquery.NewDocument(q.Url)
	if err != nil {
		fmt.Printf("%v\n", err)
	}
	PkgList, err := parseDoc(doc)
	return PkgList, err
}

func stripTaps(str string) string {
	str = strings.Replace(str, "\t", "", -1)
	str = strings.Replace(str, "\n", "", -1)
	return str
}

func parseDoc(doc *goquery.Document) (*[]*utils.PKGContent, error) {
	var PkgList []*utils.PKGContent

	/*
		DOM: [table]
		~ ~ ~
		|-tr
		|	|---td
		|	|	|--cls `pkg-name`
		|	|		|-a
		|	|
		|	|---td
		|		|--cls `pkg-synopsis`
		~ ~ ~
	*/
	doc.Find("table").Each(func(index int, tablehtml *goquery.Selection) {
		tablehtml.Find("tr").Each(func(indextr int, rowhtml *goquery.Selection) {
			pkgnode := utils.PKGContent{}

			rowhtml.Find(".pkg-name").Each(func(indexPkgName int, tableClsPkgName *goquery.Selection) {
				pkgnode.Name = stripTaps(tableClsPkgName.Text())

				tableClsPkgName.Find("a").Each(func(indexa int, tdcell *goquery.Selection) {
					href_text, ok := tdcell.Attr("href")
					if ok {
						pkgnode.Href = stripTaps(href_text)
						parentArray := strings.Split(pkgnode.Href, "/")
						lenParentArray := len(parentArray)
						/*
							archive/ ==> [archive ] <-len(2)
							parent would be self, loc = 0
							archive/tar/ ==> [archive tar ] <- len(3)
							parent would be "archive", loc = 0
							net/http/cgi ==> [net http cgi ] <- len(4)
							parent would be "http", loc = 1
						*/
						if lenParentArray <= 2 {
							pkgnode.Parent = parentArray[0]
						} else {
							pkgnode.Parent = parentArray[lenParentArray-3]
						}
					}
				})
			})

			rowhtml.Find(".pkg-synopsis").Each(func(indexSynpopsis int, tableClsSynopsis *goquery.Selection) {
				pkgnode.Synopsis = stripTaps(tableClsSynopsis.Text())
			})

			PkgList = append(PkgList, &pkgnode)

		})
	})
	return &PkgList, nil
}

func init() {
	inputs.Add("crawler", func() utils.Input {
		return &QueryUrl{}
	})
}
