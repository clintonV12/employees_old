FROM alpine
ADD employees /employees
ENTRYPOINT [ "/employees" ]
