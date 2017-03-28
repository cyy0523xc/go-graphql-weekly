package main

import (
	"github.com/graphql-go/graphql"
)

var mutationType = graphql.NewObject(graphql.ObjectConfig{
	Name: "Mutation",
	Fields: graphql.Fields{
		"id": &graphql.Field{
			Type: graphql.Int,
		},
	},
})

// root mutation
var rootMutation = graphql.NewObject(graphql.ObjectConfig{
	Name: "RootMutation",
	Fields: graphql.Fields{
		"createTask": &graphql.Field{
			Type:        mutationType,
			Description: "创建新的todo任务",
			Args: graphql.FieldConfigArgument{
				"content": &graphql.ArgumentConfig{
					Type: graphql.NewNonNull(graphql.String),
				},
			},
			Resolve: func(params graphql.ResolveParams) (interface{}, error) {

				// marshall and cast the argument value
				text, _ := params.Args["content"].(string)

				// figure out new id
				n := len(TaskList)
				var newId uint32
				if n == 0 {
					newId = 1
				} else {
					newId = TaskList[n-1].Id + 1
				}

				// perform mutation operation here
				// for e.g. create a Todo and save to DB.
				newTask := Task{
					Id:      newId,
					Content: text,
					Status:  StatusTodo,
				}

				TaskList = append(TaskList, newTask)

				return newTask, nil
			},
		}, // end of createTask

		"updateTask": &graphql.Field{
			Type:        mutationType,
			Description: "更新todo任务",
			Args: graphql.FieldConfigArgument{
				"id": &graphql.ArgumentConfig{
					Type: graphql.NewNonNull(graphql.Int),
				},
				"content": &graphql.ArgumentConfig{
					Type: graphql.NewNonNull(graphql.String),
				},
			},
			Resolve: func(params graphql.ResolveParams) (interface{}, error) {

				// marshall and cast the argument value
				text, _ := params.Args["content"].(string)
				id, _ := params.Args["id"].(uint32)

				var resTask Task
				for index, task := range TaskList {
					if id == task.Id {
						resTask = task
						TaskList[index].Content = text
					}
				}

				return resTask, nil
			},
		}, // end of updateTask

		"updateTaskStatus": &graphql.Field{
			Type:        mutationType,
			Description: "更新任务的状态",
			Args: graphql.FieldConfigArgument{
				"id": &graphql.ArgumentConfig{
					Type: graphql.NewNonNull(graphql.Int),
				},
				"status": &graphql.ArgumentConfig{
					Type: graphql.NewNonNull(graphql.String),
				},
			},
			Resolve: func(params graphql.ResolveParams) (interface{}, error) {

				// marshall and cast the argument value
				status, _ := params.Args["status"].(string)
				id, _ := params.Args["id"].(uint32)

				var resTask Task
				for index, task := range TaskList {
					if id == task.Id {
						resTask = task
						TaskList[index].Status = StatusDoing
						_ = status
					}
				}

				return resTask, nil
			},
		}, // end of updateTaskStatus

		"deleteTask": &graphql.Field{
			Type:        mutationType,
			Description: "删除任务",
			Args: graphql.FieldConfigArgument{
				"id": &graphql.ArgumentConfig{
					Type: graphql.NewNonNull(graphql.Int),
				},
			},
			Resolve: func(params graphql.ResolveParams) (interface{}, error) {

				// marshall and cast the argument value
				id, _ := params.Args["id"].(uint32)

				var resTask Task
				n := len(TaskList)
				for index, task := range TaskList {
					if id == task.Id {
						resTask = task
						if index < n-1 {
							TaskList = append(TaskList[0:index-1], TaskList[index+1:n-1]...)
						} else {
							TaskList = TaskList[0 : index-1]
						}
					}
				}

				return resTask, nil
			},
		}, // end of deleteTask
	},
})
