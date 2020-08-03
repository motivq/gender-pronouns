package pronouns

import "testing"

func TestSubjectPronoun(t *testing.T) {
	if They.Object() != "Them" {
		t.Errorf("Object was incorrect, got: %v, want: %v.", They.Object(), "Them")
	}
	if They.String() != "They" {
		t.Errorf("Subject was incorrect, got: %v, want: %v.", They.String(), "They")
	}
	if They.Possesive() != "Their" {
		t.Errorf("Possesive was incorrect, got: %v, want: %v.", They.Possesive(), "Their")
	}
	if They.PossesivePronoun() != "Theirs" {
		t.Errorf("PossesivePronoun was incorrect, got: %v, want: %v.", They.PossesivePronoun(), "Theirs")
	}
	if They.Reflexive() != "Themself" {
		t.Errorf("Reflexive was incorrect, got: %v, want: %v.", They.Reflexive(), "Themself")
	}
}
