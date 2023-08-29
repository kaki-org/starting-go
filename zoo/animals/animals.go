package animals

type T struct {
	Field1 int // 公開フィールド
	field2 int // 非公開フィールド
}

/* 公開メソッド */
func (t *T) Method1() int {
	return t.Field1
}

/* 非公開メソッド */
func (t *T) method2() int {
	return t.field2
}

func FooFunc(n int) int {
	return internalFunc(n)
}

func internalFunc(n int) int {
	return n + 1
}

type I interface {
	Method1() string
	method2() string
}

type T3 struct{}

func (t *T3) Method1() string {
	return "Method1()"
}

func (t *T3) method2() string {
	return "method2()"
}

func NewI() I {
	return &T3{}
}
