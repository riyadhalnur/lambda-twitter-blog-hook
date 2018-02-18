Lambda Twitter Blog Poster  
=================  
> Automagically post to Twitter timeline about new blog post using Go on AWS Lambda.  

**TL;DR** Events are generated using webhooks - Gitlab in this case because that's where I store my site. On specific commit events (the commits contain a [NEW POST] block), a Twitter status is pushed to my timeline.  

### Developing  

#### Requirements  
- [Serverless](https://serverless.com/)  
- [Go](https://golang.org/)  

#### Getting started  
1. Create a `serverless.env.yml` file in the root of the project.  
2. Declare the following variables and values in the file -  
```yaml
secret: "<gitlab-token-header-secret>"
apikey: "<twitter-consumer-api-key>"
apisecret: "<twitter-consumer-secret-key>"
accesstoken: "<twitter-app-access-token>"
accesssecret: "<twitter-app-access-secret>"
baseurl: "<url-of-blog>"
```  

### Deploy  
1. Build the program using `make build` in the root of the project.  
2. To deploy, use `sls deploy -s <env-name>`.
3. Once deployed, copy the API Gateway URL printed on the console.  

### Create a webhook  
*Only GitLab is supported at the moment*  

1. Navigate to the repo for which a webhook will be triggered.
2. From the panel on the left, go to `Settings > Integrations` to create a new webhook.
3. Paste the URL from the last section into the `URL` field and the secret from the `serverless.env.yml` file in to the `Secret Token` field.
4. Check `Push Events` under `Trigger`.
5. Click on `Add webhook` to save.  

### Triggering the function  
1. To trigger a new post to Twitter on new post publishes, commit your new posts in the form -  
```
[NEW POST]
The title of the post
permalink-of-the-post
```  
2. Profit!  

### Contributing  
Read the [CONTRIBUTING](CONTRIBUTING.md) guide for information.  

### License  
Licensed under MIT. See [LICENSE](LICENSE) for more information.  

### Issues  
Report a bug in [issues](https://github.com/riyadhalnur/lambda-twitter-blog-hook/issues).   

Made with love in Kuala Lumpur, Malaysia by [Riyadh Al Nur](https://verticalaxisbd.com)
