package utils

import "os"

func IsRunningInContainer() bool {
	if _, err := os.Stat("/.dockerenv"); err != nil {
		return false
	}
	return true
}

func IsRunningInKubernetes() bool {
	return os.Getenv("KUBERNETES_SERVICE_HOST") != ""
}

func String(s *string) string {
	if s != nil {
		return *s
	}
	return ""
}
