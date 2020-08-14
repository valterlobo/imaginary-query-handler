package query

import "fmt"

type Engine struct {
	Config         string
	Processor *Processor
}

func (queryEngine *Engine) Query(queryName string, queryRequest Resquest) Response {

	fmt.Println("\n /COMMAND ENGINE" + queryEngine.Config)
	result := queryEngine.Processor.ProcessQueryHandler(queryName, queryRequest)
	fmt.Println("\n /QUERY ENGINE" + queryEngine.Config)
	fmt.Println("\n /QUERY ENGINE RESULT ", result)
	return result
}
