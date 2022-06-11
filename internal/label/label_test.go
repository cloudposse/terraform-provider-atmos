package label

import (
	"github.com/stretchr/testify/assert"
	"testing"
)

func TestLabel(t *testing.T) {
	namespace := "eg"
	tenant := "plat"
	environment := "ue2"
	stage := "dev"
	name := "test"
	delimiter := "-"

	var label, err = CreateLabel(namespace, tenant, environment, stage, name, delimiter)
	assert.Nil(t, err)
	assert.Equal(t, "eg-plat-ue2-dev-test", label)
	t.Log(string(label))
}
