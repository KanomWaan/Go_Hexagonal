package logs

import "go.uber.org/zap"

var log *zap.Logger

func init() {

	log, _ = zap.NewProduction()

}
