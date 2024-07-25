package controller

import (
	"fmt"
	"leonardoramosantos/tfc_plan_notifier/api"
	"leonardoramosantos/tfc_plan_notifier/config"
	"time"

	"github.com/slack-go/slack"
)

const title_format = `*Plan waiting for approval <https://app.terraform.io/app/%s/workspaces/%s/runs/%s|%s - %s> for %s*`
const body_format = `Hello there,

During your work, it was detected a Terraform Plan waiting for approval. The run finished its plan on %s.

You can access and approve this run <https://app.terraform.io/app/%s/workspaces/%s/runs/%s|here>.
`
const footer_format = `This alert was emited using *Terraform Plan Notifier* - https://github.com/leonardoramosantos/tfc_plan_notifier`

func (x *controller) DispatchSlackNotifications(plan config.ConfigScan, orgObj api.Organization, wksObj api.Workspace, runObj api.Run) {
	for _, slack_notification := range plan.SlackNotifications {
		slack_handler := slack.New(slack_notification.Token)

		time_since_plan := time.Since(runObj.RunAttr.Timestamps.PlanPlannedAt)

		title_formated := fmt.Sprintf(
			title_format,
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
			body_format,
			runObj.RunAttr.Timestamps.PlanPlannedAt,
			orgObj.Attributes.Name,
			wksObj.Attributes.Name,
			runObj.Id,
		)
		body := slack.NewTextBlockObject("mrkdwn", body_formated, false, false)
		body_section := slack.NewSectionBlock(body, nil, nil)

		footer_formated := fmt.Sprintf(
			footer_format,
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

		for _, slack_channel := range slack_notification.Channels {
			log.Debugf("Sending slack notification Channel: %s, Run: %s", slack_channel, runObj.Id)
			_, _, err := slack_handler.PostMessage(
				slack_channel,
				slack.MsgOptionBlocks(blocks.BlockSet...),
			)
			if err != nil {
				log.Errorf("Error sending message %s", err)
			}
		}
	}
}
