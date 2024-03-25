<template>
  <div>
    <el-tag>
      来自工单 - {{ route.query.orderName }}
    </el-tag>
  </div>
</template>

<script setup>
import { useRoute } from 'vue-router'
import { getOrderDetail } from '@/api/workflow'

const route = useRoute()

const getOrderDetailData = async() => {
  console.log('route.query.orderId', route.query.orderId)
  const res = await getOrderDetail({ id: Number(route.query.orderId) })
  console.log('res', res.data)
}
getOrderDetailData()

</script>
