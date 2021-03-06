package go_UT_learn

import (
	"github.com/smartystreets/goconvey/convey"
	"github.com/stretchr/testify/assert"
	"testing"
)

// TestFunc
func CheckVal(val string) bool {
	valList := []string{"AAA", "BBB"}
	for index := range valList {
		if valList[index] == val {
			return true
		}
	}
	return false
}


// Convey特点：断言函数丰富，提供UT的web界面，较为常用
// Convey website: https://github.com/smartystreets/goconvey/wiki/Documentation

// Testify: 断言库，多用它提供的丰富的断言功能
func TestCheckVal(t *testing.T) {
	convey.Convey("TestCheckVal happy path", t, func(){
		res := CheckVal("AAA")
		convey.So(res, convey.ShouldBeTrue)
		assert.True(t, res)
	})

	convey.Convey("TestCheckVal unhappy path", t, func(){
		res := CheckVal("CCC")
		convey.So(res, convey.ShouldBeFalse)
		assert.False(t, res)
	})
}
