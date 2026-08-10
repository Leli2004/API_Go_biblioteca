package helpers

import "fmt"

func CacheGetKeyId(id int, module string) string {
	return fmt.Sprintf("biblioteca_%s_get_%d", module, id)
}
