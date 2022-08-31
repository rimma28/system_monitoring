package linux

import (
	"github.com/stretchr/testify/require"
	"io/ioutil"
	"testing"
)

func TestCalcCPUUsage(t *testing.T) {
	t.Run("simple positive test", func(t *testing.T) {
		data, err := CalcCPUUsage("")
		require.NoError(t, err)
		require.NotEmpty(t, data)
	})

	t.Run("negative test - broken rootFsPath", func(t *testing.T) {
		_, err := CalcCPUUsage("/path/to/nowhere")

		require.Error(t, err)
	})

	t.Run("negative test - broken LaFile", func(t *testing.T) {
		tmp := LaFile
		LaFile = "/path/to/nowhere"
		defer func() { LaFile = tmp }()

		_, err := CalcCPUUsage("")

		require.Error(t, err)
	})

	t.Run("negative test - broken CPUStatsFile", func(t *testing.T) {
		tmp := CPUStatsFile
		CPUStatsFile = "/path/to/nowhere"
		defer func() { CPUStatsFile = tmp }()

		_, err := CalcCPUUsage("")

		require.Error(t, err)
	})
}

func TestCalcDevStats(t *testing.T) {
	t.Run("simple positive test", func(t *testing.T) {
		data, err := CalcDevStats("")

		require.NoError(t, err)
		require.NotEmpty(t, data)
	})

	t.Run("negative test - broken rootFsPath", func(t *testing.T) {
		_, err := CalcDevStats("/path/to/nowhere")

		require.Error(t, err)
	})

	t.Run("negative test - broken BlockDevicesDir", func(t *testing.T) {
		tmp := BlockDevicesDir
		BlockDevicesDir = "/path/to/nowhere"
		defer func() { BlockDevicesDir = tmp }()

		_, err := CalcDevStats("")

		require.Error(t, err)
	})

	t.Run("negative test - broken DevStatsFilename", func(t *testing.T) {
		tmp := DevStatsFilename
		DevStatsFilename = "/path/to/nowhere"
		defer func() { DevStatsFilename = tmp }()

		_, err := CalcDevStats("")

		require.Error(t, err)
	})
}

func TestCalcFsUtilization(t *testing.T) {
	t.Run("simple positive test", func(t *testing.T) {
		data, err := CalcFsUtilization("")

		require.NoError(t, err)
		require.NotEmpty(t, data)
	})

	t.Run("negative test - broken rootFsPath", func(t *testing.T) {
		_, err := CalcFsUtilization("/path/to/nowhere")

		require.Error(t, err)
	})

	t.Run("negative test - broken MountinfoFile", func(t *testing.T) {
		tmp := MountinfoFile
		MountinfoFile = "/path/to/nowhere"
		defer func() { MountinfoFile = tmp }()

		_, err := CalcFsUtilization("")

		require.Error(t, err)
	})
}

func Test_parseDevStats(t *testing.T) {
	t.Run("simple positive test", func(t *testing.T) {
		devs, err := ioutil.ReadDir(BlockDevicesDir)
		require.NoError(t, err)
		for _, d := range devs {
			_, err := parseDevStats(d.Name(), BlockDevicesDir, "")
			require.NoError(t, err)
		}
	})

	t.Run("simple negative test", func(t *testing.T) {
		dir := "/path/to/nowhere"

		_, err := parseDevStats("not_exist_file", dir, "")
		require.Error(t, err)
	})
}
