package set

func isNil(vals ...interface{}) bool {
	for _, v := range vals {
		if v == nil {
			return true
		}
	}
	return false
}

// IsEmpty checks to see if the set is equal to the empty set
func IsEmpty(s Set) bool {
	return s == nil || (s != nil && s.Size() == 0)
}

// Disjointed compares both set against each other and
// returns true if neither set is contained within each other.
func Disjointed(a, b Set) bool {
	return !Contains(a, b)
}

// Contains will validate if a is within b.
func Contains(a, b Set) bool {
	// Same set being compared
	if Equal(a, b) {
		return true
	}
	if IsEmpty(b) {
		return false
	}
	if IsEmpty(a) {
		return true
	}
	for _, item := range a.Items() {
		if !b.HasItem(item) {
			return false
		}
	}
	return true
}

// Equal will compare if both sets contain the same values and return the result of that
func Equal(a, b Set) bool {
	if IsEmpty(a) && IsEmpty(b) {
		return true
	}
	if isNil(a, b) || a.Size() != b.Size() {
		return false
	}
	for _, i := range b.Items() {
		if !a.HasItem(i) {
			return false
		}
	}
	return true
}

// Union will merge two sets together into one set
func Union(a, b Set) Set {
	emptyA, emptyB := IsEmpty(a), IsEmpty(b)
	switch {
	// The union of two empty sets is the empty set
	case emptyA && emptyB:
		return nil
	case emptyA && !emptyB:
		return New(b.Items()...)
	case !emptyA && emptyB:
		return New(a.Items()...)
	}
	return New(append(a.Items(), b.Items()...)...)
}

// Intersection returns all the values that are found
// in both a and b.
func Intersection(a, b Set) Set {
	if isNil(a, b) {
		return nil
	}
	s := New()
	for _, i := range a.Items() {
		if b.HasItem(i) {
			s.Insert(i)
		}
	}
	return s
}

// Subtract (Compliment) returns the set of values from b that are not
// contained within a
func Subtract(a, b Set) Set {
	if IsEmpty(a) {
		return nil
	}
	if IsEmpty(b) {
		return New(a)
	}
	s := New()
	for _, i := range b.Items() {
		if !a.HasItem(i) {
			s.Insert(i)
		}
	}
	return s
}
