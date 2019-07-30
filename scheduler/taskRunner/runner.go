package taskRunner

//生产者，消费者模型

type Runner struct {
	Controller controlChnn
	Error      controlChnn
	Data       dataChnn
	dataSize   int
	longlived  bool
	Dispatcher fn
	Executor   fn
}

func NewRunner(size int, longlived bool, d fn, e fn) *Runner {
	return &Runner{
		Controller: make(chan string, 1),
		Error:      make(chan string, 1),
		Data:       make(chan interface{}, size),
		longlived:  longlived,
		dataSize:   size,
		Dispatcher: d,
		Executor:   e,
	}
}

func (r *Runner) startDispatch() {
	defer func() { //use longlived to determine whether reuse these resources
		if !r.longlived {
			close(r.Controller)
			close(r.Data)
			close(r.Error)
		}
	}() //use this function right now
	for { //为啥这里有个for循环
		select {
		//通过 dispatcher 和 executor 这两个互相发消息，让消息完全异步传输，解耦开来
		case c := <-r.Controller:
			if c == READY_TO_DISPATCH {
				err := r.Dispatcher(r.Data)
				if err != nil {
					r.Error <- CLOSE
				} else {
					r.Controller <- RRADY_TO_EXECUTE
				}
			}

			if c == RRADY_TO_EXECUTE {
				err := r.Executor(r.Data)
				if err != nil {
					r.Error <- CLOSE
				} else {
					r.Controller <- READY_TO_DISPATCH
				}
			}

		case e := <-r.Error:
			if e == CLOSE {
				return
			}

		default:
		}
	}
}

func (r *Runner) startAll() {
	r.Controller <- READY_TO_DISPATCH
	r.startDispatch()
}
