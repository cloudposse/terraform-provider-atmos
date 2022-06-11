package label

import "strings"

func CreateLabel(
	namespace string,
	tenant string,
	environment string,
	stage string,
	name string,
	delimiter string,
) (string, error) {
	return strings.Join([]string{namespace, tenant, environment, stage, name}, delimiter), nil
}
