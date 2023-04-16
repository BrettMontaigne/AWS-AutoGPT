package autogpt_handler

import (
	"context"
	"fmt"
	"os/exec"
	"github.com/aws/aws-lambda-go/lambda"
)

type MyEvent struct {
	Message string `json:"message"`
}

func HandleRequest(ctx context.Context, event MyEvent) (string, error) {
	cmd := exec.Command("python3", "../../scripts/main.py")

	output, err := cmd.CombinedOutput()
	if err != nil {
		fmt.Printf("Error executing Python script: %v\n", err)
		return "", err
	}

	return fmt.Sprintf("Script output: %s", output), nil
}

func main() {
	lambda.Start(HandleRequest)
}