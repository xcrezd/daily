package main

import (
	"flag"
	"log"
	"os/exec"
	"time"
)

func performJob(job string) error {
	cmd := exec.Command("bash", "-c", job)
	stdoutStderr, err := cmd.CombinedOutput()
	if err != nil {
		log.Printf("Job: %s, error: %s \n", job, err)
	}
	log.Printf("Job %s, stdout/err: %s\n", job, stdoutStderr)
	return err
}

func run(job string) {
	log.Println("perform the job!")
	for try := 3; try > 0; try-- {
		if err := performJob(job); err == nil {
			break
		}
		time.Sleep(time.Minute)
	}
}

func main() {
	var hour = flag.Int("hour", -1, "hour to run the command")
	var job = flag.String("job", "", "job, like \"echo 1\" ")
	flag.Parse()
	log.Printf("\n hour: %d \n job: %s \n", *hour, *job)

	if len(*job) == 0 {
		log.Fatal("job must be set")
	}

	if *hour == -1 {
		run(*job)
		return
	}

	if *hour < 0 || *hour > 23 {
		log.Fatal("wrong hour, must be 0-23")
	}

	c := time.Tick(time.Second)
	for now := range c {
		h, m, s := now.Clock()
		if h == *hour && m == 0 && s == 0 {
			run(*job)
		}
	}
}