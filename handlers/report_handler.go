package handlers

import (
	"net/http"

	"github.com/iqbalmahad/sistem-akademik-sahabat-alam-golang/helpers"
	"github.com/iqbalmahad/sistem-akademik-sahabat-alam-golang/models"
	"github.com/iqbalmahad/sistem-akademik-sahabat-alam-golang/repositories"

	"github.com/gofiber/fiber/v2"
)

type ReportHandler struct {
	ReportRepo *repositories.ReportRepository
}

func NewReportHandler(reportRepo *repositories.ReportRepository) *ReportHandler {
	return &ReportHandler{ReportRepo: reportRepo}
}

func (h *ReportHandler) CreateReport(c *fiber.Ctx) error {
	var report models.Report
	if err := c.BodyParser(&report); err != nil {
		return c.Status(http.StatusBadRequest).JSON(fiber.Map{"error": err.Error()})
	}

	if err := h.ReportRepo.Create(c.Context(), &report); err != nil {
		return c.Status(http.StatusInternalServerError).JSON(fiber.Map{"error": err.Error()})
	}

	return c.Status(http.StatusCreated).JSON(report)
}

func (h *ReportHandler) UpdateReport(c *fiber.Ctx) error {
	reportID, err := helpers.ParseID(c)
	if err != nil {
		return c.Status(http.StatusBadRequest).JSON(fiber.Map{"error": "Invalid report ID"})
	}

	var report models.Report
	if err := c.BodyParser(&report); err != nil {
		return c.Status(http.StatusBadRequest).JSON(fiber.Map{"error": err.Error()})
	}
	report.ID = uint(reportID)

	if err := h.ReportRepo.Update(c.Context(), &report); err != nil {
		return c.Status(http.StatusInternalServerError).JSON(fiber.Map{"error": err.Error()})
	}

	return c.Status(http.StatusOK).JSON(report)
}

func (h *ReportHandler) DeleteReport(c *fiber.Ctx) error {
	reportID, err := helpers.ParseID(c)
	if err != nil {
		return c.Status(http.StatusBadRequest).JSON(fiber.Map{"error": "Invalid report ID"})
	}

	if err := h.ReportRepo.Delete(c.Context(), uint(reportID)); err != nil {
		return c.Status(http.StatusInternalServerError).JSON(fiber.Map{"error": err.Error()})
	}

	return c.Status(http.StatusOK).JSON(fiber.Map{"message": "Report deleted successfully"})
}

func (h *ReportHandler) GetReportByID(c *fiber.Ctx) error {
	reportID, err := helpers.ParseID(c)
	if err != nil {
		return c.Status(http.StatusBadRequest).JSON(fiber.Map{"error": "Invalid report ID"})
	}

	report, err := h.ReportRepo.GetByID(c.Context(), uint(reportID))
	if err != nil {
		return c.Status(http.StatusNotFound).JSON(fiber.Map{"error": "Report not found"})
	}

	return c.Status(http.StatusOK).JSON(report)
}
