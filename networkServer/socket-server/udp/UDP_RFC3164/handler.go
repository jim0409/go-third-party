package syslog

import "go-third-party/networkServer/socket-server/udp/UDP_RFC3164/format"

type Handler interface {
	Handle(format.LogParts, int64, error)
}

type LogPartsChannel chan format.LogParts

type ChannelHandler struct {
	channel LogPartsChannel
}

func NewChannelHandler(channel LogPartsChannel) *ChannelHandler {
	return &ChannelHandler{
		channel: channel,
	}
}

func (h *ChannelHandler) Handle(LogParts format.LogParts, messageLength int64, err error) {
	h.channel <- LogParts
}
