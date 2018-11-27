package main

import (
	"strconv"
)

import (
	"fmt"
	"os"
	"time"
)

func main() {
	if len(os.Args) < 2 {
		fmt.Println(`
usage:
	tsconv [unixstamp|time string]
	
	time string must format like: 2006-01-02T15:04:05
		`)

		os.Args = append(os.Args, strconv.Itoa(int(time.Now().Unix())))
	}

	ts, err := strconv.Atoi(os.Args[1])
	if err != nil {
		// 输入可能是一个时间字符串
		t, err := time.ParseInLocation("2006-01-02T15:04:05", os.Args[1], time.Local)
		if err != nil {
			fmt.Printf("time.Parse error: %+v \n", err)
		} else {
			fmt.Printf("%s => unix timestamp[%d]\n", os.Args[1], t.Unix())
		}
	} else {
		t := time.Unix(int64(ts), 0)
		fmt.Printf("unix timestamp[%d] => %s\n", ts, t.Format("2006-01-02T15:04:05"))
	}

}
