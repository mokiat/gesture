package gesture_test

import (
	"testing"

	. "github.com/onsi/ginkgo/v2"
	. "github.com/onsi/gomega"
)

func TestGesture(t *testing.T) {
	RegisterFailHandler(Fail)
	RunSpecs(t, "Gesture Suite")
}
