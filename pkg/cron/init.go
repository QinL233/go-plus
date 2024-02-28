package cron

import (
	"github.com/robfig/cron/v3"
	"sort"
)

func Init() {
	if len(tasks) < 1 {
		return
	}
	// sort
	sort.Slice(tasks, func(i, j int) bool {
		return tasks[i].Sort < tasks[j].Sort
	})
	c := cron.New()
	for _, task := range tasks {
		//lazy handler
		if !task.Lazy {
			task.Handler()
		}
		//c.AddFunc("@every 1s", func() { fmt.Println("Every hour thirty, starting an hour thirty from now") })
		c.AddFunc(task.Cron, task.Handler)
	}
	c.Start()
}
