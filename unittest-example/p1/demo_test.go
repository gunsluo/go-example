package p1

import (
	"fmt"
	"testing"

	"github.com/gunsluo/go-example/unittest-example/unittest"
	"github.com/stretchr/testify/suite"
)

type demoSuite struct {
	p unittest.PostgresSuite
	suite.Suite
}

// before all test
func (s *demoSuite) SetupSuite() {
	//fmt.Println("running before all test")
	s.p.Setup()
}

// all the tests in the suite have been run.
func (s *demoSuite) TearDownSuite() {
	//fmt.Println("running after all test")
	s.p.TearDown()
}

// before each test
func (s *demoSuite) SetupTest() {
	//fmt.Println("running before each test")
}

// after each test
func (s *demoSuite) TearDownTest() {
	//fmt.Println("running after each test")
}

func TestDemo(t *testing.T) {
	s := &demoSuite{}
	suite.Run(t, s)
}

func (s *demoSuite) TestCase1() {
	fmt.Println("--->test1")
}

func (s *demoSuite) TestCase2() {
	fmt.Println("--->test2")
}
