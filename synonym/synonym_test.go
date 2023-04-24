package main

import (
	"reflect"
	"testing"
)

func TestRemoveDuplicates(t *testing.T) {
	// test removeDuplicates

	ans := removeDuplicates([]string{"a", "b", "a", "c", "b"})
	if !reflect.DeepEqual(ans, []string{"a", "b", "c"}) {
		t.Errorf("RemoveDuplicates failed")
	}
}

func TestUpdateSynonyms(t *testing.T) {
	words := []string{"a", "b"}

	updateSynonyms(&safeSynonymsMap, words)
	if !reflect.DeepEqual(synonymsMap, map[string][]string{"a": {"b"}, "b": {"a"}}) {
		t.Errorf("First UpdateSynonyms failed")
	}
	wordsNew := []string{"b", "c"}
	updateSynonyms(&safeSynonymsMap, wordsNew)
	if !reflect.DeepEqual(synonymsMap, map[string][]string{"a": {"b", "c"}, "b": {"c", "a"}, "c": {"b", "a"}}) {
		t.Errorf("Second UpdateSynonyms failed")
	}
}
