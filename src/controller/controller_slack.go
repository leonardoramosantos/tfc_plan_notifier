package controller

import (
	"fmt"
	"leonardoramosantos/tfc_plan_notifier/api"
	"leonardoramosantos/tfc_plan_notifier/config"
	"time"

	"github.com/slack-go/slack"
)

const waiting_run_title_format = `*Plan waiting for approval <https://app.terraform.io/app/%s/workspaces/%s/runs/%s|%s - %s> for %s*`
const waiting_run_body_format = `Hello there,

During your work, it was detected a Terraform Plan waiting for approval. The run finished its plan on %s.

You can access and approve this run <https://app.terraform.io/app/%s/workspaces/%s/runs/%s|here>.
`
const waiting_run_footer_format = `This alert was emited using *Terraform Plan Notifier* - https://github.com/leonardoramosantos/tfc_plan_notifier`

const errored_run_title_format = `*Plan Errored <https://app.terraform.io/app/%s/workspaces/%s/runs/%s|%s - %s> without correction %s*`
const errored_run_body_format = `Hello there,

During your work, it was detected a Terraform Plan that contains errors and was not corrected. The run errored on %s.

You can view the information about the run <https://app.terraform.io/app/%s/workspaces/%s/runs/%s|here>.
`
const errored_run_footer_format = `This alert was emited using *Terraform Plan Notifier* - https://github.com/leonardoramosantos/tfc_plan_notifier`

func (x *controller) DispatchSlackErroredRunNotification(plan config.ConfigScan, orgObj api.Organization, wksObj api.Workspace, runObj api.Run) {
	time_since_plan := time.Since(runObj.RunAttr.Timestamps.PlanErroredAt)

	title_formated := fmt.Sprintf(
		errored_run_title_format,
		orgObj.Attributes.Name,
		wksObj.Attributes.Name,
		runObj.Id,
		wksObj.Attributes.Name,
		runObj.Id,
		time_since_plan,
	)
	header := slack.NewTextBlockObject("mrkdwn", title_formated, false, false)
	headerSection := slack.NewSectionBlock(header, nil, nil)

	divider := slack.NewDividerBlock()

	body_formated := fmt.Sprintf(
		errored_run_body_format,
		runObj.RunAttr.Timestamps.PlanErroredAt,
		orgObj.Attributes.Name,
		wksObj.Attributes.Name,
		runObj.Id,
	)
	body := slack.NewTextBlockObject("mrkdwn", body_formated, false, false)
	body_section := slack.NewSectionBlock(body, nil, nil)

	footer_formated := fmt.Sprintf(
		errored_run_footer_format,
	)
	footer := slack.NewTextBlockObject("mrkdwn", footer_formated, false, false)
	footer_section := slack.NewSectionBlock(footer, nil, nil)

	blocks := slack.Blocks{
		BlockSet: []slack.Block{
			headerSection,
			divider,
			body_section,
			divider,
			footer_section,
		},
	}

	dispatchSlackNotifications(plan, blocks)
}

func (x *controller) DispatchSlackWaitingApprovalNotification(plan config.ConfigScan, orgObj api.Organization, wksObj api.Workspace, runObj api.Run) {
	time_since_plan := time.Since(runObj.RunAttr.Timestamps.PlanPlannedAt)

	title_formated := fmt.Sprintf(
		waiting_run_title_format,
		orgObj.Attributes.Name,
		wksObj.Attributes.Name,
		runObj.Id,
		wksObj.Attributes.Name,
		runObj.Id,
		time_since_plan,
	)
	header := slack.NewTextBlockObject("mrkdwn", title_formated, false, false)
	headerSection := slack.NewSectionBlock(header, nil, nil)

	divider := slack.NewDividerBlock()

	body_formated := fmt.Sprintf(
		waiting_run_body_format,
		runObj.RunAttr.Timestamps.PlanPlannedAt,
		orgObj.Attributes.Name,
		wksObj.Attributes.Name,
		runObj.Id,
	)
	body := slack.NewTextBlockObject("mrkdwn", body_formated, false, false)
	body_section := slack.NewSectionBlock(body, nil, nil)

	footer_formated := fmt.Sprintf(
		waiting_run_footer_format,
	)
	footer := slack.NewTextBlockObject("mrkdwn", footer_formated, false, false)
	footer_section := slack.NewSectionBlock(footer, nil, nil)

	blocks := slack.Blocks{
		BlockSet: []slack.Block{
			headerSection,
			divider,
			body_section,
			divider,
			footer_section,
		},
	}

	dispatchSlackNotifications(plan, blocks)
}

// Sends a Slack Notification
func dispatchSlackNotifications(plan config.ConfigScan, message_blocks slack.Blocks) {
	for _, slack_notification := range plan.SlackNotifications {
		slack_handler := slack.New(slack_notification.Token)
		for _, slack_channel := range slack_notification.Channels {
			log.Debugf("Sending slack notification Channel: %s", slack_channel)
			_, _, err := slack_handler.PostMessage(
				slack_channel,
				slack.MsgOptionBlocks(message_blocks.BlockSet...),
			)
			if err != nil {
				log.Errorf("Error sending message %s", err)
			}
		}
	}
}
