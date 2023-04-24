package main

import (
	"os"
	"reflect"
	"testing"
)

func TestMain(m *testing.M) {
	setup()
	code := m.Run()
	os.Exit(code)
}

func setup() {
	safeSynonymsMap.clear()
}

func TestRemoveDuplicates(t *testing.T) {
	ans := removeDuplicates([]string{"a", "b", "a", "c", "b"})
	if !reflect.DeepEqual(ans, []string{"a", "b", "c"}) {
		t.Errorf("RemoveDuplicates failed")
	}
}

func TestUpdateSynonyms(t *testing.T) {
	words := []string{"a", "b"}

	safeSynonymsMap.updateSynonyms(words)
	if !reflect.DeepEqual(synonymsMap, map[string][]string{"a": {"b"}, "b": {"a"}}) {
		t.Errorf("First UpdateSynonyms failed")
	}
	wordsNew := []string{"b", "c"}
	safeSynonymsMap.updateSynonyms(wordsNew)
	if !reflect.DeepEqual(synonymsMap, map[string][]string{"a": {"b", "c"}, "b": {"c", "a"}, "c": {"b", "a"}}) {
		t.Errorf("Second UpdateSynonyms failed")
	}
}

func TestClearMap(t *testing.T) {
	key := "a"
	safeSynonymsMap.sm[key] = []string{"b"}
	safeSynonymsMap.clear()
	_, exists := safeSynonymsMap.sm[key]
	if exists {
		t.Errorf("Method Clear of a Safe Map failed")
	}
}
