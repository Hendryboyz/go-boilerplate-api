package helpers_test

import (
	"go-boilerplate-api/internal/helpers"
	"testing"

	"github.com/ozontech/allure-go/pkg/framework/provider"
	"github.com/ozontech/allure-go/pkg/framework/suite"
	"github.com/stretchr/testify/assert"
)

type HelpersTestSuite struct {
	suite.Suite
}

func (s *HelpersTestSuite) TestGetRatio(t provider.T) {
	t.Run("GivenTheSamePortionAndTotal_WhenGetRatio_ThenReturnOne", func(t provider.T) {
		// Arrange
		portion := 10.0
		total := 10.0

		// Act
		ratio := helpers.GetRatio(portion, total, 2)

		// Assert
		assert.Equal(t, ratio, 1.0)
	})

	t.Run("GivenNearTheSamePortionAndTotal_WhenGetRatio_ThenReturnPointNineNineNine", func(t provider.T) {
		// Arrange
		portion := 117985.0
		total := 118030.0

		// Act
		ratio := helpers.GetRatio(portion, total, 3)

		// Assert
		assert.Equal(t, ratio, 0.999)
	})

	t.Run("GivenZeroTotals_WhenGetRatio_ThenReturnZero", func(t provider.T) {
		// Arrange
		portion := 10.0
		total := 0.0

		// Act
		ratio := helpers.GetRatio(portion, total, 2)

		// Assert
		assert.Equal(t, ratio, 0.0)
	})

	t.Run("GivenNoPrecision_WhenGetRatio_ThenReturnRatio", func(t provider.T) {
		// Arrange
		portion := 7.8
		total := 10.0

		// Act
		ratio := helpers.GetRatio(portion, total, 0)

		// Assert
		assert.Equal(t, ratio, 0.78)
	})
}

func TestHelpersTestSuite(t *testing.T) {
	suite.RunSuite(t, new(HelpersTestSuite))
}
