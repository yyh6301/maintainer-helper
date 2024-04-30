<!-- eslint-disable vue/prop-name-casing -->
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
            v-model="searchInfo.circleName"
            placeholder="流转名称"
          />
        </el-form-item>
        <el-form-item
          label="源状态"
          style="width: 18%;"
        >
          <el-select
            v-model="searchInfo.sourceID"
            placeholder="源状态"
            style="width: 100%"
          >
            <el-option
              v-for="item in CircleOptions"
              :key="item.value"
              :label="item.label"
              :value="item.value"
            />
          </el-select>
        </el-form-item>
        <el-form-item
          label="目标状态"
          style="width: 18%;"
        >
          <el-select
            v-model="searchInfo.targetID"
            placeholder="源状态"
            style="width: 100%"
          >
            <el-option
              v-for="item in CircleOptions"
              :key="item.value"
              :label="item.label"
              :value="item.value"
            />
          </el-select>
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
        @click="openDialog('addCircle')"
      >新建流转</el-button>
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
          label="流转ID"
          min-width="60"
          prop="ID"
        />
        <el-table-column
          align="left"
          label="流转名称"
          min-width="100"
          prop="circleName"
        />

        <el-table-column
          align="left"
          label="源状态"
          min-width="100"
          prop="sourceID"
        >
          <template #default="{ row }">
            <el-tag :type="getCircle(row.sourceID)">{{ getCircleLabel(row.sourceID) }}</el-tag>
          </template>
        </el-table-column>
        <el-table-column
          align="left"
          label="目标状态"
          min-width="100"
          prop="targetID"
        >
          <template #default="{ row }">
            <el-tag :type="getCircle(row.targetID)">{{ getCircleLabel(row.targetID) }}</el-tag>
          </template>
        </el-table-column>
        <el-table-column
          align="left"
          label="属性类型"
          min-width="100"
        >
          <template #default="{ row }">
            <el-tag :type="getAttributeType(row.attributeType)">{{ getAttributeTypeLabel(row.attributeType) }}</el-tag>
          </template>
        </el-table-column>

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
        ref="CircleForm"
        :model="form"
        :rules="rules"
        label-width="80px"
      >
        <el-form-item
          label="流转名称"
          prop="circleName"
        >
          <el-input
            v-model="form.circleName"
            autocomplete="off"
          />
        </el-form-item>
        <el-form-item
          label="源状态"
          prop="sourceID"
        >
          <el-select
            v-model="form.sourceID"
            placeholder="请选择"
            style="width:100%"
          >
            <el-option
              v-for="item in CircleOptions"
              :key="item.value"
              :label="`${item.label}`"
              :value="item.value"
            />
          </el-select>
        </el-form-item>
        <el-form-item
          label="目标状态"
          prop="targetID"
        >
          <el-select
            v-model="form.targetID"
            placeholder="请选择"
            style="width:100%"
          >
            <el-option
              v-for="item in CircleOptions"
              :key="item.value"
              :label="`${item.label}`"
              :value="item.value"
            />
          </el-select>
        </el-form-item>
        <el-form-item
          label="属性类型"
          prop="attributeType"
        >
          <el-select
            v-model="form.attributeType"
            placeholder="请选择"
            style="width:100%"
          >
            <el-option
              v-for="item in AttributeTypeOptions"
              :key="item.value"
              :label="`${item.label}`"
              :value="item.value"
            />
          </el-select>
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
import { ref, defineProps, defineExpose, toRefs } from 'vue'
import { ElMessage, ElMessageBox } from 'element-plus'
import {
  createCircle,
  getCircleById,
  updateCircle,
  deleteCircle,
  getCircleList,
  getTemplateStatusList,
} from '@/api/workflow'
defineOptions({
  name: 'FlowCircle'
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

const rules = ref({
  circleName: [{ required: true, message: '请输入流转名称', trigger: 'blur' }],
  sourceID: [{ required: true, message: '请输入源状态', trigger: 'blur' }],
  targetID: [{ required: true, message: '请输入目标状态', trigger: 'blur' }],
  attributeType: [{ required: true, message: '请输入属性类型', trigger: 'blur' }]
})

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

// // 分页
const handleSizeChange = (val) => {
  pageSize.value = val
  getTableData()
}

const handleCurrentChange = (val) => {
  page.value = val
  getTableData()
}

const getCircleOptions = async() => {
  const res = await getTemplateStatusList({ page: 1, pageSize: 999, templateID: Number(templateid.value) })

  CircleOptions.value = res.data.list.map(item => {
    return {
      value: item.ID,
      label: item.statusName
    }
  })
}
getCircleOptions()

// 查询
const getTableData = async() => {
  searchInfo.value.templateID = Number(templateid.value)
  const table = await getCircleList({ page: page.value, pageSize: pageSize.value, ...searchInfo.value })

  if (table.code === 0) {
    tableData.value = table.data.list
    total.value = table.data.total
    page.value = table.data.page
    pageSize.value = table.data.pageSize
  }
}

getTableData()

const handleDelete = (row) => {
  ElMessageBox.confirm('是否删除该过程?', '提示', {
    confirmButtonText: '确定',
    cancelButtonText: '取消',
    type: 'warning'
  })
    .then(async() => {
      const res = await deleteCircle(row)
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

const CircleOptions = ref([])

const AttributeTypeOptions = ref([
  {
    value: 0,
    label: '同意',
    color: 'success' // 设置颜色
  },
  {
    value: 1,
    label: '拒绝',
    color: 'danger' // 设置颜色
  },
])

// 弹窗相关
const form = ref({
  circleName: '',
  sourceID: '',
  targetID: '',
  templateID: Number(templateid.value)
})

const type = ref('')
const CircleForm = ref(null)

const initForm = () => {
  CircleForm.value.resetFields()
  form.value = {
    circleName: '',
    sourceID: '',
    targetID: '',
    templateID: Number(templateid.value),
  }
}

const getAttributeType = (type) => {
  const item = AttributeTypeOptions.value.find(item => item.value === type)
  return item ? item.color : ''
}

const getAttributeTypeLabel = (type) => {
  const item = AttributeTypeOptions.value.find(item => item.value === type)
  return item ? item.label : ''
}

const getCircle = (type) => {
  console.log(CircleOptions.value, type)
  const item = CircleOptions.value.find(item => item.value === type)
  return item ? item.color : ''
}

const getCircleLabel = (type) => {
  const item = CircleOptions.value.find(item => item.value === type)
  return item ? item.label : ''
}

const dialogTitle = ref('新增状态')
const dialogFormVisible = ref(false)
const openDialog = (key) => {
  switch (key) {
    case 'addCircle':
      dialogTitle.value = '新增状态'
      break
    case 'editCircle':
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
  const res = await getCircleById({ id: row.ID })
  form.value = res.data
  openDialog('editCircle')
}

const enterDialog = async() => {
  CircleForm.value.validate(async valid => {
    if (valid) {
      switch (type.value) {
        case 'addCircle':
          {
            searchInfo.value.templateID = Number(templateid.value)

            const res = await createCircle(form.value)
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
        case 'editCircle':
          {
            const res = await updateCircle(form.value)
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

defineExpose({
  getCircleOptions
})

</script>
