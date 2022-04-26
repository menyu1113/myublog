import Vue from 'vue'
import axios from 'axios'
import VueAxios from 'vue-axios'

const Url = 'http://127.0.0.1:9000/api/v1'

axios.defaults.baseURL = Url
Vue.use(VueAxios, axios)

axios.interceptors.request.use(config => {
  config.headers.Authorization = `Bearer ${window.sessionStorage.getItem('token')}`
  return config
})

export { Url }
