package testing

import (
	"fmt"
	"strings"

	"github.com/onsi/ginkgo"
	"github.com/onsi/ginkgo/reporters"
)

func RunSpecWithJUnit(packageName string, t ginkgo.GinkgoTestingT) {
	junitReporter := reporters.NewJUnitReporter(fmt.Sprintf("%s_junit.xml", strings.ToLower(packageName)))
	ginkgo.RunSpecsWithDefaultAndCustomReporters(t, fmt.Sprintf("%s Suite", packageName), []ginkgo.Reporter{junitReporter})
}
