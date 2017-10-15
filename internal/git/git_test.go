package git

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestGit(t *testing.T) {
	out, err := Run("status")
	assert.NoError(t, err)
	assert.NotEmpty(t, out)

	out, err = Run("command-that-dont-exist")
	assert.Error(t, err)
	assert.Empty(t, out)
	assert.Equal(
		t,
		"git: 'command-that-dont-exist' is not a git command. See 'git --help'.\n",
		err.Error(),
	)
}

func TestClean(t *testing.T) {
	out, err := Clean("asdasd 'ssadas'\nadasd", nil)
	assert.NoError(t, err)
	assert.Equal(t, "asdasd ssadas", out)

}
