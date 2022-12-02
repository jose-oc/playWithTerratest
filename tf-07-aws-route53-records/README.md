# Terraform with Terratest

## Working with records in AWS Route53

With this code we create new Route53 DNS records so that we link a name to an IP. 

The test simply check terraform outputs is compound as expected.

1. `cd test`
2. `go test -v -timeout 30m`
3. If you want to run the test skipping one stage add the environment variable `SKIP_{stage-name}=true` and run the above command. For instance: `SKIP_cleanup_terraform=true go test -v -timeout 30m`

### Prerequisites

Having a Route53 zone defined.
