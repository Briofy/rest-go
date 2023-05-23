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
	metas          interface{}
        additional     interface{}
    	responseMessage string
    	hasErrors       bool
}

// Set language of responses
//
// @author Alireza Josheghani <josheghani.dev@gmail.com>
// @since 6 Jun 2021
// @return *Respond
func (r *Respond) Language(lang string) *Respond {
	r.lang = lang
	return r
}

// New respond type with custom writer
//
// @author Alireza Josheghani <josheghani.dev@gmail.com>
// @since 6 Jun 2021
// @return *Respond
func NewWithWriter(w http.ResponseWriter) *Respond {
	return &Respond{writer: w, messages: NewMessages()}
}

// Get message type
//
// @author Alireza Josheghani <josheghani.dev@gmail.com>
// @since 15 Mar 2018
// @return *Message
func (r *Respond) Messages() *Messages {
	if r.lang != "" {
		r.messages.Lang = r.lang
	}
	r.messages.load()
	return r.messages
}

// Set status code of response and set default value as 0
//
// @author Alireza Josheghani <josheghani.dev@gmail.com>
// @since 15 Mar 2018
// @return Respond
// @param code int
func (r *Respond) SetStatusCode(code int) *Respond {
	r.statusCode = code
	return r
}

// Set status text of response
//
// @author Alireza Josheghani <josheghani.dev@gmail.com>
// @since 15 Mar 2018
// @return Respond
// @param text string
func (r *Respond) SetStatusText(text string) *Respond {
	r.statusText = text
	return r
}

// Set status code of response and set default value as 0
//
// @author Alireza Josheghani <josheghani.dev@gmail.com>
// @since 15 Mar 2018
// @return Respond
// @param code int
func (r *Respond) SetErrorCode(code int) *Respond {
	r.errorCode = code
	return r
}

// Write json data to http.ResponseWriter
//
// @author Alireza Josheghani <josheghani.dev@gmail.com>
// @since 6 Jun 2021
// @param data interface{}
// @return error
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

// Pass response with result data like this array
//
//      array := map[string]interface{} {
//        "status": respond.statusText,
//        "result": result,
//      }
//
// @author Alireza Josheghani <josheghani.dev@gmail.com>
// @since 15 Mar 2018
// @param result map[string]interface{}
// @return error
func (r *Respond) RespondWithResult(result interface{}) {
	r.writer.WriteHeader(r.statusCode)
	r.writeJSON(map[string]interface{}{
		"succeed": !r.hasErrors,
		"result": result,
		"metas": r.metas,
		"additional": r.additional,
	})
}

// Pass response with message text as string
//
// @author Alireza Josheghani <josheghani.dev@gmail.com>
// @since 15 Mar 2018
// @param message interface{}
// @return error
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

// return notfound result
//
// @author Alireza Josheghani <josheghani.dev@gmail.com>
// @since 15 Mar 2018
// @return (statuscode int, result interface{})
func (r *Respond) NotFound() {
	r.Error(404, 5404)
}

// return success result with data
//
//      data := map[string]interface{} {
//        "data": "somedata"
//      }
//
// @author Alireza Josheghani <josheghani.dev@gmail.com>
// @since 15 Mar 2018
// @param data map[string]interface{}
func (r *Respond) Succeed(data interface{}) {
	r.SetStatusCode(http.StatusOK).
		SetHasErrors(false).
		RespondWithResult(data)
}

// Insert action is succeed
//
// @author Alireza Josheghani <josheghani.dev@gmail.com>
// @since 15 Mar 2018
func (r *Respond) InsertSucceeded() {
	message := r.Messages().Errors["success"]
	r.SetStatusCode(http.StatusOK).
		SetStatusText(r.Messages().Success).
		RespondWithMessage(message["insert"])
}

// Insert action is failed
//
// @author Alireza Josheghani <josheghani.dev@gmail.com>
// @since 15 Mar 2018
func (r *Respond) InsertFailed() {
	message := r.Messages().Errors["failed"]
	r.SetStatusCode(448).
		SetStatusText(r.Messages().Failed).
		RespondWithMessage(message["insert"])
}

// Delete action is succeed
//
// @author Alireza Josheghani <josheghani.dev@gmail.com>
// @since 15 Mar 2018
// @return (statuscode int, result interface{})
func (r *Respond) DeleteSucceeded() {
	message := r.Messages().Errors["success"]
	r.SetStatusCode(200).
		SetStatusText(r.Messages().Success).
		RespondWithMessage(message["delete"])
}

