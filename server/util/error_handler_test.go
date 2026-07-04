package util_test

import (
	"bytes"
	"encoding/json"
	"errors"
	"go-playground/server/domain"
	"net/http"
	"net/http/httptest"
	"testing"

	"github.com/gin-gonic/gin"
	"github.com/rs/zerolog"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"

	"go-playground/server/util"
)

func TestLogError_EmitsStructuredFields(t *testing.T) {
	var buf bytes.Buffer
	logger := zerolog.New(&buf)
	util.SetLogErrorLogger(logger)
	defer util.SetLogErrorLogger(zerolog.Nop())

	gin.SetMode(gin.TestMode)
	w := httptest.NewRecorder()
	c, _ := gin.CreateTestContext(w)
	c.Request = httptest.NewRequest(http.MethodGet, "/api/points/earn", nil)

	err := errors.New("db connection refused")
	util.LogError(c, err)

	var entry map[string]interface{}
	require.NoError(t, json.Unmarshal(buf.Bytes(), &entry))

	assert.Equal(t, "error", entry["level"])
	assert.Equal(t, "db connection refused", entry["error"])
	assert.Equal(t, "/api/points/earn", entry["route"])
}

func TestLogError_SystemError_EmitsCode(t *testing.T) {
	var buf bytes.Buffer
	logger := zerolog.New(&buf)
	util.SetLogErrorLogger(logger)
	defer util.SetLogErrorLogger(zerolog.Nop())

	gin.SetMode(gin.TestMode)
	w := httptest.NewRecorder()
	c, _ := gin.CreateTestContext(w)
	c.Request = httptest.NewRequest(http.MethodPost, "/api/earn", nil)

	sysErr := domain.NewSystemError("PointsService.Earn", errors.New("timeout"), "failed to earn points")
	util.LogError(c, sysErr)

	var entry map[string]interface{}
	require.NoError(t, json.Unmarshal(buf.Bytes(), &entry))

	assert.Equal(t, "error", entry["level"])
	assert.NotEmpty(t, entry["error"])
}

func TestHandleError_5xx_GoesToLogError(t *testing.T) {
	var buf bytes.Buffer
	logger := zerolog.New(&buf)
	util.SetLogErrorLogger(logger)
	defer util.SetLogErrorLogger(zerolog.Nop())

	gin.SetMode(gin.TestMode)
	w := httptest.NewRecorder()
	c, _ := gin.CreateTestContext(w)
	c.Request = httptest.NewRequest(http.MethodGet, "/api/test", nil)

	util.HandleError(c, errors.New("unexpected boom"))

	assert.Equal(t, http.StatusInternalServerError, w.Code)

	var entry map[string]interface{}
	require.NoError(t, json.Unmarshal(buf.Bytes(), &entry))
	assert.Equal(t, "error", entry["level"])
}
