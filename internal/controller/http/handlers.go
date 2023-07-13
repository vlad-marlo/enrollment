package http

import (
	"github.com/gofiber/fiber/v2"
	"github.com/vlad-marlo/enrollment/internal/model"
	"go.uber.org/zap"
)

// HandleCreateRecord creates new record into service.
//
//	@Tags		records-controller
//	@Summary	Создание новых записей
//	@Accept		json,xml,x-www-form-urlencoded
//	@Produce	json,xml,x-www-form-urlencoded
//	@Param		request	body		model.CreateRecordRequest	true	"Records"
//	@Success	201		{object}	model.CreateRecordResponse	"Created"
//	@Failure	400		{object}	model.BadRequestResponse	"Bad Request"
//	@Router		/api/records/ [post]
func (srv *Controller) HandleCreateRecord(ctx *fiber.Ctx) error {
	req := new(model.CreateRecordRequest)
	if err := ctx.BodyParser(req); err != nil {
		srv.log.Error("got error while handling request", zap.Error(err))
		return srv.handleError(ctx, err)
	}
	srv.log.Debug("handled request", zap.String("user", req.User), zap.String("msg_type", req.MsgType))
	return ctx.Format(req)
}
