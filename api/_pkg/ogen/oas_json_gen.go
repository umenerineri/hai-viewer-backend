// Code generated by ogen, DO NOT EDIT.

package ogen

import (
	"math/bits"
	"strconv"

	"github.com/go-faster/errors"
	"github.com/go-faster/jx"

	"github.com/ogen-go/ogen/validate"
)

// Encode implements json.Marshaler.
func (s *ErrResp) Encode(e *jx.Encoder) {
	e.ObjStart()
	s.encodeFields(e)
	e.ObjEnd()
}

// encodeFields encodes fields.
func (s *ErrResp) encodeFields(e *jx.Encoder) {
	{
		if s.Error.Set {
			e.FieldStart("error")
			s.Error.Encode(e)
		}
	}
}

var jsonFieldsNameOfErrResp = [1]string{
	0: "error",
}

// Decode decodes ErrResp from json.
func (s *ErrResp) Decode(d *jx.Decoder) error {
	if s == nil {
		return errors.New("invalid: unable to decode ErrResp to nil")
	}

	if err := d.ObjBytes(func(d *jx.Decoder, k []byte) error {
		switch string(k) {
		case "error":
			if err := func() error {
				s.Error.Reset()
				if err := s.Error.Decode(d); err != nil {
					return err
				}
				return nil
			}(); err != nil {
				return errors.Wrap(err, "decode field \"error\"")
			}
		default:
			return d.Skip()
		}
		return nil
	}); err != nil {
		return errors.Wrap(err, "decode ErrResp")
	}

	return nil
}

// MarshalJSON implements stdjson.Marshaler.
func (s *ErrResp) MarshalJSON() ([]byte, error) {
	e := jx.Encoder{}
	s.Encode(&e)
	return e.Bytes(), nil
}

// UnmarshalJSON implements stdjson.Unmarshaler.
func (s *ErrResp) UnmarshalJSON(data []byte) error {
	d := jx.DecodeBytes(data)
	return s.Decode(d)
}

// Encode encodes string as json.
func (o OptString) Encode(e *jx.Encoder) {
	if !o.Set {
		return
	}
	e.Str(string(o.Value))
}

// Decode decodes string from json.
func (o *OptString) Decode(d *jx.Decoder) error {
	if o == nil {
		return errors.New("invalid: unable to decode OptString to nil")
	}
	o.Set = true
	v, err := d.Str()
	if err != nil {
		return err
	}
	o.Value = string(v)
	return nil
}

// MarshalJSON implements stdjson.Marshaler.
func (s OptString) MarshalJSON() ([]byte, error) {
	e := jx.Encoder{}
	s.Encode(&e)
	return e.Bytes(), nil
}

// UnmarshalJSON implements stdjson.Unmarshaler.
func (s *OptString) UnmarshalJSON(data []byte) error {
	d := jx.DecodeBytes(data)
	return s.Decode(d)
}

// Encode implements json.Marshaler.
func (s *ViewGetBadRequest) Encode(e *jx.Encoder) {
	e.ObjStart()
	s.encodeFields(e)
	e.ObjEnd()
}

// encodeFields encodes fields.
func (s *ViewGetBadRequest) encodeFields(e *jx.Encoder) {
	{
		if s.Error.Set {
			e.FieldStart("error")
			s.Error.Encode(e)
		}
	}
}

var jsonFieldsNameOfViewGetBadRequest = [1]string{
	0: "error",
}

// Decode decodes ViewGetBadRequest from json.
func (s *ViewGetBadRequest) Decode(d *jx.Decoder) error {
	if s == nil {
		return errors.New("invalid: unable to decode ViewGetBadRequest to nil")
	}

	if err := d.ObjBytes(func(d *jx.Decoder, k []byte) error {
		switch string(k) {
		case "error":
			if err := func() error {
				s.Error.Reset()
				if err := s.Error.Decode(d); err != nil {
					return err
				}
				return nil
			}(); err != nil {
				return errors.Wrap(err, "decode field \"error\"")
			}
		default:
			return d.Skip()
		}
		return nil
	}); err != nil {
		return errors.Wrap(err, "decode ViewGetBadRequest")
	}

	return nil
}

