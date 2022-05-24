Looks like you're having a CORS issue! CORS is a security measure that your browser has that prevent requests from the wrong domain.
It protects you from requests generated on `https://notYourBank.com/` being sent to and and from `https://actuallyYourBank.com/` from actually working.

CORS is controlled by the `Access-Control-Allow-Origin` header that is __sent by the server__. Each client is individually responsible for respecting and implementing CORS (in practice, this means only browsers do).
Read more here: https://developer.mozilla.org/en-US/docs/Web/HTTP/CORS