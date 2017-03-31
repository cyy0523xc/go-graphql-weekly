package main

import (
	"github.com/graphql-go/graphql"
)

var taskStatusType = graphql.NewEnum(graphql.EnumConfig{
	Name:        "TaskStatus",
	Description: "任务状态类型",
	Values: graphql.EnumValueConfigMap{
		"todo": &graphql.EnumValueConfig{
			Value: StatusTodo,
		},
		"doing": &graphql.EnumValueConfig{
			Value: StatusDoing,
		},
		"done": &graphql.EnumValueConfig{
			Value: StatusDone,
		},
	},
})

// taskType 任务的返回值结构
var taskType = graphql.NewObject(graphql.ObjectConfig{
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
