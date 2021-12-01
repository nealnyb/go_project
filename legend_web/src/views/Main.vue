<template>
 <el-container style="height: 100%">
  <el-aside width="auto"><common-aside></common-aside></el-aside>
  <el-container>
    <el-main><router-view/></el-main>
  </el-container>
</el-container>
</template>

<script>
// @ is an alias to /src

import CommonAside from '../components/CommonAside.vue'
export default {
  name: 'Main',
  components: { CommonAside }
}
</script>

<style lang="scss" scoped>
.el-main {
  padding-top: 0;
}
</style>
