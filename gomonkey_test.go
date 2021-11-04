package go_UT_learn

import (
	"reflect"
	"testing"

	"github.com/agiledragon/gomonkey"
	"github.com/smartystreets/goconvey/convey"
)

// Gomonkey: 十分常用的打桩工具,是在monkey框架基础上个人改进的一款UT框架，比较好奇公司喜欢用这类个人库而不是monkey这种偏官方的库写UT
// Gomonkey支持多种打桩特性，但在平常开发中多用来对全局变量打桩；对函数打桩；对方法打桩；对函数变量进行打桩

// 1、为函数打桩 ApplyFunc(origin func, stub func)
func CalVal(a, b int) int {
	return a + b
}

func TestCalVal(t *testing.T) {
	convey.Convey("happy path", t, func() {
		patch := gomonkey.ApplyFunc(CalVal, func(a, b int) int {
			return 3
		})
		defer patch.Reset()

		res := CalVal(1, 2)
		convey.So(res, convey.ShouldEqual, 3) // or use assert
	})
}

// 2、为全局变量打桩 ApplyGlobalVar（&var, stubvalue)

var temp = 10
func TestGlobal(t *testing.T) {
	convey.Convey("happy path", t, func() {
		patch := gomonkey.ApplyGlobalVar(&temp, 15)
		defer patch.Reset()
		convey.So(temp, convey.ShouldEqual, 15)
	})
}

// 3、对函数变量打桩 ApplyFuncVar (&func, stubfunc)   equal to ApplyFunc, but it is the address of func var

// 4、对方法打桩 ApplyMethod（reflect.TypeOf, methodName, stubfunc)

type User struct {
	money int
}

func(u *User) AddMoney(temp int) int{
	u.money += temp
	return u.money
}

func TestAddMoney(t *testing.T) {
	temp := User{}
	patch := gomonkey.ApplyMethod(reflect.TypeOf(&temp), "AddMoney", func(_ *User, temp int) int{
		return 100
	})
	defer patch.Reset()

	convey.Convey("happy path", t, func() {
		Bob := &User{
			money: 20,
		}
		convey.So(Bob.AddMoney(80),convey.ShouldEqual, 100)
	})

}

