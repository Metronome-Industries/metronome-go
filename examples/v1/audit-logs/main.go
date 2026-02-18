package main

import (
	"context"
	"fmt"

	"github.com/Metronome-Industries/metronome-go/v3"
)

func main() {
	client := metronome.NewClient() // uses METRONOME_BEARER_TOKEN env var

	listAuditLogs(client)
}

// listAuditLogs lists audit log entries with optional filtering.
func listAuditLogs(client metronome.Client) {
	page, err := client.V1.AuditLogs.List(context.TODO(), metronome.V1AuditLogListParams{
		Limit: metronome.Int(10),
		Sort:  metronome.V1AuditLogListParamsSortDateDesc,
	})
	if err != nil {
		panic(err)
	}
	for _, entry := range page.Data {
		fmt.Printf("Audit log: %+v\n", entry)
	}

	// To automatically paginate, use ListAutoPaging:
	//
	// iter := client.V1.AuditLogs.ListAutoPaging(context.TODO(), metronome.V1AuditLogListParams{})
	// for iter.Next() {
	//     entry := iter.Current()
	//     fmt.Printf("Entry: %+v\n", entry)
	// }
	// if err := iter.Err(); err != nil {
	//     panic(err)
	// }
}
