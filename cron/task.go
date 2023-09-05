package cron

var tasks []Handler

type Handler struct {
	//cron表达式
	//(1) */2 * * * * *     每个偶数分执行text
	//(2) */2 */2 * * * *   每个偶数分 并且是偶数秒执行text
	//(3) 2,4,6 * * * * *   每分钟的2,4,6这三个秒执行text
	//(4) 2,4,6 5-6 * * * 1 每周一的5点2,4,6秒和6点2,4,6秒执行text
	//(5) @daily            每天一次
	//(6) @midnight         同上
	//(7) @every 1m30s      定时1分30秒执行text
	Cron string
	//实际处理方法
	Handler func()
	//执行顺序
	Sort int
	//初始化是否执行（懒加载）
	Lazy bool
}

// Task 用于tasks注册任务
func Task(handler ...Handler) {
	if tasks == nil {
		tasks = make([]Handler, 0)
	}
	for _, task := range handler {
		tasks = append(tasks, task)
	}
}