// MarshalJSON implements stdjson.Marshaler.
func (s *ViewGetBadRequest) MarshalJSON() ([]byte, error) {
	e := jx.Encoder{}
	s.Encode(&e)
	return e.Bytes(), nil
}

// UnmarshalJSON implements stdjson.Unmarshaler.
func (s *ViewGetBadRequest) UnmarshalJSON(data []byte) error {
	d := jx.DecodeBytes(data)
	return s.Decode(d)
}

// Encode implements json.Marshaler.
func (s *ViewGetNotFound) Encode(e *jx.Encoder) {
	e.ObjStart()
	s.encodeFields(e)
	e.ObjEnd()
}

// encodeFields encodes fields.
func (s *ViewGetNotFound) encodeFields(e *jx.Encoder) {
	{
		if s.Error.Set {
			e.FieldStart("error")
			s.Error.Encode(e)
		}
	}
}

var jsonFieldsNameOfViewGetNotFound = [1]string{
	0: "error",
}

// Decode decodes ViewGetNotFound from json.
func (s *ViewGetNotFound) Decode(d *jx.Decoder) error {
	if s == nil {
		return errors.New("invalid: unable to decode ViewGetNotFound to nil")
	}

	if err := d.ObjBytes(func(d *jx.Decoder, k []byte) error {
		switch string(k) {
		case "error":
			if err := func() error {
				s.Error.Reset()
				if err := s.Error.Decode(d); err != nil {
					return err
				}
				return nil
			}(); err != nil {
				return errors.Wrap(err, "decode field \"error\"")
			}
		default:
			return d.Skip()
		}
		return nil
	}); err != nil {
		return errors.Wrap(err, "decode ViewGetNotFound")
	}

	return nil
}

// MarshalJSON implements stdjson.Marshaler.
func (s *ViewGetNotFound) MarshalJSON() ([]byte, error) {
	e := jx.Encoder{}
	s.Encode(&e)
	return e.Bytes(), nil
}

// UnmarshalJSON implements stdjson.Unmarshaler.
func (s *ViewGetNotFound) UnmarshalJSON(data []byte) error {
	d := jx.DecodeBytes(data)
	return s.Decode(d)
}

// Encode implements json.Marshaler.
func (s *ViewGetOK) Encode(e *jx.Encoder) {
	e.ObjStart()
	s.encodeFields(e)
	e.ObjEnd()
}

// encodeFields encodes fields.
func (s *ViewGetOK) encodeFields(e *jx.Encoder) {
	{
		e.FieldStart("result")
		e.ArrStart()
		for _, elem := range s.Result {
			elem.Encode(e)
		}
		e.ArrEnd()
	}
}

var jsonFieldsNameOfViewGetOK = [1]string{
	0: "result",
}

// Decode decodes ViewGetOK from json.
func (s *ViewGetOK) Decode(d *jx.Decoder) error {
	if s == nil {
		return errors.New("invalid: unable to decode ViewGetOK to nil")
	}
	var requiredBitSet [1]uint8

	if err := d.ObjBytes(func(d *jx.Decoder, k []byte) error {
		switch string(k) {
		case "result":
			requiredBitSet[0] |= 1 << 0
			if err := func() error {
				s.Result = make([]ViewGetOKResultItem, 0)
				if err := d.Arr(func(d *jx.Decoder) error {
					var elem ViewGetOKResultItem
					if err := elem.Decode(d); err != nil {
						return err
					}
					s.Result = append(s.Result, elem)
					return nil
				}); err != nil {
					return err
				}
				return nil
			}(); err != nil {
				return errors.Wrap(err, "decode field \"result\"")
			}
		default:
			return d.Skip()
		}
		return nil
	}); err != nil {
		return errors.Wrap(err, "decode ViewGetOK")
	}
	// Validate required fields.
	var failures []validate.FieldError
	for i, mask := range [1]uint8{
		0b00000001,
	} {
		if result := (requiredBitSet[i] & mask) ^ mask; result != 0 {
			// Mask only required fields and check equality to mask using XOR.
			//
			// If XOR result is not zero, result is not equal to expected, so some fields are missed.
			// Bits of fields which would be set are actually bits of missed fields.
			missed := bits.OnesCount8(result)
			for bitN := 0; bitN < missed; bitN++ {
				bitIdx := bits.TrailingZeros8(result)
				fieldIdx := i*8 + bitIdx
				var name string
				if fieldIdx < len(jsonFieldsNameOfViewGetOK) {
					name = jsonFieldsNameOfViewGetOK[fieldIdx]
				} else {
					name = strconv.Itoa(fieldIdx)
				}
				failures = append(failures, validate.FieldError{
					Name:  name,
					Error: validate.ErrFieldRequired,
				})
				// Reset bit.
				result &^= 1 << bitIdx
			}
		}
	}
	if len(failures) > 0 {
		return &validate.Error{Fields: failures}
	}

	return nil
}

