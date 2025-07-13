package cfg

import (
	"fmt"

	"github.com/reiver/oxio-20250712/env"
)

func WebServerTCPAddress() string {
	return fmt.Sprintf(":%s", env.TcpPort)
}
