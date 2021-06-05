package format

import (
	"bufio"
	"time"

	"github.com/jimweng/networkServer/socket-server/udp/UDP_RFC3164/internal/syslogparser"
)

type LogParts map[string]interface{}

type LogParser interface {
	Parse() error
	Dump() LogParts
	Location(*time.Location)
}

type Format interface {
	GetParser([]byte) LogParser
	GetSplitFunc() bufio.SplitFunc
}

type parserWrapper struct {
	syslogparser.LogParser
}

func (w *parserWrapper) Dump() LogParts {
	return LogParts(w.LogParser.Dump())
}
