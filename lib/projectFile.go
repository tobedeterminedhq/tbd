package lib

import (
	"fmt"
	"io"

	servicev1 "github.com/benfdking/tbd/proto/gen/go/tbd/service/v1"
	"sigs.k8s.io/yaml"
)

func ParseProjectFile(reader io.Reader) (*servicev1.ProjectFile, error) {
	bs, err := io.ReadAll(reader)
	if err != nil {
		return nil, fmt.Errorf("reading file: %w", err)
	}
	var c servicev1.ProjectFile
	if err := yaml.Unmarshal(bs, &c); err != nil {
		return nil, fmt.Errorf("unmarshaling yaml: %w", err)
	}
	return &c, nil
}
