package handler

import (
	"net/http"
	"regexp"
	"strings"

	"github.com/dwarvesf/gerr"
	"github.com/gin-gonic/gin"
	"github.com/go-playground/validator/v10"
	"github.com/hieuphq/backend-example/pkg/config"
	"github.com/hieuphq/backend-example/pkg/constant"
	"github.com/hieuphq/backend-example/pkg/util"
	"github.com/hieuphq/backend-example/translation"
)

// Handler for app
type Handler struct {
	log        gerr.Log
	cfg        config.Config
	translator translation.Helper
}

// NewHandler make handler
func NewHandler(cfg config.Config, l gerr.Log, th translation.Helper) *Handler {
	return &Handler{
		log:        l,
		cfg:        cfg,
		translator: th,
	}
}

func (h *Handler) handleError(c *gin.Context, err error) {
	if err == nil {
		return
	}
	locale := c.GetString(constant.LanguageKey)
	tr := h.translator.GetTranslator(locale)

	var parsedErr gerr.Error
	switch arg := err.(type) {
	case validator.ValidationErrors:
		ds := []gerr.Error{}
		childrens := []gerr.CombinedItem{}
		for _, currErr := range arg {
			msg := currErr.Translate(tr)
			targetStr := util.RemoveFirstElementBySeparator(currErr.Namespace(), ".")
			targets := makeKeysFromTarget(targetStr)
			targetCombined := strings.Join(targets, ".")
			ds = append(ds, gerr.E(msg, gerr.Target(targetCombined)))

			childrens = append(childrens, gerr.CombinedItem{
				Keys:    targets,
				Message: msg,
			})
		}

		badRequest := "bad request"
		msg, err := tr.T(badRequest)
		if err != nil {
			msg = badRequest
		}
		rs := gerr.CombinedE(
			http.StatusBadRequest,
			msg,
			childrens,
		)
		parsedErr = *rs.ToError()

	case gerr.Error:
		parsedErr = arg

	case *gerr.Error:
		parsedErr = *arg

	case error:
		str := arg.Error()
		if str == "EOF" {
			parsedErr = gerr.E("bad request", http.StatusBadRequest)
			break
		}
		parsedErr = gerr.E(arg.Error(), http.StatusInternalServerError)
	}

	// log data to console
	logDataRaw, ok := c.Get(constant.LogDataKey)
	traceID := ""
	if ok {
		if ld, parsed := logDataRaw.(gerr.LogInfo); parsed {
			traceID = ld.GetTraceID()
		}
	}
	h.log.Log(logDataRaw, parsedErr)

	c.AbortWithStatusJSON(parsedErr.StatusCode(), parsedErr.ToResponseError(traceID))
}

func makeKeysFromTarget(target string) []string {
	keys := strings.Split(target, ".")
	rs := []string{}
	reg, _ := regexp.Compile("(.+)\\[(.+)\\]")
	for idx := range keys {
		k := keys[idx]
		itms := reg.FindStringSubmatch(k)
		if len(itms) <= 0 {
			rs = append(rs, k)
			continue
		}
		rs = append(rs, itms[1:]...)

	}
	return rs
}

// Healthz handler
// Return "OK"
func (h *Handler) Healthz(c *gin.Context) {
	c.Header("Content-Type", "text/plain")
	c.Writer.WriteHeader(http.StatusOK)
	c.Writer.Write([]byte("OK"))
}
