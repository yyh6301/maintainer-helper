<template>
  <div class="card-container">
    <div
      v-for="template in cardList"
      :key="template.id"
      class="card-wrapper"
      @click="handleCardClick(template)"
    >
      <el-card
        class="card-item"
        shadow="always"
        body-style="height:100%;width:100%;"
      >
        <!-- 居中显示模板信息 -->
        <div class="template-info">
          <div>{{ template.flowName }}</div>
          <!-- 其他模板信息 -->
        </div>
      </el-card>
    </div>

    <el-dialog
      v-model="dialogVisible"
      title="创建工单"
      :close-on-click-modal="false"
    >
      <!-- 这里展示模板的详细信息 -->
      <el-tag style="margin-bottom: 20px ;">
        {{ currentTemplate.flowName }}
      </el-tag>

      <el-form

        :model="orderData"
        :rules="rules"
      >
        <el-form-item
          label="标题"
          prop="title"
        >
          <el-input
            v-model="orderData.title"
            placeholder="请输入工单标题"
          />
        </el-form-item>
      </el-form>
      <v-form-render
        ref="vFormRef"
        :form-json="formJson"
        :form-data="formData"
        :option-data="optionData"
      />
      <div style="text-align: center; margin-top: 20px;">
        <el-button
          type="default"
          @click="dialogVisible = false"
        >取消</el-button>
        <el-button
          type="primary"
          style="margin-left: 20px;"
          @click="submitForm(currentTemplate)"
        >发起工单</el-button>
      </div>
    </el-dialog>
  </div>
</template>

<script setup>

import { ref, reactive } from 'vue'
import {
  getWorkflowTemplateList,
  createWorkflowOrder,
} from '@/api/workflow'
import { VFormRender } from 'vform3-builds'
import { useUserStore } from '@/pinia/modules/user'
import { ElMessage } from 'element-plus'
import { useRouter } from 'vue-router'
const router = useRouter()
const userStore = useUserStore()
const cardList = ref([])
const dialogVisible = ref(false)
const currentTemplate = ref({})

const rules = ref({
  title: [{ required: true, message: '请输入标题', trigger: 'blur' }],
})

const formJson = reactive({ 'widgetList': [], 'formConfig': { 'modelName': 'formData', 'refName': 'vForm', 'rulesName': 'rules', 'labelWidth': 80, 'labelPosition': 'left', 'size': '', 'labelAlign': 'label-left-align', 'cssCode': '', 'customClass': '', 'functions': '', 'layoutType': 'PC', 'jsonVersion': 3, 'onFormCreated': '', 'onFormMounted': '', 'onFormDataChange': '', 'onFormValidate': '' }})

const formData = reactive({})
const optionData = reactive({})
const vFormRef = ref(null)

const orderData = ref({
  title: '',
  templateID: 0,
  orderDetail: {},
  orderCreator: '',
  orderModifier: '',
})

const getCardList = async() => {
  try {
    const res = await getWorkflowTemplateList({ page: 0, pageSize: 20 })
    if (res.code === 0) {
      cardList.value = res.data.list
    }
  } catch (error) {
    console.error('Failed to fetch card list:', error)
  }
}

const handleCardClick = (template) => {
  currentTemplate.value = template
  setTimeout(() => {
    if (vFormRef.value !== null) {
      vFormRef.value.setFormJson(currentTemplate.value.flowFormDetail)
    }
  }, 100) // 延迟 1000 毫秒后执行
  dialogVisible.value = true
}

getCardList()

const submitForm = (currentTemplate) => {
  vFormRef.value.getFormData().then(async(formData) => {
    orderData.value.templateID = currentTemplate.ID
    orderData.value.orderDetail = JSON.stringify(formData)
    orderData.value.orderCreator = userStore.userInfo.userName
    orderData.value.orderModifier = userStore.userInfo.userName
    console.log('order Data:', orderData.value)
    console.log('userInfo', userStore.userInfo)
    const res = await createWorkflowOrder(orderData.value)
    console.log('res:', res)
    if (res.code === 0) {
      ElMessage.success('工单创建成功')
      router.push({ name: 'allOrder' })
    }
  }).catch(error => {
    // Form Validation failed
    ElMessage.error(error)
  })
}
</script>

<style  scoped>
.card-container {
    display: flex;
    flex-wrap: wrap;
    justify-content: space-between; /* 横向间距均分 */
  }

  .card-wrapper{
    cursor: pointer; /* 添加指针样式以表示可点击 */
    width: calc(25% - 30px); /* 每个卡片宽度为 1/3 屏幕宽度减去间距 */
    margin-bottom: 20px; /* 卡片下边距 */
    height:200px;
  }
  .card-item{
    height:100%;
  }

  .template-info {
    display: flex;
    flex-direction: column;
    justify-content: center; /* 垂直居中 */
    align-items: center; /* 水平居中 */
    height: 80%; /* 卡片内部高度撑满 */
    width: 80%;
  }
</style>
