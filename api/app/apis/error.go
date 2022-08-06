
type APIError struct {
    Code int
    Message string
}

func JSONErrorHandler(err error, context *echo.Context) {
    code := http.StatusInternalServerError
    msg := http.StatusText(code)

    if he, ok := err.(*HTTPError); ok {
        code = he.Code
        msg = he.Message
    }
    if e.debug {
        msg = err.Error()
    }

    var apierr APIError
    apierr.Code    = code
    apierr.Message = msg

    if !c.Response().Committed() {
        c.JSON(code, apierr)
    }
    e.logger.Debug(err)
}