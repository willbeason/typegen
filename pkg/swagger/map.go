// Copyright 2019 Google LLC
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

package swagger

import (
	"github.com/willbeason/typegen/pkg/definition"
	"github.com/willbeason/typegen/pkg/maps"
)



func (p parser) newMap(meta definition.Meta, o map[string]interface{}) definition.Map {
	additionalPropertiesMap := maps.GetRequiredMap("additionalProperties", o)
	// TODO(b/141928662): Handle map of nested type. Uncommon case.
	if isObject(additionalPropertiesMap) {
		return definition.Map{
			Values: definition.Empty{},
		}
	}
	return definition.Map{
		Values: p.parseType(meta, additionalPropertiesMap),
	}
}