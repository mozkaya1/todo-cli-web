package main

import (
	"log"
	"net/http"
	"os"
	"strconv"
	"time"

	"github.com/fsnotify/fsnotify"
	"github.com/labstack/echo/v4"
	"github.com/mozkaya1/todo-cli/internal"
	"github.com/mozkaya1/todo-cli/storage"
	"github.com/mozkaya1/todo-cli/view"
)

func RunServer(data *internal.Todos, storage *storage.Storage[internal.Todos]) {
	s := echo.New()

	// Middleware to make todos and storage available in context
	s.Use(func(next echo.HandlerFunc) echo.HandlerFunc {
		return func(c echo.Context) error {
			c.Set("todos", data)
			c.Set("storage", storage)
			return next(c)
		}
	})

	// Routes
	s.GET("/", func(c echo.Context) error {
		data := c.Get("todos").(*internal.Todos)
		return view.List(*data).Render(c.Request().Context(), c.Response())
	})

	// Json list route
	s.GET("/list", HandleList(data))
	s.Use(func(next echo.HandlerFunc) echo.HandlerFunc {
		return func(c echo.Context) error {
			c.Set("todos", data)
			return next(c)
		}
	})

	// API routes
	s.POST("/todos", handleAddTodo)
	s.PUT("/todos/:id", handleEditTodo)
	s.DELETE("/todos/:id", handleDeleteTodo)
	s.PATCH("/todos/:id/toggle", handleToggleTodo)

	s.Logger.Fatal(s.Start(":3000"))
}

// Handlers with storage persistence
func handleAddTodo(c echo.Context) error {
	data := c.Get("todos").(*internal.Todos)
	storage := c.Get("storage").(*storage.Storage[internal.Todos])

	title := c.FormValue("title")
	if title == "" {
		return echo.NewHTTPError(http.StatusBadRequest, "Title cannot be empty")
	}

	data.Add(title)

	if err := storage.SaveFile(*data); err != nil {
		log.Printf("Failed to save todos: %v", err)
		return echo.NewHTTPError(http.StatusInternalServerError, "Failed to save todos")
	}

	return c.Redirect(http.StatusSeeOther, "/")
}

func handleEditTodo(c echo.Context) error {
	data := c.Get("todos").(*internal.Todos)
	storage := c.Get("storage").(*storage.Storage[internal.Todos])

	id, err := strconv.Atoi(c.Param("id"))
	if err != nil {
		return echo.NewHTTPError(http.StatusBadRequest, "Invalid ID")
	}

	title := c.FormValue("title")
	if title == "" {
		return echo.NewHTTPError(http.StatusBadRequest, "Title cannot be empty")
	}

	if err := data.Edit(id, title); err != nil {
		return echo.NewHTTPError(http.StatusNotFound, "Todo not found")
	}

	if err := storage.SaveFile(*data); err != nil {
		log.Printf("Failed to save todos: %v", err)
		return echo.NewHTTPError(http.StatusInternalServerError, "Failed to save todos")
	}

	return c.Redirect(http.StatusSeeOther, "/")
}

func handleDeleteTodo(c echo.Context) error {
	data := c.Get("todos").(*internal.Todos)
	storage := c.Get("storage").(*storage.Storage[internal.Todos])

	id, err := strconv.Atoi(c.Param("id"))
	if err != nil {
		return echo.NewHTTPError(http.StatusBadRequest, "Invalid ID")
	}

	if err := data.Delete(id); err != nil {
		return echo.NewHTTPError(http.StatusNotFound, "Todo not found")
	}

	if err := storage.SaveFile(*data); err != nil {
		log.Printf("Failed to save todos: %v", err)
		return echo.NewHTTPError(http.StatusInternalServerError, "Failed to save todos")
	}

	return c.Redirect(http.StatusSeeOther, "/")
}

func handleToggleTodo(c echo.Context) error {
	data := c.Get("todos").(*internal.Todos)
	storage := c.Get("storage").(*storage.Storage[internal.Todos])

	id, err := strconv.Atoi(c.Param("id"))
	if err != nil {
		return echo.NewHTTPError(http.StatusBadRequest, "Invalid ID")
	}

	if err := data.Toggle(id); err != nil {
		return echo.NewHTTPError(http.StatusNotFound, "Todo not found")
	}

	if err := storage.SaveFile(*data); err != nil {
		log.Printf("Failed to save todos: %v", err)
		return echo.NewHTTPError(http.StatusInternalServerError, "Failed to save todos")
	}

	return c.Redirect(http.StatusSeeOther, "/")
}

func HandleList(data *internal.Todos) echo.HandlerFunc {
	return func(c echo.Context) error {

		return c.JSON(http.StatusOK, *data)
	}
}

func main() {
	t := internal.Todos{}
	s := storage.Storage[internal.Todos]{Filename: "todo.json"}

	// Load initial data
	if err := s.LoadFile(&t); err != nil {
		// If file doesn't exist, create an empty one
		if os.IsNotExist(err) {
			if err := s.SaveFile(t); err != nil {
				log.Fatalf("Failed to create initial todo file: %v", err)
			}
		} else {
			log.Fatal(err)
		}
	}

	go RunServer(&t, &s)

	// Watch for file changes
	watcher, err := fsnotify.NewWatcher()
	if err != nil {
		log.Fatal(err)
	}
	defer watcher.Close()

	if err := watcher.Add(s.Filename); err != nil {
		log.Fatal(err)
	}

	log.Println("Watching for changes in", s.Filename)

	for {
		select {
		case event, ok := <-watcher.Events:
			if !ok {
				return
			}
			if event.Op&fsnotify.Write == fsnotify.Write || event.Op&fsnotify.Rename == fsnotify.Rename {
				time.Sleep(100 * time.Millisecond)
				if err := s.LoadFile(&t); err != nil {
					log.Printf("Failed to reload file: %v", err)
				} else {
					log.Println("File reloaded successfully!")
				}
			}

		case err, ok := <-watcher.Errors:
			if !ok {
				return
			}
			log.Println("Watcher error:", err)
		}
	}
}
