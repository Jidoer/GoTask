package worker

import (
	"GoTask/runer"
	"bufio"
	"fmt"
	"io"
	"log"
	"os/exec"
	"sync"
)

type Job interface {
	Do() error
}



//定义一个实现Job接口的数据
type PrintNum struct {
	Num int
}

type RunShell struct {
	Shell string
}

type Runmain struct {
	Name string
}

//定义对数据的处理
func (s *PrintNum) Do() error {
	fmt.Println("num:", s.Num)
	return nil
}

func (s *RunShell) Do() error {
	ss :=runer.Cmd(s.Shell)
	log.Print(ss)
	return nil
}


func (s *Runmain) Do() error {
	//print(s.Name)
	log.Print(s.Name)
	Command(s.Name)
	return nil
}




func Command(cmd string) error {
	//c := exec.Command("cmd", "/C", cmd) 	// windows
	c := exec.Command("bash", "-c", cmd)  // mac or linux
	stdout, err := c.StdoutPipe()
	if err != nil {
		return err
	}
	var wg sync.WaitGroup
	wg.Add(1)
	go func() {
		defer wg.Done()
		reader := bufio.NewReader(stdout)
		for {
			readString, err := reader.ReadString('\n')
			if err != nil || err == io.EOF {
				return
			}
			fmt.Print(readString)
		}
	}()
	err = c.Start()
	wg.Wait()
	return err
}






