<template>
  <div>
    <v-form-designer
      ref="vfDesigner"
      :field-list-api="fieldListApi"
      :banned-widgets="testBanned"
      :designer-config="designerConfig"
    >
      <!-- 自定义按钮插槽演示 -->
      <template #customToolButtons>
        <el-button
          type="text"
          @click="saveFormJson"
        >保存</el-button>
      </template>
    </v-form-designer>
  </div>

</template>

<script setup>
const vfDesigner = ref(null)
const fieldListApi = reactive({
  URL: 'https://www.fastmock.site/mock/2de212e0dc4b8e0885fea44ab9f2e1d0/vform/listField',
  labelKey: 'fieldLabel',
  nameKey: 'fieldName'
})
const testBanned = ref([
  // 'sub-form',
  // 'alert',
])
const designerConfig = reactive({
  languageMenu: true,
  // externalLink: false,
  // formTemplates: false,
  // eventCollapse: false,
  // clearDesignerButton: false,
  // previewFormButton: false,

  // presetCssCode: '.abc { font-size: 16px; }',
})

const saveFormJson = () => {
  const formJson = vfDesigner.value.getFormJson()
  // TODO: 将formJson提交给后端保存接口，需自行实现！！
  ElMessage.success('表单已保存！')
}
</script>

<style scoped lang="scss">
</style>
