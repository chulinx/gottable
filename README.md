# Gottable
## Golang Terminal Table
> Easy terminal table for golang

### How to use

```go
package main

import "github.com/chulinx/gottable"

var (
	data = [][]string{
	{"name","source","sex","age"},
	{"Meter","80","man",""},
	{"Peter","92","man","21"},
	{"xiMing","86","woman","18"},
   }
	
)

func main(){
    table := gottable.New(data,"left","")
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
    table := gottable.New(data,"right","")
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
    table := gottable.New(data,"center","")
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
    table := gottable.New(data,"center","simplicity")
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
}

