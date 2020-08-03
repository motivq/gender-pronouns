package pronouns

// SubjectPronoun holds the Subject convertion funcs
type SubjectPronoun string

// Subject Pronouns
const (
	Fae  SubjectPronoun = "Fae"
	Ae   SubjectPronoun = "Ae"
	E    SubjectPronoun = "E"
	Ey   SubjectPronoun = "Ey"
	He   SubjectPronoun = "He"
	Per  SubjectPronoun = "Per"
	She  SubjectPronoun = "She"
	They SubjectPronoun = "They"
	Ve   SubjectPronoun = "Ve"
	Xe   SubjectPronoun = "Xe"
	Ze   SubjectPronoun = "Ze"
	Zie  SubjectPronoun = "Zie"
)

// Subject returns the string representation
func (p SubjectPronoun) String() string {
	return string(p)
}

// Object returns the object for of the Subject as a string
func (p SubjectPronoun) Object() string {
	switch p {
	case Fae:
		return "Faer"
	case Ae:
		return "Aer"
	case E, Ey:
		return "Em"
	case He:
		return "Him"
	case Per:
		return "Per"
	case She:
		return "Her"
	case They:
		return "Them"
	case Ve:
		return "Ve"
	case Xe:
		return "Xe"
	case Ze, Zie:
		return "Hir"
	}
	return ""
}

// Possesive returns the Possesive for of the Subject as a string
func (p SubjectPronoun) Possesive() string {
	switch p {
	case Fae:
		return "Faer"
	case Ae:
		return "Aer"
	case E, Ey:
		return "Eir"
	case He:
		return "His"
	case Per:
		return "Pers"
	case She:
		return "Her"
	case They:
		return "Their"
	case Ve:
		return "Vis"
	case Xe:
		return "Xyr"
	case Ze, Zie:
		return "Hir"
	}
	return ""
}

// PossesivePronoun returns the Possesive Pronoun for of the Subject as a string
func (p SubjectPronoun) PossesivePronoun() string {
	switch p {
	case Fae:
		return "Faers"
	case Ae:
		return "Aers"
	case E, Ey:
		return "Eirs"
	case He:
		return "His"
	case Per:
		return "Pers"
	case She:
		return "Hers"
	case They:
		return "Theirs"
	case Ve:
		return "Vis"
	case Xe:
		return "Xyrs"
	case Ze, Zie:
		return "Hirs"
	}
	return ""
}

// Reflexive returns the Reflexive for of the Subject as a string
func (p SubjectPronoun) Reflexive() string {
	switch p {
	case Fae:
		return "Faerself"
	case Ae:
		return "Aerself"
	case E, Ey:
		return "Eirself"
	case He:
		return "Himself"
	case Per:
		return "Perself"
	case She:
		return "Herself"
	case They:
		return "Themself"
	case Ve:
		return "Verself"
	case Xe:
		return "Xemself"
	case Ze, Zie:
		return "Hirself"
	}
	return ""
}
