# AWS_challenge_1

## api
> URL of api ==> [https://849g0lddal.execute-api.us-west-2.amazonaws.com/api/](https://849g0lddal.execute-api.us-west-2.amazonaws.com/api/)

## folder structure of each part 
  - **mock** : this section provides mock of dependency for testing.
  - **model** : this section provides input and output structure in lambda function.
  - **serviceHandler** : this section contains lambda handler, and it tests file.
  - **main.go** : this go files run lambda function.
  
## environment variable
>This environment variable must be set in lambda configuration
  - **TABLE_NAME** : name of dynamoDB table that data save to it
  - **ACCESS_TOKEN** : Access_token_id for access to aws service
  - **SECRET_KEY** : Secret_key_access for access to aws service

## create device
  >This lambda get device info form *apiGetway* and add new device to *dynamoDB* and returen response to *apiGetway* again.
  
  When lambda is called, first `init()` function is invoked and create new session to AWS and connect to *dynamoDB* service and set to `dynamoDB` variable then `main()` function
  is invoked, in `main()` first create a new handler object by `NewCore()` function in `serviceHandler` package.
  
  we must pass to `NewCore()` two dependency :
  - dynamoDB instance ---> for communicate with the database
  - marshalMap method ---> for marshal input struct to map for pass to `PutItem()` in dynamoDB service.
  
  In `handelr` function, first prepair input data to pass it to `PutItem()` function in dynamoDB by using marshal function and save device info in databse then return output.
## get device
 >This lambda get id of device *apiGetway* and get device info from *dynamoDB* and returen response to *apiGetway* again.
 
  When lambda is called, first `init()` function is invoked and create new session to AWS and connect to *dynamoDB* service and set to `dynamoDB` variable then `main()` function
  is invoked, in `main()` first create a new handler onject by `NewCore()` function in `serviceHandler` package.
  
  we must pass to `NewCore()` two dependency :
  - dynamoDB instance ---> for communicate with the database
  - unmarshal method ---> for preparing output data.
  
  In `handelr()` function we get device using id that came from the input, if device was existed return device input otherwise return not found error
  
  
  ## Test
  >There are test file in `serviceHnadler` package.
  
 there are several test cases in each test function that cover different state  then for each test, first we mock dependency and create handler by using it dependency.
 then we execute handler function and in end check result 
 
