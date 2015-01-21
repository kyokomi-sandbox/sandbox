package sample

import (
	"errors"
	"fmt"
	"io"

	"github.com/benbjohnson/megajson/scanner"
)

type MediaJSONDecoder struct {
	s scanner.Scanner
}

func NewMediaJSONDecoder(r io.Reader) *MediaJSONDecoder {
	return &MediaJSONDecoder{s: scanner.NewScanner(r)}
}

func NewMediaJSONScanDecoder(s scanner.Scanner) *MediaJSONDecoder {
	return &MediaJSONDecoder{s: s}
}

func (e *MediaJSONDecoder) Decode(ptr **Media) error {
	s := e.s
	if tok, tokval, err := s.Scan(); err != nil {
		return err
	} else if tok == scanner.TNULL {
		*ptr = nil
		return nil
	} else if tok != scanner.TLBRACE {
		return fmt.Errorf("Unexpected %s at %d: %s; expected '{'", scanner.TokenName(tok), s.Pos(), string(tokval))
	}

	// Create the object if it doesn't exist.
	if *ptr == nil {
		*ptr = &Media{}
	}
	v := *ptr

	// Loop over key/value pairs.
	index := 0
	for {
		// Read in key.
		var key string
		tok, tokval, err := s.Scan()
		if err != nil {
			return err
		} else if tok == scanner.TRBRACE {
			return nil
		} else if tok == scanner.TCOMMA {
			if index == 0 {
				return fmt.Errorf("Unexpected comma at %d", s.Pos())
			}
			if tok, tokval, err = s.Scan(); err != nil {
				return err
			}
		}

		if tok != scanner.TSTRING {
			return fmt.Errorf("Unexpected %s at %d: %s; expected '{' or string", scanner.TokenName(tok), s.Pos(), string(tokval))
		} else {
			key = string(tokval)
		}

		// Read in the colon.
		if tok, tokval, err := s.Scan(); err != nil {
			return err
		} else if tok != scanner.TCOLON {
			return fmt.Errorf("Unexpected %s at %d: %s; expected colon", scanner.TokenName(tok), s.Pos(), string(tokval))
		}

		switch key {

		case "display_url":
			v := &v.DisplayURL

			if err := s.ReadString(v); err != nil {
				return err
			}

		case "expanded_url":
			v := &v.ExpandedURL

			if err := s.ReadString(v); err != nil {
				return err
			}

		case "id":
			v := &v.ID

			if err := s.ReadInt(v); err != nil {
				return err
			}

		case "id_str":
			v := &v.IDStr

			if err := s.ReadString(v); err != nil {
				return err
			}

		case "media_url":
			v := &v.MediaURL

			if err := s.ReadString(v); err != nil {
				return err
			}

		case "media_url_https":
			v := &v.MediaURLHTTPS

			if err := s.ReadString(v); err != nil {
				return err
			}

		case "type":
			v := &v.Type

			if err := s.ReadString(v); err != nil {
				return err
			}

		case "url":
			v := &v.URL

			if err := s.ReadString(v); err != nil {
				return err
			}

		}

		index++
	}

	return nil
}

func (e *MediaJSONDecoder) DecodeArray(ptr *[]*Media) error {
	s := e.s
	if tok, _, err := s.Scan(); err != nil {
		return err
	} else if tok != scanner.TLBRACKET {
		return errors.New("Expected '['")
	}

	slice := make([]*Media, 0)

	// Loop over items.
	index := 0
	for {
		tok, tokval, err := s.Scan()
		if err != nil {
			return err
		} else if tok == scanner.TRBRACKET {
			*ptr = slice
			return nil
		} else if tok == scanner.TCOMMA {
			if index == 0 {
				return fmt.Errorf("Unexpected comma in array at %d", s.Pos())
			}
			if tok, tokval, err = s.Scan(); err != nil {
				return err
			}
		}
		s.Unscan(tok, tokval)

		item := &Media{}
		if err := e.Decode(&item); err != nil {
			return err
		}
		slice = append(slice, item)

		index++
	}
}
