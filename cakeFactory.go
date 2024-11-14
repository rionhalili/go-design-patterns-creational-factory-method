package main

import "fmt"

func getCake(name string) (ICake, error) {
	if name == "Apple" {
		return newAppleCake(), nil
	}
	if name == "Orange" {
		return newOrangeCake(), nil
	}
	return nil, fmt.Errorf("selected cake does not exist")
}
