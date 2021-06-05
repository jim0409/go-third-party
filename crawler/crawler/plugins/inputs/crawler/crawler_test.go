package crawler

import (
	"strings"
	"testing"

	"github.com/PuerkitoBio/goquery"
	"github.com/stretchr/testify/assert"
)

const demo_html = `
<!DOCTYPE html>
<html>
<body>
					<table>
						<tr>
							<th class="pkg-name">Name</th>
							<th class="pkg-synopsis">Synopsis</th>
						</tr>
							<tr>
									<td class="pkg-name" style="padding-left: 0px;">
										<a href="archive/">archive</a>
									</td>
								<td class="pkg-synopsis">
								</td>
							</tr>
							<tr>
									<td class="pkg-name" style="padding-left: 0px;">
										<a href="time/">time</a>
									</td>
								<td class="pkg-synopsis">
									Package time provides functionality for measuring and displaying time.
								</td>
							</tr>
							<tr>
									<td class="pkg-name" style="padding-left: 0px;">
										<a href="unicode/">unicode</a>
									</td>
								<td class="pkg-synopsis">
									Package unicode provides data and functions to test some properties of Unicode code points.
								</td>
							</tr>
					</table>
</body>
</html>
`

const demo_string = `
line1	1	2	3
line2	4	5	6
`

func TestSomething(t *testing.T) {
	doc, err := goquery.NewDocumentFromReader(strings.NewReader(demo_html))
	assert.Nil(t, err)

	testPkgList, err := parseDoc(doc)
	assert.Nil(t, err)
	for _, j := range *testPkgList {
		assert.NotEqual(t, "", j)

	}
}

func TestStripTaps(t *testing.T) {
	trimStr := stripTaps(demo_string)
	assert.Equal(t, "line1123line2456", trimStr)
}
