//  Copyright (C) 晓白齐齐,版权所有.

package opt

import (
	"fmt"

	"github.com/bqqsrc/gson"
)

type Int32 int32

func (v *Int32) UnmarshalByKey(key string, w *gson.Lexer) error {
	return fmt.Errorf("a Int32-type can't has a key")
}

func (v *Int32) UnmarshalByIndex(index int, w *gson.Lexer) error {
	if index == 0 {
		tv, err := w.Int32()
		*v = Int32(tv)
		return err
	} else {
		// 容错模式：超过1个值时，后面的值舍弃
		return nil
		// 数量严格模式：只能有一个值，否则报错
		//return fmt.Errorf("too many field got for a int32-type, a int32-type can only has a field")
	}
}
