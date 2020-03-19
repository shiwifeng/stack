## 栈工具库
### 接口
- Clear()   // 清空栈
- IsEmpty() bool // 是否为空
- Size() int    // 获取栈容量大小
- Push(e ...interface{}) error  //入栈
- Pop() (interface{}, error)  //出栈
- Top() (interface{}, error)  //获取栈顶元素

### 示例
```
$ go test -bench=. -benchmem
goos: windows
goarch: amd64
pkg: stack
BenchmarkStack_Push-4              96829             12190 ns/op              64 B/op          2 allocs/op
BenchmarkStack_Pop-4             2818506               450 ns/op              72 B/op          3 allocs/op
PASS
ok      stack   3.860s

```
```
package stack

import (
	"fmt"
	"math/rand"
	"testing"
	"time"
)

var dd1, dd2 *Stack

func TestStack_Push(t *testing.T) {
	dd1 = NewInit()
	dd1.Push(1, 2, 3)
	dd2 = NewInit()
	dd2.Push(4, 5, 6, 7)
	fmt.Println("dd1", dd1.Size())
	fmt.Println("dd2", dd2.Size())
}

func TestStack_Top(t *testing.T) {
	ret1, _ := dd1.Top()
	ret2, _ := dd2.Top()
	fmt.Println("dd1", ret1)
	fmt.Println("dd2", ret2)
}

func TestStack_Pop(t *testing.T) {
	for {
		ret, err := dd1.Pop()
		if err != nil {
			break
		}
		fmt.Println("dd1", ret)
	}
	for {
		ret, err := dd2.Pop()
		if err != nil {
			break
		}
		fmt.Println("dd2", ret)
	}
}

func BenchmarkStack_Push(b *testing.B) {
	s1 := []int{1, 2, 3}
	m1 := make(map[int]int, 3)
	m1[11] = 11
	m1[22] = 22
	m1[33] = 33
	data := []interface{}{
		s1,
		m1,
		99,
		"加油",
	}
	dd2 = NewInit()
	b.ResetTimer()
	for i := 0; i < b.N; i++ {
		rand.Seed(time.Now().UnixNano() + int64(i)) //随机种子
		dd2.Push(data[rand.Intn(4)])
	}
}

func BenchmarkStack_Pop(b *testing.B) {
	for i := 0; i < b.N; i++ {
		_, err := dd2.Pop()
		if err != nil {
			continue
		}
	}
}

```