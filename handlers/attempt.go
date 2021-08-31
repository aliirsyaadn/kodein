package handlers

import (
	"github.com/aliirsyaadn/kodein/entity"
	"github.com/aliirsyaadn/kodein/internal/log"
	"github.com/aliirsyaadn/kodein/services/attempt"
	"github.com/gofiber/fiber/v2"
	_ "github.com/lib/pq"
)

const attemptHandlers = "Attempt Handlers"

type attemptHandlerImpl struct {
	attemptService attempt.Service
}

func NewAttemptHandler(attemptService attempt.Service) *attemptHandlerImpl {
	return &attemptHandlerImpl{
		attemptService: attemptService,
	}
}

func (h *attemptHandlerImpl) Register(r fiber.Router) {
	r = r.Group("/attempt")

	r.Get("/member/:id", h.GetAttemptsByMemberID)
	r.Get("/:id", h.GetAttemptByID)
	r.Post("/", h.CreateAttempt)
	r.Post("/:id", h.UpdateAttempt)
	r.Delete("/:id", h.DeleteAttempt)
}

func (h *attemptHandlerImpl) GetAttemptsByMemberID(c *fiber.Ctx) error {
	attempts, err := h.attemptService.GetAttemptsByMemberID(c.Context(), c.Params("id"))

	if err != nil {
		log.ErrorDetail(attemptHandlers, "error from services %v", err)
		return err
	}

	return c.JSON(attempts)
}

func (h *attemptHandlerImpl) GetAttemptByID(c *fiber.Ctx) error {
	id := c.Params("id")
	attempt, err := h.attemptService.GetAttemptByID(c.Context(), id)

	if err != nil {
		log.ErrorDetail(attemptHandlers, "error from services %v", err)
		return err
	}

	return c.JSON(attempt)
}

func (h *attemptHandlerImpl) CreateAttempt(c *fiber.Ctx) error {
	req := new(entity.CreateAttemptRequest)
	if err := c.BodyParser(req); err != nil {
		log.ErrorDetail(attemptHandlers, "error parse body %v", err)
		return err
	}

	res, err := h.attemptService.CreateAttempt(c.Context(), *req)
	if err != nil {
		log.ErrorDetail(attemptHandlers, "error from services %v", err)
		return err
	}

	return c.JSON(res)
}

func (h *attemptHandlerImpl) UpdateAttempt(c *fiber.Ctx) error {
	id := c.Params("id")
	req := new(entity.UpdateAttemptRequest)
	if err := c.BodyParser(req); err != nil {
		log.ErrorDetail(attemptHandlers, "error parse body %v", err)
		return err
	}

	res, err := h.attemptService.UpdateAttempt(c.Context(), *req, id)
	if err != nil {
		log.ErrorDetail(attemptHandlers, "error from services %v", err)
		return err
	}

	return c.JSON(res)
}

func (h *attemptHandlerImpl) DeleteAttempt(c *fiber.Ctx) error {
	id := c.Params("id")
	res, err := h.attemptService.DeleteAttempt(c.Context(), id)
	if err != nil {
		log.ErrorDetail(attemptHandlers, "error from services %v", err)
		return err
	}

	return c.JSON(res)
}
