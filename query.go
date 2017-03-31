package main

import (
	"github.com/graphql-go/graphql"
)

// root query
// we just define a trivial example here, since root query is required.
// Test with curl
// curl -g 'http://localhost:8080/graphql?query={lastTodo{id,text,done}}'
var rootQuery = graphql.NewObject(graphql.ObjectConfig{
	Name: "RootQuery",
	Fields: graphql.Fields{
		// 获取任务列表
		"taskList": &graphql.Field{
			Type:        graphql.NewList(taskType),
			Description: "获取任务列表",
			Args: graphql.FieldConfigArgument{
				"status": &graphql.ArgumentConfig{
					Type:         taskStatusType,
					DefaultValue: StatusTodo,
					Description:  "状态，注意这里是enum类型，不要传字符串",
				},
			},
			Resolve: func(params graphql.ResolveParams) (interface{}, error) {
				if params.Args["status"] != nil {
					taskList := make([]Task, 0)
					status := params.Args["status"].(TaskStatus)

					for _, task := range TaskList {
						if task.Status == status {
							taskList = append(taskList, task)
						}
					}
					return taskList, nil
				}

				return TaskList, nil
			},
		},

		// 按周获取任务列表
		"taskWeekList": &graphql.Field{
			Type:        graphql.NewList(taskType),
			Description: "获取按周汇总的任务列表（周报）",
			Args: graphql.FieldConfigArgument{
				"week": &graphql.ArgumentConfig{
					Type:        graphql.NewNonNull(graphql.Int),
					Description: "年份周次，如2017年第5周，则值为201705",
				},
			},
			Resolve: func(params graphql.ResolveParams) (interface{}, error) {
				week := params.Args["week"].(int)

				taskList := make([]Task, 0)
				for _, task := range TaskList {
					if task.FinishWeek == week {
						taskList = append(taskList, task)
					}
				}

				return taskList, nil
			},
		},
	},
})
