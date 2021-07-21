package gottable

import (
	"fmt"
	"log"
	"strings"
	"testing"
)

var (
	data = [][]string{
		{"Name","Source","Sex","Age"},
		{"Meter","80","man",""},
		{"Peter","92","man","21"},
		{"xiMing","86","woman","18"},
}
//	data = [][]string{
//		{"Name","Source","Sex","Age"},
//	}
	table = New(data,"","simplicity")
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

func Test_tableHead(t *testing.T) {
	mas := []int{4,6,3,7}
	data := tableHead(mas)
	fmt.Println(data)
	if len(data) == 33 {
		t.Logf("Success")
		return
	}
	t.Fatalf("Failed")
}

func Test_tableData(t *testing.T) {
	hData:=[]string{"服务名","主机","状态","启动时间"}
	mas:=make([]int,0)
	for _,d := range hData {
		mas = append(mas, len(d)+2)
	}
	data := tableData(hData,mas)
	fmt.Println(data)
	if strings.Join(data,"") == "│服务名    │主机    │状态    │启动时间    │"{
		t.Logf("Success")
		return
	}
	log.Fatalf("Failed")
}

func TestTable_Print(t *testing.T) {
	table.Print()
}