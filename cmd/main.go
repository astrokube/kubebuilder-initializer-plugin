package main

import (
	"bufio"
	"encoding/json"
	"fmt"
	"io"
	"log"
	"os"

	"github.com/astrokube/layout-kubebuilder-plugin/internal/command"
	"sigs.k8s.io/kubebuilder/v3/pkg/plugin"
	"sigs.k8s.io/kubebuilder/v3/pkg/plugin/external"
)

func main() {
	request := &external.PluginRequest{}
	if err := readPluginRequest(request); err != nil {
		processError(request, err)
	}

	response := command.Run(request)

	output, err := json.Marshal(response)
	if err != nil {
		processError(request, fmt.Errorf("encountered error marshaling output: %w | OUTPUT: %s", err, output))
	}

	fmt.Printf("%s", output)
}

func readPluginRequest(request *external.PluginRequest) error {
	reader := bufio.NewReader(os.Stdin)

	input, err := io.ReadAll(reader)
	if err != nil {
		return fmt.Errorf("encountered error reading STDIN: %w", err)
	}

	err = json.Unmarshal(input, request)
	if err != nil {
		return fmt.Errorf("encountered error unmarshaling STDIN: %w", err)
	}

	return nil
}

func processError(request *external.PluginRequest, err error) {
	errResponse := external.PluginResponse{
		APIVersion: request.APIVersion,
		Command:    request.Command,
		Metadata:   plugin.SubcommandMetadata{},
		Universe:   nil,
		Error:      true,
		ErrorMsgs: []string{
			err.Error(),
		},
		Flags: nil,
	}
	output, err := json.Marshal(errResponse)
	if err != nil {
		log.Fatalf("encountered error marshaling output: %s | OUTPUT: %s", err.Error(), output)
	}

	fmt.Printf("%s", output)
}
