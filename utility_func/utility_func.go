package utility_func

import "log"

func FailOnError(err error, msg string) {
	if err != nil {
		log.Fatalf("%s: %s", msg, err)
		//panic(fmt.Sprintf("%s: %s", msg, err))
	}
}
