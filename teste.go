package main

import (
	"fmt"
	"imaginary-query-handler/query"
)

func main() {

	processor := query.NewQueryProcessor()
	processor.AddQueryHandler(HelloQueryHandler{Database: "123 teste" , Name: "hello.ola" } )
	/*error := queryProcessor.AddQueryHandler(HelloQueryHandler{Database: "123 teste" , Name: "ola" } )
	if error !=nil {
		println("ERROR:" , error.Error() )
	}*/
	queryEngine := query.Engine{Processor: processor}

	//mapRow := map[string]interface{}{"nome": "SR ZE", "fone": 995877668, "email": "srze@gmail.com", "valor": 25455.50}

	//ver uma forma de representar as querys ?  > 0  ==  like    in  etc

	var d query.Direction ="ASCSSS"
	if err := d.IsValid(); err != nil{
	   println(err.Error())
	}

	sort1 := query.SortParameter{Field: "nome" , Direction: query.DESC}
	sort2 := query.SortParameter{Field: "data_cadastro" , Direction: query.ASC}
	filter1 :=query.FilterParameter{Field: "nome" , Operator: query.EQ , Value: "valter"}
	filter2 :=query.FilterParameter{Field: "data_cadastro" , Operator: query.NE , Value: "2020-01-20"}
	filter3 :=query.FilterParameter{Field: "valor" , Operator: query.GT , Value: "20800.67"}
	queryRequest := query.Resquest{UUID: "dsdsdsdsds"  , Domain: "XYZ"  , Namespace: "contato" , QueryType: "contatos" ,
		Sort:[]query.SortParameter{sort1,sort2}, Filter: []query.FilterParameter{filter1,filter2,filter3} , Page: 1 , Size: 5 }
	result := queryEngine.Query("hello.ola", queryRequest )
	fmt.Println(result)
	//fmt.Println(error)
}

type HelloQueryHandler struct {
	Database string
	Name     string
}

func (ola HelloQueryHandler) Handle(queryRequest query.Resquest) query.Response {

	ola.Database = "TESTE"
	fmt.Println("HelloQueryHandler")
	fmt.Println(queryRequest)
	fmt.Println(queryRequest.Sort)
	fmt.Println(queryRequest.Filter)
	var queryResponse = query.Response{UUID: "2332323" , RequestUUID: queryRequest.UUID}
	return queryResponse
}

func (ola HelloQueryHandler) GetName()  string {

	return ola.Name
}
