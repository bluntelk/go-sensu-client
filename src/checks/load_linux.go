package checks

import (
	"io/ioutil"
	"strings"
	"fmt"
)

// PLATFORMS
//   Linux

func (load *LoadStats) createLoadAveragePayload(timestamp uint) (string, error) {
	var payload string
	content, err := ioutil.ReadFile("/proc/loadavg")
	if nil != err {
		return payload, err
	}

	bits := strings.Split(string(content), " ")

	payload = fmt.Sprintf("loadavg.1 %s %d\n", bits[0], timestamp)
	payload += fmt.Sprintf("loadavg.5 %s %d\n", bits[1], timestamp)
	payload += fmt.Sprintf("loadavg.15 %s %d\n", bits[2], timestamp)

	return payload, nil
}
