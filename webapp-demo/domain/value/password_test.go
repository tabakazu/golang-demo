package value

import "testing"

func TestGenerateDigestFromPassword(t *testing.T) {
	p := Password("foobar")
	digest, err := p.GenerateDigest()
	if err != nil {
		t.Errorf("Password.GenerateDigest() error = %v", err)
	}
	if !digest.IsCorrectPassword(p) {
		t.Errorf("Password.GenerateDigest() returns incorrect password")
	}
}

func TestGenerateDigestFromBlankPassword(t *testing.T) {
	p := Password("")
	digest, err := p.GenerateDigest()
	if err != nil {
		t.Errorf("Password.GenerateDigest() error = %v", err)
	}
	if !digest.IsCorrectPassword(p) {
		t.Errorf("Password.GenerateDigest() returns incorrect password")
	}
}

func TestForceGenerateDigestFromPassword(t *testing.T) {
	p := Password("foobar")
	if !p.ForceGenerateDigest().IsCorrectPassword(p) {
		t.Errorf("Password.GenerateDigest() returns incorrect password")
	}
}
