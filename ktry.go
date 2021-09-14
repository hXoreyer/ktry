package ktry

func Try(fun func()) CatchHandler {
	c := &catchHandler{}
	defer func() {
		defer func() {
			if err := recover(); err != nil {
				c.err = err.(error)
			}
		}()
		fun()
	}()
	return c
}
