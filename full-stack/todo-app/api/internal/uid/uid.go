package uid

import (
	"fmt"

	"github.com/segmentio/ksuid"
)

// NewUID new uid
func NewUID(prefix string) string {
	return fmt.Sprintf("%s_%s", prefix, ksuid.New().String())
}
