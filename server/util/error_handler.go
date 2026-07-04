package util

import (
	"go-playground/pkg/logging"
	"go-playground/server/domain"
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/rs/zerolog"
)

// logErrorLogger is the logger used by LogError; replaceable in tests via SetLogErrorLogger.
var logErrorLogger zerolog.Logger

func init() {
	logErrorLogger = logging.GetLogger()
}

// SetLogErrorLogger replaces the logger used by LogError (test helper).
func SetLogErrorLogger(l zerolog.Logger) {
	logErrorLogger = l
}

// ErrorResponse represents a standardized error response
type ErrorResponse struct {
	Status  string      `json:"status"`
	Message string      `json:"message"`
	Code    string      `json:"code,omitempty"`
	Details interface{} `json:"details,omitempty"`
}

// HandleError is a helper function to handle different types of errors and return appropriate HTTP status codes
func HandleError(c *gin.Context, err error) {
	var response ErrorResponse
	var statusCode int

	switch e := err.(type) {
	case domain.ValidationError:
		statusCode = http.StatusBadRequest
		response = ErrorResponse{
			Status:  "error",
			Message: e.Error(),
			Code:    "VALIDATION_ERROR",
			Details: map[string]string{
				"field": e.Field,
			},
		}

	case domain.ResourceNotFoundError:
		statusCode = http.StatusNotFound
		response = ErrorResponse{
			Status:  "error",
			Message: e.Error(),
			Code:    "NOT_FOUND",
			Details: map[string]string{
				"resource": e.Resource,
				"id":       e.ID,
			},
		}

	case domain.AuthenticationError:
		statusCode = http.StatusUnauthorized
		response = ErrorResponse{
			Status:  "error",
			Message: e.Error(),
			Code:    "UNAUTHORIZED",
		}

	case domain.AuthorizationError:
		statusCode = http.StatusForbidden
		response = ErrorResponse{
			Status:  "error",
			Message: e.Error(),
			Code:    "FORBIDDEN",
		}

	case domain.ResourceConflictError:
		statusCode = http.StatusConflict
		response = ErrorResponse{
			Status:  "error",
			Message: e.Error(),
			Code:    "CONFLICT",
			Details: map[string]string{
				"resource": e.Resource,
			},
		}

	case domain.RateLimitError:
		statusCode = http.StatusTooManyRequests
		response = ErrorResponse{
			Status:  "error",
			Message: e.Error(),
			Code:    "RATE_LIMIT_EXCEEDED",
		}

	case domain.BusinessLogicError:
		// Business logic errors are returned as 200 with error status
		statusCode = http.StatusOK
		response = ErrorResponse{
			Status:  "error",
			Message: e.Error(),
			Code:    e.Code,
		}

	case domain.SystemError:
		statusCode = http.StatusInternalServerError
		LogError(c, e)
		response = ErrorResponse{
			Status:  "error",
			Message: "An internal server error occurred",
			Code:    "INTERNAL_SERVER_ERROR",
		}

	default:
		statusCode = http.StatusInternalServerError
		LogError(c, err)
		response = ErrorResponse{
			Status:  "error",
			Message: "An internal server error occurred",
			Code:    "INTERNAL_SERVER_ERROR",
		}
	}

	c.JSON(statusCode, response)
}

// LogError emits a structured zerolog error entry with route and correlation ID.
func LogError(c *gin.Context, err error) {
	route := ""
	correlationID := ""
	if c != nil && c.Request != nil {
		route = c.Request.URL.Path
		correlationID = c.GetHeader("X-Correlation-ID")
		if correlationID == "" {
			correlationID = c.GetString("correlation_id")
		}
	}

	event := logErrorLogger.Error().Err(err).Str("route", route)
	if correlationID != "" {
		event = event.Str("correlation_id", correlationID)
	}
	event.Msg("internal server error")
}

// EmptyResponse returns a 200 OK with empty data when no results are found
func EmptyResponse(c *gin.Context) {
	c.JSON(http.StatusOK, ErrorResponse{
		Status:  "success",
		Message: "No results found",
		Details: []interface{}{},
	})
}
