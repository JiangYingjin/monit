<template>
  <div>
    <warning-bar title="注：右上角头像下拉可切换角色" />
    <div class="gva-table-box">
      <div class="gva-btn-list">
        <el-button
          type="primary"
          icon="plus"
          @click="addMachine"
        >新增主机</el-button>
      </div>
      <el-table
        :data="tableData"
        row-key="ID"
      >
        <el-table-column
          align="left"
          label="ID"
          min-width="50"
          prop="ID"
        />
        <el-table-column
          align="left"
          label="名称"
          min-width="150"
          prop="name"
        />
        <el-table-column
          align="left"
          label="描述"
          min-width="150"
          prop="description"
        />
        <el-table-column
          align="left"
          label="dataType"
          min-width="180"
          prop="valueType"
        />

        <!--        <el-table-column-->
        <!--          align="left"-->
        <!--          label="启用"-->
        <!--          min-width="150"-->
        <!--        >-->
        <!--          <template #default="scope">-->
        <!--            <el-switch-->
        <!--              v-model="scope.row.enable"-->
        <!--              inline-prompt-->
        <!--              :active-value="1"-->
        <!--              :inactive-value="2"-->
        <!--              @change="()=>{switchEnable(scope.row)}"-->
        <!--            />-->
        <!--          </template>-->
        <!--        </el-table-column>-->

        <el-table-column
          label="操作"
          min-width="250"
          fixed="right"
        >
          <template #default="scope">
            <el-button
              type="primary"
              link
              icon="delete"
              @click="deleteMachineFunc(scope.row)"
            >删除</el-button>
            <el-button
              type="primary"
              link
              icon="edit"
              @click="openEdit(scope.row)"
            >编辑</el-button>
            <el-button
              type="primary"
              link
              icon="magic-stick"
              @click="resetPasswordFunc(scope.row)"
            >重置密码</el-button>
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
      v-model="addMachineDialog"
      title="用户"
      :show-close="false"
      :close-on-press-escape="false"
      :close-on-click-modal="false"
    >
      <div style="height:60vh;overflow:auto;padding:0 12px;">
        <el-form
          ref="machineForm"
          :rules="rules"
          :model="machineInfo"
          label-width="80px"
        >
          <el-form-item
            v-if="dialogFlag === 'add'"
            label="主机id"
            prop="ID"
          >
            <el-input-number v-model="machineInfo.ID" />
          </el-form-item>
          <el-form-item
            v-if="dialogFlag === 'add'"
            label="密码"
            prop="password"
          >
            <el-input v-model="machineInfo.password" />
          </el-form-item>
          <el-form-item
            label="昵称"
            prop="name"
          >
            <el-input v-model="machineInfo.name" />
          </el-form-item>
          <el-form-item
            label="描述"
            prop="description"
          >
            <el-input v-model="machineInfo.description" />
          </el-form-item>
          <el-form-item
            label="valueType"
            prop="valueType"
          >
            <el-input v-model="machineInfo.valueType" />
          </el-form-item>
        </el-form>

      </div>

      <template #footer>
        <div class="dialog-footer">
          <el-button @click="closeaddMachineDialog">取 消</el-button>
          <el-button
            type="primary"
            @click="enteraddMachineDialog"
          >确 定</el-button>
        </div>
      </template>
    </el-dialog>
  </div>
</template>

<script setup>

import {
  getUserList,
  setUserAuthorities,
  register,
  deleteUser
} from '@/api/user'

import {
  createMachine, getMachineList, deleteMachine, deleteMachineByIds
} from '@/api/machine'

import { getAuthorityList } from '@/api/authority'
import CustomPic from '@/components/customPic/index.vue'
import ChooseImg from '@/components/chooseImg/index.vue'
import WarningBar from '@/components/warningBar/warningBar.vue'
import { setUserInfo, resetPassword } from '@/api/user.js'

import { nextTick, ref, watch } from 'vue'
import { ElMessage, ElMessageBox } from 'element-plus'

defineOptions({
  name: 'Machine',
})

const path = ref(import.meta.env.VITE_BASE_API + '/')
// 初始化相关
const setAuthorityOptions = (AuthorityData, optionsData) => {
  AuthorityData &&
        AuthorityData.forEach(item => {
          if (item.children && item.children.length) {
            const option = {
              authorityId: item.authorityId,
              authorityName: item.authorityName,
              children: []
            }
            setAuthorityOptions(item.children, option.children)
            optionsData.push(option)
          } else {
            const option = {
              authorityId: item.authorityId,
              authorityName: item.authorityName
            }
            optionsData.push(option)
          }
        })
}

const page = ref(1)
const total = ref(0)
const pageSize = ref(10)
const tableData = ref([])
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
  const table = await getMachineList({ page: page.value, pageSize: pageSize.value })
  if (table.code === 0) {
    tableData.value = table.data.list
    total.value = table.data.total
    page.value = table.data.page
    pageSize.value = table.data.pageSize
  }
}

watch(() => tableData.value, () => {
  setAuthorityIds()
})

const initPage = async() => {
  getTableData()
  const res = await getAuthorityList({ page: 1, pageSize: 999 })
  setOptions(res.data.list)
}

initPage()

