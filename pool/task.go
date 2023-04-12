package pool

type Task struct {
	Run    func(v ...interface{})
	params []interface{}
}
