package gottable

import (
	"fmt"
	"github.com/gookit/color"
	"strings"
)

const (
	topLeftAngle      = "┌"
	topRightAngle     = "┐"
	topCenterSplit    = "┬"
	horizontalLine    = "─"
	verticalLine      = "│"
	leftBorder        = "├"
	rightBorder       = "┤"
	centerSplit       = "┼"
	bottomLeftAngle   = "└"
	bottomRightAngle  = "┘"
	bottomCenterSplit = "┴"
	filling           = " "
)

// table style params const
const (
	PositionLeft   = "left"
	PositionCenter = "center"
	PositionRight  = "right"
	StyleSimple    = "simplicity"
)

type HeadStyle struct {
	IsBorder  bool
	TextColor color.Color
}
type TableParams struct {
	HeadStyle
	Position string
	Style    string
}

type Table struct {
	// Data give table head and table content
	Data [][]string
	// Position define table content position
	// have three value center/right/left
	Position  string
	Style     string
	HeadStyle HeadStyle
}

func New(data [][]string, tableParams TableParams) *Table {
	return &Table{
		Data:      data,
		Position:  tableParams.Position,
		Style:     tableParams.Style,
		HeadStyle: tableParams.HeadStyle,
	}
}

// MaxCol 获取每列据最长的元素长度列表
func (t *Table) MaxCol() []int {
	headerColLen := len(t.Data[0]) // 获取表头元素数
	mas := make([]int, 0)          // 每一列最长元素的长度数组
	for j := 0; j < headerColLen; j++ {
		// 每一列元素数组
		temp := make([]string, 0)
		for _, d := range t.Data {
			temp = append(temp, d[j])
		}
		// 获取最长元素的长度
		max := 0
		for _, i := range temp {
			if len(i) > max {
				max = len(i)
			}
		}
		mas = append(mas, max)
	}
	return mas
}

func (t *Table) Print() {
	tableDatas := make([]string, 0)
	mas := t.MaxCol()
	head := strings.Join(tableHead(mas), "")
	bottom := strings.Join(tableBottom(mas), "")
	center := strings.Join(tableCenter(mas), "")
	tableDatas = append(tableDatas, head)
	for i, datas := range t.Data {
		var tmp string
		tmp = t.styleData(i, datas, mas)
		tableDatas = append(tableDatas, tmp)
		if i == len(t.Data)-1 {
			continue
		}
		tableDatas = append(tableDatas, center)
	}
	tableDatas = append(tableDatas, bottom)
	for _, t := range tableDatas {
		fmt.Println(t)
	}
}

func (t *Table) styleData(i int, datas []string, mas []int) string {
	if i == 0 {
		if t.HeadStyle.IsBorder {
			return strings.Join(tableData(datas, mas, color.Bold, t.HeadStyle.TextColor), "")
		}
		return strings.Join(tableData(datas, mas, t.HeadStyle.TextColor), "")
	}
	return strings.Join(tableData(datas, mas), "")
}

func (t *Table) Render() {
	lines := []string{}
	// 渲染表格
	for ri, row := range t.Data {
		// println(t.data)
		srowcontent := []string{}
		// 清空lines
		lines = []string{}
		if t.Position == PositionCenter {
			srowcontent, lines = t.center(row, srowcontent, lines, ri)
		} else if t.Position == PositionRight {
			srowcontent, lines = t.right(row, srowcontent, lines, ri)
		} else {
			lines = t.left(row, srowcontent, lines, ri)
		}
	}
	// 最后打印一行
	fmt.Println(strings.Join(lines, ""))
}

func (t *Table) left(row []string, srowcontent []string, lines []string, ri int) []string {
	for i, c := range row {
		wrow := t.MaxCol()[i] + 1
		line := strings.Repeat("-", wrow+2)
		filling := strings.Repeat(" ", wrow-len(c)+2)
		if len(srowcontent) > 0 {
			srowcontent = append(srowcontent, c+filling+"|")
			lines = append(lines, line+"+")
		} else {
			srowcontent = append(srowcontent, "|"+c+filling+"|")
			lines = append(lines, "+"+line+"+")
		}
	}
	if ri > 1 && t.Style == StyleSimple {
		fmt.Println(strings.Join(srowcontent, ""))
	} else {
		for _, p := range [][]string{lines, srowcontent} {
			fmt.Println(strings.Join(p, ""))
		}
	}
	return lines
}

