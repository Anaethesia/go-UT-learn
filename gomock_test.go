package go_UT_learn

import (
	"fmt"
	"github.com/golang/mock/gomock"
	"testing"
	mock_go_UT_learn "ut/mock"
)

// gomock 多用于接口打桩，结合mockgen工具自动生成打桩代码，方便调用

type Repo interface {
	GetName(id int) string
}

//type JD struct {
//	id   int
//	name string
//}
//
//func (p *JD) GetName(id int) string {
//	if id == p.id {
//		return p.name
//	}
//	return ""
//}

func TestGetName(t *testing.T) {
	// new mockController
	ctrl := gomock.NewController(t)
	defer ctrl.Finish()

	// mock Repo interface
	mock := mock_go_UT_learn.NewMockRepo(ctrl)
	mock.EXPECT().GetName(gomock.Eq(1)).Return("JD_1")
	if val := mock.GetName(1); val != "JD_1" {
		t.Fatal("expected JD_1, but got", val)
	} else {
		fmt.Println("the result is ", val)
	}

}