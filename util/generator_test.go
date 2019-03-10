package util

import (
	//"fmt"
	"testing"
)

func TestGeneratePlayerID(t *testing.T) {
	num := 100
	list := make([]string, num)

	for i := 0; i < len(list); i++ {
		newID := GeneratePlayerID()
		for j := 0; j < i; j++ {
			oldID := list[j]
			if newID == oldID {
				t.Errorf("two IDs are equal")
			} else {
				//fmt.Printf("old: %s, new: %s\n", oldID, newID)
			}
		}
	}
}
