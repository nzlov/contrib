// Copyright 2019-present Facebook
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

package schema

import (
	"entgo.io/contrib/entgql/internal/todo/ent/schema"
	"entgo.io/ent"
	"entgo.io/ent/schema/field"
	"github.com/google/uuid"
)

// BillProduct holds the schema definition for the BillProduct entity.
type BillProduct struct {
	ent.Schema
}

// Mixin returns BillProduct mixed-in schema.
func (BillProduct) Mixin() []ent.Mixin {
	return []ent.Mixin{
		// Reuse the fields and edges from base example.
		schema.BillProduct{},
	}
}

// Fields returns BillProduct fields.
func (BillProduct) Fields() []ent.Field {
	return []ent.Field{
		field.UUID("id", uuid.UUID{}).
			Default(uuid.New),
	}
}
