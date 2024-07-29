# Terraform Plan Notifier

![Terraform](https://img.shields.io/badge/terraform-%235835CC.svg?style=for-the-badge&logo=terraform&logoColor=white)![Slack](https://img.shields.io/badge/Slack-4A154B?style=for-the-badge&logo=slack&logoColor=white)![Go](https://img.shields.io/badge/go-%2300ADD8.svg?style=for-the-badge&logo=go&logoColor=white)![Docker](https://img.shields.io/badge/docker-%230db7ed.svg?style=for-the-badge&logo=docker&logoColor=white)

Project to notify technology teams about [Terraform Cloud](https://www.hashicorp.com/products/terraform) plans with errors and waiting for approval.
Multiple Organizations and Slack channels can be configured on the application.
The application also makes possible set a minimum waiting time for the plan to be warned.

[Slack](https://slack.com/) notifications are supported at the moment, but Teams and E-mail notifications are in the Roadmap.

# Configuration

## Terraform Token

The terraform Token can be set using the `TERRAFORM_TOKEN` environment variable, or inside the configuration file using the key `tfc-token`.

## Terraform Plan Scans

In order to perform the scans of Terraform Plans and find the ones with status `waiting for approval` and `errored` for a certain amount of time, the application tries to load a file called `config.yaml` either from the same directory of the executable or from `/etc`.
The structure of the file is described here:
```yaml
# -- Not required. The plain text token to access Terraform API. If not specified, a environment variable called TERRAFORM_TOKEN must be set
tfc-token: ""
# -- Plan to run agains terraform. Multiple plans can be specified
scans:
    # -- ISO 8601 Duration string specifying how old a waiting plan should be to warn. If Empty, waiting plans will not be warned
  - waiting-approval-interval: ""
    # -- ISO 8601 Duration string specifying how old a errored plan should be to warn. If Empty, errored plans will not be warned
    errored-plan-interval: ""
    # -- Not required. RegExp to filter Terraform Organizations
    organization: ".*"
    # -- Not required. RegExp to filter Terraform Workspaces
    workspace: ".*"
    slack-notifications:
        # -- Slack token
      - token: <SLACK_TOKEN>
        # -- List of string names of the channels to send warnings
        channels:
        - <SLACK_CHANNEL_TO_POST>
        - <SLACK_CHANNEL_TO_POST>
    # Different plans can be set, as it is a list
  - ...
```

# Running the application

By default, the *Terraform Plan Notifier* is serverd as [Docker](https://www.docker.com/) Image, but a Helm Chart makes it easier to configure and run on [kubernetes](https://kubernetes.io/) clusters.

## Locally

Create a configuration file `config.yaml` folowing the structure described above. Run the docker image:
```
docker run -v ${PWD}:/config.yaml leonardoramosantos/tfc_plan_notifier:latest
```

## K8S CronJob

A Helm Chart is provided to simplify the process of deploying it to Kubernetes Clusters, stored here https://github.com/leonardoramosantos/helm-charts.

To use this Helm Chart:

```
helm repo add ls-public-helm-charts https://leonardoramosantos.github.io/helm-charts/

helm repo update

helm install tfc-plan-notifier ls-public-helm-charts/tfc-plan-notifier
```

# Building the image

## For development

`docker-compose` can be used to run the development environment:
```
cd tfc_plan_notifier
docker-compose -f docker/docker-compose.yml build
docker-compose -f docker/docker-compose.yml up
```

## For production

To build the image for production, just run the Docker build command:
```
cd tfc_plan_notifier
docker build . -f docker/Dockerfile -t tfc_plan_notifier
```

# LICENSE

[![Licence](https://img.shields.io/github/license/Ileriayo/markdown-badges?style=for-the-badge)](./LICENSE)
