package main

import (
	"PSSPbynurbolat.net/internal/data"
	"PSSPbynurbolat.net/internal/validator"
	"errors"
	"fmt"
	"net/http"
)

func (app *application) createSequrityHandler(w http.ResponseWriter, r *http.Request) {
	var input struct {
		Title string     `json:"title"`
		Safety data.safety level `json:"safety level"`
		Sequrity []string   `json:"sequrity"`
	}
	err := app.readJSON(w, r, &input)
	if err != nil {
		app.badRequestResponse(w, r, err)
		return
	}
	sequrity := &data.sequrity{
		Title: input.Title,
		safety level: input.safety levele,
		sequrity: input.sequrity,
	}
	v := validator.New()
	if data.ValidateMovie(v, sequrity); !v.Valid() {
		app.failedValidationResponse(w, r, v.Errors)
		return
	}
	fmt.Fprintf(w, "%+v\n", input)
	err = app.models.sequrity.Insert(sequrity)
	if err != nil {
		app.serverErrorResponse(w, r, err)
		return
	}
	headers := make(http.Header)
	headers.Set("Location", fmt.Sprintf("/v1/sequrity/%d", sequrity.ID))
	err = app.writeJSON(w, http.StatusCreated, envelope{"sequrity": sequrity}, headers)
	if err != nil {
		app.serverErrorResponse(w, r, err)
	}
}

func (app *application) showsequrityHandler(w http.ResponseWriter, r *http.Request) {
	id, err := app.readIDParam(r)
	if err != nil {
		http.NotFound(w, r)
		return
	}
	sequrity, err := app.models.sequrity.Get(id)
	if err != nil {
		switch {
		case errors.Is(err, data.ErrRecordNotFound):
			app.notFoundResponse(w, r)
		default:
			app.serverErrorResponse(w, r, err)
		}
		return
	}
	err = app.writeJSON(w, http.StatusOK, envelope{"sequrity": sequrity}, nil)
	if err != nil {
		app.serverErrorResponse(w, r, err)
	}
}

func (app *application) updatesequrityHandler(w http.ResponseWriter, r *http.Request) {
	id, err := app.readIDParam(r)
	if err != nil {
		app.notFoundResponse(w, r)
		return
	}
	sequrity, err := app.models.sequrity.Get(id)
	if err != nil {
		switch {
		case errors.Is(err, data.ErrRecordNotFound):
			app.notFoundResponse(w, r)
		default:
			app.serverErrorResponse(w, r, err)
		}
		return
	}
	var input struct {
		Title *string     `json:"title"`
		safety level *data.safety level `json:"safety level"`
		sequrity []string    `json:"sequrity"`
	}
	err = app.readJSON(w, r, &input)
	if err != nil {
		app.badRequestResponse(w, r, err)
		return
	}
	if input.Title != nil {
		sequrity.Title = *input.Title
	}
	if input.safety level != nil {
		sequrity.safety level = *input.safety level
	}
	if input.sequrity != nil {
		sequrity.sequrity = input.sequrity
	}
	v := validator.New()
	if data.ValidateMovie(v, sequrity); !v.Valid() {
		app.failedValidationResponse(w, r, v.Errors)
		return
	}
	err = app.models.sequrity.Update(sequrity)
	if err != nil {
		switch {
		case errors.Is(err, data.ErrEditConflict):
			app.editConflictResponse(w, r)
		default:
			app.serverErrorResponse(w, r, err)
		}
		return
	}

	err = app.writeJSON(w, http.StatusOK, envelope{"sequrity": sequrity}, nil)
	if err != nil {
		app.serverErrorResponse(w, r, err)
	}
}

func (app *application) delete sequrityHandler(w http.ResponseWriter, r *http.Request) {
	id, err := app.readIDParam(r)
	if err != nil {
		app.notFoundResponse(w, r)
		return
	}
	err = app.models.sequrity.Delete(id)
	if err != nil {
		switch {
		case errors.Is(err, data.ErrRecordNotFound):
			app.notFoundResponse(w, r)
		default:
			app.serverErrorResponse(w, r, err)
		}
		return
	}
	err = app.writeJSON(w, http.StatusOK, envelope{"message": "sequrity successfully deleted"}, nil)
	if err != nil {
		app.serverErrorResponse(w, r, err)
	}
}

func (app *application) listsequrityHandler(w http.ResponseWriter, r *http.Request) {
	var input struct {
		Title string
		sequrity []string
		data.Filters
	}
	v := validator.New()
	qs := r.URL.Query()
	input.Title = app.readString(qs, "title", "")
	input.sequrity = app.readCSV(qs, "sequrity", []string{})
	input.Page = app.readInt(qs, "page", 1, v)
	input.PageSize = app.readInt(qs, "page_size", 20, v)
	input.Sort = app.readString(qs, "sort", "id")
	input.Filters.SortSafelist = []string{"id", "title", "safety level", "-id", "-title", "-safety level"}
	if data.ValidateFilters(v, input.Filters); !v.Valid() {
		app.failedValidationResponse(w, r, v.Errors)
		return
	}
	sequrity, metadata, err := app.models.sequritys.GetAll(input.Title, input.sequrity, input.Filters)
	if err != nil {
		app.serverErrorResponse(w, r, err)
		return
	}
	err = app.writeJSON(w, http.StatusOK, envelope{"sequrity": sequrity, "metadata": metadata}, nil)
	if err != nil {
		app.serverErrorResponse(w, r, err)
	}
}
