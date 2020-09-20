package rab

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

type Skill struct {
}
