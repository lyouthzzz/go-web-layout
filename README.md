工程目录example


### [目录文章设计](https://younman.com/2022/07/09/%E5%86%8D%E8%B0%88Go%E5%B7%A5%E7%A8%8B%E7%9B%AE%E5%BD%95/)

### build
```bash
docker build --build-arg APP_NAME=go-web-layout --build-arg PROJECT_NAME=go-web-layout  -f Dockerfile -t weblayout .
```

### run
```bash
docker run --rm weblayout
```
