package Domain

import "fmt"

type Group struct {
	id   int
	name string
}

func newGroup(id int, name string) (*Group, error) {
	if len(name) > 250 {
		return nil, fmt.Errorf("Name exceeds the maximum available length of 250 chars")
	}

	return &Group{
		id:   id,
		name: name,
	}, nil
}
