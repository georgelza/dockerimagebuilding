FROM georgelza/ubuntu20.04-openjdk11:1.0.1

USER root

RUN echo "--> set environment variable"
ENV JAVA_HOME=/usr/lib/jvm/java-11-openjdk-arm64
ENV HADOOP_VERSION=3.3.5
ENV HADOOP_HOME=/opt/hadoop-3.3.5
ENV HADOOP_CONF_DIR=/etc/hadoop
ENV MULTIHOMED_NETWORK=1
ENV PATH=$HADOOP_HOME/bin/:$PATH

WORKDIR /tmp/

COPY stage/hadoop-3.3.5-aarch64.tar.gz /tmp/

RUN echo "--> install Apache Hadoop 3.3.5" && \
    tar -xzvf /tmp/hadoop-3.3.5-aarch64.tar.gz && \
    mv /tmp/hadoop-3.3.5 /opt/ && \
    rm /tmp/hadoop-*

    # Intel/AMD64 version 
    # https://archive.apache.org/dist/hadoop/common/hadoop-3.3.5/hadoop-3.3.5.tar.gz

RUN ln -s $HADOOP_HOME/etc/hadoop $HADOOP_CONF_DIR


RUN echo "--> make Apache Hadoop directories" && \
    mkdir -p $HADOOP_HOME/logs && \
    mkdir /hadoop-data

ADD bin/entrypoint.sh /entrypoint.sh
RUN chmod a+x /entrypoint.sh

ENTRYPOINT ["/entrypoint.sh"]