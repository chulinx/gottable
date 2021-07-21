package gottable

import (
	"fmt"
	"strings"
)

const (
	topLeftAngle = "┌"
	topRightAngle = "┐"
	topCenterSplit = "┬"
	horizontalLine="─"
	verticalLine = "│"
	leftBorder = "├"
	rightBorder = "┤"
	centerSplit = "┼"
	bottomLeftAngle = "└"
	bottomRightAngle = "┘"
	bottomCenterSplit = "┴"
	filling = " "
)

type Table struct {
	Data        [][]string
	Position    string
	Style       string
	BorderColor string
}

func New(data [][]string,posi string,style string) *Table {
	return &Table{
		Data: data,
		Position: posi,
		Style:    style,
	}
}

// MaxCol 获取每列据最长的元素长度列表
func (t *Table) MaxCol() []int {
	headerColLen := len(t.Data[0]) // 获取表头元素数
	mas := make([]int, 0) // 每一列最长元素的长度数组
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

func (t *Table)Print()  {
	tableDatas := make([]string,0)
	mas := t.MaxCol()
	head := strings.Join(tableHead(mas),"")
	bottom := strings.Join(tableBottom(mas),"")
	center := strings.Join(tableCenter(mas),"")
	tableDatas = append(tableDatas, head)
	for i,datas := range t.Data {
		tmp:=strings.Join(tableData(datas,mas),"")
		tableDatas = append(tableDatas, tmp)
		if i == len(t.Data)-1 {
			continue
		}
		tableDatas = append(tableDatas, center)
	}
	tableDatas = append(tableDatas, bottom)
	for _,t := range tableDatas {
		fmt.Println(t)
	}
}

func (t *Table) Render() {
	lines := []string{}
	// 渲染表格
	for ri, row := range t.Data {
		// println(t.data)
		srowcontent := []string{}
		// 清空lines
		lines = []string{}
		if t.Position == "center" {
			srowcontent, lines = t.center(row, srowcontent, lines, ri)
		} else if t.Position == "right" {
			srowcontent, lines = t.right(row, srowcontent, lines, ri)
		} else {
			lines = t.left(row, srowcontent, lines, ri)
		}
	}
	// 最后打印一行
	fmt.Println(strings.Join(lines, ""))
}

func (t *Table) silkyLeft(row []string, srowcontent []string, lines []string, ri int) []string {
	for i, c := range row {
		wrow := t.MaxCol()[i] + 1
		//line := strings.Repeat("─", wrow+2)
		line := []string{}
		for i := 0; i < wrow+2; i++ {
			line = append(line, "─")
		}

		filling := strings.Repeat(" ", wrow-len(c)+2)
		if len(srowcontent) > 0 {
			srowcontent = append(srowcontent, c+filling+"│")
			line = append(line, "┬")
			lines = append(lines, line...)
		} else {
			srowcontent = append(srowcontent, "│"+c+filling+"│")
			line = append([]string{"┬"}, line...)
			line = append(line, "┬")
			lines = append(lines, line...)
		}
	}
	if ri > 1 && t.Style == "simplicity" {
		fmt.Println(strings.Join(srowcontent, ""))
	} else {
		for _, p := range [][]string{lines, srowcontent} {
			fmt.Println(strings.Join(p, ""))
		}
	}
	for i, l := range lines {
		lines[0] = "└"
		lines[len(lines)-1] = "┘"
		if l == "┬" {
			lines[i] = "┴"
		}
	}
	return lines
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
	if ri > 1 && t.Style == "simplicity" {
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
	if ri > 1 && t.Style == "simplicity" {
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
	if ri > 1 && t.Style == "simplicity" {
		fmt.Println(strings.Join(srowcontent, ""))
	} else {
		for _, p := range [][]string{lines, srowcontent} {
			fmt.Println(strings.Join(p, ""))

		}
	}
	return srowcontent, lines
}

// tableHead 表头部
// 根据每一列元素最长元素的长度生成对应长度的表头部
func tableHead(mas []int) []string {
	return tableBorder(mas,topLeftAngle,horizontalLine,topCenterSplit,topRightAngle)
}

//
func tableCenter(mas []int)  []string {
	return tableBorder(mas,leftBorder,horizontalLine,centerSplit,rightBorder)
}

func tableBottom(mas []int) []string  {
	return tableBorder(mas,bottomLeftAngle,horizontalLine,bottomCenterSplit,bottomRightAngle)

}

func tableBorder(mas []int,left,line,split,right string) []string {
	headArr := make([]string,1)
	headArr[0]=left
	for j,max := range mas {
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
func tableData(hData []string,mas []int) []string {
	dataArr := make([]string,1)
	dataArr[0] = verticalLine
	for h,data := range hData {
		nullSpaceCount := mas[h] - len(data)
		nullSpaceStr := strings.Repeat(filling,nullSpaceCount+2)
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