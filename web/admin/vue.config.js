const { defineConfig } = require('@vue/cli-service')
module.exports = defineConfig({
  transpileDependencies: true,
  lintOnSave: false,
  css: {
    loaderOptions: {
      less: {
        lessOptions: {
          modifyVars: {
            'primary-color': '#1DA57A', // 要重新定义的less变量【无@符】
            'link-color': '#1DA57A',
            'border-radius-base': '2px'
          },
          javascriptEnabled: true // 不可省略必须为true否则不生效！
        }
      }
    }
  }
})
