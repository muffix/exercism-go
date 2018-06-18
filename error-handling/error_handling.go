// Package erratum contains error handling
package erratum

// Use uses the given resource
func Use(o ResourceOpener, input string) (err error) {
	r, err := o()

	if err != nil {
		if _, ok := err.(TransientError); ok {
			return Use(o, input)
		}
		return
	}

	defer func() {
		if rec := recover(); rec != nil {
			switch panicError := rec.(type) {
			case FrobError:
				r.Defrob(FrobError(panicError).defrobTag)
				err = panicError
			case error:
				err = panicError
			}
		}
		r.Close()
	}()

	r.Frob(input)

	return
}
