package connectors

import (
	"bytes"
	"fmt"
	"os/exec"

	"github.com/microlib/simple"
)

// Connections struct - all backend connections in a common object
type Connectors struct {
	Logger *simple.Logger
}

func NewClientConnections(logger *simple.Logger) Clients {
	return &Connectors{Logger: logger}
}

func (c *Connectors) Error(msg string, val ...interface{}) {
	c.Logger.Error(fmt.Sprintf(msg, val...))
}

func (c *Connectors) Info(msg string, val ...interface{}) {
	c.Logger.Info(fmt.Sprintf(msg, val...))
}

func (c *Connectors) Debug(msg string, val ...interface{}) {
	c.Logger.Debug(fmt.Sprintf(msg, val...))
}

func (c *Connectors) Trace(msg string, val ...interface{}) {
	c.Logger.Trace(fmt.Sprintf(msg, val...))
}

func (c *Connectors) ExecOS(path string, command string, params []string, trim bool) (string, error) {
	var stdout, stderr bytes.Buffer
	var out string
	cmd := exec.Command(command, params...)
	cmd.Dir = path
	cmd.Stdout = &stdout
	cmd.Stderr = &stderr
	err := cmd.Run()
	outStr := stdout.String()
	if err != nil {
		return stderr.String(), err
	}
	if trim {
		if len(outStr) > 1 {
			out = outStr[:len(outStr)-1]
		} else {
			out = outStr
		}
	} else {
		out = outStr
	}
	return out, nil
}
