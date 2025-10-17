package sm

import (
	"bytes"
	"compress/gzip"
	"encoding/base64"
	"encoding/json"
	"fmt"
)

// ToMermaidLiveURL compresses and encodes a Mermaid diagram to a pako URL
func ToMermaidLiveURL(mermaidDiagram string) string {
	// Wrap the Mermaid code in a JSON object with key "code"
	payload := map[string]string{
		"code": mermaidDiagram,
	}
	jsonBytes, err := json.Marshal(payload)
	if err != nil {
		panic(err)
	}

	// Compress the JSON bytes with gzip
	var buf bytes.Buffer
	gz := gzip.NewWriter(&buf)
	_, err = gz.Write(jsonBytes)
	if err != nil {
		panic(err)
	}
	gz.Close()

	// Base64 URL encode
	encoded := base64.RawURLEncoding.EncodeToString(buf.Bytes())

	return "https://mermaid.live/edit#pako:" + encoded
}

// GenerateMermaidDiagram generates the Mermaid flowchart string from a StateMachine
func GenerateMermaidDiagram(sm *StateMachine) string {
	var sb bytes.Buffer
	sb.WriteString("flowchart TD\n")

	for _, t := range sm.transitions {
		srcName := t.Source.Name
		if srcName == "" {
			srcName = "State_" + fmt.Sprintf("%p", t.Source)
		}
		tgtName := t.Target.Name
		if tgtName == "" {
			tgtName = "State_" + fmt.Sprintf("%p", t.Target)
		}
		label := t.Name
		sb.WriteString("    " + srcName + " -->|" + label + "| " + tgtName + "\n")
	}

	return sb.String()
}
