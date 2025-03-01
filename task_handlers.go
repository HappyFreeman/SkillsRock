package main

import (
	"github.com/HappyFreeman/SkillsRock/internal/database"
	"github.com/gofiber/fiber/v2"
	"github.com/jackc/pgx/v5/pgtype"
)

func TaskHandlers(route fiber.Router, db *database.Queries) {
	route.Route("/tasks", func(api fiber.Router) {

		api.Get("/", func(c *fiber.Ctx) error {
			tasks, err := db.GetTasks(c.Context()) // Передаём контекст запроса
			if err != nil {
				return c.Status(fiber.StatusInternalServerError).JSON(fiber.Map{
					"error": "Не удалось получить задачи",
				})
			}
			return c.JSON(tasks) // Возвращаем задачи в формате JSON
		}).Name("index") // /tasks/ (name: tasks.index)


		//api.Get("/create", handler).Name("create") // /tasks/create (name: tasks.create) // для api не надо

		api.Post("/", func(c *fiber.Ctx) error {
			type parameters struct {
				Title       string `json:"title"`
				Description pgtype.Text `json:"description"`
				Status      pgtype.Text `json:"status"`
			}
			params := parameters{}

			if err := c.BodyParser(&params); err != nil {
				return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{"error": "Неверный формат данных " + err.Error()})
			}

			if !params.Status.Valid {
				params.Status = pgtype.Text{String: "new", Valid: true}
			}

			task, err := db.CreateTask(c.Context(), database.CreateTaskParams{
				Title:       params.Title,
                Description: params.Description,
                Status:      params.Status,
			})

			if err != nil {
				return c.Status(fiber.StatusInternalServerError).JSON(fiber.Map{"error": "Ошибка создания задачи " + err.Error()})
			}
		
			return c.JSON(task)
		}).Name("store") // /tasks/ (name: tasks.store)

		api.Get("/:task", func(c *fiber.Ctx) error {
			id, _ := c.ParamsInt("task")
			task, err := db.GetTaskById(c.Context(), int32(id))
			if err != nil {
				return c.Status(fiber.StatusNotFound).JSON(fiber.Map{"error": "Задача не найдена"})
			}
			return c.JSON(task)
		}).Name("show") // /tasks/ (name: tasks.show)

		//api.Get("/:task/edit", handler).Name("edit") // /tasks/{task}/edit (name: tasks.edit) // для api не надо
		
		api.Patch("/:task", func(c *fiber.Ctx) error {
			type parameters struct {
				Title       string `json:"title"`
				Description pgtype.Text `json:"description"`
				Status      pgtype.Text `json:"status"`
			}

			params := parameters{}

			if err := c.BodyParser(&params); err != nil {
				return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{"error": "Неверный формат данных"})
			}

			id, err := c.ParamsInt("task")

			if err != nil {
				return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{"error": "Некорректный id"})
			}

			task, err := db.GetTaskById(c.Context(), int32(id))

			if err != nil {
				return c.Status(fiber.StatusNotFound).JSON(fiber.Map{"error": "Задача не найдена"})
			}

			//-- может есть вариант получше?
			if !params.Status.Valid {
                params.Status = task.Status
            }
			if !params.Description.Valid {
                params.Description = task.Description
            }
			if params.Title == "" {
                params.Title = task.Title
            }
			//--

			updatedTask, err := db.UpdateTask(c.Context(), database.UpdateTaskParams{
				ID: int32(id),
				Title: params.Title,
				Description: params.Description,
                Status:      params.Status,
			}) 

			if err != nil {
				return c.Status(fiber.StatusInternalServerError).JSON(fiber.Map{"error": "Ошибка обновления"})
			}
		
			return c.JSON(updatedTask)
		}).Name("update") // /tasks/{task} (name: tasks.update) // Put меняет объект целиком | Patch только те поля которые передали
		
		api.Delete("/:task", func(c *fiber.Ctx) error {
			id, err := c.ParamsInt("task")
			if err != nil {
                return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{"error": "Некорректный id"})
            }

			_, err = db.GetTaskById(c.Context(), int32(id))
			if err != nil {
				return c.Status(fiber.StatusNotFound).JSON(fiber.Map{"error": "Задача не найдена"})
			}
			if err := db.DeleteTask(c.Context(), int32(id)); err != nil {
				return c.Status(fiber.StatusInternalServerError).JSON(fiber.Map{"error": "Ошибка удаления"})
			}
		
			return c.Status(fiber.StatusNoContent).Send(nil)
			
		}).Name("destroy") // /tasks/{task} (name: tasks.destroy)
		
	}, "tasks.")
}