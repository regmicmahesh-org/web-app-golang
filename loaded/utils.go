package main

func _panic(err error) {
	if err != nil {
		panic(err)
	}
}
