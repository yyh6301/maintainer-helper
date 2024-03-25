<template>
  <div>
    <div class="gva-search-box">
      <el-form
        ref="searchForm"
        :inline="true"
        :model="searchInfo"
      >
        <el-form-item
          style="width: 20%;"
          label="云厂商"
        >
          <!-- <el-input
            v-model="searchInfo.CloudType"
            placeholder="云厂商"
          /> -->
          <el-select
            v-model="searchInfo.CloudType"
            placeholder="云厂商"
            clearable
          >
            <el-option
              v-for="item in cloudTypeList"
              :key="item.value"
              :label="item.value"
              :value="item.value"
            />
          </el-select>
        </el-form-item>
        <el-form-item label="实例名称">
          <el-input
            v-model="searchInfo.InstanceName"
            placeholder="实例名称"
          />
        </el-form-item>
        <el-form-item label="公网IP">
          <el-input
            v-model="searchInfo.PublicIP"
            placeholder="公网IP"
          />
        </el-form-item>
        <el-form-item label="私网IP">
          <el-input
            v-model="searchInfo.PrivateIP"
            placeholder="私网IP"
          />
        </el-form-item>
        <el-form-item>
          <el-button
            type="primary"
            icon="search"
            @click="onSubmit"
          >查询</el-button>
          <el-button
            icon="refresh"
            @click="onReset"
          >重置</el-button>
        </el-form-item>
      </el-form>
    </div>
    <div class="gva-table-box">
      <el-table
        :data="tableData"
        @selection-change="handleSelectionChange"
      >
        <el-table-column
          type="selection"
          width="55"
        />
        <el-table-column
          align="left"
          label="id"
          min-width="60"
          prop="ID"
        />

        <el-table-column
          align="left"
          label="实例名称"
          min-width="100"
          prop="instanceName"
        />
        <el-table-column
          align="left"
          label="实例类型"
          min-width="100"
          prop="instanceType"
        />
        <el-table-column
          align="left"
          label="公网IP"
          min-width="100"
          prop="publicIP"
        />
        <el-table-column
          align="left"
          label="私网IP"
          min-width="100"
          prop="privateIP"
        />
        <el-table-column
          align="left"
          label="云厂商"
          min-width="100"
          prop="cloudType"
        />
      </el-table>
      <div class="gva-pagination">
        <el-pagination
          :current-page="page"
          :page-size="pageSize"
          :page-sizes="[10, 30, 50, 100]"
          :total="total"
          layout="total, sizes, prev, pager, next, jumper"
          @current-change="handleCurrentChange"
          @size-change="handleSizeChange"
        />
      </div>
    </div>
  </div>
</template>

<script setup>
import {
  getAssetsList,
} from '@/api/cmdb'
import { ref } from 'vue'
// import { ElMessage, ElMessageBox } from 'element-plus'

defineOptions({
  name: 'Assets',
})

const form = ref({
  CloudType: '',
  InstanceID: '',
  InstanceName: '',
  InstanceType: '',
  PublicIP: '',
  PrivateIP: '',
})

const rules = ref({
  // path: [{ required: true, message: '请输入api路径', trigger: 'blur' }],
  // apiGroup: [
  //   { required: true, message: '请输入组名称', trigger: 'blur' }
  // ],
  // method: [
  //   { required: true, message: '请选择请求方式', trigger: 'blur' }
  // ],
  // description: [
  //   { required: true, message: '请输入api介绍', trigger: 'blur' }
  // ]
})

const cloudTypeList = ref([
  { value: '阿里云' },
  { value: '腾讯云' },
  { value: 'AWS' },
])

const page = ref(1)
const total = ref(0)
const pageSize = ref(10)
const tableData = ref([])
const searchInfo = ref({})

const onReset = () => {
  searchInfo.value = {}
}
// 搜索
const onSubmit = () => {
  page.value = 1
  pageSize.value = 10
  getTableData()
}

// 分页
const handleSizeChange = (val) => {
  pageSize.value = val
  getTableData()
}

const handleCurrentChange = (val) => {
  page.value = val
  getTableData()
}

// 查询
const getTableData = async() => {
  const table = await getAssetsList({ page: page.value, pageSize: pageSize.value, ...searchInfo.value })
  if (table.code === 0) {
    tableData.value = table.data.list
    total.value = table.data.total
    page.value = table.data.page
    pageSize.value = table.data.pageSize
  }
}

getTableData()

</script>

<style scoped lang="scss">
.warning {
  color: #dc143c;
}
</style>
