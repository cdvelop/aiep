package app

type dbAdapter interface {
	GetForId(id string) (data string)
}

func AppRun(db dbAdapter) {
	println("AppRun")
}
