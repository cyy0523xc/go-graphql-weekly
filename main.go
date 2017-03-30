package main

import (
	"encoding/json"
	"fmt"
	"net/http"
	"time"

	"github.com/graphql-go/graphql"
)

// 端口号
const port = "8088"

type Task struct {
	Id         uint32     `json:"id"`
	Content    string     `json:"content"`
	Remark     string     `json:"remark"`
	Status     TaskStatus `json:"status"`
	CreatedAt  time.Time  `json:"created_at"`  // 创建时间
	UpdatedAt  time.Time  `json:"updated_at"`  // 更新时间
	StartAt    time.Time  `json:"start_at"`    // 开始时间
	FinishAt   time.Time  `json:"finish_at"`   // 完成时间
	FinishWeek int        `json:"finish_week"` // 完成的周次，例如2017年第5周完成，则对应的值为201705
}

// 任务状态类型
type TaskStatus uint8

// 任务状态定义
const (
	StatusTodo  TaskStatus = 0
	StatusDoing TaskStatus = 1
	StatusDone  TaskStatus = 2
)

// query结构
type queryParams struct {
	Query         string                 `json:"query"`
	OperationName string                 `json:"operationName"`
	Variables     map[string]interface{} `json:"variables"`
}

var TaskList []Task

func init() {
	TaskList = make([]Task, 0)
}

// define schema, with our rootQuery and rootMutation
var schema, _ = graphql.NewSchema(graphql.SchemaConfig{
	Query:    rootQuery,
	Mutation: rootMutation,
})

func executeQuery(params queryParams, schema graphql.Schema) *graphql.Result {
	result := graphql.Do(graphql.Params{
		Schema:         schema,
		RequestString:  params.Query,
		VariableValues: params.Variables,
		OperationName:  params.OperationName,
	})
	if len(result.Errors) > 0 {
		fmt.Printf("wrong result, unexpected errors: %v", result.Errors)
	}
	return result
}

func main() {
	http.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) {
		w.Write([]byte(usageHelp))
	})

	// 使用url参数来传递参数
	http.HandleFunc("/graphql", func(w http.ResponseWriter, r *http.Request) {
		handle("url", w, r)
	})

	// 使用json的格式，通过post等方式来传递参数
	http.HandleFunc("/graphql-json", func(w http.ResponseWriter, r *http.Request) {
		handle("json", w, r)
	})

	fmt.Println("Now server is running on port " + port)
	fmt.Println(usageHelp)

	http.ListenAndServe(":"+port, nil)
}

func handle(queryType string, w http.ResponseWriter, r *http.Request) {
	w.Header().Add("Access-Control-Allow-Origin", "*")
	w.Header().Add("Access-Control-Allow-Methods", "POST,GET,OPTIONS,DELETE")
	w.Header().Add("Access-Control-Allow-Headers", "x-requested-with,content-type")
	if "OPTIONS" == r.Method {
		// 解决跨域的问题
		w.WriteHeader(http.StatusAccepted)
		w.Write([]byte(""))
	} else {
		var params queryParams
		if queryType == "url" {
			params = getQueryStringByURL(r)
		} else {
			params = getQueryStringByJson(r)
		}
		result := executeQuery(params, schema)
		json.NewEncoder(w).Encode(result)
	}
}

func getQueryStringByURL(r *http.Request) (q queryParams) {
	q.Query = r.URL.Query()["query"][0]
	return q
}

func getQueryStringByJson(r *http.Request) (q queryParams) {
	_ = json.NewDecoder(r.Body).Decode(&q)
	return q
}
