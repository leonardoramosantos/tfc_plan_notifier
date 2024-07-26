package controller

import (
	"leonardoramosantos/tfc_plan_notifier/api"
	"leonardoramosantos/tfc_plan_notifier/config"
	"leonardoramosantos/tfc_plan_notifier/utils"
	"os"
	"regexp"
	"time"

	"github.com/op/go-logging"
)

type controller struct {
	api           *api.TFCApi
	config_plan   *config.ConfigPlan
	organizations []api.Organization
}

var log = logging.MustGetLogger("tfc_plan_notifier")

func (x *controller) planVerifyRuns(plan config.ConfigScan, org api.Organization, wks api.Workspace) {
	var runs = x.api.GetRuns(wks.Id)

	for _, run := range runs {
		log.Debugf("Testing Run: %s Status: %s Time: %s", run.Id, run.RunAttr.Status, run.RunAttr.Timestamps.PlanPlannedAt)

		var duration, _ = utils.ParseISODuration(plan.TimeInterval)
		if (run.RunAttr.Status == "planned") && run.RunAttr.Timestamps.PlanPlannedAt.Before(time.Now().Add(-duration)) {
			log.Debugf("Plan matches Run: %s", run.Id)
			x.DispatchSlackNotifications(plan, org, wks, run)
		}
	}
}

func (x *controller) planVerifyWorkspaces(plan config.ConfigScan, org api.Organization) {
	var workspaces = x.api.GetWorkspaces(org.Id)

	for _, wks := range workspaces {
		var wks_match, _ = regexp.MatchString(plan.WorkspaceMatchExpr, wks.Id)
		if wks_match {
			log.Debugf("Searching " + wks.Id)

			x.planVerifyRuns(plan, org, wks)
		} else {
			log.Debugf("Not Matching Workspace Wks: %s, Str: %s, Match: %s", wks.Id, plan.WorkspaceMatchExpr, wks_match)
		}
	}
}

func (x *controller) planVerifyOrganizations(plan config.ConfigScan) {
	for _, org := range x.organizations {
		var org_match, _ = regexp.MatchString(plan.OrganizationMatchExpr, org.Id)
		if org_match {
			log.Debugf("Searching " + org.Id)

			x.planVerifyWorkspaces(plan, org)
		} else {
			log.Debugf("Not Matching Organization Org: %s, Str: %s, Match: %s", org.Id, plan.OrganizationMatchExpr, org_match)
		}
	}
}

func (x *controller) StartPlans() {
	x.organizations = x.api.GetOrganizations()

	for _, plan := range x.config_plan.Plans {
		x.planVerifyOrganizations(plan)
	}
}

func GetController() *controller {
	var contr = controller{}

	contr.config_plan = config.GetConfigPlan("")

	var token string
	if contr.config_plan.TFCToken != "" {
		token = contr.config_plan.TFCToken
	} else {
		token = os.Getenv("TERRAFORM_TOKEN")
	}

	contr.api = api.GetTFCApi(token)

	return &contr
}
