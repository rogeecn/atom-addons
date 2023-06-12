package k8s

import (
	"os"
	"path/filepath"

	"github.com/pkg/errors"
	"github.com/rogeecn/atom/container"
	"github.com/rogeecn/atom/providers/log"
	"github.com/rogeecn/atom/utils/opt"
	"k8s.io/client-go/kubernetes"
	"k8s.io/client-go/tools/clientcmd"
)

func DefaultProvider() container.ProviderContainer {
	return container.ProviderContainer{
		Provider: Provide,
		Options:  []opt.Option{},
	}
}

func Provide(opts ...opt.Option) error {
	o := opt.New(opts...)
	return container.Container.Provide(func() (*kubernetes.Clientset, error) {
		config := filepath.Join(os.Getenv("HOME"), ".kube", "config")
		if _, err := os.Stat(config); err != nil {
			config = ""
		} else {
			log.Debugf("using kube config: %s", config)
		}

		clientConfig, err := clientcmd.BuildConfigFromFlags("", config)
		if err != nil {
			return nil, errors.Wrap(err, "failed to build config")
		}

		return kubernetes.NewForConfig(clientConfig)
	}, o.DiOptions()...)
}
