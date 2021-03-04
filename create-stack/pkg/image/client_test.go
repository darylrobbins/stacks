package image_test

import (
	"fmt"
	"github.com/google/go-containerregistry/pkg/name"
	"github.com/google/go-containerregistry/pkg/v1/daemon"
	"github.com/google/go-containerregistry/pkg/v1/random"
	"github.com/paketo-buildpacks/stacks/create-stack/pkg/image"
	"github.com/sclevine/spec"
	"github.com/sclevine/spec/report"
	assertpkg "github.com/stretchr/testify/assert"
	requirepkg "github.com/stretchr/testify/require"
	"io/ioutil"
	"os"
	"testing"
)

func TestImageClient(t *testing.T) {
	spec.Run(t, "ImageClient", testImageClient, spec.Report(report.Terminal{}))
}

func testImageClient(t *testing.T, when spec.G, it spec.S) {
	var (
		assert      = assertpkg.New(t)
		require     = requirepkg.New(t)
		imageClient image.Client
		tag         = "stack-test-image"
	)

	it.Before(func() {
		imageClient = image.Client{}

		localTag, err := name.NewTag(tag)
		require.NoError(err)

		image, err := random.Image(1, 1)
		require.NoError(err)

		_, err = daemon.Write(localTag, image)
		require.NoError(err)
	})

	it("can set labels", func() {
		err := imageClient.SetLabel(tag, "some-key", "some-value")
		require.NoError(err)

		labels := getLabels(tag, t)

		assert.Equal("some-value", labels["some-key"])
	})

	it("can build images", func() {

		dir, err := ioutil.TempDir("", "dockerfile-test")

		file, err := os.Create(fmt.Sprintf("%s/%s", dir, "Dockerfile"))
		require.NoError(err)

		file.WriteString(`FROM alpine
ARG test_build_arg
LABEL testing.key=some-value
LABEL testing.build.arg.key=$test_build_arg`)

		defer os.RemoveAll(dir)

		err = file.Close()
		require.NoError(err)

		err = imageClient.Build(tag, dir, "test_build_arg=1")
		require.NoError(err)

		labels := getLabels(tag, t)
		assert.Equal("some-value", labels["testing.key"])
		assert.Equal("1", labels["testing.build.arg.key"])
	})

}

func getLabels(tag string, t *testing.T) map[string]string {
	ref, err := name.ParseReference(tag)
	requirepkg.NoError(t, err)

	image, err := daemon.Image(ref)
	requirepkg.NoError(t, err)

	configFile, err := image.ConfigFile()
	requirepkg.NoError(t, err)

	return configFile.Config.Labels
}
