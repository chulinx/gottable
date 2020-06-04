package gottable

import (
	"fmt"
	"strings"
)

type Table struct {
	Data     [][]string
	Position string
	Style    string
}

func New(data [][]string,posi string,style string) *Table {
	return &Table{
		Data: data,
		Position: posi,
		Style: style,
	}
}

// @desc 获取每一列最大字符串长度
func (t *Table) MaxCol() []int {
	headerColLen := len(t.Data[0]) // 获取表头元素数
	maxs := make([]int, 0)
	for j := 0; j < headerColLen; j++ {
		// 每一列元素数组
		temp := make([]string, 0)
		for _, d := range t.Data {
			temp = append(temp, d[j])
		}
		max := 0
		for _, i := range temp {
			if len(i) > max {
				max = len(i)
			}
		}
		maxs = append(maxs, max)
	}
	return maxs
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
			filling string
			cfilling = ""
			wrow int
			nc string
		)
		if t.MaxCol()[i] %2 ==1 {
			wrow = t.MaxCol()[i] + 1
		}else {
			wrow = t.MaxCol()[i] +2
		}

		line := strings.Repeat("-", wrow)
		if len(c) % 2 == 1{
			nc = c+" "
			filling = strings.Repeat(" ", wrow-1-len(c))
		}else {
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