// MarshalJSON implements stdjson.Marshaler.
func (s *ViewGetOK) MarshalJSON() ([]byte, error) {
	e := jx.Encoder{}
	s.Encode(&e)
	return e.Bytes(), nil
}

// UnmarshalJSON implements stdjson.Unmarshaler.
func (s *ViewGetOK) UnmarshalJSON(data []byte) error {
	d := jx.DecodeBytes(data)
	return s.Decode(d)
}

// Encode implements json.Marshaler.
func (s *ViewGetOKResultItem) Encode(e *jx.Encoder) {
	e.ObjStart()
	s.encodeFields(e)
	e.ObjEnd()
}

// encodeFields encodes fields.
func (s *ViewGetOKResultItem) encodeFields(e *jx.Encoder) {
	{
		e.FieldStart("position")
		s.Position.Encode(e)
	}
	{
		e.FieldStart("url")
		e.Str(s.URL)
	}
}

var jsonFieldsNameOfViewGetOKResultItem = [2]string{
	0: "position",
	1: "url",
}

// Decode decodes ViewGetOKResultItem from json.
func (s *ViewGetOKResultItem) Decode(d *jx.Decoder) error {
	if s == nil {
		return errors.New("invalid: unable to decode ViewGetOKResultItem to nil")
	}
	var requiredBitSet [1]uint8

	if err := d.ObjBytes(func(d *jx.Decoder, k []byte) error {
		switch string(k) {
		case "position":
			requiredBitSet[0] |= 1 << 0
			if err := func() error {
				if err := s.Position.Decode(d); err != nil {
					return err
				}
				return nil
			}(); err != nil {
				return errors.Wrap(err, "decode field \"position\"")
			}
		case "url":
			requiredBitSet[0] |= 1 << 1
			if err := func() error {
				v, err := d.Str()
				s.URL = string(v)
				if err != nil {
					return err
				}
				return nil
			}(); err != nil {
				return errors.Wrap(err, "decode field \"url\"")
			}
		default:
			return d.Skip()
		}
		return nil
	}); err != nil {
		return errors.Wrap(err, "decode ViewGetOKResultItem")
	}
	// Validate required fields.
	var failures []validate.FieldError
	for i, mask := range [1]uint8{
		0b00000011,
	} {
		if result := (requiredBitSet[i] & mask) ^ mask; result != 0 {
			// Mask only required fields and check equality to mask using XOR.
			//
			// If XOR result is not zero, result is not equal to expected, so some fields are missed.
			// Bits of fields which would be set are actually bits of missed fields.
			missed := bits.OnesCount8(result)
			for bitN := 0; bitN < missed; bitN++ {
				bitIdx := bits.TrailingZeros8(result)
				fieldIdx := i*8 + bitIdx
				var name string
				if fieldIdx < len(jsonFieldsNameOfViewGetOKResultItem) {
					name = jsonFieldsNameOfViewGetOKResultItem[fieldIdx]
				} else {
					name = strconv.Itoa(fieldIdx)
				}
				failures = append(failures, validate.FieldError{
					Name:  name,
					Error: validate.ErrFieldRequired,
				})
				// Reset bit.
				result &^= 1 << bitIdx
			}
		}
	}
	if len(failures) > 0 {
		return &validate.Error{Fields: failures}
	}

	return nil
}

