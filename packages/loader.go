package ManuUtils

type Loader struct {
	label string
	level int
}

func GenerateLoader(label string) *Loader {

	loader := new(Loader)

	loader.label = label
	loader.level = 0

	return loader

}


