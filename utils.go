package main

import (
	"encoding/json"
	"fmt"
	"os"
)

func parseFiles(constructors *Constructors, constructorsFullInfo *ConstructorsFullInfo) error {

	basic_info_file, err := os.Open("data/f1_constructors_basic_info.json")
	if err != nil {
		return fmt.Errorf("error occured while opening basic info file :- %w", err)
	}

	err = json.NewDecoder(basic_info_file).Decode(&constructors)
	if err != nil {
		return fmt.Errorf("error occured while decoding basic info file :- %w", err)
	}

	details_file, err := os.Open("data/f1_constructors_details.json")
	if err != nil {
		return fmt.Errorf("error occured while opening details file :- %w", err)
	}

	err = json.NewDecoder(details_file).Decode(&constructorsFullInfo)
	if err != nil {
		return fmt.Errorf("error occured while decoding details file :- %w", err)
	}

	return nil
}
