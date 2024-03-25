package connectors

import (
	"fmt"

	"github.com/kroulyhq/krouly/playbooks"
)

type Connector struct {
	// fields required for conection, url, creds, etc.
}

func NewConnector() *Connector {
	// pre initi and config
	return &Connector{}
}

func (c *Connector) ExtractData() error {
	// logic to extract data
	err := playbooks.RunPlaybook("../playbooks/krouly.sample.yaml")
	if err != nil {
		return fmt.Errorf("error running playbook: %v", err)
	}
	return nil
}
