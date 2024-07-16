package errplus

import (
	"fmt"
)

// Merge is useful for passing into logger.Log()
//
//	logger.Log(errplus.Merge(errplus.ToError(err).Args,
//			"level", "error",
//			"msg", "something went wrong")...)
func Merge(args1 []any, args2 ...any) []any {
	args := make([]any, 0, len(args1)+len(args2))

	args = append(args, args1...)
	args = append(args, args2...)

	keys := map[interface{}]int{}
	for i := 0; i < len(args); i += 2 {
		key := args[i]
		count, found := keys[key]
		if found {
			args[i] = fmt.Sprintf("%s-duplicate-%d", key, count)
		}

		keys[key] = count + 1
	}

	return args
}