// MarshalJSON implements stdjson.Marshaler.
func (s *ViewGetOKResultItem) MarshalJSON() ([]byte, error) {
	e := jx.Encoder{}
	s.Encode(&e)
	return e.Bytes(), nil
}

// UnmarshalJSON implements stdjson.Unmarshaler.
func (s *ViewGetOKResultItem) UnmarshalJSON(data []byte) error {
	d := jx.DecodeBytes(data)
	return s.Decode(d)
}

// Encode implements json.Marshaler.
func (s *ViewGetOKResultItemPosition) Encode(e *jx.Encoder) {
	e.ObjStart()
	s.encodeFields(e)
	e.ObjEnd()
}

// encodeFields encodes fields.
func (s *ViewGetOKResultItemPosition) encodeFields(e *jx.Encoder) {
	{
		e.FieldStart("x")
		e.Int(s.X)
	}
	{
		e.FieldStart("y")
		e.Int(s.Y)
	}
}

var jsonFieldsNameOfViewGetOKResultItemPosition = [2]string{
	0: "x",
	1: "y",
}

// Decode decodes ViewGetOKResultItemPosition from json.
func (s *ViewGetOKResultItemPosition) Decode(d *jx.Decoder) error {
	if s == nil {
		return errors.New("invalid: unable to decode ViewGetOKResultItemPosition to nil")
	}
	var requiredBitSet [1]uint8

	if err := d.ObjBytes(func(d *jx.Decoder, k []byte) error {
		switch string(k) {
		case "x":
			requiredBitSet[0] |= 1 << 0
			if err := func() error {
				v, err := d.Int()
				s.X = int(v)
				if err != nil {
					return err
				}
				return nil
			}(); err != nil {
				return errors.Wrap(err, "decode field \"x\"")
			}
		case "y":
			requiredBitSet[0] |= 1 << 1
			if err := func() error {
				v, err := d.Int()
				s.Y = int(v)
				if err != nil {
					return err
				}
				return nil
			}(); err != nil {
				return errors.Wrap(err, "decode field \"y\"")
			}
		default:
			return d.Skip()
		}
		return nil
	}); err != nil {
		return errors.Wrap(err, "decode ViewGetOKResultItemPosition")
	}
	// Validate required fields.
	var failures []validate.FieldError
	for i, mask := range [1]uint8{
		0b00000011,
	} {
		if result := (requiredBitSet[i] & mask) ^ mask; result != 0 {
			// Mask only required fields and check equality to mask using XOR.
			//
			// If XOR result is not zero, result is not equal to expected, so some fields are missed.
			// Bits of fields which would be set are actually bits of missed fields.
			missed := bits.OnesCount8(result)
			for bitN := 0; bitN < missed; bitN++ {
				bitIdx := bits.TrailingZeros8(result)
				fieldIdx := i*8 + bitIdx
				var name string
				if fieldIdx < len(jsonFieldsNameOfViewGetOKResultItemPosition) {
					name = jsonFieldsNameOfViewGetOKResultItemPosition[fieldIdx]
				} else {
					name = strconv.Itoa(fieldIdx)
				}
				failures = append(failures, validate.FieldError{
					Name:  name,
					Error: validate.ErrFieldRequired,
				})
				// Reset bit.
				result &^= 1 << bitIdx
			}
		}
	}
	if len(failures) > 0 {
		return &validate.Error{Fields: failures}
	}

	return nil
}

// MarshalJSON implements stdjson.Marshaler.
func (s *ViewGetOKResultItemPosition) MarshalJSON() ([]byte, error) {
	e := jx.Encoder{}
	s.Encode(&e)
	return e.Bytes(), nil
}

// UnmarshalJSON implements stdjson.Unmarshaler.
func (s *ViewGetOKResultItemPosition) UnmarshalJSON(data []byte) error {
	d := jx.DecodeBytes(data)
	return s.Decode(d)
}