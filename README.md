# Terraform Plan Notifier

Project to notify teams about Terraform plans waiting for approval.


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
