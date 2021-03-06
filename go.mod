module github.com/cloudfoundry/python-buildpack

require (
	github.com/Dynatrace/libbuildpack-dynatrace v1.4.1
	github.com/blang/semver v3.5.1+incompatible
	github.com/cloudfoundry/libbuildpack v0.0.0-20211201073524-877368d188fd
	github.com/elazarl/goproxy v0.0.0-20191011121108-aa519ddbe484 // indirect
	github.com/elazarl/goproxy/ext v0.0.0-20191011121108-aa519ddbe484 // indirect
	github.com/fsnotify/fsnotify v1.5.1 // indirect
	github.com/golang/mock v1.6.0
	github.com/kr/text v0.2.0
	github.com/onsi/ginkgo v1.16.5
	github.com/onsi/gomega v1.16.0
	golang.org/x/crypto v0.0.0-20201002170205-7f63de1d35b0
	golang.org/x/net v0.0.0-20210805182204-aaa1db679c0d // indirect
	golang.org/x/sys v0.0.0-20211124211545-fe61309f8881 // indirect
	gopkg.in/ini.v1 v1.62.0
)

go 1.13

replace github.com/cloudfoundry/libbuildpack => ../libbuildpack