package detector

import (
	"math"
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestIsConfiguration(t *testing.T) {
	configs := []string{
		"test.xml",
		"test.json",
		"test.toml",
		"test.yaml",
		"test.yml",
		"test.conf",
		"test.ini",
	}
	nonConfigs := []string{
		"test.c",
		"test.cpp",
		"test.h",
		"test.go",
		"test.js",
		"test.html",
		"test.rs",
	}
	for _, testConf := range configs {
		assert.True(t, IsConfiguration(testConf))
	}
	for _, nonTestConf := range nonConfigs {
		assert.False(t, IsConfiguration(nonTestConf))
	}
}

func TestIsDotFile(t *testing.T) {
	assert.True(t, IsDotFile(".test"))
	assert.False(t, IsDotFile("test"))
}

func TestIsDocumentation(t *testing.T) {
	docs := []string{
		"CHANGELOG.md",
		"CODE_OF_CONDUCT.md",
		"CONTRIBUTING.md",
		"MAINTAINERS.md",
		"PULL_REQUEST_TEMPLATE.md",
		"README.md",
		"ROADMAP.md",
		"SECURITY.md",
		"Documentation",
		"examples",
		"demos",
		"docs",
		"man",
	}
	nondocs := []string{
		"bin",
		"boot",
		"dev",
		"etc",
		"home",
		"lib",
		"lib64",
		"mnt",
		"opt",
		"proc",
		"root",
		"run",
		"sbin",
		"srv",
		"sys",
		"tmp",
		"usr",
		"var",
	}
	for _, doc := range docs {
		assert.True(t, IsDocumentation(doc))
	}
	for _, nondoc := range nondocs {
		assert.False(t, IsDocumentation(nondoc))
	}
}

func TestIsVendor(t *testing.T) {
	vendors := []string{
		"vendor",
		".vscode",
		".github",
		".gitignore",
		"Godeps",
		"gradlew",
	}
	for _, vendor := range vendors {
		assert.True(t, IsVendor(vendor))
	}
}

func TestGetLanguage(t *testing.T) {
	result := GetLanguage(".", -1)
	assert.Equal(t, result, "Go", "this should be a go project")
	result = GetLanguage(".", 1)
	assert.Equal(t, result, "Go", "this should be a go project")
	result = GetLanguage(".", math.MaxInt)
	assert.Equal(t, result, "Go", "this should be a go project")
	result = GetLanguage(".", math.MinInt)
	assert.Equal(t, result, "Go", "this should be a go project")
	result = GetLanguage("/dev/null", 1)
	assert.Equal(t, result, "None", "this should NOT be a go project")
	result = GetLanguage("/dev/null", math.MaxInt)
	assert.Equal(t, result, "None", "this should NOT be a go project")
	result = GetLanguage("/dev/null", math.MinInt)
	assert.Equal(t, result, "None", "this should NOT be a go project")
}
