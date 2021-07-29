# URL-SHORTNER
url-shortner is a golang based backend for the app which takes long URL as input from the user and generates the short URL for it. Short URL is generated by hashing the long URL.

## PREREQUISITE
You should have the docker environment if you want to deploy the app using docker image. Alternatively if you want to build binary using the source repo, then go should be installed on your system.

## Setting up Application
You can deploy application in two ways :

### 1. Docker based deployment

* Pull the docker image
```shell
docker pull itsmeashish/url-shortner:latest
```
* use the docker run command to run the image
```shell
 docker run -p 9999:9999 itsmeashish/url-shortner:latest
```
* After this your docker container will be up and running.You can start using application now at port:9999

### 2. Build executable from source code

* Clone the repository in your GOPATH or use go get command 
```shell
go get -u github.com/ash3798/url-shortner
```
* Run the build command at root of the project to create the executable of the application
```
go build
```
* Once created, you can run the executable to start using your application. 

## USAGE GUIDE
> Application runs by default on the port number : 9999
**Note : Server is HTTP. So your requests should use http protocol.**
```shell
http:://<hostname>:<port>/<path>
```
### API available
1. **Generate Short URL** : User can use this API to generate the short-URL for given long-URL
```json
POST  /url

Body :   {
           "url" : "https://www.testurl.com"
         }
         
Response : {
             "url": "https://www.testurl.com",
             "short_url": "https://8da3432a12b13e"
           }
```
> * Response consists of the original-URL as well as the generated short-URL for it.
> * If same URL is passed again and again, the URL that was generated first time will be returned for all these requests.

2. **Get Short URL** : User can use this api to view the short URL created previously for the long URL
```json
GET  /url

Body :   {
           "url" : "https://www.testurl.com"
         }
         
Response : {
             "url": "https://www.testurl.com",
             "short_url": "https://8da3432a12b13e"
           }
```
> * Response consists of the original-URL as well as the generated short-URL for it.
> * If user passes the URL for whom the shortURL has not been generated previously, error will be returned.
