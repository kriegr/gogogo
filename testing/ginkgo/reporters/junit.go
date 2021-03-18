package reporters

import (
	"fmt"
	"strings"

	"github.com/onsi/ginkgo"
	"github.com/onsi/ginkgo/reporters"
)

// RunSpecsWithJunit runs specs with junit reporter enabled. It creates a *_junit.xml.
func RunSpecsWithJUnit(packageName string, t ginkgo.GinkgoTestingT) {
	junitReporter := reporters.NewJUnitReporter(fmt.Sprintf("%s_junit.xml", strings.ToLower(packageName)))
	ginkgo.RunSpecsWithDefaultAndCustomReporters(t, fmt.Sprintf("%s Suite", packageName), []ginkgo.Reporter{junitReporter})
}
