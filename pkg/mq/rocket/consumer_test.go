package rocket

import (
	"fmt"
	"testing"
)

func TestName(t *testing.T) {
	Consumer(
		Message{"test-group", "test-topic",
			func(msg string) {
				fmt.Printf("msg %s", msg)
			},
		},
		Message{"test2-group", "test2-topic",
			func(msg string) {
				fmt.Printf("msg %s", msg)
			},
		},
	)
}
