FROM amazonlinux:2

LABEL maintainer.name="Afonso Rodrigues"
LABEL maintainer.email=afonsoaugustoventura@gmail.com  

ARG TERRAFORM_VERSION

ENV TERRAFORM_VERSION=0.12.24

RUN yum -y update python

RUN yum -y install python3 \
    python3-pip \
    shadow-utils

RUN adduser ci && \
    yum -y install make \
    unzip \
    wget \
    ruby \
    git \
    tar \
    pip3 install ansible==2.9.7 && \
    yum clean all

RUN amazon-linux-extras install docker -y && \
    usermod -a -G docker ci

RUN curl https://releases.hashicorp.com/terraform/${TERRAFORM_VERSION}/terraform_${TERRAFORM_VERSION}_linux_amd64.zip -o "terraform_${TERRAFORM_VERSION}_linux_amd64.zip"  && \
    unzip terraform_${TERRAFORM_VERSION}_linux_amd64.zip && \
    rm -rf terraform_${TERRAFORM_VERSION}_linux_amd64.zip && \
    chown ci:ci terraform && \
    mv  terraform /bin/terraform

RUN curl "https://awscli.amazonaws.com/awscli-exe-linux-x86_64.zip" -o "awscliv2.zip" && \
    unzip awscliv2.zip && \
    rm -rf awscliv2.zip && \
    bash ./aws/install

RUN curl -sfL https://raw.githubusercontent.com/aquasecurity/trivy/master/contrib/install.sh | sh -s -- -b /usr/local/bin

WORKDIR /home/ci/

USER ci

CMD ["bash"]
