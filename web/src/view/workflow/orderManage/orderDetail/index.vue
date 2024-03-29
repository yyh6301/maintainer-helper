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

    <el-form>
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

    <el-table
      :data="tableData"
      border
      style="width: 100%"
    >
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
            {{ formatDate(row.UpdateAt) }}
          </div>
          <div v-else>
            <el-tag :type="'info'">等待处理</el-tag>
          </div>
        </template>
      </el-table-column>
      <el-table-column
        prop="status"
        label="处理结果"
      >
        <template #default="{row}">
          <el-tag :type="getStatusColor(row.status)">
            {{ getStatusLabel(row.status) }}
          </el-tag></template>
      </el-table-column>
    </el-table>

  </div>

</template>

<script setup>
import { useRoute } from 'vue-router'
import { getOrderDetail } from '@/api/workflow'
import { ref, onMounted, reactive } from 'vue'
import { VFormRender } from 'vform3-builds'
import { formatDate } from '@/utils/format.js'
const route = useRoute()

const form = ref({})
const tableData = ref([])

const getOrderDetailData = async() => {
  console.log('route.query.orderId', route.query.orderId)
  const res = await getOrderDetail({ id: Number(route.query.orderId) })
  if (res.code === 0) {
    form.value = res.data
    tableData.value = form.value.workFlowOrderLog
    console.log('form', form.value)
    console.log('tableData', tableData.value)
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
  option: '',
  result: '',
})

const handleOrder = (type) => {
  console.log('type', type)
  switch (type) {
    case 'reject':
      handleParams

      break
    case 'pass':
      break
    default:
      break
  }
}

</script>
