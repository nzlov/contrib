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

package entgql_test

import (
	"testing"

	"github.com/facebook/ent/schema"
	"github.com/facebookincubator/ent-contrib/entgql"
	"github.com/stretchr/testify/require"
)

func TestAnnotation(t *testing.T) {
	require.Implements(t, (*schema.Annotation)(nil), entgql.Annotation{})

	annotation := entgql.OrderField("foo")
	require.Equal(t, "foo", annotation.OrderField)

	annotation = entgql.Bind()
	require.True(t, annotation.Bind)
	require.Empty(t, annotation.Mapping)

	names := []string{"foo", "bar", "baz"}
	annotation = entgql.MapsTo(names...)
	require.False(t, annotation.Bind)
	require.ElementsMatch(t, names, annotation.Mapping)
}