package handler

import (
	"errors"
	"fmt"
	"github.com/gin-gonic/gin"
	"log/slog"
	"net/http"
	"strconv"
)

var (
	ErrCreateGoods        = errors.New("create error")
	ErrUpdateGoods        = errors.New("update error")
	ErrRemoveGoods        = errors.New("remove error")
	ErrReprioritiizeGoods = errors.New("reprioritiize error")
)

type goodCreateReq struct {
	Name string `json:"name"`
}

func (h *Handler) create(c *gin.Context) {
	const op = "handler.create"

	log := h.log.With(
		slog.String("op", op),
	)

	projectId, err := strconv.Atoi(c.Query("projectId"))
	if err != nil {
		newErrorResponse(c, http.StatusBadRequest, InvalidInputBodyErr)
		log.Error("error get id:", fmt.Errorf("%s", InvalidInputBodyErr))
		return
	}

	var input goodCreateReq

	if err := c.BindJSON(&input); err != nil {
		newErrorResponse(c, http.StatusBadRequest, InvalidInputBodyErr)
		log.Error("error bind json:", fmt.Errorf("%s", InvalidInputBodyErr))
		return
	}

	g, err := h.Create(c.Request.Context(), input.Name, projectId)
	if err != nil {
		newErrorResponse(c, http.StatusInternalServerError, ErrCreateGoods.Error())
		log.Error("failed to create", ErrCreateGoods.Error())
		return
	}

	log.Info("Handler create successfully")

	c.JSON(http.StatusOK, map[string]interface{}{
		"id":          g.ID,
		"projectID":   g.ProjectID,
		"name":        g.Name,
		"description": g.Description,
		"priority":    g.Priority,
		"removed":     g.Removed,
		"created_at":  g.CreatedAt,
	})
}

type updateReq struct {
	Name        string `json:"name"`
	Description string `json:"description"`
}

func (h *Handler) update(c *gin.Context) {
	const op = "handler.update"

	log := h.log.With(
		slog.String("op", op),
	)

	id, projectId, err := parseParams(c)
	if err != nil {
		newErrorResponse(c, http.StatusBadRequest, InvalidInputBodyErr)
		log.Error("error get ids:", fmt.Errorf("%s", InvalidInputBodyErr))
		return
	}

	var input updateReq

	if err := c.BindJSON(&input); err != nil {
		newErrorResponse(c, http.StatusBadRequest, InvalidInputBodyErr)
		log.Error("error bind json:", fmt.Errorf("%s", InvalidInputBodyErr))
		return
	}

	g, err := h.Update(c.Request.Context(), input.Name, input.Description, id, projectId)
	if err != nil {
		newErrorResponse(c, http.StatusInternalServerError, ErrUpdateGoods.Error())
		log.Error("failed to create", ErrUpdateGoods.Error())
		return
	}

	log.Info("Handler update successfully")

	c.JSON(http.StatusOK, map[string]interface{}{
		"id":          g.ID,
		"projectID":   g.ProjectID,
		"name":        g.Name,
		"description": g.Description,
		"priority":    g.Priority,
		"removed":     g.Removed,
		"created_at":  g.CreatedAt,
	})
}

func (h *Handler) remove(c *gin.Context) {
	const op = "handler.remove"

	log := h.log.With(
		slog.String("op", op),
	)

	id, projectId, err := parseParams(c)
	if err != nil {
		newErrorResponse(c, http.StatusBadRequest, InvalidInputBodyErr)
		log.Error("error get ids:", fmt.Errorf("%s", InvalidInputBodyErr))
		return
	}

	g, err := h.Remove(c.Request.Context(), id, projectId)
	if err != nil {
		newErrorResponse(c, http.StatusInternalServerError, ErrRemoveGoods.Error())
		log.Error("failed to create", ErrRemoveGoods.Error())
		return
	}

	log.Info("Handler remove successfully")

	c.JSON(http.StatusOK, map[string]interface{}{
		"id":         g.ID,
		"name":       g.Name,
		"created_at": g.CreatedAt,
	})
}

type reprioritiizeReq struct {
	NewPriority int `json:"newPriority"`
}

func (h *Handler) reprioritiize(c *gin.Context) {
	const op = "handler.reprioritiize"

	log := h.log.With(
		slog.String("op", op),
	)

	id, projectId, err := parseParams(c)
	if err != nil {
		newErrorResponse(c, http.StatusBadRequest, InvalidInputBodyErr)
		log.Error("error get ids:", fmt.Errorf("%s", InvalidInputBodyErr))
		return
	}

	var input reprioritiizeReq

	if err := c.BindJSON(&input); err != nil {
		newErrorResponse(c, http.StatusBadRequest, InvalidInputBodyErr)
		log.Error("error bind json:", fmt.Errorf("%s", InvalidInputBodyErr))
		return
	}

	p, err := h.Reprioritize(c.Request.Context(), input.NewPriority, id, projectId)
	if err != nil {
		newErrorResponse(c, http.StatusInternalServerError, ErrReprioritiizeGoods.Error())
		log.Error("failed to create", ErrReprioritiizeGoods.Error())
		return
	}

	log.Info("Handler reprioriitize successfully")

	c.JSON(http.StatusOK, map[string]interface{}{
		"priorities": p,
	})
}

func parseParams(c *gin.Context) (int, int, error) {
	projectId, err := strconv.Atoi(c.Query("projectId"))
	if err != nil {
		_ = c.AbortWithError(http.StatusBadRequest, err)
		return 0, 0, err
	}

	id, err := strconv.Atoi(c.Query("id"))
	if err != nil {
		_ = c.AbortWithError(http.StatusBadRequest, err)
		return 0, 0, err
	}
	return id, projectId, nil
}
