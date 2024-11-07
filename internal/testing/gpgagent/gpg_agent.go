package gpgagent

import (
	"os"
	"os/exec"
)

// Kill the running gpg-agent to drop unlocked keys.
// This is useful to ensure tests don’t leave processes around (in TestMain), or for testing handling of invalid passphrases.
func KillGPGAgent(gpgHomeDir string) error {
	cmd := exec.Command("gpgconf", "--kill", "gpg-agent")
	cmd.Env = append(os.Environ(), "GNUPGHOME="+gpgHomeDir)
	return cmd.Run()
}
