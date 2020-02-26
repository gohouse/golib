package rules

import (
	"errors"
	"github.com/gohouse/golib/gopass"
	"github.com/gohouse/golib/t"
	"strings"
)

// Max ...
func Max() gopass.ValidatorHandler {
	return func(v *gopass.Validator) {
		v.Register("max", func(data interface{}, rule ...string) error {
			if len(rule) == 0 {
				return errors.New("max规则格式有误")
			}
			if !strings.Contains(rule[0], ":") {
				return errors.New("max规则格式有误")
			}
			rules := strings.Split(rule[0], ":")
			// 获取data长度
			if len(t.New(data).String()) > t.New(rules[1]).Int() {
				return errors.New("参数长度超限")
			}
			return nil
		})
	}
}
