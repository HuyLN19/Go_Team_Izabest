package main

import (
	"fmt"
	"sync"
	"time"
)

func HandleCompany(c *chan int, wg *sync.WaitGroup) {
	wg.Add(1)
	defer wg.Done()
	time.Sleep(time.Second)
	v := company[0]
	company = company[1:]
	fmt.Println("HANDLE " + v.data)
	channelJob := make(chan int, 2)
	for i := 0; i < len(job); i++ {
		channelJob <- 1
		go HandleJob(&channelJob, wg, i)
	}
	time.Sleep(time.Second)
	<-*c
}
func HandleJob(c *chan int, wg *sync.WaitGroup, i int) {
	wg.Add(1)
	defer wg.Done()
	time.Sleep(time.Second)
	if len(job) > 0 {
		v := job[0]
		job = job[1:]
		fmt.Println("HANDLE JOB ", v)

	}
	time.Sleep(time.Second)
	<-*c
}
func main() {
	channelCompany := make(chan int, 3)
	var wg sync.WaitGroup
	for i := 0; i < len(company); i++ {
		channelCompany <- 1
		go HandleCompany(&channelCompany, &wg)
	}
	wg.Wait()
}

type row struct {
	data string
}

var company = []row{
	row{
		data: "DacTech",
	},
	row{
		data: "Fsoft",
	},
	row{
		data: "Mgm",
	},
	row{
		data: "Sun*",
	},
	row{
		data: "Datahouse",
	},
	row{
		data: "Openweb",
	},
	row{
		data: "Kms",
	},
	row{
		data: "Google Inc",
	},
	row{
		data: "Facebook Inc",
	},
}

var job = []row{
	row{
		data: "C/C++ Emblemed",
	},
	row{
		data: "Java web/Spring boot",
	},
	row{
		data: "PHP/Laravel",
	},
	row{
		data: "Front end vuejs/reactjs",
	},
	row{
		data: "Android kotlin/java Flutter",
	},
	row{
		data: "Backoffice",
	},
	row{
		data: "Front end",
	},
	row{
		data: "Google Inc",
	},
	row{
		data: "Facebook Inc",
	},
	row{
		data: "C/C++ Emblemed",
	},
	row{
		data: "Java web/Spring boot",
	},
	row{
		data: "PHP/Laravel",
	},
	row{
		data: "Front end vuejs/reactjs",
	},
	row{
		data: "Android kotlin/java Flutter",
	},
	row{
		data: "Backoffice",
	},
	row{
		data: "Front end",
	},
	row{
		data: "Google Inc",
	},
	row{
		data: "Facebook Inc",
	},
	row{
		data: "C/C++ Emblemed",
	},
	row{
		data: "Java web/Spring boot",
	},
	row{
		data: "PHP/Laravel",
	},
	row{
		data: "Front end vuejs/reactjs",
	},
	row{
		data: "Android kotlin/java Flutter",
	},
	row{
		data: "Backoffice",
	},
	row{
		data: "Front end",
	},
	row{
		data: "Google Inc",
	},
	row{
		data: "Facebook Inc",
	},
	row{
		data: "C/C++ Emblemed",
	},
	row{
		data: "Java web/Spring boot",
	},
	row{
		data: "PHP/Laravel",
	},
	row{
		data: "Front end vuejs/reactjs",
	},
	row{
		data: "Android kotlin/java Flutter",
	},
	row{
		data: "Backoffice",
	},
	row{
		data: "Front end",
	},
	row{
		data: "Google Inc",
	},
	row{
		data: "Facebook Inc",
	},
}
