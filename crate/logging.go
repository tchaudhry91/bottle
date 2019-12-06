package crate

import (
	"github.com/go-kit/kit/log"
	"github.com/tchaudhry91/bottle/bottle"
	"time"
)

type LoggingMiddleware struct {
	logger log.Logger
	next   CrateService
}

func NewLoggingMiddleware(logger log.Logger, next CrateService) *LoggingMiddleware {
	return &LoggingMiddleware{
		logger: logger,
		next:   next,
	}
}

func (mw *LoggingMiddleware) StoreBottle(b *bottle.Bottle) (err error) {
	defer func(begin time.Time) {
		mw.logger.Log(
			"method", "store",
			"id", b.ID,
			"err", err,
			"took", time.Since(begin),
		)
	}(time.Now())
	err = mw.next.StoreBottle(b)
	return
}

func (mw *LoggingMiddleware) DrainBottle(id string) (b *bottle.Bottle, err error) {
	defer func(begin time.Time) {
		mw.logger.Log(
			"method", "drain",
			"id", id,
			"err", err,
			"took", time.Since(begin),
		)
	}(time.Now())
	b, err = mw.next.DrainBottle(id)
	return
}

func (mw *LoggingMiddleware) PourBottle(id string) (b *bottle.Bottle, err error) {
	defer func(begin time.Time) {
		mw.logger.Log(
			"method", "pour",
			"id", id,
			"err", err,
			"took", time.Since(begin),
		)
	}(time.Now())
	b, err = mw.next.PourBottle(id)
	return
}
