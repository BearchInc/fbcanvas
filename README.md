# Lets Change The World

# Running the App

1. `$ npm install -g bower`
2. `$ npm install http-serve`
3. `$ bower install`
4. `http-serve ./public [-p 8000 -a 0.0.0.0]`

# About Polymer

1. [What is Polymer?](https://www.polymer-project.org/1.0/docs/start/what-is-polymer.html)
2. [Polymer Components Catalog](https://elements.polymer-project.org)
3. [Polymer Articles](https://www.polymer-project.org/1.0/articles)

# Deploy to Prod notes

1. It's necessary to use Juliana Ferrari Paypal app, change the API Credentials in file hellofbcanvas.go to this one:
Client ID: ATRub8NK5m1iZV1EFPcs2Ad_lcKx6A7yasQaRSj6wdjKEDPBpzZ1UZBUr4qQtxg45fG-zO8OlZ85fJx4
Secret: EF4fNq7M9l_VztubdFCLsTsUnqGAoSj12WTnGWuguyQKisAC2aneCVNuXDAusmwE5EjDit67YYTMev3z
2. Change the ExperienceProfileID when calling paypal api to XP-H5BE-78MM-5XMU-LZDX. This web-profile was created for Juliana live app already (info: https://developer.paypal.com/docs/api/#payment-experience)
3. Change ReturnURL and CancelURL to production values in hellofbcanvas.go
4. Follow the normal deploy process