# go-to-authnz

Inspired by

* OpenIDC - https://github.com/coreos/dex
* Docker Auth - https://github.com/cesanta/docker_auth
* OAuth2 & OpenIDC - https://github.com/ory/hydra  https://github.com/ory/hydra-consent-app-go
* Docker Distribution - https://github.com/docker/distribution
* Gitlab - https://gitlab.com/gitlab-org/gitlab-ce

### How to

Go install
```
fanhonglingdeMacBook-Pro:go-to-authnz fanhongling$ go install -v ./
###sippets###
github.com/tangfeixiong/go-to-authnz/cmd
github.com/tangfeixiong/go-to-authnz
```

Help
```
fanhonglingdeMacBook-Pro:go-to-authnz fanhongling$ go-to-authnz 
Usage:
  go-to-docker [command]

Available Commands:
  cesantadas  
  dex         
  hydra       Hydra is a cloud native high throughput OAuth2 and OpenID Connect provider

Flags:
  -h, --help   help for go-to-docker

Use "go-to-docker [command] --help" for more information about a command.
```

### Example

Dex

Hydra
```
http://localhost:4445/consent?challenge=eyJhbGciOiJSUzI1NiIsInR5cCI6IkpXVCJ9.eyJhdWQiOiJkZW1vIiwiZXhwIjoxNTAxNjIwMDI5LCJqdGkiOiIyMDVmMDJiNy05YjY5LTRlZDAtOTUxOC1kOGYyZGEyZDg0NTUiLCJyZWRpciI6Imh0dHA6Ly9sb2NhbGhvc3Q6NDQ0NC9vYXV0aDIvYXV0aD9jbGllbnRfaWQ9ZGVtb1x1MDAyNnJlZGlyZWN0X3VyaT1odHRwJTNBJTJGJTJGbG9jYWxob3N0JTNBNDQ0NSUyRmNhbGxiYWNrXHUwMDI2cmVzcG9uc2VfdHlwZT1jb2RlXHUwMDI2c2NvcGU9b2ZmbGluZStvcGVuaWRcdTAwMjZzdGF0ZT1kZW1vc3RhdGVkZW1vc3RhdGVkZW1vXHUwMDI2bm9uY2U9ZGVtb3N0YXRlZGVtb3N0YXRlZGVtbyIsInNjcCI6WyJvZmZsaW5lIiwib3BlbmlkIl19.sGjS4I_rpxtRUy1TgVW_ikWjjwLwMrL-OdbtsIi0Wffz-QHODVICJX0Zih73zOiJcGLnW24x_LKEE-7I6rpptAxwwNBRAPj9MAXo_eEJugi5d4Non-T1PnQS0WJXYwiAR61mfVqv8pXKXNYrNRMbN0p81WF0Rm_rqnP7jPwzfMj7hiq94rM8ydNLVw7rBL6oMUT_seZVqLBymtmBu0k2-OY6-27VVwWhAfqH288TBi0fCoCaBLtoQ7nIiV24ErVF5ImUuSN0_GjOlnUufSSXJEOrG-RZ6LAIl7PCWFd5R7eB7fuXAoFL3KEsbQPZkwjIPoFJrLnikqqX5scwJNSgomD77XxJU12IPA6mxT1X--kfE_YZqglCmyKGNyApdHARssGqsZFoxb2BTDYFAJ29W3j_qOJgqK0FJFhVw8MVJHIWwJbfcUUVdyiAzU2UMXMdygQgb_1dFueQeKEZ9x8-41qeQ4919nG3Q_F1846QGDKqzueu-fpH2Sh9ylicNLhRUiy_ib742_c8J8cQ3Rm3jSBUpdCrAgcySF1bCJpXtTDf2diKwWNL3yUICGrY7QxBXI40411I85UbE4J63hVjBBSb6NGVESZI8BqeHuGH5WZREI-vI7k4GOUMfJaniFwSy6nueTq9BKsz7BsgxdCUQBGfgt2-keoTq3Ylz2jfdd4

http://localhost:4445/callback?code=BC4s-THxeIUPEMWcznMLBITy8sbViiT1EDgBRxjAwvw.otaxIAu5qSEtje-GTb6kgHYK_tfIFAn5jy31xi8Gi2M&scope=openid&state=demostatedemostatedemo
```

Cesenta docker_auth auth_server example


