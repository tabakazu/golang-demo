package value

import "testing"

func TestCheackCorrectPasswordWithPasswordDigest(t *testing.T) {
	p := Password("foobar")
	digest := p.ForceGenerateDigest()

	if !digest.IsCorrectPassword(p) {
		t.Errorf("Password.GenerateDigest() returns incorrect password")
	}
}
