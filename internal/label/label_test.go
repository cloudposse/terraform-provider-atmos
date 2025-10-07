package label

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestLabel(t *testing.T) {
	namespace := "eg"
	tenant := "plat"
	environment := "ue2"
	stage := "dev"
	name := "test"
	delimiter := "-"

	label, err := CreateLabel(namespace, tenant, environment, stage, name, delimiter)
	assert.Nil(t, err)
	assert.Equal(t, "eg-plat-ue2-dev-test", label)
	t.Log(label)
}
