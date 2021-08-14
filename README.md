# AWS_challenge_1

## folder structure of each part 
  - **mock** : this section provide mocke of dependency for testing.
  - **model** : this section provide input and output structure in lambda function.
  - **serviceHandler** : this section contains lambda handler and it test file.
  - **main.go** : this go file run lambda faunction.
  
  
## create device
  >This lambda get device info form *apiGetway* and add new device to *dynamoDB* and returen response to *apiGetway* again.
  
  When lambda is called, first `init()` function is invoked and create new session to AWS and connect to *dynamoDB* service and set to `dynamoDB` variable then `main()` function
  is invoked, in `main()` first create a new handler onject by `NewCore()` function in `serviceHandler` package.
  
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
  - unmarshalmap method ---> for preparing output data.
  
  In `handelr()` function we get device using id that came from the input, if device was exist return device input otherwise return not found error
  
  
  ## Test
  >There are test file in `serviceHnadler` package.
  
 there are serveral test cases in each test function that cover different state  then for each test, first we mock dependency and create handler by using it dependency.
 then we execute handelr function and in end check rusult 
 
