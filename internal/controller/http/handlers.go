package http

import (
	"github.com/gofiber/fiber/v2"
	"github.com/vlad-marlo/enrollment/internal/model"
	"go.uber.org/zap"
	"net/http"
)

// HandleCreateRecord creates new record into service.
//
//	@Tags		records-controller
//	@Summary	Создание новых записей
//	@Accept		json,xml,x-www-form-urlencoded
//	@Produce	json
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
	resp, err := srv.srv.CreateRecord(ctx.UserContext(), req)
	if err != nil {
		return srv.handleError(ctx, err)
	}
	ctx.Status(http.StatusCreated)
	return ctx.JSON(resp)
}

// HandleGetRecord returns record by id.
//
//	@Tags		records-controller
//	@Summary	Получение записи
//	@Accept		json
//	@Produce	json
//	@Param		record_id	path		int							true	"record identifier"
//	@Success	200			{object}	model.GetRecordResponse		"OK"
//	@Failure	400			{object}	model.BadRequestResponse	"Bad Request"
//	@Router		/api/records/{record_id} [get]
func (srv *Controller) HandleGetRecord(ctx *fiber.Ctx) error {
	id := ctx.Params("id")
	resp, err := srv.srv.GetRecord(ctx.UserContext(), id)
	if err != nil {
		return srv.handleError(ctx, err)
	}
	return ctx.JSON(resp)
}

// HandleGetUserRecords returns records of user.
//
//	@Tags		records-controller
//	@Summary	Получение записей пользователя
//	@Accept		json
//	@Produce	json
//	@Param		user_id	path		string							true	"user identifier"
//	@Success	200		{object}	model.GetUserRecordsResponse	"OK
//	@Failure	400		{object}	model.GetRecordResponse			"Bad Request"
//	@Router		/api/users/records/{user_id} [get]
func (srv *Controller) HandleGetUserRecords(ctx *fiber.Ctx) error {
	user := ctx.Params("user")
	resp, err := srv.srv.GetUser(ctx.UserContext(), user)
	if err != nil {
		return srv.handleError(ctx, err)
	}
	return ctx.JSON(resp)
}

func (srv *Controller) HandleGetAllRecords(ctx *fiber.Ctx) error {
	resp, err := srv.srv.GetRecords(ctx.UserContext())
	if err != nil {
		return srv.handleError(ctx, err)
	}
	return ctx.JSON(resp)
}
