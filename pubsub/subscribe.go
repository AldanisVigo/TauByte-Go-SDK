package pubsub

import (
	"fmt"

	pubsubSym "bitbucket.org/taubyte/go-sdk-symbols/pubsub"
)

// Subscribe tells a node to subscribe to a given pub-sub channel
// returns an error
func (c *ChannelObject) Subscribe() error {
	err := pubsubSym.SetSubscriptionChannel(c.name)
	if err != 0 {
		return fmt.Errorf("Set subscribe to channel `%s` failed with: %s", c.name, err)
	}

	return nil
}
