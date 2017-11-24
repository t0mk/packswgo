# Generate Golang client code for Packet from Swagger api-docs

## Prerequiesitues

Install this repo:
```
$ go get github.com/t0mk/packswgo
```

Install Java runtime.



Obtain `swagger-codegen-cli.jar`, maybe from https://swagger.io/docs/swagger-tools/#installation-from-maven-central-14. For example

```
wget -O swagger-codegen-cli.jar https://oss.sonatype.org/content/repositories/releases/io/swagger/swagger-codegen-cli/2.2.1/swagger-codegen-cli-2.2.1.jar
```

## Usage

To generate Golang client code for Packet, just run inside the repo dir

```
java -jar ~/swagger-codegen-cli.jar generate -i https://api.packet.net/api-docs -l go -o ./ -c c.json
```
