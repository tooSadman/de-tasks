# Task 2 - Function as a Service sample

**The goal**: try in the wild real FaaS offering.

We expect you to implement HTTP-triggered Function, which stores incoming HTTP request content in a cloud Object Storage as json file.

### Requirements

- Function was deployed to your cloud account
- Function could be triggered via HTTP request
- Function does not allow anonymous access
- Content of a request is stored as json file in Object Store
- New requests does not override any existing file in Object Store

### Deliveries

- Reference to git repository with FaaS sources
- Cloud web-interface screenshots to prove function successful invocation:
    - function invocation logs from cloud interface
    - created files in Object Storage

## Steps

1. Set the project with `gcloud config set project <project_id>`
2. In current directory run `make` command
3. For testing the function, get its URL after creation and invoke the following command:  
    ```
    curl -m 70 -X POST <your_function_url> \
    -H "Authorization:bearer $(gcloud auth print-identity-token)" \
    -H "Content-Type:application/json" \
    -d '{
      "name": "Hello World"
    }'
    ```  
    That should create a cloud bucket (if it does not exist) and write new object to it.
