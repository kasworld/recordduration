// Copyright 2015,2016,2017,2018,2019 SeukWon Kang (kasworld@gmail.com)
// Licensed under the Apache License, Version 2.0 (the "License");
// you may not use this file except in compliance with the License.
// You may obtain a copy of the License at
//    http://www.apache.org/licenses/LICENSE-2.0
// Unless required by applicable law or agreed to in writing, software
// distributed under the License is distributed on an "AS IS" BASIS,
// WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
// See the License for the specific language governing permissions and
// limitations under the License.

package recordduration

import (
	"testing"
	"time"
)

func TestNew(t *testing.T) {
	rd1 := New("recordtest")
	defer t.Logf("run dur is %v", rd1)
	time.Sleep(time.Second)
}

func TestCompact(t *testing.T) {
	defer println("run dur is %v", New("recordtest"))
	time.Sleep(time.Second)
}
