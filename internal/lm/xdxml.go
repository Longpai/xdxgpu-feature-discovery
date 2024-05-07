package lm

import (
	"fmt"
	"strconv"
	"strings"

	"github.com/NVIDIA/gpu-feature-discovery/internal/resource"
	spec "github.com/NVIDIA/k8s-device-plugin/api/config/v1"
)

func NewXDXMLLabeler(manager resource.Manager, config *spec.Config) (Labeler, error) {
	if err := manager.Init(); err != nil {
		return nil, fmt.Errorf("failed to initialize XDXML: %v", err)
	}
}