package main

import (
	"fmt"
	"sync"
	"time"
)

func HandleCompany(c *chan int, wg *sync.WaitGroup, r row) {
	wg.Add(1)
	defer func() {
		<-*c
	}()
	defer wg.Done()
	time.Sleep(time.Second)
	fmt.Println("-----HANDLE COMPANY " + r.data)
	channelJob := make(chan int, 2)
	var jobWG sync.WaitGroup
	for _, row := range r.job {
		channelJob <- 1
		go HandleJob(&channelJob, &jobWG, row)
	}
	time.Sleep(time.Second)
	jobWG.Wait()
}
func HandleJob(c *chan int, wg *sync.WaitGroup, job string) {
	wg.Add(1)
	defer func() {
		<-*c
	}()
	defer wg.Done()
	time.Sleep(time.Second)
	fmt.Println("HANDLE JOB ", job)
	time.Sleep(time.Second)
}
func main() {
	channelCompany := make(chan int, 3)
	var wg sync.WaitGroup
	for _, row := range company {
		channelCompany <- 1
		go HandleCompany(&channelCompany, &wg, row)
	}
	wg.Wait()
}

type row struct {
	data string
	job  []string
}

//32 JOB IN 9 COMPANY
var company = []row{
	row{
		data: "DacTech",
		job:  []string{"C/C++ Emblemed", "Java web/Spring boot", "PHP/Laravel", "Android kotlin/java Flutter", "Front end", "Adtech"},
	},
	row{
		data: "Fsoft",
		job:  []string{"PHP/Laravel", "Machine learning", "AI", "Mobile"},
	},
	row{
		data: "Mgm",
		job:  []string{"Spring", "React", "AI", "Mobile"},
	},
	row{
		data: "Sun*",
		job:  []string{"Reactjs", "Ruby on rail", "AI", "Mobile"},
	},
	row{
		data: "Datahouse",
		job:  []string{"AI", "Machine learning"},
	},
	row{
		data: "Openweb",
		job:  []string{"PHP/Laravel", "Mobile"},
	},
	row{
		data: "Kms",
		job:  []string{"PHP/Laravel", "AI", "Mobile"},
	},
	row{
		data: "Google Inc",
		job:  []string{"Machine learning", "AI", "Mobile"},
	},
	row{
		data: "Facebook Inc",
		job:  []string{"PHP/Laravel", "Machine learning", "AI", "Mobile"},
	},
}
