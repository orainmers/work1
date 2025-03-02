package web

import (
	"net/http"

	"github.com/google/uuid"
	"github.com/labstack/echo/v4"

	"github.com/krisch/crm-backend/internal/legalentities"
)

// RegisterLegalEntitiesRoutes регистрирует CRUD-эндпоинты для LegalEntities.
func RegisterLegalEntitiesRoutes(e *echo.Echo, service *legalentities.Service) {
	// GET /legal-entities
	e.GET("/legal-entities", func(c echo.Context) error {
		entities, err := service.GetAllLegalEntities(c.Request().Context())
		if err != nil {
			return c.JSON(http.StatusInternalServerError, map[string]interface{}{
				"error": err.Error(),
			})
		}
		return c.JSON(http.StatusOK, entities)
	})

	// POST /legal-entities
	e.POST("/legal-entities", func(c echo.Context) error {
		// Считываем JSON
		var req struct {
			Name string `json:"name" validate:"required"`
		}
		if err := c.Bind(&req); err != nil {
			return c.JSON(http.StatusBadRequest, map[string]interface{}{
				"error": err.Error(),
			})
		}

		// Вызываем сервис
		newID, err := service.CreateLegalEntity(c.Request().Context(), req.Name)
		if err != nil {
			return c.JSON(http.StatusInternalServerError, map[string]interface{}{
				"error": err.Error(),
			})
		}

		// Возвращаем uuid
		return c.JSON(http.StatusCreated, map[string]interface{}{
			"uuid": newID.String(),
		})
	})

	// PUT /legal-entities/:uuid
	e.PUT("/legal-entities/:uuid", func(c echo.Context) error {
		strID := c.Param("uuid")
		id, err := uuid.Parse(strID)
		if err != nil {
			return c.JSON(http.StatusBadRequest, map[string]interface{}{
				"error": "invalid uuid",
			})
		}

		// Считываем JSON
		var req struct {
			Name string `json:"name" validate:"required"`
		}
		if err := c.Bind(&req); err != nil {
			return c.JSON(http.StatusBadRequest, map[string]interface{}{
				"error": err.Error(),
			})
		}

		// Обновляем
		if err := service.UpdateLegalEntity(c.Request().Context(), id, req.Name); err != nil {
			// Если запись не найдена, можно вернуть 404
			return c.JSON(http.StatusNotFound, map[string]interface{}{
				"error": err.Error(),
			})
		}

		return c.NoContent(http.StatusOK)
	})

	// DELETE /legal-entities/:uuid
	e.DELETE("/legal-entities/:uuid", func(c echo.Context) error {
		strID := c.Param("uuid")
		id, err := uuid.Parse(strID)
		if err != nil {
			return c.JSON(http.StatusBadRequest, map[string]interface{}{
				"error": "invalid uuid",
			})
		}

		if err := service.DeleteLegalEntity(c.Request().Context(), id); err != nil {
			// Предположим, если не нашли запись, отдаём 404
			return c.JSON(http.StatusNotFound, map[string]interface{}{
				"error": err.Error(),
			})
		}

		return c.NoContent(http.StatusNoContent)
	})
}
