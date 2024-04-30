<template>
  <div>
    <div class="gva-search-box">
      <el-form
        ref="searchForm"
        :inline="true"
        :model="searchInfo"
      >
        <el-form-item label="转让人">
          <el-input
            v-model="searchInfo.Owner"
            placeholder="转让人"
          />
        </el-form-item>
        <el-form-item label="接收人">
          <el-input
            v-model="searchInfo.ToOwner"
            placeholder="接收人"
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
      >转让机器</el-button>
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
          label="转让人"
          min-width="100"
          prop="owner"
        />
        <el-table-column
          align="left"
          label="接收人"
          min-width="100"
          prop="toOwner"
        />
        <el-table-column
          align="left"
          label="转让数量 "
          min-width="100"
          prop="count"
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
            label="云厂商"
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
          <el-form-item
            prop="owner"
            label="转让人"
            style="width: 30%;"
          >
            <el-select
              v-model="form.owner"
              placeholder="请选择转让人"
            >
              <el-option
                v-for="(item, index) in UserList"
                :key="index"
                :label="item.userName"
                :value="item.userName"
              />
            </el-select>
          </el-form-item>
          <el-form-item
            prop="toOwner"
            label="接收人"
            style="width: 30%;"
          >
            <el-select
              v-model="form.toOwner"
              placeholder="请选择接收人"
            >
              <el-option
                v-for="(item, index) in UserList"
                :key="index"
                :label="item.userName"
                :value="item.userName"
              />
            </el-select>
          </el-form-item>
        </div>
        <el-table
          ref="multipleTableRef"
          :data="dialogTableData"
          style="
          margin-top:20px;
          width: 600px;
          height:300px"
          @selection-change="handleSelectionChange"
        >
          <el-table-column
            type="selection"
            width="55"
          />
          <el-table-column
            label="实例名称"
            property="instanceName"
          />
          <el-table-column
            property="instanceType"
            label="实例类型"
          />

        </el-table>

        <el-form-item
          style="margin-top:20px"
          label="转让原因"
        >
          <el-input
            v-model="form.applyReason"
            type="textarea"
            placeholder="请输入转让原因"
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
  getTransferList,
  createTransfer,
  deleteTransfer,
  getAssetsList,
} from '@/api/cmdb'
import {
  getUserList
} from '@/api/user'
import { useRouter } from 'vue-router'
import { ref, onMounted } from 'vue'
import { ElMessageBox, ElMessage } from 'element-plus'
import { useUserStore } from '@/pinia/modules/user'
import { formatDate } from '@/utils/format'

const userStore = useUserStore()
const multipleTableRef = ref(null)
const multipleSelection = ref([])
const dialogForm = ref(null)

const handleSelectionChange = (val) => {
  multipleSelection.value = val
}

const dialogTableData = ref([])
const UserList = ref([])

const rules = ref({
  cloudType: [{ required: true, message: '请输入云厂商', trigger: 'blur' }],
  owner: [{ required: true, message: '请输入转让人', trigger: 'blur' }],
  toOwner: [{ required: true, message: '请输入接收人', trigger: 'blur' }],
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
  name: 'Transfer',
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
  const table = await getTransferList({ page: page.value, pageSize: pageSize.value, ...searchInfo.value })

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
  owner: '',
  toOwner: '',
  applyReason: '',
  count: '',
  orderID: 0,
  instanceList: [],
  workFlowOrder: {
    orderDetail: '',
    templateID: 19,
  }
})

const closeDialog = () => {
  form.value = {
    cloudType: '',
    owner: '',
    toOwner: '',
    applyReason: '',
    count: '',
    instanceList: [],
    orderID: 0,
    workFlowOrder: {
      orderDetail: '',
      templateID: 19,
    }
  }
  dialogFormVisible.value = false
}

const enterDialog = async() => {
  dialogForm.value.validate(async valid => {
    if (valid) {
      form.value.count = multipleSelection.value.length
      if (form.value.count === 0) {
        ElMessage.warning('请选择转让的机器')
        return
      }
      form.value.applyer = userStore.userInfo.userName
      form.value.instanceList = multipleSelection.value
      const res = await createTransfer(form.value)
      if (res.code === 0) {
        ElMessage.success(res.msg)
      }
      closeDialog()
      getTableData()
    }
  })
}

onMounted(async() => {
  const res = await getUserList({ page: 1, pageSize: 999 })
  if (res.code === 0) {
    UserList.value = res.data.list
  }
})

const getDialogTableData = async() => {
  const res = await getAssetsList({ page: 1, pageSize: 999 })
  if (res.code === 0) {
    dialogTableData.value = res.data.list
  }
}

const handleAdd = async() => {
  dialogFormVisible.value = true
  dialogTitle.value = '新建转让'
  await getDialogTableData()
}

const handleEdit = (row) => {
  var orderName = row.cloudType + '机器转让-' + row.instanceName
  router.replace({
    name: 'orderDetail',
    query: {
      orderId: row.orderID,
      orderName: orderName
    }
  })
}

const handleDelete = (row) => {
  ElMessageBox.confirm('是否撤销此次转让?', '提示', {
    confirmButtonText: '确定',
    cancelButtonText: '取消',
    type: 'warning'
  })
    .then(async() => {
      console.log(row)
      const res = await deleteTransfer(row)
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
