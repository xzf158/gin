// Copyright 2014 Manu Martinez-Almeida.  All rights reserved.
// Use of this source code is governed by a MIT style
// license that can be found in the LICENSE file.

package gin

import (
	"bytes"
	"fmt"
)

const (
	ErrorTypePrivate = 1 << iota
	ErrorTypePublic  = 1 << iota
)

const (
	ErrorMaskAny = 0xffffffff
)

// Used internally to collect errors that occurred during an http request.
type errorMsg struct {
	Error error       `json:"error"`
	Type  int         `json:"-"`
	Meta  interface{} `json:"meta"`
}

type errorMsgs []errorMsg

func (a errorMsgs) ByType(typ int) errorMsgs {
	if len(a) == 0 {
		return a
	}
	result := make(errorMsgs, 0, len(a))
	for _, msg := range a {
		if msg.Type&typ > 0 {
			result = append(result, msg)
		}
	}
	return result
}

func (a errorMsgs) Errors() []string {
	if len(a) == 0 {
		return []string{}
	}
	errors := make([]string, len(a))
	for i, err := range a {
		errors[i] = err.Error.Error()
	}
	return errors
}

func (a errorMsgs) String() string {
	if len(a) == 0 {
		return ""
	}
	var buffer bytes.Buffer
	for i, msg := range a {
		fmt.Fprintf(&buffer, "Error #%02d: %s\n     Meta: %v\n", (i + 1), msg.Error, msg.Meta)
	}
	return buffer.String()
}
