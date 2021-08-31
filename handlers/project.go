package handlers

import (
	"github.com/aliirsyaadn/kodein/entity"
	"github.com/aliirsyaadn/kodein/internal/log"
	"github.com/aliirsyaadn/kodein/services/project"
	"github.com/gofiber/fiber/v2"
	_ "github.com/lib/pq"
)

const projectHandlers = "Project Handlers"

type projectHandlerImpl struct {
	projectService project.Service
}

func NewProjectHandler(projectService project.Service) *projectHandlerImpl {
	return &projectHandlerImpl{
		projectService: projectService,
	}
}

func (h *projectHandlerImpl) Register(r fiber.Router) {
	r = r.Group("/project")

	r.Get("/member/:id", h.GetProjectsByMemberID)
	r.Get("/:id", h.GetProjectByID)
	r.Post("/", h.CreateProject)
	r.Post("/:id", h.UpdateProject)
	r.Delete("/:id", h.DeleteProject)
}

func (h *projectHandlerImpl) GetProjectsByMemberID(c *fiber.Ctx) error {
	projects, err := h.projectService.GetProjectsByMemberID(c.Context(), c.Params("id"))

	if err != nil {
		log.ErrorDetail(projectHandlers, "error from services %v", err)
		return err
	}

	return c.JSON(projects)
}

func (h *projectHandlerImpl) GetProjectByID(c *fiber.Ctx) error {
	id := c.Params("id")
	project, err := h.projectService.GetProjectByID(c.Context(), id)

	if err != nil {
		log.ErrorDetail(projectHandlers, "error from services %v", err)
		return err
	}

	return c.JSON(project)
}

func (h *projectHandlerImpl) CreateProject(c *fiber.Ctx) error {
	req := new(entity.CreateProjectRequest)
	if err := c.BodyParser(req); err != nil {
		log.ErrorDetail(projectHandlers, "error parse body %v", err)
		return err
	}

	res, err := h.projectService.CreateProject(c.Context(), *req)
	if err != nil {
		log.ErrorDetail(projectHandlers, "error from services %v", err)
		return err
	}

	return c.JSON(res)
}

func (h *projectHandlerImpl) UpdateProject(c *fiber.Ctx) error {
	id := c.Params("id")
	req := new(entity.UpdateProjectRequest)
	if err := c.BodyParser(req); err != nil {
		log.ErrorDetail(projectHandlers, "error parse body %v", err)
		return err
	}

	res, err := h.projectService.UpdateProject(c.Context(), *req, id)
	if err != nil {
		log.ErrorDetail(projectHandlers, "error from services %v", err)
		return err
	}

	return c.JSON(res)
}

func (h *projectHandlerImpl) DeleteProject(c *fiber.Ctx) error {
	id := c.Params("id")
	res, err := h.projectService.DeleteProject(c.Context(), id)
	if err != nil {
		log.ErrorDetail(projectHandlers, "error from services %v", err)
		return err
	}

	return c.JSON(res)
}
