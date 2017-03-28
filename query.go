package main

import (
	"github.com/graphql-go/graphql"
)

var taskQueryType = graphql.NewObject(graphql.ObjectConfig{
	Name: "Task",
	Fields: graphql.Fields{
		"id": &graphql.Field{
			Type: graphql.Int,
		},
		"content": &graphql.Field{
			Type: graphql.String,
		},
		"status": &graphql.Field{
			Type: graphql.String,
		},
		"updated_at": &graphql.Field{
			Type: graphql.String,
		},
		"start_at": &graphql.Field{
			Type: graphql.String,
		},
		"finish_at": &graphql.Field{
			Type: graphql.String,
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
					Type: graphql.String,
				},
			},
			Resolve: func(p graphql.ResolveParams) (interface{}, error) {
				return TaskList, nil
			},
		},

		// 按周获取任务列表
		"taskWeekList": &graphql.Field{
			Type:        graphql.NewList(taskQueryType),
			Description: "获取按周汇总的任务列表（周报）",
			Args: graphql.FieldConfigArgument{
				"week": &graphql.ArgumentConfig{
					Type: graphql.Int,
				},
			},
			Resolve: func(p graphql.ResolveParams) (interface{}, error) {
				return TaskList, nil
			},
		},
	},
})
