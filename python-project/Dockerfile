FROM python:slim-bullseye

WORKDIR /app

COPY . .

RUN pip install --no-cache-dir -r requirements/backend.in

CMD ["uvicorn", "spaceship.main:app", "--host=0.0.0.0", "--port=8080"]
