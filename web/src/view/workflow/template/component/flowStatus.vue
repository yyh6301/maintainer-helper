<template>
  <div>
    <div class="gva-search-box">
      <el-form
        ref="searchsubForm"
        :inline="true"
        :model="searchInfo"
      >
        <el-form-item label="状态名称">
          <el-input
            v-model="searchInfo.statusName"
            placeholder="流程名称"
          />
        </el-form-item>
        <el-form-item label="审批人">
          <el-input
            v-model="searchInfo.approvalUser"
            placeholder="审批人"
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
        @click="openDialog('addStatus')"
      >新建状态</el-button>
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
          label="状态ID"
          min-width="60"
          prop="ID"
        />

        <el-table-column
          align="left"
          label="状态名称"
          min-width="100"
          prop="statusName"
        />
        <el-table-column
          align="left"
          label="状态类型"
          min-width="100"
        >
          <template #default="{ row }">
            <el-tag :type="getStatusType(row.statusType)">{{ getStatusLabel(row.statusType) }}</el-tag>
          </template>

        </el-table-column>

        <el-table-column
          align="left"
          label="审批类型"
          min-width="100"
        >
          <template #default="{ row }">
            <el-tag :type="getApprovalType(row.approvalType)">{{ getApprovalLabel(row.approvalType) }}</el-tag>
          </template>
        </el-table-column>

        <el-table-column
          align="left"
          label="审批人"
          min-width="100"
          prop="approvalUser"
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
              编辑
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
        ref="statusForm"
        :model="statusForm"
        :rules="rules"
        label-width="80px"
      >
        <el-form-item
          label="状态名称"
          prop="statusName"
        >
          <el-input
            v-model="form.statusName"
            autocomplete="off"
          />
        </el-form-item>
        <el-form-item
          label="状态类型"
          prop="statusType"
        >
          <el-select
            v-model="form.statusType"
            placeholder="请选择"
            style="width:100%"
          >
            <el-option
              v-for="item in statusTypeOptions"
              :key="item.value"
              :label="`${item.label}`"
              :value="item.value"
            />
          </el-select>
        </el-form-item>
        <el-form-item
          label="排序"
          prop="orderNumber"
          style="width: 100%;"
        >
          <el-input-number
            v-model="form.orderNumber"
            autocomplete="off"
          />
        </el-form-item>

        <el-form-item
          label="审批类型"
          prop="approvalType"
        >
          <el-select
            v-model="form.approvalType"
            placeholder="请选择"
            style="width:100%"
          >
            <el-option
              v-for="item in approvalTypeOptions"
              :key="item.value"
              :label="`${item.label}`"
              :value="item.value"
            />
          </el-select>
        </el-form-item>
        <el-form-item
          label="审批人"
          prop="approvalUser"
        >
          <el-input
            v-model="form.approvalUser"
            autocomplete="off"
          />
        </el-form-item>
        {{ form }}
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
import { ref, defineProps, toRefs } from 'vue'
import { ElMessage, ElMessageBox } from 'element-plus'
import {
  createTemplateStatus,
  getTemplateStatusById,
  updateTemplateStatus,
  deleteTemplateStatus,
  getTemplateStatusList,
} from '@/api/workflow'
defineOptions({
  name: 'FlowStatus'
})

const props = defineProps({
  // 子组件接收父组件传递过来的值
  templateid: {
    type: String,
    default: ''
  },
})
// 使用父组件传递过来的值
const { templateid } = toRefs(props)

const tableData = ref([])
const page = ref(1)
const total = ref(0)
const pageSize = ref(10)
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
  searchInfo.value.templateID = Number(templateid.value)
  const table = await getTemplateStatusList({ page: page.value, pageSize: pageSize.value, ...searchInfo.value })

  if (table.code === 0) {
    tableData.value = table.data.list
    total.value = table.data.total
    page.value = table.data.page
    pageSize.value = table.data.pageSize
  }
}

getTableData()