func (t *Table) right(row []string, srowcontent []string, lines []string, ri int) ([]string, []string) {
	for i, c := range row {
		wrow := t.MaxCol()[i] + 1
		line := strings.Repeat("-", wrow+2)
		filling := strings.Repeat(" ", wrow-len(c)+2)
		if len(srowcontent) > 0 {
			srowcontent = append(srowcontent, filling+c+"|")
			lines = append(lines, line+"+")
		} else {
			srowcontent = append(srowcontent, "|"+filling+c+"|")
			lines = append(lines, "+"+line+"+")
		}
	}
	if ri > 1 && t.Style == StyleSimple {
		fmt.Println(strings.Join(srowcontent, ""))
	} else {
		for _, p := range [][]string{lines, srowcontent} {
			fmt.Println(strings.Join(p, ""))
		}
	}
	return srowcontent, lines
}

func (t *Table) center(row []string, srowcontent []string, lines []string, ri int) ([]string, []string) {
	for i, c := range row {
		var (
			filling  string
			cfilling = ""
			wrow     int
			nc       string
		)
		if t.MaxCol()[i]%2 == 1 {
			wrow = t.MaxCol()[i] + 1
		} else {
			wrow = t.MaxCol()[i] + 2
		}

		line := strings.Repeat("-", wrow)
		if len(c)%2 == 1 {
			nc = c + " "
			filling = strings.Repeat(" ", wrow-1-len(c))
		} else {
			nc = c
			filling = strings.Repeat(" ", wrow-len(c))
		}

		//if len(filling)%2 == 1 {
		//	cfilling = strings.Repeat(" ", (len(filling)+3)/2)
		//	line = strings.Repeat("-", len(line)+2)
		//} else {
		cfilling = strings.Repeat(" ", (len(filling)+2)/2)
		line = strings.Repeat("-", len(line)+2)
		//}
		if len(srowcontent) > 0 {
			srowcontent = append(srowcontent, cfilling+nc+cfilling+"|")
			//srowcontent = append(srowcontent, filling+c+filling+"|")
			lines = append(lines, line+"+")
		} else {
			srowcontent = append(srowcontent, "|"+cfilling+nc+cfilling+"|")
			//srowcontent = append(srowcontent, "|"+filling+c+filling+"|")
			lines = append(lines, "+"+line+"+")
		}
	}
	if ri > 1 && t.Style == StyleSimple {
		fmt.Println(strings.Join(srowcontent, ""))
	} else {
		for _, p := range [][]string{lines, srowcontent} {
			fmt.Println(strings.Join(p, ""))

		}
	}
	return srowcontent, lines
}

// tableHead 表头部渲染
// 根据每一列元素最长元素的长度生成对应长度的表头部
func tableHead(mas []int) []string {
	return tableBorder(mas, topLeftAngle, horizontalLine, topCenterSplit, topRightAngle)
}

// tableCenter 表中心区域渲染
func tableCenter(mas []int) []string {
	return tableBorder(mas, leftBorder, horizontalLine, centerSplit, rightBorder)
}

// tableBottom 表底部渲染
func tableBottom(mas []int) []string {
	return tableBorder(mas, bottomLeftAngle, horizontalLine, bottomCenterSplit, bottomRightAngle)

}

func tableBorder(mas []int, left, line, split, right string) []string {
	headArr := make([]string, 1)
	headArr[0] = left
	for j, max := range mas {
		for i := 0; i < max+3; i++ {
			headArr = append(headArr, line)
		}
		// 最后一列不添加topCenterSplit
		if j == len(mas)-1 {
			continue
		}
		headArr = append(headArr, split)
	}
	headArr = append(headArr, right)
	return headArr
}

// tableData 数据行
func tableData(hData []string, mas []int, styles ...color.Color) []string {
	dataArr := make([]string, 1)
	dataArr[0] = verticalLine
	for h, data := range hData {
		nullSpaceCount := mas[h] - len(data)
		data = styleRender(data, styles...)
		nullSpaceStr := strings.Repeat(filling, nullSpaceCount+2)
		dataArr = append(dataArr, filling+data+nullSpaceStr)
		//dataArr = append(dataArr, nullSpaceStr)
		if h == len(hData)-1 {
			continue
		}
		dataArr = append(dataArr, verticalLine)
	}
	dataArr = append(dataArr, verticalLine)
	return dataArr
}

func styleRender(s string, styles ...color.Color) string {
	if len(styles) == 0 {
		return s
	}
	return color.New(styles...).Render(s)
}
