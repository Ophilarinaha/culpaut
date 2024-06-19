import (
	"context"
	"fmt"
	"io"

	logging "cloud.google.com/go/logging/apiv2"
	"cloud.google.com/go/logging/apiv2/loggingpb"
	"google.golang.org/api/iterator"
)

// readEntries reads log entries from the given log.
func readEntries(w io.Writer, projectID, logID string) error {
	// projectID := "my-project-id"
	// logID := "my-log-id"
	ctx := context.Background()
	client, err := logging.NewClient(ctx, projectID)
	if err != nil {
		return fmt.Errorf("logging.NewClient: %v", err)
	}
	defer client.Close()

	it := client.Entries(ctx, &loggingpb.ListLogEntriesRequest{
		ResourceNames: []string{fmt.Sprintf("projects/%s/logs/%s", projectID, logID)},
	})
	for {
		entry, err := it.Next()
		if err == iterator.Done {
			break
		}
		if err != nil {
			return fmt.Errorf("it.Next: %v", err)
		}
		fmt.Fprintf(w, "Log entry: %v\n", entry)
	}
	return nil
}
  
