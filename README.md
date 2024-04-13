
## Libraries
- 使用air进行热更新和启动
```
air
```

- 使用cobra构建应用，执行可执行文件时可以加参数选项
```
make
./_output/miniblog
/_output/miniblog -h
/_output/miniblog test
```

- 使用viper读取配置文件，并配置

- 基于zap定制日志包

- 使用verflag添加版本号
```
$ make
$ _output/miniblog --version -c configs/miniblog.yaml
  gitVersion: f9bcff9
   gitCommit: f9bcff93cc300b8e00de1789d89eb5ffd2b13dfe
gitTreeState: clean
   buildDate: 2022-12-12T11:59:45Z
   goVersion: go1.19.4
    compiler: gc
    platform: linux/amd64
$ _output/miniblog --version=raw -c configs/miniblog.yaml
version.Info{GitVersion:"f9bcff9", GitCommit:"f9bcff93cc300b8e00de1789d89eb5ffd2b13dfe", GitTreeState:"clean", BuildDate:"2022-12-12T11:59:45Z", GoVersion:"go1.19.4", Compiler:"gc", Platform:"linux/amd64"}
```

- 使用gin构建简单web服务


## 日志规范

### 日志规范

- 日志包统一使用 `github.com/marmotedu/miniblog/internal/pkg/log`;
- 使用结构化的日志打印格式：`log.Infow`, `log.Warnw`, `log.Errorw` 等; 例如：`log.Infow("Update post function called")`;
- 日志均以大写开头，结尾不跟 `.`，例如：`log.Infow("Update post function called")`;
- 使用过去时，例如：`Could not delete B` 而不是 `Cannot delete B`;
- 遵循日志级别规范：
  - Debug 级别的日志使用 `log.Debugw`;
  - Info 级别的日志使用 `log.Infow`;
  - Warning 级别的日志使用 `log.Warnw`;
  - Error 级别的日志使用 `log.Errorw`;
  - Panic 级别的日志使用 `log.Panicw`;
  - Fatal 级别的日志使用 `log.Fatalw`.
- 日志设置：
  - 开发测试环境：日志级别设置为 `debug`、日志格式可根据需要设置为 `console` / `json`、开启 caller；
  - 生产环境：日志级别设置为 `info`、日志格式设置为 `json`、开启 caller。（注意：上线初期，为了方便现网排障，日志级别可以设置为 `debug`）
- 在记录日志时，不要输出一些敏感信息，例如密码、密钥等。
- 如果在具有 `context.Context` 参数的函数/方法中，调用日志函数，建议使用 `log.L(ctx).Infow()` 进行日志记录。