const resetPasswordFunc = (row) => {
  ElMessageBox.confirm(
    '是否将此用户密码重置为123456?',
    '警告',
    {
      confirmButtonText: '确定',
      cancelButtonText: '取消',
      type: 'warning',
    }
  ).then(async() => {
    const res = await resetPassword({
      ID: row.ID,
    })
    if (res.code === 0) {
      ElMessage({
        type: 'success',
        message: res.msg,
      })
    } else {
      ElMessage({
        type: 'error',
        message: res.msg,
      })
    }
  })
}
const setAuthorityIds = () => {
  tableData.value && tableData.value.forEach((machine) => {
    machine.authorityIds = machine.authorities && user.ID.map(i => {
      return i.authorityId
    })
  })
}

const chooseImg = ref(null)
const openHeaderChange = () => {
  chooseImg.value.open()
}

const authOptions = ref([])
const setOptions = (authData) => {
  authOptions.value = []
  setAuthorityOptions(authData, authOptions.value)
}

const deleteMachineFunc = async(row) => {
  ElMessageBox.confirm('确定要删除吗?', '提示', {
    confirmButtonText: '确定',
    cancelButtonText: '取消',
    type: 'warning'
  }).then(async() => {
    const res = await deleteMachineByIds({ id: row.ID })
    if (res.code === 0) {
      ElMessage.success('删除成功')
      await getTableData()
    }
  })
}

// 弹窗相关
const machineInfo = ref({
  ID: 0,
  password: '',
  name: '',
  description: '',
  valueType: '',
})

// const rules = ref({
//   userName: [
//     { required: true, message: '请输入用户名', trigger: 'blur' },
//     { min: 5, message: '最低5位字符', trigger: 'blur' }
//   ],
//   password: [
//     { required: true, message: '请输入用户密码', trigger: 'blur' },
//     { min: 6, message: '最低6位字符', trigger: 'blur' }
//   ],
//   nickName: [
//     { required: true, message: '请输入用户昵称', trigger: 'blur' }
//   ],
//   phone: [
//     { pattern: /^1([38][0-9]|4[014-9]|[59][0-35-9]|6[2567]|7[0-8])\d{8}$/, message: '请输入合法手机号', trigger: 'blur' },
//   ],
//   email: [
//     { pattern: /^([0-9A-Za-z\-_.]+)@([0-9a-z]+\.[a-z]{2,3}(\.[a-z]{2})?)$/g, message: '请输入正确的邮箱', trigger: 'blur' },
//   ],
//   authorityId: [
//     { required: true, message: '请选择用户角色', trigger: 'blur' }
//   ]
// })


const mustUint = (rule, value, callback) => {
  if (!/^[0-9]*[1-9][0-9]*$/.test(value)) {
    return callback(new Error('请输入正整数'))
  }
  return callback()
}

const rules = ref({
  ID: [
    { required: true, message: '请输入ID', trigger: 'blur' },
    { validator: mustUint, trigger: 'blur', message: '必须为正整数' }
  ],
  password: [
    { required: true, message: '请输入密码', trigger: 'blur' },
    { min: 6, message: '最低6位字符', trigger: 'blur' }
  ],
  name: [
    { required: true, message: '请输入名称', trigger: 'blur' }
  ],
  description: [
    { required: true, message: '请输入描述', trigger: 'blur' },
    { min: 6, message: '最低6位字符', trigger: 'blur' },
  ],
  dataType: [
    { required: true, message: '请输入数据类型', trigger: 'blur' },
    { min: 6, message: '最低6位字符', trigger: 'blur' },
  ],
})

const machineForm = ref(null)
const enteraddMachineDialog = async() => {
  machineForm.value.ID = Number(machineForm.value.ID)
  machineForm.value.validate(async valid => {
    if (valid) {
      const req = {
        ...machineInfo.value
      }
      if (dialogFlag.value === 'add') {
        const res = await createMachine(req)
        if (res.code === 0) {
          ElMessage({ type: 'success', message: '创建成功' })
          await getTableData()
          closeaddMachineDialog()
        }
      }
      if (dialogFlag.value === 'edit') {
        const res = await setMachineInfo(req)
        if (res.code === 0) {
          ElMessage({ type: 'success', message: '编辑成功' })
          await getTableData()
          closeaddMachineDialog()
        }
      }
    }
  })
}

const addMachineDialog = ref(false)
const closeaddMachineDialog = () => {
  machineForm.value.resetFields()
  addMachineDialog.value = false
}

const dialogFlag = ref('add')

const addMachine = () => {
  dialogFlag.value = 'add'
  addMachineDialog.value = true
}

const tempAuth = {}

const openEdit = (row) => {
  dialogFlag.value = 'edit'
  machineInfo.value = JSON.parse(JSON.stringify(row))
  addMachineDialog.value = true
}

const switchEnable = async(row) => {
  machineInfo.value = JSON.parse(JSON.stringify(row))
  await nextTick()
  const req = {
    ...machineInfo.value
  }
  const res = await setMachineInfo(req)
  if (res.code === 0) {
    ElMessage({ type: 'success', message: `${req.enable === 2 ? '禁用' : '启用'}成功` })
    await getTableData()
  }
}

</script>

<style lang="scss">
  .header-img-box {
    @apply w-52 h-52 border border-solid border-gray-300 rounded-xl flex justify-center items-center cursor-pointer;
 }
</style>
