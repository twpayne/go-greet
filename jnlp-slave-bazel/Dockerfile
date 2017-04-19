FROM jenkinsci/jnlp-slave:3.7-1
USER root
RUN echo "deb [arch=amd64] http://storage.googleapis.com/bazel-apt stable jdk1.8" > /etc/apt/sources.list.d/bazel.list
RUN curl -s https://bazel.build/bazel-release.pub.gpg | apt-key add -
RUN apt-get update
RUN apt-get install -y bazel
USER jenkins
