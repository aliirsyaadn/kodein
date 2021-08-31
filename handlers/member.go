package handlers

import (
	"github.com/aliirsyaadn/kodein/entity"
	"github.com/aliirsyaadn/kodein/internal/log"
	"github.com/aliirsyaadn/kodein/services/member"
	"github.com/gofiber/fiber/v2"
	_ "github.com/lib/pq"
)

const memberHandlers = "Member Handlers"

type memberHandlerImpl struct {
	memberService member.Service
}

func NewMemberHandler(memberService member.Service) *memberHandlerImpl {
	return &memberHandlerImpl{
		memberService: memberService,
	}
}

func (h *memberHandlerImpl) Register(r fiber.Router) {
	r = r.Group("/member")

	r.Get("/", h.GetMembers)
	r.Get("/:id", h.GetMemberByID)
	r.Post("/", h.CreateMember)
	r.Post("/:id", h.UpdateMember)
	r.Delete("/:id", h.DeleteMember)
}

func (h *memberHandlerImpl) GetMembers(c *fiber.Ctx) error {
	members, err := h.memberService.GetMembers(c.Context())

	if err != nil {
		log.ErrorDetail(memberHandlers, "error from services %v", err)
		return err
	}

	return c.JSON(members)
}

func (h *memberHandlerImpl) GetMemberByID(c *fiber.Ctx) error {
	id := c.Params("id")
	member, err := h.memberService.GetMemberByID(c.Context(), id)

	if err != nil {
		log.ErrorDetail(memberHandlers, "error from services %v", err)
		return err
	}

	return c.JSON(member)
}

func (h *memberHandlerImpl) CreateMember(c *fiber.Ctx) error {
	req := new(entity.CreateMemberRequest)
	if err := c.BodyParser(req); err != nil {
		log.ErrorDetail(memberHandlers, "error parse body %v", err)
		return err
	}

	res, err := h.memberService.CreateMember(c.Context(), *req)
	if err != nil {
		log.ErrorDetail(memberHandlers, "error from services %v", err)
		return err
	}

	return c.JSON(res)
}

func (h *memberHandlerImpl) UpdateMember(c *fiber.Ctx) error {
	id := c.Params("id")
	req := new(entity.UpdateMemberRequest)
	if err := c.BodyParser(req); err != nil {
		log.ErrorDetail(memberHandlers, "error parse body %v", err)
		return err
	}

	res, err := h.memberService.UpdateMember(c.Context(), *req, id)
	if err != nil {
		log.ErrorDetail(memberHandlers, "error from services %v", err)
		return err
	}

	return c.JSON(res)
}

func (h *memberHandlerImpl) DeleteMember(c *fiber.Ctx) error {
	id := c.Params("id")
	res, err := h.memberService.DeleteMember(c.Context(), id)
	if err != nil {
		log.ErrorDetail(memberHandlers, "error from services %v", err)
		return err
	}

	return c.JSON(res)
}
