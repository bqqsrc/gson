//  Copyright (C) 晓白齐齐,版权所有.

package opt

import (
	"fmt"

	"github.com/bqqsrc/gson"
)

type String string

func (v *String) UnmarshalByKey(key string, w *gson.Lexer) error {
	return fmt.Errorf("a String-type can't has a key")
}

func (v *String) UnmarshalByIndex(index int, w *gson.Lexer) error {
	if index == 0 {
		tv := w.String()
		*v = String(tv)
		return nil
	} else {
		// 容错模式：超过1个值时，后面的值舍弃
		return nil
		// 数量严格模式：只能有一个值，否则报错
		//return fmt.Errorf("too many field got for a string-type, a string-type can only has a field")
	}
}
