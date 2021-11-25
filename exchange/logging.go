package exchange

import (
	"github.com/streamingfast/logging"
	"go.uber.org/zap"
)

var zlog *zap.Logger

func init() {
	logging.Register("github.com/biswapcom/sparkle-biswap/exchange", &zlog)
}
