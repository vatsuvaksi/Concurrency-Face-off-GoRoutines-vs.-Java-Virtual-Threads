# Start with the official OpenJDK 21 image as the base image
FROM eclipse-temurin:21-jdk

WORKDIR /app

COPY target/Java-Virtual-Thread-1.0.jar /app/Java-Virtual-Thread-1.0.jar

CMD ["java", "-jar", "Java-Virtual-Thread-1.0.jar"]