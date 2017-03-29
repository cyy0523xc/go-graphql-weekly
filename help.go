package main

const (
	usageHelp = `

Help:
    curl -g 'http://localhost:port/'

查看Schema（其中queryType可以换成mutationType）:
    curl -g 'http://localhost:port/graphql?query={__schema{queryType{fields{name,description,type{description},args{type{name},description}}}}}'   

查看所有的types：
    curl -g 'http://localhost:port/graphql?query={__schema{types{name,description}}}'

查看特定的Type：
    curl -g 'http://localhost:port/graphql?query={__type(name:"Task"){fields{name,description,type{name,description}}}}'
    curl -g 'http://localhost:port/graphql?query={__type(name:"RootQuery"){fields{name,description,type{name,description}}}}'

Get task list: 
    curl -g 'http://localhost:port/graphql?query={taskList(status:todo){id,content,status,updated_at,start_at,finish_at}}'

Get done task list with week: 
    curl -g 'http://localhost:port/graphql?query={taskWeekList(week:201705){id,content,status,finish_at}}'

Create new task: 
    curl -g 'http://localhost:port/graphql?query=mutation+_{createTask(content:"My+new+todo"){id}}'

Update task: 
    curl -g 'http://localhost:port/graphql?query=mutation+_{updateTask(id:1,content:"my+new+content"){id}}'

Update task status: 
    curl -g 'http://localhost:port/graphql?query=mutation+_{updateTaskStatus(id:1,status:doing){id}}'

Delete task: 
    curl -g 'http://localhost:port/graphql?query=mutaion+_{deleteTask(id:1){id}}'
`
)
