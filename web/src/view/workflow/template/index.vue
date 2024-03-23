<template>
  <div>
    <div class="gva-search-box">
      <el-form
        ref="searchForm"
        :inline="true"
        :model="searchInfo"
      >
        <el-form-item label="流程名称">
          <el-input
            v-model="searchInfo.flowName"
            placeholder="流程名称"
          />
        </el-form-item>
        <el-form-item label="流程类型">
          <el-input
            v-model="searchInfo.flowType"
            placeholder="流程类型"
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
      <el-button
        type="primary"
        icon="plus"
        @click="handleAdd"
      >新建模版</el-button>
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
          label="流程名称"
          min-width="100"
          prop="flowName"
        />
        <el-table-column
          align="left"
          label="流程描述"
          min-width="100"
          prop="flowDesc"
        />
        <el-table-column
          align="left"
          label="流程创建人"
          min-width="100"
          prop="flowCreator"
        />
        <el-table-column
          align="left"
          label="流程修改人"
          min-width="100"
          prop="flowModifier"
        />
        <el-table-column
          align="left"
          fixed="right"
          label="操作"
          width="200"
        >
          <template #default="scope">
            <el-button
              icon="edit"
              type="primary"
              link
              @click="handleEdit(scope.row)"
            >
              详情
            </el-button>
            <el-button
              icon="delete"
              type="primary"
              link
              @click="handleDelete(scope.row)"
            >
              删除
            </el-button>
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
  getWorkflowTemplateList,
  deleteWorkflowTemplate,
} from '@/api/workflow'

import { useRouter } from 'vue-router'
import { ref } from 'vue'
import { ElMessageBox, ElMessage } from 'element-plus'

// import { ElMessage, ElMessageBox } from 'element-plus'

const router = useRouter()

defineOptions({
  name: 'WorkflowTemplate',
})

// const form = ref({
//   flowName: '',
//   flowDesc: '',
//   flowFormDetail: '',
//   flowDetail: '',
//   flowCreator: '',
//   flowModifier: ''
// })

// const rules = ref({
//   // path: [{ required: true, message: '请输入api路径', trigger: 'blur' }],
//   // apiGroup: [
//   //   { required: true, message: '请输入组名称', trigger: 'blur' }
//   // ],
//   // method: [
//   //   { required: true, message: '请选择请求方式', trigger: 'blur' }
//   // ],
//   // description: [
//   //   { required: true, message: '请输入api介绍', trigger: 'blur' }
//   // ]
// })

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

// // 分页
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
  const table = await getWorkflowTemplateList({ page: page.value, pageSize: pageSize.value, ...searchInfo.value })

  if (table.code === 0) {
    tableData.value = table.data.list
    total.value = table.data.total
    page.value = table.data.page
    pageSize.value = table.data.pageSize
  }
}

getTableData()

const handleAdd = () => {
  router.replace({ name: 'templateDetail' })
}

const handleEdit = (row) => {
  router.replace({
    name: 'templateDetail',
    query: {
      id: row.ID,
      name: row.flowName
    }
  })
}

const handleDelete = (row) => {
  ElMessageBox.confirm('是否删除该工作流模板?', '提示', {
    confirmButtonText: '确定',
    cancelButtonText: '取消',
    type: 'warning'
  })
    .then(async() => {
      const res = await deleteWorkflowTemplate(row)
      if (res.code === 0) {
        ElMessage({
          type: 'success',
          message: '删除成功!'
        })
        if (tableData.value.length === 1 && page.value > 1) {
          page.value--
        }
        getTableData()
      }
    })
}

</script>

  <style scoped lang="scss">
  .warning {
    color: #dc143c;
  }
  </style>
