package utils

import "strings"

func GenNSQtopicName(topic string) string {
	parts := strings.SplitN(topic, "_", 2)
	return parts[0]
}

//  run(uinqId, addrNsqdHTTP string, addrNsqdTCP, addrNsqlookupd *string) {
func Run(uinqId, addrNsqdHTTP string, addrNsqdTCP, addrNsqlookupd *string) {
	go h.run(uinqId, addrNsqdHTTP, addrNsqdTCP, addrNsqlookupd)
}
