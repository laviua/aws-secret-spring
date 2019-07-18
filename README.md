# aws-secret-spring
AWS Secrets to Spring config

We want to use AWS Secrets in the Spring Boot applications.
As an option we can use External configurations:
https://docs.spring.io/spring-boot/docs/current/reference/html/boot-features-external-config.html

Before you need to store AWS credentials

    AWS_REGION=eu-central-1;
    AWS_ACCESS_KEY_ID=AKIAEXAMPLE;
    AWS_SECRET_ACCESS_KEY=dr6uMO2EXAMPLE
    
Then run application with secrets with optional parameter: app

Secrets will be filtered by the tag - app
   
    #!/bin/bash
    export SPRING_APPLICATION_JSON=`/opt/secrets/aws-secret-spring <app>`
    java -Dspring.application.json=$SPRING_APPLICATION_JSON -jar myapp.jar



It should write secrets into the json and run application with the externalized secret, because storing secrets into the system variable is not safe.

