package main

import (
	"bytes"
	"log"
	"os/exec"
	"strconv"
	"strings"
	"fmt"
	"time"
)

type Process struct {
	pid int
	cpu float64
}

func main1() *Process {
	cmd := exec.Command("ps", "aux")
	var out bytes.Buffer
	cmd.Stdout = &out
	err := cmd.Run()
	if err != nil {
		log.Fatal(err)
	}
	processes := make([]*Process, 0)
	for {
		line, err := out.ReadString('\n')
		if err!=nil {
			break;
		}
		tokens := strings.Split(line, " ")
		ft := make([]string, 0)
		for _, t := range(tokens) {
			if t!="" && t!="\t" {
				ft = append(ft, t)
			}
		}
		// log.Println(len(ft), ft)
		pid, err := strconv.Atoi(ft[1])
		if err!=nil {
			continue
		}
		cpu, err := strconv.ParseFloat(ft[2], 64)
		if err!=nil {
			log.Fatal(err)
		}
		processes = append(processes, &Process{pid, cpu})
	}
	//for _, p := range(processes) {
	//	log.Println("Process ", p.pid, " takes ", p.cpu, " % of the CPU")
	//}

	maxcpu := processes[0]

	for _, p := range(processes) {
		if p == maxcpu {
			continue
		}
		if p.cpu > maxcpu.cpu {
			maxcpu = p
		}
	}

	fmt.Println("max cpu:", maxcpu.pid, maxcpu.cpu)

	return maxcpu
}

func main() {
	pid := 0
	cnt := 0

	maxcpu := 50.0
	maxcnt := 10

	for true {
		time.Sleep(time.Millisecond * 300)
		cur := main1()
		if cur.cpu > maxcpu {
			if cur.pid == pid {
				cnt++
			} else {
				cnt = 0
			}
			pid = cur.pid
		}
		if cnt > maxcnt {
			cnt = 0
			cmd := exec.Command("kill", strconv.Itoa(pid))
			fmt.Println("cmd kill:", strconv.Itoa(pid))
			cmd.Run()
		}
	}
}
