FROM openjdk:11.0.3-jdk-stretch as builder

WORKDIR /root
COPY gradle* ./
COPY build.gradle ./
COPY src ./src
COPY gradle ./gradle
RUN ./gradlew build shadowJar --no-daemon
RUN ls /root/build/libs/

# Runner stage only needs JRE and JAR
FROM openjdk:11.0.3-jre-slim-stretch

EXPOSE 8081
WORKDIR /opt/imageHost
COPY --from=builder /root/build/libs/*-all.jar ./app.jar
CMD ["java", "-server", "-jar", "app.jar"]
