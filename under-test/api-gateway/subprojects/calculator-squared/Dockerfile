FROM gradle:6.0.0-jdk8 as builder
COPY . ./
RUN ./gradlew bootJar

FROM openjdk:8-alpine
COPY --from=builder home/gradle/build/libs/calculator-squared-0.0.1-SNAPSHOT.jar squared.jar
EXPOSE 80
ENTRYPOINT ["java", "-jar", "squared.jar"]
