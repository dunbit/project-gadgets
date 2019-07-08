// Copyright 2019 Authors of protobuf-gadgets
//
// Licensed under the Apache License, Version 2.0 (the "License");
// you may not use this file except in compliance with the License.
// You may obtain a copy of the License at
//
//     http://www.apache.org/licenses/LICENSE-2.0
//
// Unless required by applicable law or agreed to in writing, software
// distributed under the License is distributed on an "AS IS" BASIS,
// WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
// See the License for the specific language governing permissions and
// limitations under the License.

package comment

import "strings"

// Add prepends the linedata with comments
func Add(comment string, data string) string {
	var str strings.Builder
	str.WriteString(comment)
	if data != "" {
		str.WriteString(" ")
		str.WriteString(data)
	}
	return str.String()
}

// Strip removes the comments from the linedata
func Strip(comment string, data string) string {
	data = strings.TrimPrefix(data, comment)
	data = strings.TrimPrefix(data, " ")
	return data
}
