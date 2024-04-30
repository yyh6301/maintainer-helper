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

    <el-dialog
      v-model="dialogFormVisible"
      :before-close="closeDialog"
      :title="dialogTitle"
    >

      <el-form
        ref="templateForm"
        :model="form"
        :rules="rules"
        label-width="80px"
      >
        <el-form-item
          label="流程名称"
          prop="flowName"
        >
          <el-input
            v-model="form.flowName"
            autocomplete="off"
            placeholder="请输入流程名称"
          />
        </el-form-item>
        <el-form-item
          label="流程描述"
          prop="flowDesc"
        >
          <el-input
            v-model="form.flowDesc"
            placeholder="请输入流程描述"
          />
        </el-form-item>
      </el-form>
      <template #footer>
        <div class="dialog-footer">
          <el-button @click="closeDialog">取 消</el-button>
          <el-button
            type="primary"
            @click="enterDialog"
          >确 定</el-button>
        </div>
      </template>
    </el-dialog>
  </div>
</template>

<script setup>
import {
  getWorkflowTemplateList,
  createWorkflowTemplate,
  deleteWorkflowTemplate,
} from '@/api/workflow'

import { useRouter } from 'vue-router'
import { ref } from 'vue'
import { useUserStore } from '@/pinia/modules/user'
import { ElMessageBox, ElMessage } from 'element-plus'

const userStore = useUserStore()

// import { ElMessage, ElMessageBox } from 'element-plus'

const rules = ref({
  flowName: [{ required: true, message: '请输入流程名称', trigger: 'blur' }],
})
const router = useRouter()

defineOptions({
  name: 'WorkflowTemplate',
})

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

const dialogTitle = ref('新增状态')
const dialogFormVisible = ref(false)
const templateForm = ref(null)
const form = ref({
  flowName: '',
  flowDesc: '',
  flowFormDetail: '',
  flowDetail: '',
  flowCreator: '',
  flowModifier: ''
})

const handleAdd = () => {
  // router.replace({ name: 'templateDetail' })
  dialogFormVisible.value = true
}

const closeDialog = () => {
  dialogFormVisible.value = false
  templateForm.value.resetFields()
  form.value = {
    flowName: '',
    flowDesc: '',
    flowFormDetail: '',
    flowDetail: '',
    flowCreator: '',
    flowModifier: ''
  }
}

const enterDialog = async() => {
  templateForm.value.validate(async valid => {
    if (valid) {
      // 新增工作流
      form.value.flowCreator = userStore.userInfo.userName
      form.value.flowModifier = userStore.userInfo.userName

      const res = await createWorkflowTemplate(form.value)
      if (res.code === 0) {
        ElMessage.success('新增工作流模板成功')
      }
      getTableData()
      closeDialog()
    }
  })
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
