package dex

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

// Class of the actor
// It's static and can't be change in game
//
type Class struct {
	id   int
	name string

	initLvl int
	maxLvl  int

	skills []*Skill
}

func (c *Class) ID() int {
	return c.id
}

func (c *Class) Name() string {
	return c.name
}

func (c *Class) Level() (initLvl int, maxLvl int) {
	return c.initLvl, c.maxLvl
}

func (c *Class) Skills() []*Skill {
	return c.skills
}

type Skill struct {
}
