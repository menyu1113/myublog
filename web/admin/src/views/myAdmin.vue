<template>
  <a-layout class="container">
    <myNav></myNav>
    <a-layout>
      <a-layout-header class="headerBtn">
        <myHeader></myHeader>
      </a-layout-header>
      <a-layout-content :style="{ margin: '24px 16px 0', height: '100%' }">
        <div
          :style="{ padding: '24px', background: '#fff', minHeight: '360px' }"
        >
          <router-view :key="$route.fullPath"></router-view>
        </div>
        <a-layout-footer style="textalign: center">
          <myFooter></myFooter>
        </a-layout-footer>
      </a-layout-content>
    </a-layout>
  </a-layout>
</template>
<script>
import myFooter from '../components/admin/myFooter'
import myHeader from '../components/admin/myHeader'
import myNav from '../components/admin/myNav'
export default {
  components: { myFooter, myHeader, myNav }
}
</script>

<style scoped>
.container {
  height: 100%;
}
.container .logo {
  height: 32px;
  background: rgba(255, 255, 255, 0.2);
  margin: 16px;
}
.headerBtn {
  display: flex;
  justify-content: flex-end;
  align-items: center;
}
</style>
