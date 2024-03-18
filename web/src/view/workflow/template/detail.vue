<template>
  <div>
    <bpmn-modeler
      ref="refNode"
      :xml="xml"
      :users="users"
      :groups="groups"
      :categorys="categorys"
      :is-view="false"
      @save="save"
    />
  </div>
</template>

<script setup>
import { bpmnModeler } from 'workflow-bpmn-modeler'

defineOptions({
  name: 'TemplateDetail',
})

const xml = ref('')
const users = ref([])
const groups = ref([])
const categorys = ref([])
const refNode = ref(null)

</script>

<style scoped lang="scss">
.warning {
  color: #dc143c;
}
</style>
