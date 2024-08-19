# Docker Image building

## Basic idea.

We're going to show how to build images layer by layer... In the end producing a Apache Hadoop cluster providing HDFS services comprised out of a Namenode, a Datanode and a resourcemanager, nodemanager and a historyserver.

The steps below detail how we start from a fresh Ubunto 20.04 OS image to the end.

Various diagrams can be found in blog-doc/diagrams as well as the Word version of the www.medium.com article.

Each of the directories in the "project" have a local README.md file with more details, in addition to a local Makefile.

The "entire" project/s can be build by executing 

- ```make pullall```
- ```make buildall```

I will also explain a little bit about my docker compose.yaml files and the Makefile's utilised.

From the root of the project, the following commands can be executed:
(see the root Makefile for the commands utilised).

### 1. Pull source images and binaries

    - make pullall          -> pull all the source images and binaries..
    - make pullos           -> Pull only the OS image.
    - make pullsource       -> Pull only the source application (Hadoop) binaries/distribution..

### 2. -> build-ubuntu-os-20.04

    - make buildos          -> use the base OS image and add utilities.

### 3. -> build-ubuntu-os-openjdk11

    - make buildjdk11       -> Use the image build in #2 to build the JDK image.

### 4. -> build-hadoop-openjdk11
    
    - make buildhdfsbase    -> Build the base image used to build the HDFS cluster.
    - make buildhdfs        -> Build the various HDFS nodes.

    - make buildall         -> Do everything above in one go.

### 5. Re openjdk:11.*

    I know there is a prebuild openjdk:11-jdk-slim-buster (and others) images available, but I wanted to build my own for the purpose of the research, learning opportunity... 


### Some random usefull commands:
    
    - docker image instpect <image name:tag>
    - docker image history <image name:tag>
    - docker build --target <stage name> -t <image name:tag> .


### References 

Below are some good links to follow, providing a good initial starting point for docker and docker-compose etc.

- [Docker Crash Course for Absolute Beginners [NEW]](https://www.youtube.com/watch?v=pg19Z8LL06w)
- [Ultimate Docker Compose Tutorial](https://www.youtube.com/watch?v=SXwC9fSwct8)

- [Learning Docker // Getting started! ](https://www.youtube.com/watch?v=Nm1tfmZDqo8)
- [This Docker Compose UI is amazing! // Dockge](https://www.youtube.com/watch?v=HEklvsr7q54)

- [Best Docker Containers for Home Server!](https://www.youtube.com/watch?v=9uF2us2PabM)
- [Best Docker Container Monitoring Tools - Free and open source](https://www.youtube.com/watch?v=zxAmqY63eJE)

- [Multi Stage Docker image build](https://docs.docker.com/build/guide/multi-stage/)
- [Bulti Platform docker image build](https://docs.docker.com/build/building/multi-platform/)

- [Docker ENTRYPOINT vs CMD With Examples - Docker Development Tips & Tricks](https://www.youtube.com/watch?v=U1P7bqVM7xM) 