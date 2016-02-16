FROM scratch
MAINTAINER Stig Kleppe-Jørgensen "from.github@nisgits.net"

COPY dist/counter-linux /
COPY static /static
COPY templates /templates


ENV MONGO_URL ""
ENV REDIS_URL ""
ENV POSTGRES_URL ""

ENV PORT 80
EXPOSE 80

CMD ["/counter-linux"]
