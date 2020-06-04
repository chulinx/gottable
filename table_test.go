package gottable

import "testing"

var (
	data = [][]string{
	{"name","source","sex","age"},
	{"Meter","80","man",""},
	{"Peter","92","man","21"},
	{"xiMing","86","woman","18"},
}
	table = New(data,"center","simplicity")
)

func TestTable_MaxCol(t *testing.T) {
	maxCol:=table.MaxCol()
	if maxCol[0]==6 && maxCol[1]==6 && maxCol[2]==5 && maxCol[3]==3{
		t.Logf("Success %v",maxCol)
		return
	}
	t.Fatalf("Failed %v",maxCol)
}

func TestTable_Render(t *testing.T) {
	table.Render()
}