// Copyright 2018 Drone.IO Inc
// Use of this software is governed by the Business Source License
// that can be found in the LICENSE file.

package google

import "testing"

func TestDefaults(t *testing.T) {
	p := New().(*provider)
	if got, want := p.image, "ubuntu-1510-wily-v20151114"; got != want {
		t.Errorf("Want image %q, got %q", want, got)
	}
	if got, want := p.zone, "us-central1-a"; got != want {
		t.Errorf("Want region %q, got %q", want, got)
	}
	if got, want := p.size, "n1-standard-1"; got != want {
		t.Errorf("Want size %q, got %q", want, got)
	}
	if got, want := p.key, ""; got != want {
		t.Errorf("Want key %q, got %q", want, got)
	}
	if got, want := p.token, ""; got != want {
		t.Errorf("Want token %q, got %q", want, got)
	}
	if got, want := len(p.tags), 0; got != want {
		t.Errorf("Want %d tags, got %d", want, got)
	}
}
