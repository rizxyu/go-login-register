package helper

func IfError(err error) {
	if err != nil {
		panic(err)
	}
}
