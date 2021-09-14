package ktry

import "reflect"

type catchHandler struct {
	err      error
	hasCatch bool
}

func (c *catchHandler) CheckCatch() bool {
	if c.hasCatch {
		return false
	}
	if c.err == nil {
		return false
	}

	return true
}

func (c *catchHandler) Catch(e error, handler func(err error)) CatchHandler {
	if !c.CheckCatch() {
		return c
	}
	if reflect.TypeOf(e) == reflect.TypeOf(c.err) {
		handler(e)
		c.hasCatch = true
	}
	return c
}

func (c *catchHandler) CatchAll(handler func(err error)) CatchHandler {
	if !c.CheckCatch() {
		return c
	}

	handler(c.err)
	c.hasCatch = true
	return c
}

func (c *catchHandler) Finally(handlers ...func()) {
	for _, handler := range handlers {
		defer handler()
	}

	e := c.err
	if e != nil && !c.hasCatch {
		panic(e)
	}
}
