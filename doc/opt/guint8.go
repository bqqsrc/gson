//  Copyright (C) 晓白齐齐,版权所有.

package opt

import (
	"fmt"

	"github.com/bqqsrc/gson"
)

type Uint8 uint8

func (v *Uint8) UnmarshalByKey(key string, w *gson.Lexer) error {
	return fmt.Errorf("a Uint8-type can't has a key")
}

func (v *Uint8) UnmarshalByIndex(index int, w *gson.Lexer) error {
	if index == 0 {
		tv, err := w.Uint8()
		*v = Uint8(tv)
		return err
	} else {
		// 容错模式：超过1个值时，后面的值舍弃
		return nil
		// 数量严格模式：只能有一个值，否则报错
		//return fmt.Errorf("too many field got for a uint8-type, a uint8-type can only has a field")
	}
}
