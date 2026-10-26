// Package ping answers health probes.
package ping

// Reply is what Handler writes.
const Reply = "pong"

// Check reports whether the service is healthy.
func Check() string {
	return Reply
}
