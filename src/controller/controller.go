package controller

import (
	"leonardoramosantos/tfc_plan_notifier/api"
	"leonardoramosantos/tfc_plan_notifier/config"
	"regexp"
	"time"

	"github.com/op/go-logging"
)

type controller struct {
	config        *config.TFCApi
	config_plan   *config.ConfigPlan
	organizations []api.Organization
}

var log = logging.MustGetLogger("tfc_plan_notifier")

func (x *controller) planVerifyRuns(plan config.ConfigScan, wks api.Workspace) {
	//var t = time.Now().Add(-time.Minute * 5)

	var runs = api.GetRuns(x.config, wks.Id)

	for _, run := range runs {
		log.Debugf("Testing Run: %s Status: %s Time: %s", run.Id, run.RunAttr.Status, run.RunAttr.Timestamps.PlanQueueableAt)

		if (run.RunAttr.Status == "planned") && run.RunAttr.Timestamps.PlanQueueableAt.Before(time.Now().Add(-plan.TimeInterval)) {
			log.Debugf("Plan matches Run: %s", run.Id)
		}
	}
}

func (x *controller) planVerifyWorkspaces(plan config.ConfigScan, org api.Organization) {
	var workspaces = api.GetWorkspaces(x.config, org.Id)

	for _, wks := range workspaces {
		var wks_match, _ = regexp.MatchString(plan.WorkspaceMatchExpr, wks.Id)
		if wks_match {
			log.Debugf("Searching " + wks.Id)

			x.planVerifyRuns(plan, wks)
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
	x.organizations = api.GetOrganizations(x.config)

	for _, plan := range x.config_plan.Plans {
		x.planVerifyOrganizations(plan)
	}
}

func GetController() *controller {
	var contr = controller{}

	contr.config = config.GetConfig()
	contr.config_plan = config.GetConfigPlan("")

	return &contr
}