// Delete action is failed
//
// @author Alireza Josheghani <josheghani.dev@gmail.com>
// @since 15 Mar 2018
// @return (statuscode int, result interface{})
func (r *Respond) DeleteFailed() {
	message := r.Messages().Errors["failed"]
	r.SetStatusCode(447).
		SetStatusText(r.Messages().Failed).
		RespondWithMessage(message["delete"])
}

// Update action is succeed
//
// @author Alireza Josheghani <josheghani.dev@gmail.com>
// @since 15 Mar 2018
// @return (statuscode int, result interface{})
func (r *Respond) UpdateSucceeded() {
	message := r.Messages().Errors["success"]
	r.SetStatusCode(200).
		SetStatusText(r.Messages().Success).
		RespondWithMessage(message["update"])
}

// Update action is failed
//
// @author Alireza Josheghani <josheghani.dev@gmail.com>
// @since 15 Mar 2018
// @return (statuscode int, result interface{})
func (r *Respond) UpdateFailed() {
	message := r.Messages().Errors["failed"]
	r.SetStatusCode(449).
		SetStatusText(r.Messages().Failed).
		RespondWithMessage(message["update"])
}

// Wrong parameters are entered
//
// @author Alireza Josheghani <josheghani.dev@gmail.com>
// @since 15 Mar 2018
func (r *Respond) WrongParameters() {
	r.Error(406, 5406)
}

// Wrong parameters are entered
//
// @author Alireza Josheghani <josheghani.dev@gmail.com>
// @since 15 Mar 2018
func (r *Respond) MethodNotAllowed() {
	r.Error(405, 5405)
}

// There ara validation translations
//
// @author Alireza Josheghani <josheghani.dev@gmail.com>
// @since 15 Mar 2018
// @param translations map[string]interface{}
func (r *Respond) ValidationErrors(errors interface{}) {
	r.SetStatusCode(420).
		SetHasErrors(true).
		SetErrorCode(5420).
		RespondWithResult(errors)
}

// The request field is not found
//
// @author Alireza Josheghani <josheghani.dev@gmail.com>
// @since 15 Mar 2018
// @return (statuscode int, result interface{})
func (r *Respond) RequestFieldNotfound() {
	r.Error(446, 1001)
}

// The request field is duplicated
//
// @author Alireza Josheghani <josheghani.dev@gmail.com>
// @since 15 Mar 2018
// @return (statuscode int, result interface{})
func (r *Respond) RequestFieldDuplicated() {
	r.Error(400, 1004)
}

// The error message
//
// @author Alireza Josheghani <josheghani.dev@gmail.com>
// @since 15 Mar 2018
// @param statusCode int,errorCode string
// @return (statuscode int, result interface{})
func (r *Respond) Error(statusCode int, errorCode int) {
	message := r.Messages().Errors[strconv.Itoa(errorCode)]
	r.SetStatusCode(statusCode).
		SetStatusText(r.Messages().Failed).
		SetErrorCode(errorCode).
		RespondWithMessage(message["message"])
}

// set additional of response
func (r *Respond) SetAdditional(data map[string]interface{}) *Respond {
    r.additional = data
    return r
}

//set metadata for response
func (r *Respond) SetMetas(metas []interface{}) *Respond {
    r.metas = metas
    return r
}

// get message of response,  default is successful
func (r *Respond) Message() string {
    message := r.Messages().Errors["general"]
    r.responseMessage = message["successful_message"].(string)
    return r.responseMessage
}

// set message of response
func (r *Respond) SetMessage(message string) *Respond {
    r.responseMessage = message
    return r
}

//set haseeror
func (r *Respond) SetHasErrors(hasErrors bool) *Respond {
    r.hasErrors = hasErrors
    return r
}

func (r *Respond) Respond(result interface{}) {
	r.writer.WriteHeader(r.statusCode)
	r.writeJSON(map[string]interface{}{
		"succeed": !r.hasErrors,
		"result": result,
		"metas": r.metas,
		"additional": r.additional,
	})
}
func (r *Respond) RespondWithError(statusCode int, errorCode int) {
	message := r.Messages().Errors["general"]
	r.SetStatusCode(statusCode).
		SetHasErrors(true).
		SetErrorCode(errorCode).
		Respond(message["error"])
}

func (r *Respond) SetCreatedMessage() *Respond {
    message:= r.Messages().Errors["general"]
    r.responseMessage = message["successful_message"].(string)
    return r
}

func (r *Respond) SetUpdatedMessage() *Respond {
    r.responseMessage = r.Messages().Errors["general"]["successful_update_message"].(string)
    return r
}

