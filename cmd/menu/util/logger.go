package util

import (
    "os"
    "log"
    "container/list"
)

type Logger struct {
    file os.File
    name string
    q *list.List
}

func CreateLogger() Logger {
    name := "menu.log"
    f, err : = os.OpenFile(name, os.O_RDWR|os.O_CREATE|os.O_APPEND, 0666)
    if err != nil {
        log.Fatal(err)
    }

    l := Logger {
        file: f,
        name: name,
        q: list.New()
    }

    return l
}

func (l *Logger) queue(msg string) {
    l.q.PushBack(msg)
}

func (l *Logger) dequeue(msg string) string {
    if (l.q.empty()) {
        val := l.q.Front()
        l.q.Remove(val)
        return val
    } else {
        return nil
    }
}

func (l *Logger) empty() {
    return (l.q.Len() == 0)
}

func (l *Logger) dump(msg string) {
    WriteFile(l.name, msg)
}
