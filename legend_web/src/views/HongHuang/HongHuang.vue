<template>
  <el-table
    :data="tableData"
    style="width: 100%"
    border
    >
    <el-table-column
       v-for="(val,key) in labelData"
       :key="key"
       :prop="key"
      :label="val"
      width="160">
      <!-- <template slot-scope="scope">
        <span style="margin-left: 10px">{{ scope.row.date }}</span>
      </template> -->
    </el-table-column>

    <el-table-column label="操作">
      <template slot-scope="scope">
        <!-- <el-button
          size="mini"
          type="success"
          :value="scope.row.counttime"
          >{{scope.row.counttime}}</el-button> -->
        <el-button
          size="mini"
          type="info"
          @click="Input(scope.$index, scope.row)">输入时间</el-button>
        <el-button
          size="mini"
          type="primary"
          @click="Refesh(scope.row)">普通刷新</el-button>
          <el-button
          size="mini"
          type="primary"
          @click="HalfRefesh(scope.row)">节日刷新</el-button>
      </template>
    </el-table-column>
  </el-table>
</template>

<script>
  export default {
    data() {
      return {
        tableData: [],
        labelData:{
            id:"编号",
            name:"怪物名字",
            addr:"地址",
            refresh_time:"刷新周期(s)",
            refresh:"刷新时间",
            path:"路线"
        }
      }
    },
    mounted() {
      this.GetData();
    },
    methods: {
      GetData(){
      var self = this;
      this.$http.get('/honghuang/query').then(function (response) {
      self.tableData = response.data["data"]
  }) 
  .catch(function (error) {
    console.log(error);
  });
},
    UpdateData(data){
      this.$http.post('/honghuang/refresh/update',data).then(function (response){
        if (response.data["code"] == 0){
          console.log("请求成功");
        }else{
          console.log("请求失败");
        }
      })
    },
      Input(index, row) {
          this.$prompt("请输入刷新时间(s)","提示",{
            confirmButtonText: '确定',
            cancelButtonText: '取消',
            inputPattern:/^\d+$|^\d+[.]?\d+$/,
            inputErrorMessage: '格式不正确,只能为数字'
          }).then(({ value }) => {
            var timestamp = Date.parse(new Date()) / 1000 + parseInt(value);
            var myTime = new Date(parseInt(timestamp) * 1000).toLocaleString('zh',{hour12:false})
            row.refresh = myTime
            let data = {
          "id":row.id,
          "refresh":myTime
        }
        this.UpdateData(data);
            this.$message({
            type: 'success',
            message: '你输入的时间是: ' + value + "s"
          });
          }).catch(() => {
              this.$message({
            type: 'info',
            message: '取消输入'
          });    
          })
      },
      Refesh(row) {
        var timestamp = Date.parse(new Date()) / 1000 + row.refresh_time;
        var myTime = new Date(parseInt(timestamp) * 1000).toLocaleString('zh',{hour12:false})
        row.refresh = myTime
        let data = {
          "id":row.id,
          "refresh":myTime
        }
        this.UpdateData(data);
        // row.refresh=row.refresh_time;
        // localStorage.setItem("myTime",JSON.stringify(row.refresh));
        // clearInterval(this.interval)
        // this.Time(row)
      },
    HalfRefesh(row){
        var timestamp = Date.parse(new Date()) / 1000 + (row.refresh_time / 2);
        var myTime = new Date(parseInt(timestamp) * 1000).toLocaleString('zh',{hour12:false})
        row.refresh = myTime
        let data = {
          "id":row.id,
          "refresh":myTime
        }
        this.UpdateData(data);
    },
    }
  }
</script>
    <div>
        真是洪荒页面
    </div>
</template>