// Copyright 2014 beego Author. All Rights Reserved.
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

package toolbox_test

import (
	"fmt"
	"sync"
	"testing"
	"time"

	"github.com/airdb/sailor/toolbox"
)

func TestParse(t *testing.T) {
	tk := toolbox.NewTask("taska", "0/30 * * * * *", func() error { fmt.Println("hello world"); return nil })
	err := tk.Run()
	if err != nil {
		t.Fatal(err)
	}
	toolbox.AddTask("taska", tk)
	toolbox.StartTask()
	time.Sleep(6 * time.Second)
	toolbox.StopTask()
}

func TestSpec(t *testing.T) {
	wg := &sync.WaitGroup{}
	wg.Add(2)
	tk1 := toolbox.NewTask("tk1", "0 12 * * * *", func() error { fmt.Println("tk1"); return nil })
	tk2 := toolbox.NewTask("tk2", "0,10,20 * * * * *", func() error { fmt.Println("tk2"); wg.Done(); return nil })
	tk3 := toolbox.NewTask("tk3", "0 10 * * * *", func() error { fmt.Println("tk3"); wg.Done(); return nil })

	toolbox.AddTask("tk1", tk1)
	toolbox.AddTask("tk2", tk2)
	toolbox.AddTask("tk3", tk3)
	toolbox.StartTask()

	defer toolbox.StopTask()

	select {
	case <-time.After(200 * time.Second):
		t.FailNow()
	case <-wait(wg):
	}
}

func wait(wg *sync.WaitGroup) chan bool {
	ch := make(chan bool)

	go func() {
		wg.Wait()
		ch <- true
	}()

	return ch
}
