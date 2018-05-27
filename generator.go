package main

// Generator di file JSON per la creazione dei dati tabellari della GUI
func Generator(data []Commits) {
	var nomi []string
	for i := 0; i < len(data); i++ {
		nomi[i] = data[i].name
	}

}
