package jstream

import (
	"encoding/json"
	"io"
)

type Array[T any] struct {
	pos    position
	elems  []T
	commit func([]T)
}

func (a *Array[T]) Next(dec *json.Decoder) (err error) {
	switch a.pos {
	case posFirst:
		if err = requireToken(dec, arrayStart); err != nil {
			return err
		}
		a.pos = posDecoding
	case posDecoding:
		var elem T
		for dec.More() {
			err = dec.Decode(&elem)
			if err != nil {
				return err
			}
			a.elems = append(a.elems, elem)
		}
		a.pos = posLast
	case posLast:
		if err = requireToken(dec, arrayEnd); err != nil {
			return err
		}
		a.pos = posEOF
	case posEOF:
		a.commit(a.elems)
		return io.EOF
	}
	return nil
}
