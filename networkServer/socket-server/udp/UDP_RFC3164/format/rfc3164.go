package format

import (
	"bufio"

	"github.com/jimweng/networkServer/socket-server/udp/UDP_RFC3164/internal/syslogparser/rfc3164"
)

type RFC3164 struct{}

func (f *RFC3164) GetParser(line []byte) LogParser {
	return &parserWrapper{rfc3164.NewParser(line)}
}

func (f *RFC3164) GetSplitFunc() bufio.SplitFunc {
	return nil
}
