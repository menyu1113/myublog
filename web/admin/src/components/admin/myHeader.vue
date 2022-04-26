<template>
  <!-- <a-layout-header :style="{ background: '#fff', padding: 0 }" /> -->
  <div>
    <a-button type="danger" @click="loginOut">退出</a-button>
  </div>
</template>

<script>
export default {
  methods: {
    loginOut () {
      window.sessionStorage.clear('token')
      this.$router.push('/login')
    }
  }
}
</script>
