FROM continuumio/miniconda3:latest

RUN apt-get update && apt-get install -y sshpass && \
    apt-get clean

WORKDIR /app

COPY runner/environment.yml .

RUN conda env create -f environment.yml && conda clean -afy

SHELL ["conda", "run", "-n", "runner-env", "/bin/bash", "-c"]

COPY runner/src/. .

COPY proto/runner.proto protobuf/runner.proto

RUN python -m grpc_tools.protoc -Iproto=protobuf --python_out=. --pyi_out=. --grpc_python_out=. protobuf/runner.proto

EXPOSE 50051

ENTRYPOINT ["conda", "run", "--no-capture-output", "-n", "runner-env", "python", "main.py"]
