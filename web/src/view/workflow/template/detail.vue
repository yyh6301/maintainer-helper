<template>
  <div id="app">
    <el-tag
      v-if="route.query.name === undefined"
      style="margin-bottom:10px;"
    > 创建工单模板</el-tag>
    <el-tag
      v-else
      style="margin-bottom: 10px;"
    >
      来自工单模板 - {{ route.query?.name }}
    </el-tag>

    <el-tabs type="border-card">
      <!-- ===============================基本信息============================= -->
      <el-tab-pane label="基本信息">
        <el-form :model="form">
          <div style="display:flex">
            <el-form-item label="工作流名称">
              <el-input
                v-model="form.flowName"
                placeholder="请输入工作流名称"
              />
            </el-form-item>
            <el-form-item
              label="流程描述"
              style="margin-left:20%;
              width:60%;"
            >
              <el-input
                v-model="form.flowDesc"
                placeholder="请输入流程描述"
              />
            </el-form-item>
          </div>
          <el-form-item>
            <el-button
              type="primary"
              size="default"
              style=""
              @click="submitTemplate"
            >
              确认
            </el-button>
          </el-form-item>
        </el-form>
      </el-tab-pane>

      <!-- ===============================自定义表单============================= -->
      <el-tab-pane label="自定义表单">

        <v-form-designer
          ref="vfDesigner"
          :banned-widgets="testBanned"
          :designer-config="designerConfig"
        >
          <!-- 自定义按钮插槽演示
                <template #customToolButtons>
                  <el-button
                    type="text"
                    @click="saveFormJson"
                  >保存</el-button>
                </template> -->
        </v-form-designer>
      </el-tab-pane>

      <!-- ===============================状态============================= -->
      <el-tab-pane label="状态">
        <FlowStatus :templateid="route.query.id" />
      </el-tab-pane>
      <!-- ===============================流转============================= -->
      <el-tab-pane label="流转">
        <FlowCircle :templateid="route.query.id" />
      </el-tab-pane>

    </el-tabs>

  </div>
</template>

<script setup>
import { ref, reactive, onMounted, nextTick } from 'vue'
import { ElMessage } from 'element-plus'
import {
  createWorkflowTemplate,
  getWorkflowTemplateById,
  updateWorkflowTemplate,
} from '@/api/workflow'
import { useRouter, useRoute } from 'vue-router'
import { useUserStore } from '@/pinia/modules/user'
import { VFormDesigner } from 'vform3-builds'
import FlowStatus from '@/view/workflow/template/component/flowStatus.vue'
import FlowCircle from '@/view/workflow/template/component/flowCircle.vue'
const userStore = useUserStore()
const router = useRouter()
const route = useRoute()

var isCreate = route.query.id === undefined
var nullJson = '{"formConfig": {"size": "", "cssCode": "", "refName": "vForm", "functions": "", "modelName": "formData", "rulesName": "rules", "labelAlign": "label-left-align", "labelWidth": 80, "layoutType": "PC", "customClass": [], "jsonVersion": 3, "labelPosition": "left", "onFormCreated": "", "onFormMounted": "", "onFormDataChange": ""}, "widgetList": []}'

// ============工作流基本信息========
const form = ref({
  flowName: '',
  flowDesc: '',
  flowFormDetail: '',
  flowDetail: '',
  flowCreator: '',
  flowModifier: ''
})

onMounted(async() => {
  vfDesigner.value.setFormJson(nullJson)
  if (!isCreate) {
    var id = Number(route.query.id)
    const res = await getWorkflowTemplateById({ ID: id })
    if (res.code === 0) {
      form.value = res.data
      nextTick(() => {
        vfDesigner.value.setFormJson(form.value.flowFormDetail)
      })
    }
  }
})

const submitTemplate = async() => {
  form.value.flowFormDetail = JSON.stringify(vfDesigner.value.getFormJson())
  form.value.flowDetail = JSON.stringify('{"test": "test"}')

  if (isCreate) {
    // 新增工作流
    form.value.flowCreator = userStore.userInfo.nickName
    form.value.flowModifier = userStore.userInfo.nickName
    const res = await createWorkflowTemplate(form.value)
    if (res.code === 0) {
      ElMessage.success('新增工作流模板成功')
      router.push({ name: 'template' })
    }
  } else {
    // 更改工作流
    form.value.flowModifier = userStore.userInfo.nickName
    const res = await updateWorkflowTemplate(form.value)
    if (res.code === 0) {
      ElMessage.success('更新工作流模板成功')
      router.replace({ name: 'template' })
    }
  }
}

// ========表单设计=======
const vfDesigner = ref(null)
// const fieldListApi = reactive({
//   URL: 'https://www.fastmock.site/mock/2de212e0dc4b8e0885fea44ab9f2e1d0/vform/listField',
//   labelKey: 'fieldLabel',
//   nameKey: 'fieldName'
// })
const testBanned = ref([
  // 'sub-form',
  // 'alert',
])
const designerConfig = reactive({
  languageMenu: true,
  externalLink: false,
  formTemplates: true,
  eventCollapse: true,
  toolbarMaxWidth: 250,
  exportCodeButton: false,
  exportJsonButton: true,
  importJsonButton: true,
  generateSFCButton: false,
  // clearDesignerButton: true,
  // previewFormButton: false,

  // presetCssCode: '.abc { font-size: 16px; }',
})

</script>

<style lang="scss">
  #app {
    height: 100%;
  }
</style>
