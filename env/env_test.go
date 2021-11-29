package env

import (
	"bytes"
	"github.com/stretchr/testify/assert"
	"log"
	"os"
	"testing"
)

func TestRequire(t *testing.T) {
	init := func() {
		_ = os.Setenv("REQUIRED_KEY", "value")
	}
	reset := func() {
		_ = os.Unsetenv("REQUIRED_KEY")
		resetState()
	}

	t.Run("Happy", func(t *testing.T) {
		init()
		val := Require("REQUIRED_KEY")
		assert.Equal(t, val, "value")
		assert.Len(t, panickingMessages, 0)
		reset()
	})

	t.Run("Unhappy, env not defined", func(t *testing.T) {
		init()
		val := Require("REQUIRED_KEY_2")
		val2 := Require("REQUIRED_KEY_3", "this env is important")

		assert.Equal(t, val, "")
		assert.Equal(t, val2, "")
		assert.Equal(t, panickingMessages[0], "REQUIRED_KEY_2 env is required.")
		assert.Equal(t, panickingMessages[1], "REQUIRED_KEY_3 env is required. (this env is important)")
		reset()
	})
}

func TestWarnIfEmpty(t *testing.T) {
	init := func() {
		_ = os.Setenv("WARN_KEY", "value")
	}
	reset := func() {
		_ = os.Unsetenv("WARN_KEY")
		resetState()
	}

	t.Run("Happy", func(t *testing.T) {
		init()
		val := WarnIfEmpty("WARN_KEY")
		assert.Equal(t, val, "value")
		assert.Len(t, warningMessages, 0)
		reset()
	})

	t.Run("Unhappy, env not defined", func(t *testing.T) {
		init()
		val := WarnIfEmpty("WARN_KEY_2")
		val2 := WarnIfEmpty("WARN_KEY_3", "this env may be important")

		assert.Equal(t, val, "")
		assert.Equal(t, val2, "")
		assert.Equal(t, warningMessages[0], "WARN_KEY_2 env is empty, it may be needed.")
		assert.Equal(t, warningMessages[1], "WARN_KEY_3 env is empty, it may be needed. (this env may be important)")
		reset()
	})
}

func TestDefault(t *testing.T) {
	init := func() {
		_ = os.Setenv("DEFAULT_KEY", "value")
	}
	reset := func() {
		_ = os.Unsetenv("DEFAULT_KEY")
		resetState()
	}

	t.Run("Happy", func(t *testing.T) {
		init()
		val := Default("DEFAULT_KEY", "default_value")
		assert.Equal(t, val, "value")
		reset()
	})

	t.Run("Happy, env not defined, use default instead", func(t *testing.T) {
		init()
		val := Default("DEFAULT_KEY_2", "default_value")
		assert.Equal(t, val, "default_value")
		reset()
	})
}

func TestAssert(t *testing.T) {
	var buffer bytes.Buffer
	log.SetOutput(&buffer)

	t.Run("Happy", func(t *testing.T) {
		// Nothing to log
		Assert()
		assert.Empty(t, buffer.String())
		buffer.Reset()

		// Assign message to be logged
		preWarn("Message 1.")
		preWarn("Message 2.")
		Assert()
		assert.Contains(t, buffer.String(), "Message 1.\nMessage 2.")
		buffer.Reset()

		// Assign message to be panicked
		prePanic("Message 1.")
		prePanic("Message 2.")
		assert.Panics(t, func() {
			Assert()
		})
		assert.Contains(t, buffer.String(), "Message 1.\nMessage 2.")
		buffer.Reset()
	})
}
