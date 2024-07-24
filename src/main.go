package main

import (
	"leonardoramosantos/tfc_plan_notifier/controller"
)

func main() {
	var scanController = controller.GetController()
	scanController.StartPlans()

	//config := &tfe.Config{
	//	Token:             apiToken,
	//	RetryServerErrors: true,
	//}

	//client, err := tfe.NewClient(config)
	//if err != nil {
	//	log.Fatal(err)
	//}

	//orgs, err := client.Organizations.List(context.Background(), nil)
	//if err != nil {
	//	log.Fatal(err)
	//}
}
