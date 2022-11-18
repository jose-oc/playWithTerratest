# Terraform with Terratest

## Adding a map variable

We're adding a map variable here to see how to use it in terraform and how to pass this type of variable in terratest.

Execute this terraform code:

```shell
terraform init
terraform apply -auto-approve -var your_name=Pepe -var 'rock_band={"drums": "Ana", "bass": "Pepa", "guitar": "Alicia"}'
```