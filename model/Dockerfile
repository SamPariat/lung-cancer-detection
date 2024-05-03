ARG PYTHON_VERSION=3.10.11
FROM python:${PYTHON_VERSION}-slim as base

WORKDIR /app

COPY requirements.txt .

COPY . .

EXPOSE 1234

CMD [ "python", "__init__.py" ]
