## Misc Notes - Ubuntu 20.04 + Open JDK 11

Base Java (OpenJDK11) application server build...


- Ubuntu20.04 Image for ARM64, we specified the base os in the dockerfile.
- With some usefull tools added.
- OpenJDK11 application server, docker pulls the ARM64 version of the image based on the base os.


I'm well aware there are pre build OpenJDK11 images available (i.e.: openjdk:11-jdk-slim-buster), but I wanted to build my own for the purpose of the research, learning opportunity...