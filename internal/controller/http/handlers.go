package http

import (
	"github.com/gofiber/fiber/v2"
	"github.com/vlad-marlo/enrollment/internal/model"
	"net/http"
)

// HandleCreateRecord creates new record into service.
//
//	@Tags		records-controller
//	@Summary	Создание новых записей
//	@Accept		json
//	@Produce	json
//	@Param		request	body		model.CreateRecordRequest	true	"Records"
//	@Success	201		{object}	model.CreateRecordRequest	"OK"
//	@Failure	400		{object}	model.BadRequestResponse	"Bad Request"
//	@Router		/api/records/ [post]
func (srv *Controller) HandleCreateRecord(ctx *fiber.Ctx) error {
	req := new(model.CreateRecordRequest)
	if err := ctx.BodyParser(req); err != nil {
		return srv.handleError(ctx, err)
	}
	return ctx.SendStatus(http.StatusBadRequest)
}
