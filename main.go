package main

import (
	"flag"
	"log"
	"time"
	"os/exec"
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

func main() {
	var hour = flag.Int("hour", -1, "hour to run the command")
	var job = flag.String("job", "", "job, like \"echo 1\" ")
	flag.Parse()
	log.Printf("\n hour: %d \n job: %s \n", *hour, *job)
	
	if *hour < 0 || *hour > 23 {
		log.Fatal("wrong hour, must be 0-23")
	}
	if len(*job) == 0 {
		log.Fatal("job must be set")
	}
	
	c := time.Tick(time.Second)
	for now := range c {
		h, m, s := now.Clock()
		if h == *hour && m == 0 && s == 0 {
			log.Println("perform the job!")
			for err := performJob(*job); err!= nil; {
				time.Sleep(time.Minute)
				err = performJob(*job)
			}
		}
	}
}