const handleDelete = (row) => {
  ElMessageBox.confirm('是否删除该状态?', '提示', {
    confirmButtonText: '确定',
    cancelButtonText: '取消',
    type: 'warning'
  })
    .then(async() => {
      const res = await deleteTemplateStatus(row)
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

// 弹窗相关
const form = ref({
  statusName: '',
  statusType: '',
  approvalType: '',
  approvalUser: '',
  templateID: Number(templateid.value),
  orderNumber: 0
})

const type = ref('')
const statusForm = ref(null)

const initForm = () => {
  statusForm.value.resetFields()
  form.value = {
    statusName: '',
    statusType: '',
    approvalType: '',
    approvalUser: '',
    templateID: Number(templateid.value),
    orderNumber: 0
  }
}

const approvalTypeOptions = ref([
  {
    value: 0,
    label: '无',
    color: 'info' // 设置颜色
  },
  {
    value: 1,
    label: '审批人审批',
    color: 'primary' // 设置颜色
  },
  {
    value: 2,
    label: 'hook调用',
    color: 'primary' // 设置颜色
  }
])

const statusTypeOptions = ref([
  {
    value: 0,
    label: '开始节点',
    color: 'info' // 设置颜色
  },
  {
    value: 1,
    label: '中间节点',
    color: 'primary' // 设置颜色
  },
  {
    value: 2,
    label: '结束节点',
    color: 'danger' // 设置颜色
  }
])

const getStatusType = (type) => {
  const item = statusTypeOptions.value.find(item => item.value === type)
  return item ? item.color : ''
}

const getApprovalType = (type) => {
  const item = approvalTypeOptions.value.find(item => item.value === type)
  return item ? item.color : ''
}

const getStatusLabel = (type) => {
  const item = statusTypeOptions.value.find(item => item.value === type)
  return item ? item.label : ''
}

const getApprovalLabel = (type) => {
  const item = approvalTypeOptions.value.find(item => item.value === type)
  return item ? item.label : ''
}

const dialogTitle = ref('新增状态')
const dialogFormVisible = ref(false)
const openDialog = (key) => {
  switch (key) {
    case 'addStatus':
      dialogTitle.value = '新增状态'
      break
    case 'editStatus':
      dialogTitle.value = '编辑状态'
      break
    default:
      break
  }
  type.value = key
  dialogFormVisible.value = true
}
const closeDialog = () => {
  initForm()
  dialogFormVisible.value = false
}

const handleEdit = async(row) => {
  const res = await getTemplateStatusById({ id: row.ID })
  form.value = res.data
  openDialog('editStatus')
}

const enterDialog = async() => {
  statusForm.value.validate(async valid => {
    if (valid) {
      switch (type.value) {
        case 'addStatus':
          {
            searchInfo.value.templateID = Number(templateid.value)
            const res = await createTemplateStatus(form.value)
            if (res.code === 0) {
              ElMessage({
                type: 'success',
                message: '添加成功',
                showClose: true
              })
            }
            getTableData()
            closeDialog()
          }

          break
        case 'editStatus':
          {
            const res = await updateTemplateStatus(form.value)
            if (res.code === 0) {
              ElMessage({
                type: 'success',
                message: '编辑成功',
                showClose: true
              })
            }
            getTableData()
            closeDialog()
          }
          break
        default:
          // eslint-disable-next-line no-lone-blocks
          {
            ElMessage({
              type: 'error',
              message: '未知操作',
              showClose: true
            })
          }
          break
      }
    }
  })
}

// 定义一个函数，用于批量转换对象的数字属性为数字类型
// const convertNumberProps = (obj) => {
//   // 遍历对象的属性
//   for (const key in obj) {
//     // 检查属性值是否可以转换为数字
//     if (!isNaN(parseFloat(obj[key])) && isFinite(obj[key])) {
//       // 将属性值转换为数字类型
//       obj[key] = parseFloat(obj[key])
//     }
//   }
// }

</script>

