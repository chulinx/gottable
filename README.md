# Gottable
## Golang Terminal Table
> Easy terminal table for golang

### How to use

```go
package main

import (
	"github.com/chulinx/gottable"
	"github.com/gookit/color"
)

var (
	data = [][]string{
	{"name","source","sex","age"},
	{"Meter","80","man",""},
	{"Peter","92","man","21"},
	{"xiMing","86","woman","18"},
   }
	
)

func main(){
    tableParams := gottable.TableParams{Position: gottable.PositionLeft}
    table := gottable.New(data,tableParams)
    table.Render()
// out
/* 
go run main.go
+---------+---------+--------+------+
|name     |source   |sex     |age   |
+---------+---------+--------+------+
|Meter    |80       |man     |      |
+---------+---------+--------+------+
|Peter    |92       |man     |21    |
+---------+---------+--------+------+
|xiMing   |86       |woman   |18    |
+---------+---------+--------+------+
*/
    tableParams = gottable.TableParams{Position: gottable.PositionRight}
    table = gottable.New(data,tableParams)
    table.Render()
// out 
/*
+---------+---------+--------+------+
|     name|   source|     sex|   age|
+---------+---------+--------+------+
|    Meter|       80|     man|      |
+---------+---------+--------+------+
|    Peter|       92|     man|    21|
+---------+---------+--------+------+
|   xiMing|       86|   woman|    18|
+---------+---------+--------+------+
*/
    tableParams = gottable.TableParams{Position: gottable.PositionCenter}
    table = gottable.New(data,tableParams)
    table.Render()
// out
/*
+----------+----------+--------+------+
|   name   |  source  |  sex   | age  |
+----------+----------+--------+------+
|  Meter   |    80    |  man   |      |
+----------+----------+--------+------+
|  Peter   |    92    |  man   |  21  |
+----------+----------+--------+------+
|  xiMing  |    86    | woman  |  18  |
+----------+----------+--------+------+
*/  
    tableParams = gottable.TableParams{Position: gottable.PositionCenter,Style: gottable.StyleSimple}
    table = gottable.New(data,tableParams)
    table.Render()
// out
/*
+----------+----------+--------+------+
|   name   |  source  |  sex   | age  |
+----------+----------+--------+------+
|  Meter   |    80    |  man   |      |
|  Peter   |    92    |  man   |  21  |
|  xiMing  |    86    | woman  |  18  |
+----------+----------+--------+------+
*/

    tableParams = gottable.TableParams{HeadStyle:gottable.HeadStyle{IsBorder: true,TextColor:color.Red }}
    table = gottable.New(data,tableParams)
    table.Print()
}
```
[](images/tablePrint.jpg)