Code Challenge Submission
=====

The following are issues that I would fix/implement if this were for production
* I would have liked to add 100% test coverage given the time
* I'd add a proper logger that integrates with AWS/GCP etc
* I'd obviously have more security (DB password, tokens, authentication on making 
sure client not requesting data for another customer etc)
* I have a bug with the MongoDB client where if passed a context, it cancels the 
context prematurely and does not continue with the db request. As a quick fix I 
passed nil instead of the context passed from the request, however it may be 
leaking connections as a result but requires further debugging to resolve.
* I'd add more validation on user input