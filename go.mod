module github.com/kevin-zhaoshuai/k3s-operator

go 1.12

require (
	github.com/alexellis/k3sup v0.0.0-20190831121911-113b8298565a
	github.com/go-logr/logr v0.1.0
	github.com/golang/glog v0.0.0-20160126235308-23def4e6c14b
	github.com/mitchellh/go-homedir v1.1.0 // indirect
	github.com/morikuni/aec v0.0.0-20170113033406-39771216ff4c // indirect
	github.com/onsi/ginkgo v1.8.0
	github.com/onsi/gomega v1.5.0
	k8s.io/apimachinery v0.0.0-20190404173353-6a84e37a896d
	k8s.io/client-go v11.0.1-0.20190409021438-1a26190bd76a+incompatible
	sigs.k8s.io/controller-runtime v0.2.0-beta.4
	sigs.k8s.io/controller-tools v0.2.0-beta.4 // indirect
)
