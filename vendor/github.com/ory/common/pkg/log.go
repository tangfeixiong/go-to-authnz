package pkg

import "github.com/go-errors/errors"

// import log "github.com/Sirupsen/logrus"
import log "github.com/sirupsen/logrus"

func LogError(err error, code int) {
	stack := "not available."
	if ee, ok := err.(*errors.Error); ok {
		stack = ee.ErrorStack()
	}
	message := "no message available"
	if err != nil {
		message = err.Error()
	}
	log.WithFields(log.Fields{
		"error":        err,
		"errorMessage": message,
		"statusCode":   code,
		"stack":        stack,
	}).Warnln("Request could not be served.")
}
