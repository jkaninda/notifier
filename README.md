# Notifier
Email, Telegram, and Mattermost Notifier

This tool was developed for personal needs.
I use the binary in my Docker Toolkit image to run on Gitlab jobs or Jenkins CI/CD.
This allows me to send notifications via email, Telegram channel, or Mattermost.

I wanted to share it with you.

## Build
```sh
go build .
```

## Send a mail

```sh
./notifier sendMail --body "Your message" --subject "Your subject"
```
## Attach a file to the mail

```sh
./notifier sendMail --body "Your message" --subject "Your subject" --attach photo.jpg
```

## Send Telegram Message
```sh
notifier sendMessage --message "Your message"
```
## Run on Docker
```sh
docker run --env-file your-env-file jkaninda/notifier sendMessage --message "Your message" 
```

`.env`
```conf
## Mail
MAIL_HOST=localhost
MAIL_PORT=1025
MAIL_USERNAME=jkaninda@mailhog.local
MAIL_PASSWORD=password
MAIL_TO=test@mailhog.local
MAIL_BODY="Lorem Ipsum is simply dummy text of the printing and typesetting industry."
MAIL_SUBJECT="Greetings"
MAIL_FROM=jkaninda@mailhog.local

## Telegram
TG_TOKEN=
TG_CHAT_ID=
TG_MESSAGE="Lorem Ipsum is simply dummy text of the printing and typesetting industry."
```
## Gitlab CI

```yaml
variables:
  CI_REGISTRY: ""
  CI_REGISTRY_USER: ""
  #CI_REGISTRY_PASSWORD: "" ## From Gitlab secret, don't put the password here!!!
  CI_DOCKER_FILE: Dockerfile
  IMAGE_NAME: jkaninda/sample
  IMAGE_TAG: latest
  # Telegram
  TG_TOKEN: "xxxx"
  TG_CHAT_ID: "xxxx"
  # Mail
  MAIL_HOST: localhost
  MAIL_PORT: 1025
  MAIL_USERNAME: jkaninda@mailhog.local
  MAIL_PASSWORD: password
  MAIL_TO: test@mailhog.local
  MAIL_BODY: "Lorem Ipsum is simply dummy text of the printing and typesetting industry."
  MAIL_SUBJECT: "Gitlab CI Build Completed"
  MAIL_FROM: jkaninda@mailhog.local
  PROJECT_NAME: "Sample API"
build:
  image: jkaninda/toolkit
  stage: build
  only:
    - main
  # artifacts:
  #   paths:
  #    - VERSION
  script:
    - >
      TAG=v1.0-${CI_COMMIT_SHORT_SHA}
    # Registry login
    - docker login $CI_REGISTRY -u $CI_REGISTRY_USER -p $CI_REGISTRY_PASSWORD
    # Docker build and Push
    - docker build -f $CI_DOCKER_FILE  -t $IMAGE_NAME:$TAG .
    - docker tag $IMAGE_NAME:$TAG $IMAGE_NAME:$IMAGE_TAG
    - docker push $IMAGE_NAME:$TAG
    - docker push $IMAGE_NAME:$IMAGE_TAG
    - echo "Build and push completed, new tag is $TAG"
    ## Send Notification
    - |
      TG_MESSAGE="
      [INFO] --------------------------------------
      [INFO]   BUILD COMPLETED
      [INFO] --------------------------------------
      Repository: REPOSITORY_URL
      Build and Push Completed
      New Image Tag: ${IMAGE_NAME}:${TAG}
      State: Completed
      You will get an email when the application will start running a new version of deployments.
      Note: Full update may take 5-10 minutes, please be patient!
      #CI/CD"
    - echo "$TG_MESSAGE"  # For debugging purposes to see the message
    - notifier sendMessage --message "$TG_MESSAGE"
    - echo "========================= JOB COMPLETED ================="
  tags:
    - docker
test:
  image: jkaninda/toolkit
  stage: test
  only:
    - main
  script:
    - >
      TAG=v1.0-${CI_COMMIT_SHORT_SHA}
    # Test
    - echo "Start API Test...."
    - echo "Test Completed"
    ## Send Email
    - BUILD_DATE=$(date)
    - |
      MESSAGE="<p>Hi Team,</p>
      <p>Great news! The latest build for Project <strong> ${PROJECT_NAME} </strong> has successfully passed all tests.</p>
      <h3>Build Details:</h3>
      <ul>
      <li><strong>Project:</strong> ${PROJECT_NAME}</li>
      <li><strong>Branch:</strong> main</li>
      <li><strong>Build Number:</strong> ${CI_PIPELINE_ID}</li>
      <li><strong>Commit:</strong> ${CI_COMMIT_SHA}</li>
      <li><strong>Build Time:</strong> ${BUILD_DATE}</li>
      </ul>
      <h3>Test Summary:</h3>
      <ul>
      <li><strong>Total Tests:</strong> 150</li>
      <li><strong>Passed:</strong> 150</li>
      <li><strong>Failed:</strong> 0</li>
      <li><strong>Skipped:</strong> 0</li>
      </ul>
      
      <p>You can view the detailed build report <a href="#">here</a>.</p>
      <p>Keep up the great work!</p>
      <p>Best,</p>
      <p>CI/CD Pipeline</p>
      "
    - notifier sendMail --message "$MESSAGE" --subject "Gitlab CI Build Completed"
    -  echo "================= API TEST COMPLETED ======================================="
    - echo "========================= JOB COMPLETED ================="
  tags:
    - docker
```



👏👏👏 Enjoy! Welcome star if you like it😄