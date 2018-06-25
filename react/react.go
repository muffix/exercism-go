// Package react implements a basic reactive system
package react

type reactor struct {
	cells []*cell
}

type cell struct {
	value   int
	reactor *reactor

	c1, c2    Cell
	compute   func()
	callbacks map[callback]struct{}
}

type callback *func(int)

type canceller struct {
	cancelFn func()
}

// New returns a new Reactor that controls cells
func New() Reactor {
	return &reactor{}
}

// compute updates all cells when something has changed.
func (r *reactor) compute() {
	for _, cell := range r.cells {
		oldValue := cell.Value()
		cell.compute()
		if oldValue != cell.Value() {
			for callback := range cell.callbacks {
				(*callback)(cell.Value())
			}
		}
	}
}

func (r *reactor) CreateInput(value int) InputCell {
	newCell := cell{reactor: r, compute: func() {}}

	r.cells = append(r.cells, &newCell)
	newCell.SetValue(value)

	return &newCell
}

func (r *reactor) CreateCompute1(c Cell, fn func(int) int) ComputeCell {
	newCell := cell{reactor: r, c1: c}

	newCell.compute = func() {
		newCell.value = fn((newCell.c1).Value())
	}

	newCell.compute()

	r.cells = append(r.cells, &newCell)

	return &newCell
}

func (r *reactor) CreateCompute2(c1, c2 Cell, fn func(int, int) int) ComputeCell {
	newCell := cell{reactor: r, c1: c1, c2: c2}

	newCell.compute = func() {
		newCell.value = fn((newCell.c1).Value(), (newCell.c2).Value())
	}

	newCell.compute()

	r.cells = append(r.cells, &newCell)

	return &newCell
}

func (c *cell) Value() int {
	return c.value
}

func (c *cell) SetValue(val int) {
	c.value = val
	c.reactor.compute()
}

func (c *cell) AddCallback(fn func(int)) Canceler {
	if c.callbacks == nil {
		c.callbacks = make(map[callback]struct{})
	}

	var nothing struct{}
	c.callbacks[&fn] = nothing

	return &canceller{func() { delete(c.callbacks, &fn) }}
}

func (c *canceller) Cancel() {
	c.cancelFn()
}
