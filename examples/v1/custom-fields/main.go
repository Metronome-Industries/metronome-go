package main

import (
	"context"
	"fmt"

	"github.com/Metronome-Industries/metronome-go/v3"
)

func main() {
	client := metronome.NewClient() // uses METRONOME_BEARER_TOKEN env var

	addKey(client)
	setValues(client)
	listKeys(client)
	deleteValues(client)
	removeKey(client)
}

// addKey adds a new custom field key definition.
func addKey(client metronome.Client) {
	err := client.V1.CustomFields.AddKey(context.TODO(), metronome.V1CustomFieldAddKeyParams{
		EnforceUniqueness: true,
		Entity:            metronome.V1CustomFieldAddKeyParamsEntityCustomer,
		Key:               "x_account_id",
	})
	if err != nil {
		panic(err)
	}
	fmt.Println("Custom field key added successfully")
}

// setValues sets custom field values for an entity.
func setValues(client metronome.Client) {
	err := client.V1.CustomFields.SetValues(context.TODO(), metronome.V1CustomFieldSetValuesParams{
		Entity:   metronome.V1CustomFieldSetValuesParamsEntityCustomer,
		EntityID: "99594816-e8a5-4bca-be21-8d1de0f45120",
		CustomFields: map[string]string{
			"x_account_id": "KyVnHhSBWl7eY2bl",
		},
	})
	if err != nil {
		panic(err)
	}
	fmt.Println("Custom field values set successfully")
}

// listKeys lists all custom field key definitions.
func listKeys(client metronome.Client) {
	page, err := client.V1.CustomFields.ListKeys(context.TODO(), metronome.V1CustomFieldListKeysParams{
		Entities: []string{"customer"},
	})
	if err != nil {
		panic(err)
	}
	for _, key := range page.Data {
		fmt.Printf("Key: %+v\n", key)
	}

	// To automatically paginate, use ListKeysAutoPaging:
	//
	// iter := client.V1.CustomFields.ListKeysAutoPaging(context.TODO(), metronome.V1CustomFieldListKeysParams{})
	// for iter.Next() {
	//     key := iter.Current()
	//     fmt.Printf("Key: %+v\n", key)
	// }
	// if err := iter.Err(); err != nil {
	//     panic(err)
	// }
}

// deleteValues deletes custom field values for an entity.
func deleteValues(client metronome.Client) {
	err := client.V1.CustomFields.DeleteValues(context.TODO(), metronome.V1CustomFieldDeleteValuesParams{
		Entity:   metronome.V1CustomFieldDeleteValuesParamsEntityCustomer,
		EntityID: "99594816-e8a5-4bca-be21-8d1de0f45120",
		Keys:     []string{"x_account_id"},
	})
	if err != nil {
		panic(err)
	}
	fmt.Println("Custom field values deleted successfully")
}

// removeKey removes a custom field key definition.
func removeKey(client metronome.Client) {
	err := client.V1.CustomFields.RemoveKey(context.TODO(), metronome.V1CustomFieldRemoveKeyParams{
		Entity: metronome.V1CustomFieldRemoveKeyParamsEntityCustomer,
		Key:    "x_account_id",
	})
	if err != nil {
		panic(err)
	}
	fmt.Println("Custom field key removed successfully")
}
