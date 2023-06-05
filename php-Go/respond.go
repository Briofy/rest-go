package respond

import (
	"encoding/json"
	"net/http"
	"strconv"
)

type Respond struct {
	statusCode int
	statusText string
	errorCode  int
	lang       string
	messages   *Messages
	writer     http.ResponseWriter
}

func (r *Respond) Language(lang string) *Respond {
	r.lang = lang
	return r
}

func NewWithWriter(w http.ResponseWriter) *Respond {
	return &Respond{writer: w, messages: NewMessages()}
}

func (r *Respond) Messages() *Messages {
	if r.lang != "" {
		r.messages.Lang = r.lang
	}
	r.messages.load()
	return r.messages
}

func (r *Respond) SetStatusCode(code int) *Respond {
	r.statusCode = code
	return r
}

func (r *Respond) SetStatusText(text string) *Respond {
	r.statusText = text
	return r
}

func (r *Respond) SetErrorCode(code int) *Respond {
	r.errorCode = code
	return r
}

func (r *Respond) writeJSON(data interface{}) error {
	b, err := json.Marshal(data)
	if err != nil {
		return err
	}
	if _, err := r.writer.Write(b); err != nil {
		return err
	}
	r.writer.Header().Set("content-type", "application/json")
	return nil
}

func (r *Respond) RespondWithResult(result interface{}) {
	r.writer.WriteHeader(r.statusCode)
	r.writeJSON(map[string]interface{}{
		"status": r.statusText,
		"result": result,
	})
}

func (r *Respond) RespondWithMessage(message interface{}) {
	data := map[string]interface{}{
		"status":  r.statusText,
		"message": message,
	}
	if r.errorCode != 0 {
		data["error"] = r.errorCode
	}
	r.writer.WriteHeader(r.statusCode)
	r.writeJSON(data)
}

func (r *Respond) NotFound() {
	r.Error(404, 5404)
}

func (r *Respond) Succeed(data interface{}) {
	r.SetStatusCode(http.StatusOK).
		SetStatusText(r.Messages().Success).
		RespondWithResult(data)
}

func (r *Respond) InsertSucceeded() {
	message := r.Messages().Errors["success"]
	r.SetStatusCode(http.StatusOK).
		SetStatusText(r.Messages().Success).
		RespondWithMessage(message["insert"])
}

func (r *Respond) InsertFailed() {
	message := r.Messages().Errors["failed"]
	r.SetStatusCode(448).
		SetStatusText(r.Messages().Failed).
		RespondWithMessage(message["insert"])
}

func (r *Respond) DeleteSucceeded() {
	message := r.Messages().Errors["success"]
	r.SetStatusCode(200).
		SetStatusText(r.Messages().Success).
		RespondWithMessage(message["delete"])
}

func (r *Respond) DeleteFailed() {
	message := r.Messages().Errors["failed"]
	r.SetStatusCode(447).
		SetStatusText(r.Messages().Failed).
		RespondWithMessage(message["delete"])
}

func (r *Respond) UpdateSucceeded() {
	message := r.Messages().Errors["success"]
	r.SetStatusCode(200).
		SetStatusText(r.Messages().Success).
		RespondWithMessage(message["update"])
}

func (r *Respond) UpdateFailed() {
	message := r.Messages().Errors["failed"]
	r.SetStatusCode(449).
		SetStatusText(r.Messages().Failed).
		RespondWithMessage(message["update"])
}

func (r *Respond) WrongParameters() {
	r.Error(406, 5406)
}

func (r *Respond) MethodNotAllowed() {
	r.Error(405, 5405)
}

func (r *Respond) ValidationErrors(errors interface{}) {
	r.SetStatusCode(420).
		SetStatusText(r.Messages().Failed).
		SetErrorCode(5420).
		RespondWithResult(errors)
}

func (r *Respond) RequestFieldNotfound() {
	r.Error(446, 1001)
}

func (r *Respond) RequestFieldDuplicated() {
	r.Error(400, 1004)
}

func (r *Respond) Error(statusCode int, errorCode int) {
	message := r.Messages().Errors[strconv.Itoa(errorCode)]
	r.SetStatusCode(statusCode).
		SetStatusText(r.Messages().Failed).
		SetErrorCode(errorCode).
		RespondWithMessage(message["message"])
}
