FROM alpine
WORKDIR "/opt"
ADD migrations migrations
ADD ./pay pay
RUN chmod +x /opt/pay

ENTRYPOINT ["/opt/pay"]