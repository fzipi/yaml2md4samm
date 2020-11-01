package model

import (
	"testing"
	"yaml2md4samm/parsing"
)

var design = `
---
model: SAMM20
id: 88c296acaae841a2b2fc5314bff44cb4
name: Design
description: |
  Design concerns the processes and activities related to how an organization defines goals and creates software within development projects. In general, this will include requirements gathering, high-level architecture specification and detailed design.

color:
  #993300

logo: https://icon-library.net/icon/architecture-icon-png-26.html

order: 2
type: Function
`

func TestModelParse(t *testing.T) {
	var samm = Model{Version: 2.0, License: "CC-1.0", ExecutiveSummary: "Summary"}

	err := parsing.ParseFullModel(&samm)
	if err != nil {
		t.Error(err)
	}

	if samm.Version != 2.0 {
		t.Errorf("Model parsing incorrect, got: %s, want: %s.", samm.Version, 2.0)
	}

}
