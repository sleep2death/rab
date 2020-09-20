package rab

// Actor of the game
//
type Actor struct {
	id    int
	name  string
	class *Class
}

// ID of the actor
func (a *Actor) ID() int {
	return a.id
}

// Name of the actor
func (a *Actor) Name() string {
	return a.name
}

// Class of the actor
func (a *Actor) Class() *Class {
	return a.class
}

func (c *Class) Skills() []*Skill {
	return c.skills
}
