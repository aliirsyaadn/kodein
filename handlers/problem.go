package handlers

import (
	"github.com/aliirsyaadn/kodein/entity"
	"github.com/aliirsyaadn/kodein/internal/log"
	"github.com/aliirsyaadn/kodein/services/problem"
	"github.com/gofiber/fiber/v2"
	_ "github.com/lib/pq"
)

const problemHandlers = "Problem Handlers"

type problemHandlerImpl struct {
	problemService problem.Service
}

func NewProblemHandler(problemService problem.Service) *problemHandlerImpl {
	return &problemHandlerImpl{
		problemService: problemService,
	}
}

func (h *problemHandlerImpl) Register(r fiber.Router) {
	r = r.Group("/problem")

	r.Get("/", h.GetProblems)
	r.Get("/:id", h.GetProblemByID)
	r.Post("/", h.CreateProblem)
	r.Post("/:id", h.UpdateProblem)
	r.Delete("/:id", h.DeleteProblem)
}

func (h *problemHandlerImpl) GetProblems(c *fiber.Ctx) error {
	problems, err := h.problemService.GetProblems(c.Context())

	if err != nil {
		log.ErrorDetail(problemHandlers, "error from services %v", err)
		return err
	}

	return c.JSON(problems)
}

func (h *problemHandlerImpl) GetProblemByID(c *fiber.Ctx) error {
	id := c.Params("id")
	problem, err := h.problemService.GetProblemByID(c.Context(), id)

	if err != nil {
		log.ErrorDetail(problemHandlers, "error from services %v", err)
		return err
	}

	return c.JSON(problem)
}

func (h *problemHandlerImpl) CreateProblem(c *fiber.Ctx) error {
	req := new(entity.CreateProblemRequest)
	if err := c.BodyParser(req); err != nil {
		log.ErrorDetail(problemHandlers, "error parse body %v", err)
		return err
	}

	res, err := h.problemService.CreateProblem(c.Context(), *req)
	if err != nil {
		log.ErrorDetail(problemHandlers, "error from services %v", err)
		return err
	}

	return c.JSON(res)
}

func (h *problemHandlerImpl) UpdateProblem(c *fiber.Ctx) error {
	id := c.Params("id")
	req := new(entity.UpdateProblemRequest)
	if err := c.BodyParser(req); err != nil {
		log.ErrorDetail(problemHandlers, "error parse body %v", err)
		return err
	}

	res, err := h.problemService.UpdateProblem(c.Context(), *req, id)
	if err != nil {
		log.ErrorDetail(problemHandlers, "error from services %v", err)
		return err
	}

	return c.JSON(res)
}

func (h *problemHandlerImpl) DeleteProblem(c *fiber.Ctx) error {
	id := c.Params("id")
	res, err := h.problemService.DeleteProblem(c.Context(), id)
	if err != nil {
		log.ErrorDetail(problemHandlers, "error from services %v", err)
		return err
	}

	return c.JSON(res)
}
