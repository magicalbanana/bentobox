package dirls

import (
	"io/ioutil"
	"os"
	"testing"

	"github.com/c2fo/testify/require"
)

func TestDirLs(t *testing.T) {
	t.Run("directory exists", func(t *testing.T) {
		ff, _, err := DirLs(".")
		require.NoError(t, err)
		require.NotEmpty(t, ff)
	})

	t.Run("directory does not exist", func(t *testing.T) {
		ff, _, err := DirLs("../../bobomeme")
		require.Error(t, err)
		require.Nil(t, ff)
	})

	t.Run("directory is empty", func(t *testing.T) {
		dir, err := ioutil.TempDir("", "bento-test")
		require.NoError(t, err)
		ff, _, err := DirLs(dir)
		require.NoError(t, err)
		require.Equal(t, 1, len(ff))
		err = os.RemoveAll(dir)
		require.NoError(t, err)
	})
}

func TestSortFiles(t *testing.T) {
	t.Run("sort ASC", func(t *testing.T) {
		ff, _, err := DirLs(".")
		require.NoError(t, err)
		require.NotEmpty(t, ff)
		SortFiles(ff, ASC)
		for i, f := range ff {
			if i == 0 {
				continue
			}
			require.True(t, ff[i-1].size < f.size)
		}
	})

	t.Run("sort DESC", func(t *testing.T) {
		ff, _, err := DirLs(".")
		require.NoError(t, err)
		require.NotEmpty(t, ff)
		SortFiles(ff, DESC)
		for i, f := range ff {
			if i == 0 {
				continue
			}
			require.True(t, ff[i-1].size > f.size)
		}
	})
}
