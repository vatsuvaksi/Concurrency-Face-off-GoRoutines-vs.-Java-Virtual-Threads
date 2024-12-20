# Start with the official OpenJDK 17 image as the base image
FROM eclipse-temurin:17-jdk

WORKDIR /app

COPY target/Java-Traditional-Threads-1.0.jar /app/Java-Traditional-Threads-1.0.jar

CMD ["java", "-jar", "Java-Traditional-Threads-1.0.jar"]