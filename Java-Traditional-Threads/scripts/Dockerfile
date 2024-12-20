FROM python:3.11-slim

WORKDIR /scripts

COPY requirement.txt /scripts/
RUN pip install --no-cache-dir -r requirement.txt

COPY load_data.py /scripts/
CMD ["python", "load_data.py"]
