# gin-rating-api
Serverless API built with AWS Lambda, API Gateway and Cognito for authorization.

### TODO
- Setup circleci workflow to build, lint, test and deploy API. ✅
- Add logging library `logrus`. ✅
- Cognito authorization for API endpoints. ✅
- Design dynamo table to store information for different gins. ✅
  - Support adding new gin items to the table against a specific user.
    - Remove and update as well.
  - Support users rating gins.
  - Support aggregating gin ratings for users.
  - Store datestamp against reviews as to support 'gin of the week' and such.
- Add package(s) for DB interactions with DynamoDB. ✅
- Investigate/create package to easily make JSON responses for endpointd `render`. ✅
- Parse claims from the Authorization header JWT token. ✅
- Change logging to use use `logger.For` instead of `logger.Entry`. 🔧
- Review /gins handler unit tests.
- Review error responses in /gins endpoints.
- Review handling of duplicate gins.
  - Allow users to only update gins they have uploaded.
- Endpoints to add reviews for gins.
  - Short text review and an out of 5 star rating for each gin.
- Explore backend work for user signup and login.
- Configure DNS.
  
  MVP COMPLETE -> FRONTEND