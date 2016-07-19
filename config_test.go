package ya_go_config

import (
	"testing"
	"github.com/stretchr/testify/assert"
	"time"
)

func TestSimple(t *testing.T) {
	defaults := map[string]interface{}{
		"a.str": "default string",
		"a.int": 120,
		"a.float": 31.4,
		"b.key": "defval",
		"c.bool_t": true,
		"c.bool_f": false,
		"d.duration": "1h20m",
	}

	config := New(defaults)
	config.Load("./test.json")

	assert.Equal(t, "String value", config.GetStr("a.str"))
	assert.Equal(t, 12, config.GetInt("a.int"))
	assert.Equal(t, 3.14, config.GetFloat64("a.float"))
	assert.Equal(t, "defval", config.GetStr("b.key"))
	assert.Equal(t, true, config.GetBool("c.bool_t"))
	assert.Equal(t, false, config.GetBool("c.bool_f"))
	assert.Equal(t, time.Hour + 20*time.Minute, config.GetDuration("d.duration"))
	assert.Panics(t, func() { config.GetStr("d.non_existent_key") })


	configStrState := `{
  "a.float": 3.14,
  "a.int": 12,
  "a.str": "String value",
  "b.key": "defval",
  "c.bool_f": false,
  "c.bool_t": true,
  "d.duration": "1h20m"
}`
	assert.Equal(t, configStrState, config.String())
}

// TODO: Add some more complex tests