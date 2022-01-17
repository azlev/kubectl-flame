package utils

import (
	"fmt"
)

func GetTargetFileSystemLocation(containerId string) (string, error) {
	return fmt.Sprintf("/run/containerd/io.containerd.runtime.v1.linux/k8s.io/%s/rootfs", containerId), nil
}
