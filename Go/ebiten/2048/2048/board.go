package twenty48

import (
	"errors"
	"github.com/hajimehoshi/ebiten/v2"
)

var taskTerminated = errors.New("twenty48: task terminated")

type task func() error

type Board struct {
	size  int
	tiles map[*Tile]struct{}
	tasks []task
}

// NewBoard generates a new Board with giving size.
func NewBoard(size int) (*Board, error) {
	b := &Board{
		size:  size,
		tiles: make(map[*Tile]struct{}),
	}
	for i := 0; i < 2; i++ {
		if err := addRandomTile(b.tiles, b.size); err != nil {
			return nil, err
		}
	}
	return b, nil
}

func (b *Board) tileAt(x, y int) *Tile {
	return tileAt(b.tiles, x, y)
}

// Update updates the board state
func (b *Board) Update(input *Input) error {
	for t := range b.tiles {
		if err := t.Update(); err != nil {
			return err
		}
	}
	if 0 < len(b.tasks) {
		t := b.tasks[0]
		if err := t(); errors.Is(err, taskTerminated) {
			b.tasks = b.tasks[1:]
		} else if err != nil {
			return err
		}
		return nil
	}
	if dir, ok := input.Dir(); ok {
		if err := b.Move(dir); err != nil {
			return err
		}
	}
	return nil
}

// Move enqueues tile moving task
func (b *Board) Move(dir Dir) error {
	for t := range b.tiles {
		t.stopAnimation()
	}
	if !MoveTiles(b.tiles, b.size, dir) {
		return nil
	}
	b.tasks = append(b.tasks, func() error {
		for t := range b.tiles {
			if t.IsMoving() {
				return nil
			}
		}
		return taskTerminated
	})
	b.tasks = append(b.tasks, func() error {
		nextTiles := map[*Tile]struct{}{}
		for t := range b.tiles {
			if t.IsMoving() {
				panic("not reach")
			}
			if t.next.value != 0 {
				panic("not reach")
			}
			if t.current.value == 0 {
				continue
			}
			nextTiles[t] = struct{}{}
		}
		b.tiles = nextTiles
		if err := addRandomTile(b.tiles, b.size); err != nil {
			return err
		}
		return taskTerminated
	})
	return nil
}

// Size returns the board size.
func (b *Board) Size() (int, int) {
	x := b.size*tileSize + (b.size+1)*tileMargin
	y := x
	return x, y
}

// Draw draws the board to the given boardImage.
func (b *Board) Draw(boardImage *ebiten.Image) {
	boardImage.Fill(frameColor)
	for j := 0; j < b.size; j++ {
		for i := 0; i < b.size; i++ {
			v := 0
			op := &ebiten.DrawImageOptions{}
			x := i*tileSize + (i+1)*tileMargin
			y := j*tileSize + (j+1)*tileMargin
			op.GeoM.Translate(float64(x), float64(y))
			op.ColorScale.ScaleWithColor(tileBackgroundColor(v))
			boardImage.DrawImage(tileImage, op)
		}
	}
	animatingTiles := map[*Tile]struct{}{}
	nonAnimatingTiles := map[*Tile]struct{}{}
	for t := range b.tiles {
		if t.IsMoving() {
			animatingTiles[t] = struct{}{}
		} else {
			nonAnimatingTiles[t] = struct{}{}
		}
		for tt := range nonAnimatingTiles {
			tt.Draw(boardImage)
		}
		for tt := range animatingTiles {
			tt.Draw(boardImage)
		}
	}
}
