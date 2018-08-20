package hellomock

import (
	"testing"

	"github.com/golang/mock/gomock"
	"github.com/zhengyscn/golang-doc/test/hellomock/mock_hellomock"
)

func TestCompany_Meeting(t *testing.T) {
	person := NewPerson("zhengyscn")
	company := NewCompany(person)
	t.Log(company.Meeting("monkey"))
}

func TestCompany_Meeting2(t *testing.T) {
	ctl := gomock.NewController(t)
	mock_talker := mock_hellomock.NewMockTalker(ctl)
	mock_talker.EXPECT().SayHello(gomock.Eq("monkey")).Return("this is zhengyscn.")
	company := NewCompany(mock_talker)
	t.Log(company.Meeting("monkey"))
}
