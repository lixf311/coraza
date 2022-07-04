// Copyright 2022 Juan Pablo Tosso
//
// Licensed under the Apache License, Version 2.0 (the "License");
// you may not use this file except in compliance with the License.
// You may obtain a copy of the License at
//
//      http://www.apache.org/licenses/LICENSE-2.0
//
// Unless required by applicable law or agreed to in writing, software
// distributed under the License is distributed on an "AS IS" BASIS,
// WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
// See the License for the specific language governing permissions and
// limitations under the License.

package operators

import (
	"unicode/utf8"

	"github.com/corazawaf/coraza/v3"
)

type validateUtf8Encoding struct{}

func (o *validateUtf8Encoding) Init(options coraza.RuleOperatorOptions) error {
	return nil
}

func (o *validateUtf8Encoding) Evaluate(tx *coraza.Transaction, value string) bool {
	return utf8.ValidString(value)
}
