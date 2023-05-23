package convert

import (
    "net/http"

   // "github.com/gin-gonic/gin"
)

type Respond struct {
    Metas          interface{}
    Additional     interface{}
    ResponseMessage string
    HasErrors       bool
}



func (r *Respond) SetAdditional(data map[string]interface{}) *Respond {
    r.Additional = data
    return r
}


func (r *Response) SetMetas(metas []interface{}) *Response {
    r.metas = metas
    return r
}

func (r *Respond) RespondWithMetas(metas interface{}) {
	data := map[string]interface{}{
	    "succeed" : !hasErrors(),
		"message": message,
        "results" : data,
        "metas" : thismetas(),
	}
	if r.errorCode != 0 {
		data["error"] = r.errorCode
	}
	r.writer.WriteHeader(r.statusCode)
	r.writeJSON(data)
}


func (r *Response) Message() string {
    if r.responseMessage == "" {
        return "rest.respond.successful_message"
    }
    return r.responseMessage
}

func (r *Response) SetMessage(message string) *Response {
    r.responseMessage = message
    return r
}

func (r *Respond) RespondWithMessage(message interface{}) {
	data := map[string]interface{}{
	    "succeed" : !hasErrors(),
		"message": message,
        "results" : data,
        "metas" : thismetas(),
	}
	if r.errorCode != 0 {
		data["error"] = r.errorCode
	}
	r.writer.WriteHeader(r.statusCode)
	r.writeJSON(data)
}



func (r *Response) SetCreatedMessage() *Response {
    r.responseMessage = "rest.respond.successful_created_message"
    return r
}

func (r *Response) SetUpdatedMessage() *Response {
    r.responseMessage = "rest.respond.successful_update_message"
    return r
}

func (r *Response) HasErrors() bool {
    return r.hasErrors
}

func (r *Response) SetHasErrors(hasErrors bool) *Response {
    r.hasErrors = hasErrors
    return r
}
func (c *Controller) respond(data interface{}, statusCode int, headers map[string]string) *JsonResponse {
    if collection, ok := data.(ResourceCollection); ok {
        if paginator, ok := collection.Resource.(LengthAwarePaginatorContract); ok {
            return c.respondWithPagination(collection, []interface{}{}, headers)
        }
    }
    response :=  interface{} {
        "succeed": !c.hasErrors(),
        "message": c.message(),
        "results": data,
        "metas":   c.metas(),
    }
    for key, value := range c.additional() {
        response[key] = value
    }
    return &JsonResponse{Data: response, StatusCode: statusCode, Headers: headers}
}

func (c *Controller) respondWithPagination(resource ResourceCollection, metas []interface{}, headers map[string]string) *JsonResponse {
    data := resource.response().getData(true)
    metas["links"] = data["links"]
    return c.setMetas(append(append(data["meta"], metas...), c.metas()...)).respond(data["data"], Response::HTTP_OK, headers)
}
func respondInvalidParameters(message *string) *JsonResponse {
    return respondWithError(
        message,
        http.StatusUnprocessableEntity,
    )
}

func respondUnauthorized(exception *Exception, message *string) *JsonResponse {
    return respondWithError(
        exception,
        message,
        http.StatusUnauthorized,
    )
}

func respondForbidden(message *string) *JsonResponse {
    return respondWithError(
        message,
        http.StatusForbidden,
    )
}

func respondNotAcceptable(message *string) *JsonResponse {
    return respondWithError(
        message,
        http.StatusNotAcceptable,
    )
}

