package version_test

import (
	"fmt"
	"testing"

	"github.com/lewinloo/restful-api-demo/version"
)

func TestVersion(t *testing.T) {
	fmt.Println(version.FullVersion())
}
