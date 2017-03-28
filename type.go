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
