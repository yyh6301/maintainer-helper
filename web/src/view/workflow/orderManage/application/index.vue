<template>
  <div>
    <div class="gva-search-box">
      <el-form
        ref="searchForm"
        :inline="true"
        :model="searchInfo"
      >
        <el-form-item label="工单标题">
          <el-input
            v-model="searchInfo.title"
            placeholder="工单标题"
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
          label="工单ID"
          min-width="60"
          prop="ID"
        />

        <el-table-column
          align="left"
          label="工单标题"
          min-width="100"
          prop="title"
        />
        <el-table-column
          align="left"
          label="流程名"
          min-width="100"
        >
          <template #default="{ row }">
            <el-tag>{{ row.workFlowTemplate.flowName }}</el-tag>
          </template>
          <!-- <template slot-scope="scope">
            <el-tag>
              {{ scope.row.workFlowTemplate.flowName }}
            </el-tag>
          </template> -->
        </el-table-column>
        <el-table-column
          align="left"
          label="当前状态"
          min-width="100"
        >
          <template #default="{ row }">
            <el-tag>{{ row.workFlowStatus.statusName }}</el-tag>
          </template>
        </el-table-column>

        <el-table-column
          align="left"
          label="创建人"
          min-width="100"
          prop="orderCreator"
        />
        <el-table-column
          align="left"
          label="创建时间"
          min-width="100"
          prop="CreatedAt"
        >
          <template #default="{ row }">
            {{ beautifyTime(row.CreatedAt) }}
          </template>
        </el-table-column>
        <el-table-column
          align="left"
          label="操作"
          min-width="100"
        >
          <template #default="{ row }">
            <el-button
              type="text"
              @click="handleDetail(row)"
            >详情</el-button>
          </template>
        </el-table-column>
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
  getOrderList
} from '@/api/workflow'
import { ref } from 'vue'
import { useRouter } from 'vue-router'
import { useUserStore } from '@/pinia/modules/user'

// import { ElMessage, ElMessageBox } from 'element-plus'

const router = useRouter()
const userStore = useUserStore()
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

const handleDetail = (row) => {
  console.log('row', row)
  router.push({
    name: 'orderDetail',
    query: {
      orderName: row.title,
      orderId: row.ID
    }
  })
}

// 查询
const getTableData = async() => {
  searchInfo.value.application = userStore.userInfo.userName
  const table = await getOrderList({ page: page.value, pageSize: pageSize.value, ...searchInfo.value })
  if (table.code === 0) {
    tableData.value = table.data.list
    total.value = table.data.total
    page.value = table.data.page
    pageSize.value = table.data.pageSize
  }
}

getTableData()

const beautifyTime = (timeString) => {
  // 创建 Date 对象并解析时间字符串
  const date = new Date(timeString)

  // 获取年月日时分秒
  const year = date.getFullYear()
  const month = (date.getMonth() + 1).toString().padStart(2, '0')
  const day = date.getDate().toString().padStart(2, '0')
  const hours = date.getHours().toString().padStart(2, '0')
  const minutes = date.getMinutes().toString().padStart(2, '0')
  const seconds = date.getSeconds().toString().padStart(2, '0')

  // 格式化时间字符串
  const formattedTime = `${year}-${month}-${day} ${hours}:${minutes}:${seconds}`

  return formattedTime
}

</script>

<style scoped lang="scss">
.warning {
  color: #dc143c;
}
</style>
