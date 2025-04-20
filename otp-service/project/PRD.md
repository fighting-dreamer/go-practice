## OTP service

This service generates and sends the OTP to the end user on request from an internal service.
This service also verify the OTP is right if verified within certain time limit.

Design choice :

1. Can OTP be re-sent ?
   1. NO : the OTP is stored such that we can;t just query what is hte OTP directly.
   2. Yes : the OTP is stored directly with a TTL
