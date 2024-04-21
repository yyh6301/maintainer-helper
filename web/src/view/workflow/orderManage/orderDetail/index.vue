<template>
  <div>
    <el-tag>
      来自工单 - {{ route.query.orderName }}
    </el-tag>
    <el-form>
      <el-form-item
        style="margin-left: 20px;margin-top:20px;"
        label="工单标题"
      >
        <el-input
          v-model="form.title"
          style="width: 300px"
          readonly
        />
      </el-form-item>
    </el-form>

    <v-form-render
      ref="vFormRef"
      :form-json="formJson"
      :form-data="vFormData"
      :option-data="optionData"
    />

    <el-form v-show="UserHandle">
      <el-divider />

      <el-form-item>
        <el-tag>用户工单处理</el-tag>
      </el-form-item>
      <el-form-item>
        <el-input
          v-model="handleParams.opinion"
          type="textarea"
          style="width: 350px;"
          placeholder="请输入处理理由"
        />
      </el-form-item>

      <el-form-item>
        <el-button
          type="danger"
          @click="handleOrder('reject')"
        >
          拒绝工单
        </el-button>
        <el-button
          type="primary"
          @click="handleOrder('pass')"
        >
          通过工单
        </el-button>
      </el-form-item>
    </el-form>

    <el-divider />

    <el-tag>
      操作记录
    </el-tag>

    <el-table
      :data="tableData"
      border
      style="margin-top:10px;width: 100%"
    >
      <el-table-column
        prop="ID"
        label="序号"
        width="80"
      />
      <el-table-column
        prop="sourceStatus.statusName"
        label="源状态"
        width="180"
      />
      <el-table-column
        prop="targetStatus.statusName"
        label="目标状态"
        width="180"
      />

      <el-table-column
        prop="handler"
        label="处理人"
        width="180"
      />
      <el-table-column
        prop="UpdateAt"
        label="处理时间"
        width="180"
      >
        <template #default="{ row }">
          <div v-if="row.status !== 0">
            {{ formatDate(row.UpdatedAt) }}
          </div>
          <div v-else>
            <el-tag :type="'info'">等待处理</el-tag>
          </div>
        </template>
      </el-table-column>
      <el-table-column
        prop="status"
        label="处理结果"
        width="180"
      >
        <template #default="{row}">
          <el-tag :type="getStatusColor(row.status)">
            {{ getStatusLabel(row.status) }}
          </el-tag></template>
      </el-table-column>
      <el-table-column
        prop="opinion"
        label="处理意见"
      />
    </el-table>

  </div>

</template>

<script setup>
import { useUserStore } from '@/pinia/modules/user'
import { useRoute, useRouter } from 'vue-router'
import { getOrderDetail } from '@/api/workflow'
import { ref, onMounted, reactive } from 'vue'
import { VFormRender } from 'vform3-builds'
import { formatDate } from '@/utils/format.js'
import { handleOrderApi } from '@/api/workflow.js'
import { ElMessage } from 'element-plus'
const route = useRoute()
const router = useRouter()
const UserHandle = ref(true)
const form = ref({})
const tableData = ref([])
const userStore = useUserStore()
const orderId = ref(0)
orderId.value = Number(route.query.orderId)
const getOrderDetailData = async() => {
  console.log('route.query.orderId', route.query.orderId)
  const res = await getOrderDetail({ id: Number(route.query.orderId) })
  if (res.code === 0) {
    form.value = res.data
    tableData.value = form.value.workFlowOrderLog
    setTimeout(() => {
      if (vFormRef.value !== null) {
        vFormRef.value.setFormJson(form.value.workFlowTemplate.flowFormDetail)
        setTimeout(() => {
          vFormRef.value.setFormData(JSON.parse(form.value.orderDetail))
          vFormRef.value.disableForm()
        }, 100)
      }
    }, 100) // 延迟 1000 毫秒后执行
  }
}

onMounted(async() => {
  await getOrderDetailData()
  console.log(form.value.workFlowStatus)
  if (form.value.workFlowStatus.approvalUser !== userStore.userInfo.userName &&
  userStore.userInfo.userName !== 'admin') {
    UserHandle.value = false
  }
})

const statusOptions = ref([
  {
    value: 0,
    label: '未处理',
    color: 'info',
  },
  {
    value: 1,
    label: '同意',
    color: 'primary',
  },
  {
    value: 2,
    label: '拒绝',
    color: 'danger',
  },
])

const getStatusColor = (type) => {
  const item = statusOptions.value.find(item => item.value === type)
  return item ? item.color : ''
}

const getStatusLabel = (type) => {
  const item = statusOptions.value.find(item => item.value === type)
  return item ? item.label : ''
}

const formJson = reactive({ 'widgetList': [], 'formConfig': { 'modelName': 'formData', 'refName': 'vForm', 'rulesName': 'rules', 'labelWidth': 80, 'labelPosition': 'left', 'size': '', 'labelAlign': 'label-left-align', 'cssCode': '', 'customClass': '', 'functions': '', 'layoutType': 'PC', 'jsonVersion': 3, 'onFormCreated': '', 'onFormMounted': '', 'onFormDataChange': '', 'onFormValidate': '' }})
const vFormData = reactive({})
const optionData = reactive({})
const vFormRef = ref(null)

const handleParams = ref({
  handler: '',
  opinion: '',
  result: '',
})
handleParams.value.handler = userStore.userInfo.userName

const handleOrder = async(type) => {
  console.log('type', type)
  var res = {}
  switch (type) {
    case 'reject':
      // todo ，这个opinion让用户自己输入
      handleParams.value.result = 'false'
      res = await handleOrderApi({ id: orderId.value, ...handleParams.value })
      console.log('res', res)
      if (res.code === 0) {
        ElMessage.success('拒绝工单成功')
        router.push({ name: 'allOrder' })
      }
      break
    case 'pass':
      handleParams.value.result = 'true'
      res = await handleOrderApi({ id: orderId.value, ...handleParams.value })
      console.log('res', res)
      if (res.code === 0) {
        ElMessage.success('通过工单成功')
        router.push({ name: 'allOrder' })
      }
      break
    default:
      break
  }
}

</script>
