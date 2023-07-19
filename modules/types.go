package modules

type Request struct {
	Url    string
	Worker func(url string) WorkResult
}

type WorkResult struct {
	Requests []Request
	Items    []interface{}
}

func NilWorker(url string) WorkResult {
	return WorkResult{}
}
