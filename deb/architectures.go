/*
   Copyright 2013 Am Laher

   Licensed under the Apache License, Version 2.0 (the "License");
   you may not use this file except in compliance with the License.
   You may obtain a copy of the License at

       http://www.apache.org/licenses/LICENSE-2.0

   Unless required by applicable law or agreed to in writing, software
   distributed under the License is distributed on an "AS IS" BASIS,
   WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
   See the License for the specific language governing permissions and
   limitations under the License.
*/

package deb

import (
	"fmt"
)

//A processor architecture (ARM/x86/AMD64) - as named by Debian. At this stage: i386, armel, amd64 and 'all'. Note that 'any' is not valid for a binary package.
type Architecture string

const (
	Arch_i386  Architecture = "i386"
	Arch_armel Architecture = "armel" //TODO armhf
	Arch_amd64 Architecture = "amd64"
	Arch_all   Architecture = "all" //for binary package
)

func resolveArches(arches string) ([]Architecture, error) {
	if arches == "any" || arches == "" {
		return []Architecture{Arch_i386, Arch_armel, Arch_amd64}, nil
	} else if arches == string(Arch_i386) {
		return []Architecture{Arch_i386}, nil
	} else if arches == string(Arch_armel) {
		return []Architecture{Arch_armel}, nil
	} else if arches == string(Arch_amd64) {
	} else if arches == string(Arch_all) {
		return []Architecture{Arch_all}, nil
	}
	return nil, fmt.Errorf("Architecture %s not supported", arches)
}