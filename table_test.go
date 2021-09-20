package gottable

import (
	"github.com/gookit/color"
	"testing"
)

var (
	data = [][]string{
		{"Name", "Source", "Sex", "Age"},
		{"Meter", "80", "man", ""},
		{"Peter", "92", "man", "22"},
		{"xiMing", "86", "woman", "18"},
	}
	tableParams = &TableParams{Position: PositionLeft,
		HeadStyle: HeadStyle{IsBorder: true, TextColor: color.Red},
		Style:     StyleSimple,
	}
	table = New(data, *tableParams)
)

func TestTable_MaxCol(t *testing.T) {
	maxCol := table.MaxCol()
	if maxCol[0] == 6 && maxCol[1] == 6 && maxCol[2] == 5 && maxCol[3] == 3 {
		t.Logf("Success %v", maxCol)
		return
	}
	t.Fatalf("Failed %v", maxCol)
}

func TestTable_Render(t *testing.T) {
	table.Render()
}

func TestTable_Print(t *testing.T) {
	table.Print()
}
