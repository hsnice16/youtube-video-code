# What is a Rate Limiter?

If you are into Software Engineering from sometime, then you might have heard this term.

A `Rate Limiter` helps in maintaining the stability of our systems (or, servers). It ensures the resourcers are not exploited by a single user.

**What could happen is** that a user might send a number of requests in bulk to our servers with the intention to bring it down, so the Rate Limiter helps to avoid that.

## How does Rate Limiter help us?

There are different algorithms that we can use to implement Rate Limiter. And, Rate Limiter can be used to secure any service/resource, but in our example we will use it to secure a HTTP server.

Different algorithms that exist:

- Token Bucket Algorithm
- Leaky Bucket Algorithm
- ... (and there are more that you can search on Google)
