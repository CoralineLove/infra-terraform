package helpers

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"log"
	"os"
	"path/filepath"
	"sort"
	"strconv"
	"strings"

	"github.com/Masterminds/semver"
)

type Version struct {
	Major int
	Minor int
	Patch int
}

func GetLatestVersion(version string) (*Version, error) {
	parsed, err := semver.Parse(version)
	if err != nil {
		return nil, err
	}
	return &Version{Major: parsed.Major, Minor: parsed.Minor, Patch: parsed.Patch}, nil
}

func GetLatestTerraformVersion() (*Version, error) {
	data, err := ioutil.ReadFile(filepath.Join(os.Getenv("GOPATH"), "/src/github.com/hashicorp/terraform/VERSION"))
	if err != nil {
		return nil, err
	}
	var terraformVersion Version
	err = json.Unmarshal(data, &terraformVersion)
	if err != nil {
		return nil, err
	}
	return &terraformVersion, nil
}

func GetLatestProviderVersion(provider string) (*Version, error) {
	data, err := ioutil.ReadFile(filepath.Join(os.Getenv("GOPATH"), "/src/",
		strings.ReplaceAll(provider, "-", "/"), "/VERSION"))
	if err != nil {
		return nil, err
	}
	var providerVersion Version
	err = json.Unmarshal(data, &providerVersion)
	if err != nil {
		return nil, err
	}
	return &providerVersion, nil
}

func CompareVersions(v1 *Version, v2 *Version) int {
	if v1.Major < v2.Major {
		return -1
	}
	if v1.Major > v2.Major {
		return 1
	}
	if v1.Minor < v2.Minor {
		return -1
	}
	if v1.Minor > v2.Minor {
		return 1
	}
	if v1.Patch < v2.Patch {
		return -1
	}
	if v1.Patch > v2.Patch {
		return 1
	}
	return 0
}

func GetVersionRange(sourceVersion string, targetVersion string) string {
	v1, err := GetLatestVersion(sourceVersion)
	if err != nil {
		log.Fatal(err)
	}
	v2, err := GetLatestVersion(targetVersion)
	if err != nil {
		log.Fatal(err)
	}
	if CompareVersions(v1, v2) < 0 {
		return fmt.Sprintf(">= %s <= %s", targetVersion, sourceVersion)
	}
	if CompareVersions(v1, v2) > 0 {
		return fmt.Sprintf("< %s >= %s", sourceVersion, targetVersion)
	}
	return fmt.Sprintf("== %s", sourceVersion)
}

func GenerateVersionRange(provider string, terraformVersion string, sourceVersion string) string {
	v1, err := GetLatestVersion(sourceVersion)
	if err != nil {
		log.Fatal(err)
	}
	v2, err := GetLatestTerraformVersion()
	if err != nil {
		log.Fatal(err)
	}
	v3, err := GetLatestProviderVersion(provider)
	if err != nil {
		log.Fatal(err)
	}
	sort.Slice(v2, func(i, j int) bool {
		return CompareVersions(&v2[i], &v2[j]) < 0
	})
	sort.Slice(v3, func(i, j int) bool {
		return CompareVersions(&v3[i], &v3[j]) < 0
	})
	var versionRange string
	if CompareVersions(v1, v2) < 0 {
		versionRange = fmt.Sprintf(">= %s <= %s", v2[len(v2)-1].String(), v3[len(v3)-1].String())
	} else if CompareVersions(v1, v3) < 0 {
		versionRange = fmt.Sprintf(">= %s <= %s", v3[len(v3)-1].String(), v2[len(v2)-1].String())
	} else {
		versionRange = fmt.Sprintf(">= %s =< %s", v2[len(v2)-1].String(), v3[len(v3)-1].String())
	}
	return versionRange
}

func GetMajorMinorVersion(version string) string {
	v, err := GetLatestVersion(version)
	if err != nil {
		log.Fatal(err)
	}
	return fmt.Sprintf("%d.%d", v.Major, v.Minor)
}

func GetMajorMinorPatchVersion(version string) string {
	v, err := GetLatestVersion(version)
	if err != nil {
		log.Fatal(err)
	}
	return fmt.Sprintf("%d.%d.%d", v.Major, v.Minor, v.Patch)
}