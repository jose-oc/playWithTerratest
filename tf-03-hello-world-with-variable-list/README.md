# Terraform with Terratest

## Example with string and list variables

Execute this terraform code:

```shell
terraform init
terraform apply -auto-approve -var your_name=Pepe -var 'friends=["Michiel", "Irene", "Pilar"]'
```