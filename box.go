package golang_united_school_homework

import (
	"fmt"
	"reflect"
)

// box contains list of shapes and able to perform operations on them
type box struct {
	shapes         []Shape
	shapesCapacity int // Maximum quantity of shapes that can be inside the box.
}

// NewBox creates new instance of box
func NewBox(shapesCapacity int) *box {
	return &box{
		shapesCapacity: shapesCapacity,
	}
}

// AddShape adds shape to the box
// returns the error in case it goes out of the shapesCapacity range.
func (b *box) AddShape(shape Shape) error {
	err := fmt.Errorf("exceeds shapesCapacity %v", b.shapesCapacity)
	if len(b.shapes) >= b.shapesCapacity {
		return err
	}
	b.shapes = append(b.shapes, shape)
	return nil
}

// GetByIndex allows getting shape by index
// whether shape by index doesn't exist or index went out of the range, then it returns an error
func (b *box) GetByIndex(i int) (Shape, error) {
	err := fmt.Errorf("exceeds shapesCapacity %v", b.shapesCapacity)
	if i >= b.shapesCapacity {
		return nil, err
	}
	return b.shapes[i], nil

}

// ExtractByIndex allows getting shape by index and removes this shape from the list.
// whether shape by index doesn't exist or index went out of the range, then it returns an error
func (b *box) ExtractByIndex(i int) (Shape, error) {
	err := fmt.Errorf("exceeds shapesCapacity %v", b.shapesCapacity)
	if i >= b.shapesCapacity {
		return nil, err
	}

	res := []Shape{}
	dropeedItem := b.shapes[i]

	for numberItem, x := range b.shapes {
		if numberItem != i {
			res = append(res, x)
		}
	}
	b.shapes = res
	return dropeedItem, nil

}

// ReplaceByIndex allows replacing shape by index and returns removed shape.
// whether shape by index doesn't exist or index went out of the range, then it returns an error
func (b *box) ReplaceByIndex(i int, shape Shape) (Shape, error) {
	err := fmt.Errorf("exceeds shapesCapacity %v", b.shapesCapacity)
	if i >= b.shapesCapacity {
		return nil, err
	}

	res := []Shape{}
	dropeedItem := b.shapes[i]

	for numberItem, x := range b.shapes {
		if numberItem == i {
			res = append(res, shape)
			continue
		}
		res = append(res, x)
	}
	b.shapes = res
	return dropeedItem, nil

}

// SumPerimeter provides sum perimeter of all shapes in the list.
func (b *box) SumPerimeter() float64 {
	var res float64

	for _, x := range b.shapes {

		res = res + x.CalcPerimeter()
	}
	return res
}

// SumArea provides sum area of all shapes in the list.
func (b *box) SumArea() float64 {
	var res float64

	for _, x := range b.shapes {
		res = res + x.CalcArea()
	}
	return res

}

// RemoveAllCircles removes all circles in the list
// whether circles are not exist in the list, then returns an error
func (b *box) RemoveAllCircles() error {
	err := fmt.Errorf("no circle :(")
	circle := Circle{}

	res := []Shape{}
	dropeedItems := []Shape{}

	for _, x := range b.shapes {
		if reflect.TypeOf(x) == reflect.TypeOf(circle) {
			dropeedItems = append(dropeedItems, x)
		}
		res = append(res, x)
	}
	if len(dropeedItems) == 0 {
		return err
	}
	b.shapes = res
	return nil

}
