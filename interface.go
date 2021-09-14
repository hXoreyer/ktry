package ktry

type FinalHandler interface {
	Finally(handlers ...func())
}

type CatchHandler interface {
	Catch(e error, handler func(err error)) CatchHandler
	CatchAll(handler func(err error)) CatchHandler
	FinalHandler
}
