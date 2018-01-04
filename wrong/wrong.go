package wrong

func Println(err error) error {
	target, e := core()
	if err != nil {
		__printf__("%s: %s", target, err.Error())
	}
	return e
}
func Panicln(err error) error {
	target, e := core()
	if err != nil {
		__panicf__("%s: %s", target, err.Error())
	}
	return e
}
func Fatalln(err error) error {
	target, e := core()
	if err != nil {
		__fatalf__("%s: %s", target, err.Error())
	}
	return e
}
