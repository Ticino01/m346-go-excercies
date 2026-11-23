package main

import "fmt"

func main() {
	existing := map[string]int{"Aprhodite": 0, "Apollo": 0, "Ares": 0,
		"Artemis": 0, "Athena": 0, "Demeter": 0, "Dionysus": 0, "Hephaestus": 0,
		"Hera": 0, "Hermes": 0, "Hestia": 0, "Poseidon": 0, "Zeus": 0}
	required := map[string]int{"Apollo": 0, "Ares": 0, "Artemis": 0, "Athena": 0,
		"Dionysus": 0, "Hades": 0, "Hephaestus": 0, "Hermes": 0, "Hestia": 0,
		"Persephone": 0, "Plouton": 0, "Poseidon": 0, "Zeus": 0}

	fmt.Println("existing VMs:", existing)
	fmt.Println("required VMs:", required)

	var to_be_removed map[string]int
	var to_be_created map[string]int

	// TODO: to_be_removed = set difference between existing and required
	to_be_removed = make(map[string]int)
	for name := range existing {
		if _, exists := required[name]; !exists {
			to_be_removed[name] = existing[name]
		}
	}

	// TODO: to_be_created = set difference between required and existing
	to_be_created = make(map[string]int)
	for name := range required {
		if _, exists := existing[name]; !exists {
			to_be_created[name] = required[name]
		}
	}

	fmt.Println("to be removed VMs:", to_be_removed)
	fmt.Println("to be created VMs:", to_be_created)
}