package handler

import (
	"errors"
	"github.com/gin-gonic/gin"
	"log/slog"
	"net/http"
	"strconv"
)

var (
	ErrGetList = errors.New("error get list")
)

func (h *Handler) getList(c *gin.Context) {
	const op = "handler.getList"

	log := h.log.With(
		slog.String("op", op),
	)

	limit, err := strconv.Atoi(c.Query("limit"))
	if err != nil {
		limit = 10
		return
	}

	offset, err := strconv.Atoi(c.Query("offset"))
	if err != nil {
		offset = 1
		return
	}

	l, err := h.GetList(c.Request.Context(), limit, offset)
	if err != nil {
		newErrorResponse(c, http.StatusInternalServerError, ErrGetList.Error())
		log.Error("failed to create", ErrGetList.Error())
		return
	}

	log.Info("Handler get lists successfully")

	c.JSON(http.StatusOK, map[string]interface{}{
		"meta":  l.Meta,
		"goods": l.Goods,
	})
}
