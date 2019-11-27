package bottle

import (
	"errors"
	"io"
	"io/ioutil"
)

type EmptyBottleError struct {
	Err error
}

func (e EmptyBottleError) Error() string {
	return e.Err.Error()
}

// Bottle is a stores data from readers
type Bottle struct {
	ID       string
	Contents []byte
}

func NewBottle(id string) *Bottle {
	return &Bottle{ID: id}
}

// Fill puts the contents from the io.Reader into the bottle
func (b *Bottle) Fill(r io.Reader) (err error) {
	b.Contents, err = ioutil.ReadAll(r)
	if err != nil {
		return err
	}
	return nil
}

// Pour copies the contents of the bottle into the Writer
func (b *Bottle) Pour(w io.Writer) (err error) {
	if b.Contents == nil {
		return EmptyBottleError{Err: errors.New("Can't pour from empty bottle")}
	}
	_, err = w.Write(b.Contents)
	if err != nil {
		return err
	}
	return nil
}

// Drain removes contents from the bottle into the Writer
func (b *Bottle) Drain(w io.Writer) (err error) {
	if b.Contents == nil {
		return EmptyBottleError{Err: errors.New("Can't drain empty bottle")}
	}
	_, err = w.Write(b.Contents)
	if err != nil {
		return err
	}
	b.Contents = nil
	return nil
}
