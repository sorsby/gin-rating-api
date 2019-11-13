# gin-rating-api
Serverless API built with AWS Lambda, API Gateway and Cognito for authorization.

### TODO
- Setup circleci workflow to build, lint, test and deploy API.
- Add logging library `logrus` ✅
- Cognito authorization for API endpoints.
- Investigate/create package to easily make JSON responses for endpoints.
- Design dynamo table to store information for different gins.
  - Support adding new gin items to the table against a specific user.
    - Remove and update as well.
  - Support users rating gins.
  - Support aggregating gin ratings for users.
  - Store datestamp against reviews as to support 'gin of the week' and such.