package main

import (
	"net/http"
	"strconv"

	"github.com/Mateusz2734/wdai-project/backend/internal/response"
	"github.com/Mateusz2734/wdai-project/backend/internal/validator"
	"github.com/alexedwards/flow"
)

func (app *application) getGeneralStats(w http.ResponseWriter, r *http.Request) {
	reviewCountByStars, err := app.db.GetReviewCountByStars(r.Context())

	if err != nil {
		app.serverError(w, r, err)
		return
	}

	userCount, err := app.db.GetUserCount(r.Context())

	if err != nil {
		app.serverError(w, r, err)
		return
	}

	offerCountByCategory, err := app.db.GetOfferCountByCategory(r.Context())

	if err != nil {
		app.serverError(w, r, err)
		return
	}

	offerCountBySkill, err := app.db.GetOfferCountBySkill(r.Context())

	if err != nil {
		app.serverError(w, r, err)
		return
	}

	offerCount, err := app.db.GetOfferCount(r.Context())

	if err != nil {
		app.serverError(w, r, err)
		return
	}

	data := map[string]interface{}{
		"userCount":            userCount,
		"offerCount":           offerCount,
		"reviewCountByStars":   reviewCountByStars,
		"offerCountBySkill":    offerCountBySkill,
		"offerCountByCategory": offerCountByCategory,
	}

	err = response.JSONSuccess(w, data)

	if err != nil {
		app.serverError(w, r, err)
	}
}

func (app *application) getUserStats(w http.ResponseWriter, r *http.Request) {
	val := validator.Validator{}

	userID, err := strconv.ParseInt(flow.Param(r.Context(), "userId"), 10, 32)

	if err != nil {
		val.AddError("userID must be a number")
		app.serverError(w, r, err)
		return
	}

	userID32 := int32(userID)

	val.Check(validator.Between(userID32, 1, 2147483647), "userID must be bigger than 0")

	if val.HasErrors() {
		app.failedValidation(w, r, val)
		return
	}

	_, err = app.db.GetUserById(r.Context(), userID32)

	if err != nil {
		app.notFound(w, r)
		return
	}

	offerCount, err := app.db.GetOfferCountByUser(r.Context(), userID32)

	if err != nil {
		app.serverError(w, r, err)
		return
	}

	reviewCount, err := app.db.GetReviewCountByUser(r.Context(), userID32)

	if err != nil {
		app.serverError(w, r, err)
		return
	}
	averageStars, err := app.db.GetAverageStarsByUser(r.Context(), userID32)

	if err != nil {
		app.serverError(w, r, err)
		return
	}

	data := map[string]interface{}{
		"userId":       userID32,
		"offerCount":   offerCount,
		"reviewCount":  reviewCount,
		"averageStars": averageStars,
	}

	err = response.JSONSuccess(w, data)

	if err != nil {
		app.serverError(w, r, err)
	}
}
