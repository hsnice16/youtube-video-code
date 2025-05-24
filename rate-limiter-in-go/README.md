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

## Token Bucket Algorithm

In Token Bucket Algorithm, there are tokens in a bucket (bucket gets created for each unique user), and there is refill rate at which new tokens get added in the bucket, if the bucket is already full, the new tokens just get drop.

Now, whenever a user makes a request, we take one token out from the bucket for each request, and if the bucket is empty, we can decide if we want to put the request in the queue or want to drop the request.

In our implementation, we will drop the request, and send the error message.









Let's code the `Token Bucket`.

Let's end the video here, and connect in the next one.

Bye 👋