# tasklist
Tasklist it is a simple task manager api.</br>
You can build your Frontend with any technology you want.</br>

### API

With this API you can build your own task manager.</br>

- GET


    http://localhost:3000/tasks

- POST
  

    {
    "name": "My task",
    "description": "This is a task"
    }


- DELETE


    {
        "hash": "a6cab625-88d8-4e5d-961b-f3fe984066ce"
    }


Test our client.http file to validade this API.</br>

### How to run

- Clone this repository
- Run `go run bin/server/main.go`