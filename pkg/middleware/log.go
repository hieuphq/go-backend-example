package middleware

import (
	"github.com/hieuphq/backend-example/pkg/constant"
	"github.com/hieuphq/backend-example/pkg/middleware/random"

	"github.com/dwarvesf/gerr"
	"github.com/gin-gonic/gin"
)

const headerXRequestID = "X-Request-ID"
const headerAcceptLanguage = "Accept-Language"

// GinConfig log config for gin framework
type GinConfig struct {
	Generator          func() string
	HeaderRequestIDKey func() string
}

// NewLogDataMiddleware make a log data middleware
func NewLogDataMiddleware(service string, environment string, config ...GinConfig) gin.HandlerFunc {
	var cfg GinConfig
	if len(config) > 0 {
		cfg = config[0]
	}

	// Set config default values
	if cfg.Generator == nil {
		cfg.Generator = func() string {
			return random.String(20)
		}
	}

	// Set config default values
	if cfg.HeaderRequestIDKey == nil {
		cfg.HeaderRequestIDKey = func() string {
			return headerXRequestID
		}
	}

	return func(c *gin.Context) {
		rid := c.GetHeader(cfg.HeaderRequestIDKey())
		if rid == "" {
			rid = cfg.Generator()
		}

		dt := gerr.NewLogInfo(
			gerr.Service(service),
			gerr.Environment(environment),
			gerr.RequestInfo{
				Method:    c.Request.Method,
				Path:      c.Request.URL.Path,
				IP:        c.ClientIP(),
				UserAgent: c.Request.UserAgent(),
				TraceID:   rid,
			})
		c.Set(constant.LogDataKey, dt)
		c.Writer.Header().Set(cfg.HeaderRequestIDKey(), rid)

		lang := c.GetHeader(headerAcceptLanguage)
		c.Set(constant.LanguageKey, lang)
	}
}
