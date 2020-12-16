FROM python:3.6-alpine

WORKDIR /opt/factorial/
COPY app app/
COPY requirements.txt .
RUN pip3 install -r requirements.txt
EXPOSE 80
ENTRYPOINT ["python3"]
CMD ["app/main/app.py"]