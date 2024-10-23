# Goa project

A project for Outh 2 authentication (client credentials flow) with a token generation endpoint `/auth`



![DSL entities](image.png)

## Requirements 

* Go 1.22
* Make

## Usage 

To generate the Goa scaffolding code  
`make generate`

Running binaries
`make build-cli`
`make run-local-http`

Running client
`./client-binary token --help`
`./client-binary math --help`
`./client-binary math mul --help`


# Running some manual tests
1) Let’s get a token with the CLI.
`./server-binary token auth --body '{ "password": "password", "username": "user" }'`

`curl -X POST http://localhost:8080/auth -H 'Content-Type: application/json' --data-binary '{ "password": "password", "username": "user" }'`



2) Let’s try to multiply some numbers.
`./server-binary math mul --numbers '["4", "3.543", "-2"]' --token "TOKEN_FROM_THE_PREVIOUS_STEP"`
  
`curl http://localhost:8080/mul/4,3.543,-2 -H 'Authorization: Bearer TOKEN_FROM_THE_PREVIOUS_STEP'`



3) Finally, let’s submit a JSON document for addition of its numeric elements.
`./server-cli math sum   --token "TOKEN_FROM_THE_PREVIOUS_STEP"  --body '{"a":{"b":4},"c":2, "d":[1, 3]}'` 

```
curl -X POST http://localhost:8080/sum \
-H 'Authorization: Bearer eyJhbGciOiJSUzI1NiIsInR5cCI6IkpXVCJ9.eyJhdWQiOiJ1c2VyIiwiZXhwIjoxNjY0ODgwNTY4LCJqdGkiOiJmZDQ3Zjc2MC0xN2Q3LTRlNDYtYjdjZi0yMWM1MzBjNzc0NzgiLCJpYXQiOjE2NjQ4NzY5NjgsImlzcyI6Imh0dHA6Ly9sb2NhbGhvc3Q6ODA4MCIsInN1YiI6InVzZXIifQ.efisAM2XQdbmuNPRoYffYewH636CNgCLNh4tecnkXSGygfPvNrQhXpkeM3zA731j2VIIJqss8NeDcXPxQFbwHDdcxqmt5w0b-onNuGCFf2u7W55rWANNOCMjke5B4QSCop9waVV-eXSF70yIXPT5iNKD7SIlOv4FrNzkvNye5w4VCg7g-oZovjsZdmaLN2SfLdzyxXTBLw4TCst5SFiVzzcyhPMOevnX5mSv6p0uI_iPCO0GhdLyW20-HghazSRDI6xoj5vepuoP5_fCPdpwZUhsO75o_pl66IENhBmiovEsvpbEV5qYc0sXVKH4yk6Jcis6OCbHL3gQSVB97Sc4xTSmGkEkGWvbdyf_j4uBKRE8SVMZd93EqDJkze5Os7umKP5Nw8ws4JLuoyWab-kM1wNwTJmGNFWuhC_Tfql4blDKFZPIZOF-Yqqj67QP5f1nQ9pW1GNerbiPhTQ4noUNRtyokruyFjBhWql_0ebYf94xesvJMIa93GnVoIlSuuX9' \
--data-binary '"{\"a\":{\"b\":4},\"c\":2, \"d\":[1, 3]}"'
```