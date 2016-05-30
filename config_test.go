package ya_go_config

import (
	"testing"
	"github.com/stretchr/testify/assert"
	log "github.com/Sirupsen/logrus"
)

func TestSimple(t *testing.T) {
	log.Infof("OK")


	config := New(map[string]interface{}{
		"a.str": "default string",
		"a.int": 120,
		"a.float": 31.4,
		"b.key": "defval",
		"c.bool_t": true,
		"c.bool_f": false})

	config.Load("./test.json")

	assert.Equal(t, "String value", config.GetStr("a.str"))
	assert.Equal(t, 12, config.GetInt("a.int"))
	assert.Equal(t, 3.14, config.GetFloat64("a.float"))
	assert.Equal(t, "defval", config.GetStr("b.key"))
	assert.Equal(t, true, config.GetStr("c.bool_t"))
	assert.Equal(t, false, config.GetStr("c.bool_f"))
}


// TODO: Add some more complex tests