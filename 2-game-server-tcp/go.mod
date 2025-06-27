module game-server

go 1.24
//require 列出项目依赖的模块及其版本
//exclude 明确排除某个模块的特定版本，避免项目使用该版本
//replace 将一个模块替换为另一个模块，可用于替换本地开发模块或解决依赖冲突
//retract 用于声明某个版本的模块存在问题，建议用户不要使用。一般由模块作者在发布新版本时添加
//indirect