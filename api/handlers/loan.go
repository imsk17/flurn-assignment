package handlers

import (
	"assignment/pkg/entities"
	"assignment/pkg/loan"
	"github.com/gofiber/fiber/v2"
	"net/http"
	"strconv"
	"strings"
)

func getOneLoan(repo loan.Repository) fiber.Handler {
	return func(ctx *fiber.Ctx) error {
		id := ctx.Params("id")
		if id == "" {
			return ctx.Status(http.StatusNotFound).JSON(fiber.Map{"message": "no loan with that id found"})
		}
		uid, err := strconv.Atoi(id)
		if err != nil {
			return ctx.Status(http.StatusBadRequest).JSON(fiber.Map{"message": "invalid id found"})
		}
		loanReq, err := repo.Read(uint(uid))
		if err != nil {
			return ctx.Status(http.StatusBadRequest).JSON(fiber.Map{"message": err.Error()})
		}
		return ctx.Status(http.StatusOK).JSON(loanReq)
	}
}

func getAllLoans(repo loan.Repository) fiber.Handler {
	return func(ctx *fiber.Ctx) error {
		status := ctx.Query("status")
		loanAmountGreater := ctx.Query("loanAmountGreater")
		if status != "" && loanAmountGreater == "" {
			statuses := strings.Split(status, ",")
			var stats []entities.Status
			for _, s := range statuses {
				stats = append(stats, entities.Status(s))
			}
			loanReqs, err := repo.ReadByStatus(stats)
			if err != nil {
				return ctx.Status(http.StatusBadRequest).JSON(fiber.Map{"message": err.Error()})
			}
			return ctx.Status(http.StatusOK).JSON(loanReqs)
		}
		if status == "" && loanAmountGreater != "" {
			loanAmount, err := strconv.Atoi(loanAmountGreater)
			if err != nil {
				return ctx.Status(http.StatusBadRequest).JSON(fiber.Map{"message": err.Error()})
			}
			loanReqs, err := repo.ReadByLoanAmount(int32(loanAmount))
			if err != nil {
				return ctx.Status(http.StatusBadRequest).JSON(fiber.Map{"message": err.Error()})
			}
			return ctx.Status(http.StatusOK).JSON(loanReqs)
		}
		if status != "" && loanAmountGreater != "" {
			loanAmount, err := strconv.Atoi(loanAmountGreater)
			if err != nil {
				return ctx.Status(http.StatusBadRequest).JSON(fiber.Map{"message": err.Error()})
			}
			statuses := strings.Split(status, ",")
			var stats []entities.Status
			for _, s := range statuses {
				stats = append(stats, entities.Status(s))
			}
			loanReqs, err := repo.ReadByLoanAmountAndStatus(int32(loanAmount), stats)
			if err != nil {
				return ctx.Status(http.StatusBadRequest).JSON(fiber.Map{"message": err.Error()})
			}
			return ctx.Status(http.StatusOK).JSON(loanReqs)
		}
		loanReqs, err := repo.ReadAll()
		if err != nil {
			return ctx.Status(http.StatusBadRequest).JSON(fiber.Map{"message": err.Error()})
		}
		return ctx.Status(http.StatusOK).JSON(loanReqs)
	}
}

func createLoan(repo loan.Repository) fiber.Handler {
	return func(ctx *fiber.Ctx) error {
		var loanReq entities.LoanRequest
		if err := ctx.BodyParser(&loanReq); err != nil {
			return err
		}
		// Make sure the status is `New` by default
		loanReq.Status = entities.New
		err := repo.Create(loanReq)
		if err != nil {
			return ctx.Status(http.StatusInternalServerError).JSON(fiber.Map{"error": err.Error()})
		}
		return ctx.Status(http.StatusCreated).JSON(fiber.Map{"message": "loan request created"})
	}
}

func updateLoan(repo loan.Repository) fiber.Handler {
	type ReqBody struct {
		Status entities.Status
	}
	return func(ctx *fiber.Ctx) error {
		id := ctx.Params("id")
		if id == "" {
			return ctx.Status(http.StatusNotFound).JSON(fiber.Map{"message": "no loan with that id found"})
		}
		uid, err := strconv.Atoi(id)
		if err != nil {
			return ctx.Status(http.StatusBadRequest).JSON(fiber.Map{"message": "invalid id found"})
		}
		var reqBody ReqBody
		err = ctx.BodyParser(&reqBody)
		if err != nil {
			return ctx.Status(http.StatusBadRequest).JSON(fiber.Map{"message": "invalid request body"})
		}
		err = repo.UpdateStatus(uint(uid), reqBody.Status)
		if err != nil {
			return ctx.Status(http.StatusInternalServerError).JSON(fiber.Map{"message": err.Error()})
		}
		return ctx.Status(http.StatusOK).JSON(fiber.Map{"message": "successfully updated loan status"})
	}
}

func deleteLoan(repo loan.Repository) fiber.Handler {
	return func(ctx *fiber.Ctx) error {
		id := ctx.Params("id")
		if id == "" {
			return ctx.Status(http.StatusNotFound).JSON(fiber.Map{"message": "no loan with that id found"})
		}
		uid, err := strconv.Atoi(id)
		if err != nil {
			return ctx.Status(http.StatusBadRequest).JSON(fiber.Map{"message": "invalid id found"})
		}
		err = repo.UpdateStatus(uint(uid), entities.Cancelled)
		if err != nil {
			return ctx.Status(http.StatusInternalServerError).JSON(fiber.Map{"message": err.Error()})
		}
		return ctx.Status(http.StatusOK).JSON(fiber.Map{"message": "successfully updated loan status to cancelled"})
	}
}

func SetupLoanRoutes(app *fiber.App, repo loan.Repository) {
	app.Post("/loans", createLoan(repo))
	app.Patch("/loans/:id", updateLoan(repo))
	app.Delete("/loans/:id", deleteLoan(repo))
	app.Get("/loans", getAllLoans(repo))
	app.Get("/loans/:id", getOneLoan(repo))
}
