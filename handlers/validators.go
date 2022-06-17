package handlers

import (
	"fmt"
	"meisterwerk/entities"
	"time"
)

const layer = "2006-01-02 03:04"

func timeParamsChecker(f, t string) (err error) {
	var from, to time.Time
	if len(f) != 0 {
		from, err = time.Parse(layer, f)
		if err != nil {
			return fmt.Errorf("bad start value %s , %v", f, err)
		}
	}

	if len(t) != 0 {
		to, err = time.Parse(layer, t)
		if err != nil {
			return fmt.Errorf("bad end value %s , %v", t, err)
		}
	}

	if len(f) != 0 && len(t) != 0 && from.After(to) {
		return fmt.Errorf("to (%s) must be greater than from (%s) value", to, from)
	}

	return nil
}

func requiredFields(event entities.Event) error {
	switch {
	case len(event.Start) == 0 || len(event.End) == 0:
		return fmt.Errorf("empty time boundaries for the event is not allowed %s %s", event.Start, event.End)
	case !event.Status.Validate():
		return fmt.Errorf("bad status %s", event.Status)
	case len(event.Title) == 0:
		return fmt.Errorf("empty title %s", event.Title)
	case len(event.Address) == 0:
		return fmt.Errorf("empty address %s", event.Address)
	}

	return timeParamsChecker(event.Start, event.End)
}
