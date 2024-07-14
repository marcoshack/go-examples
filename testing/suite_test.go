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
