<template>
    <el-menu    
    class="el-menu-vertical-demo" 
    :collapse="isCollapse"
    active-text-color="red"
    >
<h3>怪物计时管理系统</h3>
  <el-menu-item :index="item.path" v-for="item in menu" :key="item.path" @click="ClickMenu(item)">
    <i class="el-icon-menu"></i>
    <span slot="title">{{item.label}}</span>
  </el-menu-item>
</el-menu>
</template>

<style>
  .el-menu-vertical-demo:not(.el-menu--collapse) {
    width: 200px;
    min-height: 400px;
  }
</style>

<script>
  export default {
    data() {
      return {
        isCollapse: false,
        menu:[
            {
                name:"home",
                label:"本区",
                path:"/",
                url:"/Home/Home"
            },
            {
                name:"shenlong",
                label:"神龙",
                path:"/shenlong",
                url:"/ShenLong/ShenLong"
            },
            {
                name:"shengyu",
                label:"圣狱",
                path:"/shengyu",
                url:"/ShengYu/ShengYu"
            },
            {
                name:"zhenying",
                label:"阵营",
                path:"/zhenying",
                url:"/ZhenYing/ZhenYing"
            },
            {
                name:"shengxiao",
                label:"生肖",
                path:"/shengxiao",
                url:"/ShengXiao/ShengXiao"
            },
            {
                name:"honghuang",
                label:"洪荒",
                path:"/honghuang",
                url:"/HongHuang/HongHuang"
            },
            {
                name:"humanoid",
                label:"人型",
                path:"/humanoid",
                url:"/Humanoid/Humanoid"
            },
        ]
      };
    },
    methods: {
      ClickMenu(item) {
          this.$router.push({name:item.name})
      }
    }
  }
</script>   