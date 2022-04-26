import Vue from 'vue'
import message from 'ant-design-vue'

message.config({
  top: '100px',
  duration: 0,
  maxCount: 3
})
Vue.prototype.$message = message
