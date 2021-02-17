# Terraform Metering

This repo contains Metering configuration

## Prerequisites

Follow these steps in order to setup your environment to run this demo.

- [Install Docker](https://docs.docker.com/)
- [Install Docker Compose](https://docs.docker.com/compose/install/)

## Folders

This terraform deployment uses multiple different folders, one for each tenancy:

- `dev`
- `staging`
- `production`

## State Management

We store state for all workspaces in the S3 bucket on defineed in `terraform/init.tf` backend
terraform projects use the same workspace nomenclature with an S3 backend using the same

## Working with this repo samples

```bash
# Initialize directories locally (run only once)
# This allows to store state locally about resources
# so we don't have to sync with an external datastore all the time
docker-compose run -w /terraform/live/dev terraform init

# To pull down state
# This is the current state on the environment. Not the
# changes we have in this branch
docker-compose run -w /terraform/live/dev terraform refresh


# Getting the latest version of the image
docker-compose build --no-cache

# Creating a plan with local changes:
# This is the diff between our state when we refreshed and our local changes
docker-compose run -w /terraform/live/dev terraform plan -out=test.tfplan

# Applying the plan (this is what mutates the actual configuration in Okta)
docker-compose run -w /terraform/live/dev terraform apply "test.tfplan"
```


## Creating the kube config file

```bash
eksctl --profile devadmin utils write-kubeconfig --name <CLUSTER_NAME> --kubeconfig ~/.kube/metering-dev
```

wher `<CLUSTER_NAME>` is the name of the output product
