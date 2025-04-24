# ToDo CLi and Webserver with GO / HTMX / Templ

Flawless ToDo app which supports cli-linux terminal and web together. It has dynamic change with htmx and beautiful bootstrap css with templ. Lists will be saving to `todo.json` file and the file will be watched and loaded dynamically every change.

Simply do your task;

- List
- Add
- Delete
- Toggle status
- Edit / only for CLi for now ...

## Components:

- Go Core
- Echo Framework for webserver
- Htmx
- Templ

## Code:

[Github Source](http://github.com/mozkaya1/todo-cli-web)

## Screenshot:

### Webserver

![Webserver](./static/webserver.png)

### Cli Terminal // Linux

![Cli](./static/cli.png)

## Usage:

### Cli --

```bash
  go  run main.go -h
```

```bash
Usage of /home/musti/.cache/go-build/42/421d5e23fb3f596265313754cf378e1797b3ce698eab60058d40dde53c314e3e-d/main:
  -Add string
        Adding new todo with title..
  -Delete int
        Delete task with ID (default -1)
  -Edit string
        Editing task with ID. id:New Title
  -List
        List all Tasks
  -Toggle int
        Toggle task with ID (default -1)

```

```bash
[musti@musti-deputyp25 todo-cli]$ go run main.go -List
┌───┬──────────────────────────┬───────────┬────────────────────────────────┬────────────────────────────────┐
│ # │          Title           │ Completed │          Created Time          │         completed_time         │
├───┼──────────────────────────┼───────────┼────────────────────────────────┼────────────────────────────────┤
│ 0 │ First Task !             │ ✅        │ Mon, 14 Apr 2025 13:45:58 CEST │ Mon, 14 Apr 2025 17:51:56 CEST │
│ 1 │ Things to be completed.. │ ⭕        │ Mon, 14 Apr 2025 13:46:13 CEST │                                │
│ 2 │ New Task created..       │ ✅        │ Mon, 14 Apr 2025 13:46:37 CEST │ Mon, 14 Apr 2025 13:46:49 CEST │
│ 3 │ This is last Task..      │ ⭕        │ Mon, 14 Apr 2025 13:46:42 CEST │                                │
└───┴──────────────────────────┴───────────┴────────────────────────────────┴────────────────────────────────┘

[musti@musti-deputyp25 todo-cli]$ go run main.go -Toggle 3
[musti@musti-deputyp25 todo-cli]$ go run main.go -List
┌───┬──────────────────────────┬───────────┬────────────────────────────────┬────────────────────────────────┐
│ # │          Title           │ Completed │          Created Time          │         completed_time         │
├───┼──────────────────────────┼───────────┼────────────────────────────────┼────────────────────────────────┤
│ 0 │ First Task !             │ ✅        │ Mon, 14 Apr 2025 13:45:58 CEST │ Mon, 14 Apr 2025 17:51:56 CEST │
│ 1 │ Things to be completed.. │ ⭕        │ Mon, 14 Apr 2025 13:46:13 CEST │                                │
│ 2 │ New Task created..       │ ✅        │ Mon, 14 Apr 2025 13:46:37 CEST │ Mon, 14 Apr 2025 13:46:49 CEST │
│ 3 │ This is last Task..      │ ✅        │ Mon, 14 Apr 2025 13:46:42 CEST │ Mon, 14 Apr 2025 17:54:09 CEST │
└───┴──────────────────────────┴───────────┴────────────────────────────────┴────────────────────────────────┘

```

### webserver --

Run webserver with below command...

```bash
go run server/server.go

```

Server is running on port : 3000

```bash
[musti@musti-deputyp25 todo-cli]$ go run server/server.go 
2025/04/14 17:46:34 Watching for changes in todo.json

   ____    __
  / __/___/ /  ___
 / _// __/ _ \/ _ \
/___/\__/_//_/\___/ v4.13.3
High performance, minimalist Go web framework
https://echo.labstack.com
____________________________________O/_______
                                    O\
⇨ http server started on [::]:3000

2025/04/14 17:51:56 File reloaded successfully!
2025/04/14 17:51:56 File reloaded successfully!

```

On Browser

`http://localhost:3000`

List - Json
`http://localhost:3000/list`

### Docker --
Repo is public in my docker cloud : `ozkaya84/todo-cli-web`
_Run docker with your initial local file_
Create an empty todo.json file:
```bash
echo "[{}]" > todo.json
```
Run docker on your folder with binding your `todo.json` file:

```bash
docker run --name   todo-web -v ./todo.json:/app/todo.json  -p 8080:3000    -d  ozkaya84/todo-cli-web:latest
```
and it is now available on  your  localhost `http://localhost:8080`
