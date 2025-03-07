FROM python:3.12-alpine

WORKDIR /app

ADD https://github.com/ufoscout/docker-compose-wait/releases/download/2.12.1/wait /
RUN chmod +x /wait

COPY requirements.txt .
RUN pip install --no-cache-dir -r requirements.txt

COPY . .

CMD ["sh", "-c", "/wait && uvicorn fastapi_server:app --host 0.0.0.0 --reload --port 8000"]
