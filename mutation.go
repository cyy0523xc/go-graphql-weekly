package main

import (
	"errors"
	"fmt"
	"time"

	"github.com/graphql-go/graphql"
)

// root mutation
var rootMutation = graphql.NewObject(graphql.ObjectConfig{
	Name: "RootMutation",
	Fields: graphql.Fields{
		"createTask": &graphql.Field{
			Type:        taskType,
			Description: "创建新的任务，默认为todo状态",
			Args: graphql.FieldConfigArgument{
				"content": &graphql.ArgumentConfig{
					Type:        graphql.NewNonNull(graphql.String),
					Description: "任务内容",
				},
			},
			Resolve: func(params graphql.ResolveParams) (interface{}, error) {
				text, _ := params.Args["content"].(string)

				n := len(TaskList)
				var newId uint32
				if n == 0 {
					newId = 1
				} else {
					newId = TaskList[n-1].Id + 1
				}

				// perform mutation operation here
				currentTime := time.Now()
				newTask := Task{
					Id:        newId,
					Content:   text,
					CreatedAt: currentTime,
					UpdatedAt: currentTime,
					Status:    StatusTodo,
				}
				TaskList = append(TaskList, newTask)

				return newTask, nil
			},
		}, // end of createTask

		"updateTask": &graphql.Field{
			Type:        taskType,
			Description: "更新任务的内容",
			Args: graphql.FieldConfigArgument{
				"id": &graphql.ArgumentConfig{
					Type:        graphql.NewNonNull(graphql.Int),
					Description: "任务ID",
				},
				"content": &graphql.ArgumentConfig{
					Type:        graphql.NewNonNull(graphql.String),
					Description: "任务内容",
				},
			},
			Resolve: func(params graphql.ResolveParams) (interface{}, error) {
				text, _ := params.Args["content"].(string)
				_id, _ := params.Args["id"].(int)
				id := uint32(_id)

				var resTask Task
				for index, task := range TaskList {
					if id == task.Id {
						resTask = task
						TaskList[index].Content = text
						TaskList[index].UpdatedAt = time.Now()
					}
				}

				if resTask.Id == 0 {
					return nil, errors.New(fmt.Sprintf("id=%d不存在", id))
				}

				return resTask, nil
			},
		}, // end of updateTask

		"updateTaskStatus": &graphql.Field{
			Type:        taskType,
			Description: "更新任务的状态",
			Args: graphql.FieldConfigArgument{
				"id": &graphql.ArgumentConfig{
					Type:        graphql.NewNonNull(graphql.Int),
					Description: "任务ID",
				},
				"status": &graphql.ArgumentConfig{
					Type:        graphql.NewNonNull(taskStatusType),
					Description: "任务状态，值为：todo, doing, done",
				},
			},
			Resolve: func(params graphql.ResolveParams) (interface{}, error) {
				status, _ := params.Args["status"].(TaskStatus)
				_id, _ := params.Args["id"].(int)
				id := uint32(_id)

				var resTask Task
				for index, task := range TaskList {
					if id == task.Id {
						resTask = task
						oldStatus := TaskList[index].Status
						if oldStatus == StatusTodo {
							if status == StatusDoing {
								// 任务开始
								TaskList[index].Status = status
								TaskList[index].StartAt = time.Now()
							} else {
								return nil, errors.New("新状态错误")
							}
						} else if oldStatus == StatusDoing {
							if status == StatusDone {
								// 任务完成
								currTime := time.Now()
								TaskList[index].Status = status
								TaskList[index].FinishAt = currTime
								year, week := currTime.ISOWeek()
								TaskList[index].FinishWeek = year*100 + week
							} else {
								return nil, errors.New("新状态错误")
							}
						} else {
							return nil, errors.New("已经完成的任务不能修改状态")
						}
					}
				}

				if resTask.Id == 0 {
					return nil, errors.New(fmt.Sprintf("id=%d不存在", id))
				}

				return resTask, nil
			},
		}, // end of updateTaskStatus

		"deleteTask": &graphql.Field{
			Type:        taskType,
			Description: "删除任务",
			Args: graphql.FieldConfigArgument{
				"id": &graphql.ArgumentConfig{
					Type:        graphql.NewNonNull(graphql.Int),
					Description: "任务ID",
				},
			},
			Resolve: func(params graphql.ResolveParams) (interface{}, error) {
				n := len(TaskList)
				if n == 0 {
					return nil, errors.New("任务列表为空")
				}

				_id, _ := params.Args["id"].(int)
				id := uint32(_id)

				var resTask Task
				for index, task := range TaskList {
					if id == task.Id {
						resTask = task
						if n == 1 {
							TaskList = TaskList[0:0]
						} else if index < n-1 {
							TaskList = append(TaskList[0:index-1], TaskList[index+1:n-1]...)
						} else {
							TaskList = TaskList[0 : index-1]
						}

						break
					}
				}

				if resTask.Id == 0 {
					return nil, errors.New(fmt.Sprintf("id=%d不存在", id))
				}

				return resTask, nil
			},
		}, // end of deleteTask
	},
})
