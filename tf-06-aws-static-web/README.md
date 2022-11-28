# Terraform with Terratest

## Working with Terratest Stages

TBC

## Working with AWS Route53

This test create the webserver with all its components, test the webserver is running (getting a page, similar to a healthcheck) and then destroy the infrastructure.

1. `cd test`
2. `go test -v -timeout 30m`
3. If you want to run the test skipping one stage add the environment variable `SKIP_{stage-name}=true` and run the above command. For instance: `SKIP_cleanup_terraform=true go test -v -timeout 30m`

### Prerequisites

Having a Route53 zone defined.


## Caveat

This code generates a new certificate to ssh into the EC2 instance, this way of doing so it's ok for development and 
temporary environments but not for production as it is not as secure as it should be. 