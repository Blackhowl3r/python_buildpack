package integration_test

import (
	. "github.com/onsi/ginkgo"
	. "github.com/onsi/gomega"

	"github.com/cloudfoundry/libbuildpack/cutlass"
	cutlass7 "github.com/cloudfoundry/libbuildpack/cutlass/v7"
)

var _ = Describe("pushing an app a second time", func() {
	var app *cutlass7.App

	BeforeEach(func() {
		if cutlass.Cached {
			Skip("but running cached tests")
		}

		if isSerialTest {
			Skip("Skipping parallel tests")
		}

		app = cutlass7.New(Fixtures("no_deps"))
		app.Buildpacks = []string{"python_buildpack"}
	})

	AfterEach(func() {
		if app != nil {
			app.Destroy()
		}
		app = nil
	})

	Regexp := `\[.*/python\_[\d\.]+\_linux\_x64\_(cflinuxfs.*_)?[\da-f]+\.tgz\]`
	DownloadRegexp := "Download " + Regexp
	CopyRegexp := "Copy " + Regexp

	It("uses the cache for manifest dependencies", func() {
		PushAppAndConfirm(app)
		Expect(app.Stdout.String()).To(MatchRegexp(DownloadRegexp))
		Expect(app.Stdout.String()).ToNot(MatchRegexp(CopyRegexp))

		app.Stdout.Reset()
		PushAppAndConfirm(app)
		Expect(app.Stdout.String()).To(MatchRegexp(CopyRegexp))
		Expect(app.Stdout.String()).ToNot(MatchRegexp(DownloadRegexp))
	})
})
