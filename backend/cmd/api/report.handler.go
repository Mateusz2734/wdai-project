package main

import (
	"net/http"
	"strconv"

	"github.com/Mateusz2734/wdai-project/backend/internal/db"
	"github.com/Mateusz2734/wdai-project/backend/internal/request"
	"github.com/Mateusz2734/wdai-project/backend/internal/response"
	"github.com/Mateusz2734/wdai-project/backend/internal/validator"
	"github.com/alexedwards/flow"
	"github.com/jackc/pgx/v5"
	"github.com/jackc/pgx/v5/pgtype"
)

func (app *application) getReports(w http.ResponseWriter, r *http.Request) {
	reports, err := app.db.GetReports(r.Context())

	if err != nil && err != pgx.ErrNoRows {
		app.serverError(w, r, err)
		return
	}

	data := map[string]interface{}{
		"reports": reports,
	}

	err = response.JSONSuccess(w, data)

	if err != nil {
		app.serverError(w, r, err)
	}
}

func (app *application) addReport(w http.ResponseWriter, r *http.Request) {
	var input struct {
		ReportedOfferID int32               `json:"reportedOfferId"`
		Reason          string              `json:"reason"`
		Description     string              `json:"description"`
		Status          string              `json:"status"`
		Validator       validator.Validator `json:"-"`
	}

	err := request.DecodeJSON(w, r, &input)

	if err != nil {
		app.badRequest(w, r, err)
		return
	}

	user := contextGetAuthenticatedUser(r)

	if user == nil {
		app.authenticationRequired(w, r)
		return
	}

	input.Validator.CheckField(input.Reason != "", "reason", "Reason is required")
	input.Validator.CheckField(input.Description != "", "description", "Description is required")
	input.Validator.CheckField(input.Status != "", "status", "Status is required")

	if input.Validator.HasErrors() {
		app.failedValidation(w, r, input.Validator)
		return
	}

	reason, err := app.db.GetReason(r.Context(), input.Reason)

	if err != nil && err != pgx.ErrNoRows {
		app.serverError(w, r, err)
		return
	}

	if err == pgx.ErrNoRows || reason == "" {
		app.errorMessage(w, r, http.StatusUnprocessableEntity, "Report reason does not exist", nil)
	}

	reportedOffer := pgtype.Int4{Int32: input.ReportedOfferID, Valid: false}

	if input.ReportedOfferID != 0 {
		reportedOffer.Valid = true
	}

	params := db.AddReportParams{
		ReportingUserID: user.UserID,
		Description:     input.Description,
		Reason:          input.Reason,
		Status:          input.Status,
		ReportedOfferID: reportedOffer,
	}

	report, err := app.db.AddReport(r.Context(), params)

	if err != nil {
		app.badRequest(w, r, err)
		return
	}

	data := map[string]interface{}{
		"report": report,
	}

	err = response.JSONSuccess(w, data)

	if err != nil {
		app.serverError(w, r, err)
	}
}

func (app *application) deleteReport(w http.ResponseWriter, r *http.Request) {
	reportID, err := strconv.ParseInt(flow.Param(r.Context(), "reportId"), 10, 32)

	if err != nil || reportID == 0 {
		app.badRequest(w, r, err)
		return
	}

	user := contextGetAuthenticatedUser(r)

	if user == nil || user.Role != "admin" {
		app.authenticationRequired(w, r)
		return
	}

	report, err := app.db.GetReportById(r.Context(), int32(reportID))

	if err != nil && err != pgx.ErrNoRows {
		app.serverError(w, r, err)
		return
	}

	if err == pgx.ErrNoRows || report == nil {
		app.errorMessage(w, r, http.StatusUnprocessableEntity, "Report does not exist", nil)
		return
	}

	err = app.db.DeleteReport(r.Context(), int32(reportID))

	if err != nil {
		app.serverError(w, r, err)
		return
	}

	err = response.JSONSuccess(w, nil)

	if err != nil {
		app.serverError(w, r, err)
	}
}
