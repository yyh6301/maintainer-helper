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
          <el-select
            v-model="searchInfo.CloudType"
            placeholder="云厂商"
            clearable
          >
            <el-option
              v-for="(item,index) in CloudTypeOption"
              :key="index"
              :label="item.value"
              :value="item.value"
            />
          </el-select>
        </el-form-item>
        <el-form-item label="实例名称">
          <el-input
            v-model="searchInfo.instanceName"
            placeholder="实例名称"
          />
        </el-form-item>
        <el-form-item label="实例类型">
          <el-input
            v-model="searchInfo.instanceType"
            placeholder="实例类型"
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
      >新建申请</el-button>
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
          label="厂商"
          min-width="100"
          prop="cloudType"
        />
        <el-table-column
          align="left"
          label="实例类型"
          min-width="100"
          prop="instanceType"
        />
        <el-table-column
          align="left"
          label="实例名称"
          min-width="100"
          prop="instanceName"
        />
        <el-table-column
          align="left"
          label="申请人"
          min-width="100"
          prop="applyer"
        />
        <el-table-column
          align="left"
          label="续费时间"
          min-width="100"
          prop="renewTime"
        />
        <el-table-column
          align="middle"
          label="申请时间"
          min-width="100"
        >
          <template #default="scope">{{ formatDate(scope.row.CreatedAt) }}</template>
        </el-table-column>
        <el-table-column
          align="left"
          label="工单状态"
          width="120"
        >
          <template #default="{ row }">
            <el-tag>{{ row.workFlowOrder.workFlowStatus.statusName }}</el-tag>
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
              查看工单详情
            </el-button>
            <el-button
              v-show="userStore.userInfo.isAdmin"
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
        ref="dialogForm"
        :rules="rules"
        :model="form"
        label-width="80px"
      >
        <div style="display: flex">

          <el-form-item
            prop="cloudType"
            label="厂商"
            style="width:30%"
          >
            <el-select
              v-model="form.cloudType"
              placeholder="请选择厂商"
            >
              <el-option
                v-for="(item, index) in CloudTypeOption"
                :key="index"
                :label="item.value"
                :value="item.value"
              />

            </el-select>
          </el-form-item>
        </div>

        <v-form-render
          ref="vFormRef"
          :form-json="formJson"
          :form-data="formData"
          :option-data="optionData"
        />

        <el-form-item
          label="续费原因"
        >
          <el-input
            v-model="form.applyReason"
            type="textarea"
            placeholder="请输入续费原因"
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
  getWorkflowTemplateById,
} from '@/api/workflow'
import {
  getRenewList,
  createRenew,
  deleteRenew,
} from '@/api/cmdb'
import { VFormRender } from 'vform3-builds'
import { useRouter } from 'vue-router'
import { ref, reactive } from 'vue'
import { ElMessageBox, ElMessage } from 'element-plus'
import { useUserStore } from '@/pinia/modules/user'
import { formatDate } from '@/utils/format'
const dialogForm = ref(null)
const userStore = useUserStore()
const rules = ref({
  cloudType: [{ required: true, message: '请输入云厂商', trigger: 'blur' }],
})

const CloudTypeOption = ref([{
  value: '腾讯云'
}, {
  value: '阿里云'
}, {
  value: 'AWS'
}])

const router = useRouter()

defineOptions({
  name: 'Renew',
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
  const table = await getRenewList({ page: page.value, pageSize: pageSize.value, ...searchInfo.value })

  if (table.code === 0) {
    tableData.value = table.data.list
    total.value = table.data.total
    page.value = table.data.page
    pageSize.value = table.data.pageSize
  }
}

getTableData()
const dialogTitle = ref('')
const dialogFormVisible = ref(false)
const form = ref({
  cloudType: '',
  instanceName: '',
  instanceType: '',
  applyer: '',
  renewTime: '',
  orderID: 0,
  workFlowOrder: {
    orderDetail: '',
    templateID: 16,
  }
})

const closeDialog = () => {
  form.value = {
    cloudType: '',
    instanceName: '',
    instanceType: '',
    applyer: '',
    renewTime: '',
    orderID: 0,
    workFlowOrder: {
      orderDetail: '',
      templateID: 16,
    }
  }
  dialogFormVisible.value = false
}

const enterDialog = () => {
  dialogForm.value.validate(async valid => {
    if (valid) {
      vFormRef.value.getFormData().then(async(formData) => {
        form.value.applyer = userStore.userInfo.userName
        form.value.workFlowOrder.orderDetail = JSON.stringify(formData)
        console.log(form.value)
        const res = await createRenew(form.value)
        if (res.code === 0) {
          ElMessage.success(res.msg)
        }
        closeDialog()
        getTableData()
      }).catch(error => {
        ElMessage.error(error)
      })
    }
  })
}

const formJson = reactive({ 'widgetList': [], 'formConfig': { 'modelName': 'formData', 'refName': 'vForm', 'rulesName': 'rules', 'labelWidth': 80, 'labelPosition': 'left', 'size': '', 'labelAlign': 'label-left-align', 'cssCode': '', 'customClass': '', 'functions': '', 'layoutType': 'PC', 'jsonVersion': 3, 'onFormCreated': '', 'onFormMounted': '', 'onFormDataChange': '', 'onFormValidate': '' }})

const formData = reactive({})
const optionData = reactive({})
const vFormRef = ref(null)
// 根据templateid找到对应模板的form json
const getVFormJson = async() => {
  const res = await getWorkflowTemplateById({ id: form.value.workFlowOrder.templateID })
  if (res.code === 0) {
    form.value.formDetail = res.data.flowFormDetail
  }
}
const handleAdd = async() => {
  dialogFormVisible.value = true
  dialogTitle.value = '新建续费'
  await getVFormJson()
  setTimeout(() => {
    if (vFormRef.value !== null) {
      vFormRef.value.setFormJson(form.value.formDetail)
    }
  }, 100) // 延迟 1000 毫秒后执行
}

const handleEdit = (row) => {
  var orderName = row.cloudType + '机器续费-' + row.instanceName
  router.replace({
    name: 'orderDetail',
    query: {
      orderId: row.orderID,
      orderName: orderName
    }
  })
}

const handleDelete = (row) => {
  ElMessageBox.confirm('是否删除此次续费?', '提示', {
    confirmButtonText: '确定',
    cancelButtonText: '取消',
    type: 'warning'
  })
    .then(async() => {
      console.log(row)
      const res = await deleteRenew(row)
      if (res.code === 0) {
        ElMessage({
          type: 'success',
          message: '删除成功!'
        })
        if (tableData.value.length === 1 && page.value > 1) {
          page.value--
        }
        getTableData()
      } else {
        ElMessage.warning(res.msg)
      }
    })
}

</script>

  <style scoped lang="scss">
  .warning {
    color: #dc143c;
  }
  </style>
