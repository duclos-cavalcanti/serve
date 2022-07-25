package util

import (
    "os"
    "fmt"
    exec "os/exec"
    "sync"
    "log"
)

func createLogger() *log.Logger {
    name := "menu.log"
    if (ExistsFile(name)) {
        _, err := exec.Command("bash", "-c", fmt.Sprintf("rm -f %s", name)).Output()
        if err != nil {
            log.Fatal(err)
        }
    }

    f, err := os.OpenFile(name, os.O_RDWR|os.O_CREATE|os.O_APPEND, 0666)
    if err != nil {
        log.Fatal(err)
    }

    l := log.New(f, "INFO: ", log.Ltime)
    return l
}

func LogEvents(debug_channel <-chan string, wait_group *sync.WaitGroup) {
    defer wait_group.Done()
    l := createLogger()

    for {
        s, more := <- debug_channel
        if !more {
            break
        }

        l.Println(s)
    }
    return
}
