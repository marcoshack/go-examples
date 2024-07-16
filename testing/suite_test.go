package testing

import (
	"strings"
	"testing"

	"github.com/stretchr/testify/require"
	"github.com/stretchr/testify/suite"
)

func TestExampleSuite(t *testing.T) {
	testSuite := &ExtendedExampleSuite{}
	testSuite.SkipContaining("Two", "Three")
	testSuite.Skip("TestExampleSuite/TestFour")
	suite.Run(t, testSuite)
}

type ExampleSuite struct {
	suite.Suite
	number         int
	text           string
	flag           bool
	ignoreContaing []string
	ignoreExact    []string
}

// SkipContaining adds the given substrings to the list of ignored tests.
// Test cases that matches any of these substrings will be skipped.
func (s *ExampleSuite) SkipContaining(substrings ...string) {
	if s.ignoreContaing == nil {
		s.ignoreContaing = make([]string, 0, len(substrings))
	}
	s.ignoreContaing = append(s.ignoreContaing, substrings...)
}

// Skip adds the given test case names to the list of ignored tests.
// Test cases wich names exact match the one added with Skip will be skipped.
func (s *ExampleSuite) Skip(testNames ...string) {
	if s.ignoreExact == nil {
		s.ignoreExact = make([]string, 0, len(testNames))
	}
	s.ignoreExact = append(s.ignoreExact, testNames...)
}

func (s *ExampleSuite) SetupSuite() {
	s.number = 10
	s.text = "foo"
	s.flag = true
}

func (s *ExampleSuite) SetupTest() {
	for _, ignored := range s.ignoreContaing {
		if strings.Contains(s.T().Name(), ignored) {
			s.T().Skip()
		}
	}
	for _, ignored := range s.ignoreExact {
		if s.T().Name() == ignored {
			s.T().Skip()
		}
	}
	if strings.Contains(s.T().Name(), "Special") {
		s.T().Log("This is a special test case")
	}
}

func (s *ExampleSuite) TestOne() {
	require.Equal(s.T(), s.text, "foo")
}

func (s *ExampleSuite) TestTwo() {
	require.True(s.T(), s.flag)
}

func (s *ExampleSuite) TestThreeSpecial() {
	require.GreaterOrEqual(s.T(), s.number, 1)
}

type ExtendedExampleSuite struct {
	ExampleSuite
}

func (s *ExtendedExampleSuite) TestFour() {
	require.True(s.T(), s.flag)
}

//
// Test multiple suites running in a single Test method.
// This can be used to share and compose test suites.
//

// BaseSuite provides common attributes and behavior to all my test suites.
type BaseSuite struct {
	suite.Suite
	SomeInterestingAttribute string
}

func (bs *BaseSuite) SetupTest() {
	if bs.SomeInterestingAttribute == "" {
		bs.SomeInterestingAttribute = "foo"
	}
}

type Suite1 struct {
	BaseSuite
}

func (s *Suite1) TestSuite1_TestCase1() {
	s.Require().Equal(s.SomeInterestingAttribute, "foo")
}

type Suite2 struct {
	BaseSuite
}

func (s *Suite2) TestSuite2_TestCase1() {
	s.Require().Equal(s.SomeInterestingAttribute, "foo")
}

type Suite3 struct {
	BaseSuite
}

func (s *Suite3) TestSuite3_TestCase1() {
	s.Require().Equal(s.SomeInterestingAttribute, "foo")
}

var (
	suites = []suite.TestingSuite{new(Suite1), new(Suite2), new(Suite3)}
)

func TestSuites(t *testing.T) {
	for _, s := range suites {
		suite.Run(t, s)
	}
}

func TestSuite1(t *testing.T) {
	suite.Run(t, new(Suite1))
}
