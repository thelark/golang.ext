package wrong

func Println(err error) error {
	if err != nil {
		__printf__("%v\n", err)
	}
	return err
}
func Panicln(err error) error {
	if err != nil {
		__printf__("%v\n", err)
	}
	return err
}
func Fatalln(err error) error {
	if err != nil {
		__printf__("%v\n", err)
	}
	return err
}
