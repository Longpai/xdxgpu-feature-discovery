package resource

import (
	spec "github.com/NVIDIA/k8s-device-plugin/api/config/v1"
	"gitlab.com/nvidia/cloud-native/go-nvlib/pkg/nvlib/info"
	"k8s.io/klog/v2"
)

// NewManager is a factory method that creates a resource Manager based on the specified config.
func NewManager(config *spec.Config) Manager {
	return WithConfig(getManager(), config)
}

// WithConfig modifies a manager depending on the specified config.
// If failure on a call to init is allowed, the manager is wrapped to allow fallback to a Null manager.
func WithConfig(manager Manager, config *spec.Config) Manager {
	if *config.Flags.FailOnInitError {
		return manager
	}

	return NewFallbackToNullOnInitError(manager)
}

// getManager returns the resource manager depending on the system configuration.
func getManager() Manager {
	// logWithReason logs the output of the has* / is* checks from the info.Interface
	logWithReason := func(f func() (bool, string), tag string) bool {
		is, reason := f()
		if !is {
			tag = "non-" + tag
		}
		klog.Infof("Detected %v platform: %v", tag, reason)
		return is
	}

	infolib := info.New()

	hasNVML := logWithReason(infolib.HasNvml, "NVML")
	isTegra := logWithReason(infolib.IsTegraSystem, "Tegra")

	// The NVIDIA container stack does not yet support the use of integrated AND discrete GPUs on the same node.
	if hasNVML && isTegra {
		klog.Warning("Disabling Tegra-based resources on NVML system")
		isTegra = false
	}

	if hasNVML {
		klog.Info("Using NVML manager")
		return NewNVMLManager()
	} else if isTegra {
		klog.Info("Using CUDA manager")
		return NewCudaManager()
	}

	klog.Warning("No valid resources detected; using empty manager.")
	return NewNullManager()
}
