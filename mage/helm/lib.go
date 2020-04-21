// Copyright DataStax, Inc.
// Please see the included license file for details.

package helm_util

import (
	"fmt"

	shutil "github.com/datastax/cass-operator/mage/sh"
)

func Install(chartPath string, releaseName string, namespace string, overrides map[string]string) error {
	args := []string{
		"install",
		fmt.Sprintf("--namespace=%s", namespace),
	}

	if overrides != nil && len(overrides) > 0 {
		var overrideString = ""
		for key, val := range overrides {
			if overrideString == "" {
				overrideString = fmt.Sprintf("%s=%s", key, val)
			} else {
				overrideString = fmt.Sprintf("%s,%s=%s", overrideString, key, val)
			}
		}

		args = append(args, "--set", overrideString)
	}

	args = append(args, releaseName, chartPath)
	return shutil.RunV("helm", args...)
}

func uninstallArgs(releaseName string, namespace string) []string {
	return []string{
		"uninstall",
		fmt.Sprintf("--namespace=%s", namespace),
		releaseName,
	}
}

func Uninstall(releaseName string, namespace string) error {
	return shutil.RunV("helm", uninstallArgs(releaseName, namespace)...)
}

func UninstallCapture(releaseName string, namespace string) (string, string, error) {
	return shutil.RunVCapture("helm", uninstallArgs(releaseName, namespace)...)
}
