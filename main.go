package main

import (
	"fmt"
	"justdocker/common/db"
	"justdocker/common/log"
	"justdocker/idocker"
	"strconv"
)

var myApps = make([]*db.Apps, 0)

func main() {

	idocker.Log.SetLogToFile("./log.txt")
	idocker.Log.SetLogLevel(log.DEBUG)
	//idocker.Log.Log2File("this is test logs!\n")
	idocker.ReadCfgFrom("./cfg.json")
	idocker.Log.CloseLogfile()
	go idocker.StartRPCserver()
	//createapp(10)
	//getapp()
	//updateapp()
	//getapp()
	//deleteapp()

	LearnArry()
	LearnArry1()

}

func createapp(num int) {
	n := num + len(myApps)
	for i := len(myApps); i < n; i++ {
		app := &db.Apps{AppName: "newName-" + strconv.Itoa(i)}
		app.AddApp()
		myApps = append(myApps, app)
	}
}

func deleteapp() {

	for _, a := range myApps {
		a.DeleteApp()
	}

}

func updateapp() {
	tmpApp := &db.Apps{}
	for index, a := range myApps {
		tmpApp.AppName = "u" + a.AppName
		l, _ := a.ModifyApp(tmpApp)

		myApps[index] = l[0]
	}

}

func getapp() {
	apps := &db.Apps{}
	list, _ := apps.GetApps()

	for _, a := range list {
		fmt.Println("app:", a)
	}

}

func LearnArry(){
	var buffer [20]int
	var strarry =[...]string{"123","qaz","fgh"}

	strarry[len(strarry)-1] = "update"
	buffer[1] = 3

	fmt.Println(strarry)
	fmt.Println(buffer)
}

func LearnArry1(){
	a := make(map[int]interface{},100)
	b := make(map[int]interface{},0)

	c := make([]int,0)
	d := make([]db.Apps,0)

	e := make(map[string]int)
	f := map[int]int{1:1,2:2,3:3,4:4,5:5,6:6}

	if _, is := f[8]; is {

		fmt.Println("found in map\n",a,b,c,d,e,f)

	}

}























