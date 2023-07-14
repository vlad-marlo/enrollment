package http

import (
	"errors"
	"github.com/gofiber/fiber/v2"
	"github.com/gofiber/swagger"
	_ "github.com/vlad-marlo/enrollment/docs"
	"github.com/vlad-marlo/enrollment/internal/controller"
	"github.com/vlad-marlo/enrollment/internal/model"
	"github.com/vlad-marlo/enrollment/internal/pkg/fielderr"
	"github.com/vlad-marlo/enrollment/internal/pkg/logger"
	"go.uber.org/zap"
	"net/http"
)

import (
	"context"
)

type Controller struct {
	engine *fiber.App
	log    *zap.Logger
	cfg    controller.Config
	srv    controller.Service
}

// New initializes application with provided objects.
func New(
	log *zap.Logger,
	cfg controller.Config,
	service controller.Service,
) (*Controller, error) {
	srv := &Controller{
		engine: fiber.New(),
		log:    log.With(zap.String(logger.EntityField, "http")),
		cfg:    cfg,
		srv:    service,
	}
	if log == nil || cfg == nil || service == nil {
		return nil, ErrNilReference
	}
	srv.configure()
	log.Info("successful initialized server")
	return srv, nil
}

func (srv *Controller) configure() {
	srv.configureMW()
	srv.configureRoutes()
}

func (srv *Controller) Start(context.Context) error {
	go func() {
		srv.log.Fatal("starting http server", zap.Error(srv.engine.Listen(srv.cfg.BindAddr())))
	}()
	srv.log.Info("starting http server", zap.String("bind_addr", srv.cfg.BindAddr()))
	return nil
}

// Stop gracefully stops server.
func (srv *Controller) Stop(context.Context) error {
	srv.log.Info("stopping http server", zap.String("bind_addr", srv.cfg.BindAddr()))
	return srv.engine.Shutdown()
}

// configureRoutes sets all needed
func (srv *Controller) configureRoutes() {
	srv.engine.Get("/swagger/*", swagger.HandlerDefault)
	srv.engine.Post("/api/records", srv.HandleCreateRecord)
	srv.engine.Get("/api/records/:id", srv.HandleGetRecord)
	srv.engine.Get("/api/user/:user/records", srv.HandleGetUserRecords)
	srv.engine.Get("/api/records", srv.HandleGetAllRecords)
}

// configureMW configures all middlewares to engine.
func (srv *Controller) configureMW() {
}

func (srv *Controller) handleError(c *fiber.Ctx, msg string, err error, fields ...zap.Field) error {
	var fieldErr *fielderr.Error
	if errors.As(err, &fieldErr) {
		srv.log.Warn(msg, append(fieldErr.Fields(), fields...)...)
		c.Status(fieldErr.CodeHTTP())
		return c.JSON(fieldErr.Data())
	}

	srv.log.Warn(msg, append(fields, zap.NamedError("checked_error", err))...)
	c.Status(http.StatusBadRequest)
	return c.JSON(model.BadRequestResponse{})
}
