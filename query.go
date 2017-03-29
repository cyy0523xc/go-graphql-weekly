package main

import (
	"github.com/graphql-go/graphql"
)

var taskQueryType = graphql.NewObject(graphql.ObjectConfig{
	Name:        "Task",
	Description: "任务的返回结构",
	Fields: graphql.Fields{
		"id": &graphql.Field{
			Type:        graphql.Int,
			Description: "任务ID",
		},
		"content": &graphql.Field{
			Type:        graphql.String,
			Description: "任务内容",
		},
		"status": &graphql.Field{
			Type:        taskStatusType,
			Description: "任务状态",
		},
		"updated_at": &graphql.Field{
			Type:        graphql.String,
			Description: "任务最后更新时间（更新内容）",
		},
		"start_at": &graphql.Field{
			Type:        graphql.String,
			Description: "任务开始时间",
		},
		"finish_at": &graphql.Field{
			Type:        graphql.String,
			Description: "任务完成时间",
		},
	},
})

// root query
// we just define a trivial example here, since root query is required.
// Test with curl
// curl -g 'http://localhost:8080/graphql?query={lastTodo{id,text,done}}'
var rootQuery = graphql.NewObject(graphql.ObjectConfig{
	Name: "RootQuery",
	Fields: graphql.Fields{
		// 获取任务列表
		"taskList": &graphql.Field{
			Type:        graphql.NewList(taskQueryType),
			Description: "获取任务列表",
			Args: graphql.FieldConfigArgument{
				"status": &graphql.ArgumentConfig{
					Type:         taskStatusType,
					DefaultValue: StatusTodo,
					Description:  "状态，注意这里是enum类型，不要传字符串",
				},
			},
			Resolve: func(params graphql.ResolveParams) (interface{}, error) {
				status := params.Args["status"].(TaskStatus)

				taskList := make([]Task, 0)
				for _, task := range TaskList {
					if task.Status == status {
						taskList = append(taskList, task)
					}
				}

				return taskList, nil
			},
		},

		// 按周获取任务列表
		"taskWeekList": &graphql.Field{
			Type:        graphql.NewList(taskQueryType),
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
