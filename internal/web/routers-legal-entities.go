package web

import (
	"net/http"

	"github.com/google/uuid"
	"github.com/labstack/echo/v4"

	"github.com/krisch/crm-backend/internal/legalentities"
)

// RegisterLegalEntitiesRoutes регистрирует CRUD-эндпоинты для LegalEntities и BankAccounts.
func RegisterLegalEntitiesRoutes(e *echo.Echo, service *legalentities.Service) {
	// GET /legal-entities - Получение всех юридических лиц
	e.GET("/legal-entities", func(c echo.Context) error {
		entities, err := service.GetAllLegalEntities(c.Request().Context())
		if err != nil {
			return c.JSON(http.StatusInternalServerError, map[string]interface{}{
				"error": err.Error(),
			})
		}
		return c.JSON(http.StatusOK, entities)
	})

	// POST /legal-entities - Создание нового юридического лица
	e.POST("/legal-entities", func(c echo.Context) error {
		var req struct {
			Name string `json:"name" validate:"required"`
		}
		if err := c.Bind(&req); err != nil {
			return c.JSON(http.StatusBadRequest, map[string]interface{}{
				"error": err.Error(),
			})
		}

		newID, err := service.CreateLegalEntity(c.Request().Context(), req.Name)
		if err != nil {
			return c.JSON(http.StatusInternalServerError, map[string]interface{}{
				"error": err.Error(),
			})
		}

		return c.JSON(http.StatusCreated, map[string]interface{}{
			"uuid": newID.String(),
		})
	})

	// PUT /legal-entities/:uuid - Обновление юр. лица
	e.PUT("/legal-entities/:uuid", func(c echo.Context) error {
		strID := c.Param("uuid")
		id, err := uuid.Parse(strID)
		if err != nil {
			return c.JSON(http.StatusBadRequest, map[string]interface{}{
				"error": "invalid uuid",
			})
		}

		var req struct {
			Name string `json:"name" validate:"required"`
		}
		if err := c.Bind(&req); err != nil {
			return c.JSON(http.StatusBadRequest, map[string]interface{}{
				"error": err.Error(),
			})
		}

		if err := service.UpdateLegalEntity(c.Request().Context(), id, req.Name); err != nil {
			return c.JSON(http.StatusNotFound, map[string]interface{}{
				"error": err.Error(),
			})
		}

		return c.NoContent(http.StatusOK)
	})

	// DELETE /legal-entities/:uuid - Удаление юр. лица
	e.DELETE("/legal-entities/:uuid", func(c echo.Context) error {
		strID := c.Param("uuid")
		id, err := uuid.Parse(strID)
		if err != nil {
			return c.JSON(http.StatusBadRequest, map[string]interface{}{
				"error": "invalid uuid",
			})
		}

		if err := service.DeleteLegalEntity(c.Request().Context(), id); err != nil {
			return c.JSON(http.StatusNotFound, map[string]interface{}{
				"error": err.Error(),
			})
		}

		return c.NoContent(http.StatusNoContent)
	})

	// --- БАНКОВСКИЕ СЧЕТА ---

	// GET /bank-accounts - Получение всех банковских счетов
	e.GET("/bank-accounts", func(c echo.Context) error {
		bankAccounts, err := service.GetAllBankAccounts(c.Request().Context(), uuid.Nil)
		if err != nil {
			return c.JSON(http.StatusInternalServerError, map[string]interface{}{
				"error": err.Error(),
			})
		}
		return c.JSON(http.StatusOK, bankAccounts)
	})

	// GET /bank-accounts - Получение всех банковских счетов для юр. лица
	e.GET("/bank-accounts", func(c echo.Context) error {
		var req struct {
			LegalEntityUuid uuid.UUID `json:"legal_entity_uuid" validate:"required"`
		}
		if err := c.Bind(&req); err != nil {
			return c.JSON(http.StatusBadRequest, map[string]interface{}{
				"error": err.Error(),
			})
		}

		bankAccounts, err := service.GetAllBankAccounts(c.Request().Context(), req.LegalEntityUuid)
		if err != nil {
			return c.JSON(http.StatusInternalServerError, map[string]interface{}{
				"error": err.Error(),
			})
		}
		return c.JSON(http.StatusOK, bankAccounts)
	})

	// GET /bank-accounts/:uuid - Получение одного банковского счета
	e.GET("/bank-accounts/:uuid", func(c echo.Context) error {
		strID := c.Param("uuid")
		bankAccountUUID, err := uuid.Parse(strID)
		if err != nil {
			return c.JSON(http.StatusBadRequest, map[string]interface{}{
				"error": "invalid uuid",
			})
		}

		bankAccount, err := service.GetBankAccountByUUID(c.Request().Context(), bankAccountUUID)
		if err != nil {
			return c.JSON(http.StatusNotFound, map[string]interface{}{
				"error": err.Error(),
			})
		}

		return c.JSON(http.StatusOK, bankAccount)
	})

	// POST /bank-accounts - Создание банковского счета
	e.POST("/bank-accounts", func(c echo.Context) error {
		var req struct {
			LegalEntityUuid uuid.UUID `json:"legal_entity_uuid" validate:"required"`
			BIC             string    `json:"bic" validate:"required"`
			BankName        string    `json:"bank_name" validate:"required"`
			Address         string    `json:"address"`
			CorrAccount     string    `json:"correspondent_account"`
			AccountNumber   string    `json:"account_number" validate:"required"`
			Currency        string    `json:"currency"`
			Comment         string    `json:"comment"`
		}
		if err := c.Bind(&req); err != nil {
			return c.JSON(http.StatusBadRequest, map[string]interface{}{
				"error": err.Error(),
			})
		}

		bankAccount := &legalentities.BankAccount{
			BIC:                  req.BIC,
			BankName:             req.BankName,
			Address:              req.Address,
			CorrespondentAccount: req.CorrAccount,
			AccountNumber:        req.AccountNumber,
			Currency:             req.Currency,
			Comment:              req.Comment,
			LegalEntityUUID:      req.LegalEntityUuid,
		}

		newID, err := service.CreateBankAccount(c.Request().Context(), bankAccount.ToDomain())
		if err != nil {
			return c.JSON(http.StatusInternalServerError, map[string]interface{}{
				"error": err.Error(),
			})
		}

		return c.JSON(http.StatusCreated, map[string]interface{}{
			"uuid": newID.String(),
		})
	})

	// PUT /bank-accounts/:uuid - Обновление банковского счета
	e.PUT("/bank-accounts/:uuid", func(c echo.Context) error {
		strID := c.Param("uuid")
		bankAccountUUID, err := uuid.Parse(strID)
		if err != nil {
			return c.JSON(http.StatusBadRequest, map[string]interface{}{
				"error": "invalid uuid",
			})
		}

		var req struct {
			BIC           string `json:"bic" validate:"required"`
			BankName      string `json:"bank_name" validate:"required"`
			Address       string `json:"address"`
			CorrAccount   string `json:"correspondent_account"`
			AccountNumber string `json:"account_number" validate:"required"`
			Currency      string `json:"currency"`
			Comment       string `json:"comment"`
		}
		if err := c.Bind(&req); err != nil {
			return c.JSON(http.StatusBadRequest, map[string]interface{}{
				"error": err.Error(),
			})
		}

		bankAccount := &legalentities.BankAccount{
			UUID:                 bankAccountUUID,
			BIC:                  req.BIC,
			BankName:             req.BankName,
			Address:              req.Address,
			CorrespondentAccount: req.CorrAccount,
			AccountNumber:        req.AccountNumber,
			Currency:             req.Currency,
			Comment:              req.Comment,
		}

		domainBankAccount := bankAccount.ToDomain()

		if err := service.UpdateBankAccount(c.Request().Context(), domainBankAccount); err != nil {
			return c.JSON(http.StatusNotFound, map[string]interface{}{
				"error": err.Error(),
			})
		}

		return c.JSON(http.StatusOK, bankAccount)
	})

	// DELETE /bank-accounts/:uuid - Удаление банковского счета
	e.DELETE("/bank-accounts/:uuid", func(c echo.Context) error {
		strID := c.Param("uuid")
		bankAccountUUID, err := uuid.Parse(strID)
		if err != nil {
			return c.JSON(http.StatusBadRequest, map[string]interface{}{
				"error": "invalid uuid",
			})
		}

		if err := service.DeleteBankAccount(c.Request().Context(), bankAccountUUID); err != nil {
			return c.JSON(http.StatusNotFound, map[string]interface{}{
				"error": err.Error(),
			})
		}

		return c.NoContent(http.StatusNoContent)
	})
}
