package main

import (
	"encoding/json"
	"fmt"
	"net/http"
	"time"

	"github.com/graphql-go/graphql"
)

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

var TaskList []Task

func init() {
	TaskList = make([]Task, 0)
}

// define schema, with our rootQuery and rootMutation
var schema, _ = graphql.NewSchema(graphql.SchemaConfig{
	Query:    rootQuery,
	Mutation: rootMutation,
})

func executeQuery(query string, schema graphql.Schema) *graphql.Result {
	result := graphql.Do(graphql.Params{
		Schema:        schema,
		RequestString: query,
	})
	if len(result.Errors) > 0 {
		fmt.Printf("wrong result, unexpected errors: %v", result.Errors)
	}
	return result
}

func main() {
	http.HandleFunc("/graphql", func(w http.ResponseWriter, r *http.Request) {
		result := executeQuery(r.URL.Query()["query"][0], schema)
		json.NewEncoder(w).Encode(result)
	})

	// Display some basic instructions
	fmt.Println("Now server is running on port 8080")
	fmt.Println("Get task list: curl -g 'http://localhost:8080/graphql?query={taskList(status:todo){id,content,status,updated_at,start_at,finish_at}}'")
	fmt.Println("Get done task list with week: curl -g 'http://localhost:8080/graphql?query={taskWeekList(week:201705){id,content,status,finish_at}}'")
	fmt.Println("Create new todo: curl -g 'http://localhost:8080/graphql?query=mutation+_{createTask(content:\"My+new+todo\"){id}}'")
	fmt.Println("Update task: curl -g 'http://localhost:8080/graphql?query=mutation+_{updateTask(id:1,content:\"my+new+content\"){id}}'")
	fmt.Println("Update task status: curl -g 'http://localhost:8080/graphql?query=mutation+_{updateTaskStatus(id:1,status:doing){id}}'")
	fmt.Println("delete task: curl -g 'http://localhost:8080/graphql?query=mutaion+_{deleteTask(id:1){id}}'")

	http.ListenAndServe(":8080", nil)
}
