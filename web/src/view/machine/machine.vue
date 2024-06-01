<template>
  <div>
    <div class="gva-search-box">
      <el-form
        ref="elSearchFormRef"
        :inline="true"
        :model="searchInfo"
        class="demo-form-inline"
        :rules="searchRule"
        @keyup.enter="onSubmit"
      >
        <el-form-item label="创建日期" prop="createdAt">
          <template #label>
            <span>
              创建日期
              <el-tooltip
                content="搜索范围是开始日期（包含）至结束日期（不包含）"
              >
                <el-icon><QuestionFilled /></el-icon>
              </el-tooltip>
            </span>
          </template>
          <el-date-picker
            v-model="searchInfo.startCreatedAt"
            type="datetime"
            placeholder="开始日期"
            :disabled-date="
              (time) =>
                searchInfo.endCreatedAt
                  ? time.getTime() > searchInfo.endCreatedAt.getTime()
                  : false
            "
          ></el-date-picker>
          —
          <el-date-picker
            v-model="searchInfo.endCreatedAt"
            type="datetime"
            placeholder="结束日期"
            :disabled-date="
              (time) =>
                searchInfo.startCreatedAt
                  ? time.getTime() < searchInfo.startCreatedAt.getTime()
                  : false
            "
          ></el-date-picker>
        </el-form-item>

        <el-form-item>
          <el-button type="primary" icon="search" @click="onSubmit"
            >查询</el-button
          >
          <el-button icon="refresh" @click="onReset">重置</el-button>
        </el-form-item>
      </el-form>
    </div>
    <div class="gva-table-box">
      <div class="gva-btn-list">
        <el-button type="primary" icon="plus" @click="openDialog"
          >新增</el-button
        >
        <el-button
          icon="delete"
          style="margin-left: 10px"
          :disabled="!multipleSelection.length"
          @click="onDelete"
          >删除</el-button
        >
      </div>
      <el-table
        ref="multipleTable"
        style="width: 100%"
        tooltip-effect="dark"
        :data="tableData"
        row-key="ID"
        @selection-change="handleSelectionChange"
      >
        <el-table-column type="selection" width="55" />

        <el-table-column align="left" label="日期" width="180">
          <template #default="scope">{{
            formatDate(scope.row.CreatedAt)
          }}</template>
        </el-table-column>

        <el-table-column align="left" label="名字" prop="name" width="120" />
        <el-table-column
          align="left"
          label="描述"
          prop="description"
          width="120"
        />
        <el-table-column
          align="left"
          label="IP地址"
          prop="ip_addr"
          width="120"
        />
        <el-table-column
          align="left"
          label="密钥"
          prop="password"
          width="120"
        />
        <!-- <el-table-column
          align="left"
          label="在线状态"
          prop="status"
          width="120"
        >
          <template #default="scope">{{
            formatBoolean(scope.row.status)
          }}</template>
        </el-table-column>
        <el-table-column align="left" label="服务" prop="service" width="120" /> -->
        <el-table-column
          align="left"
          label="操作"
          fixed="right"
          min-width="240"
        >
          <template #default="scope">
            <el-button
              type="primary"
              link
              class="table-button"
              @click="getDetails(scope.row)"
            >
              <el-icon style="margin-right: 5px"><InfoFilled /></el-icon>
              查看详情
            </el-button>
            <el-button
              type="primary"
              link
              class="table-button"
              @click="getServiceList(scope.row)"
            >
              <el-icon style="margin-right: 5px"><Menu /></el-icon>
              状态监控
            </el-button>
            <el-button
              type="primary"
              link
              icon="edit"
              class="table-button"
              @click="updateMachineFunc(scope.row)"
              >变更</el-button
            >
            <el-button
              type="primary"
              link
              icon="delete"
              @click="deleteRow(scope.row)"
              >删除</el-button
            >
          </template>
        </el-table-column>
      </el-table>
      <div class="gva-pagination">
        <el-pagination
          layout="total, sizes, prev, pager, next, jumper"
          :current-page="page"
          :page-size="pageSize"
          :page-sizes="[10, 30, 50, 100]"
          :total="total"
          @current-change="handleCurrentChange"
          @size-change="handleSizeChange"
        />
      </div>
    </div>
    <el-drawer
      size="800"
      v-model="dialogFormVisible"
      :show-close="false"
      :before-close="closeDialog"
    >
      <template #title>
        <div class="flex justify-between items-center">
          <span class="text-lg">{{ type === "create" ? "添加" : "修改" }}</span>
          <div>
            <el-button type="primary" @click="enterDialog">确 定</el-button>
            <el-button @click="closeDialog">取 消</el-button>
          </div>
        </div>
      </template>

      <el-form
        :model="formData"
        label-position="top"
        ref="elFormRef"
        :rules="rule"
        label-width="80px"
      >
        <el-form-item label="名字:" prop="name">
          <el-input
            v-model="formData.name"
            :clearable="true"
            placeholder="请输入名字"
          />
        </el-form-item>
        <el-form-item label="描述:" prop="description">
          <el-input
            v-model="formData.description"
            :clearable="true"
            placeholder="请输入描述"
          />
        </el-form-item>
        <el-form-item label="IP地址:" prop="ip_addr">
          <el-input
            v-model="formData.ip_addr"
            :clearable="true"
            placeholder="请输入IP地址"
          />
        </el-form-item>
        <el-form-item label="密钥:" prop="password">
          <el-input
            v-model="formData.password"
            :clearable="true"
            placeholder="请输入密钥"
          />
        </el-form-item>
        <!-- <el-form-item label="在线状态:" prop="status">
          <el-switch
            v-model="formData.status"
            active-color="#13ce66"
            inactive-color="#ff4949"
            active-text="是"
            inactive-text="否"
            clearable
            :disabled="true"
          ></el-switch>
        </el-form-item>
        <el-form-item label="服务:" prop="service">
          <el-input
            v-model="formData.service"
            :clearable="true"
            placeholder="请输入服务"
            :disabled="true"
          />
        </el-form-item> -->
      </el-form>
    </el-drawer>

    <el-drawer
      size="800"
      v-model="detailShow"
      :before-close="closeDetailShow"
      title="查看详情"
      destroy-on-close
    >
      <template #title>
        <div class="flex justify-between items-center">
          <span class="text-lg">查看详情</span>
        </div>
      </template>
      <el-descriptions :column="1" border>
        <el-descriptions-item label="名字">
          {{ formData.name }}
        </el-descriptions-item>
        <el-descriptions-item label="描述">
          {{ formData.description }}
        </el-descriptions-item>
        <el-descriptions-item label="IP地址">
          {{ formData.ip_addr }}
        </el-descriptions-item>
        <el-descriptions-item label="密钥">
          {{ formData.password }}
        </el-descriptions-item>
        <el-descriptions-item label="在线状态">
          {{ formatBoolean(formData.status) }}
        </el-descriptions-item>
        <el-descriptions-item label="服务">
          {{ formData.service }}
        </el-descriptions-item>
      </el-descriptions>
    </el-drawer>

    <el-drawer
      v-model="drawerVisible"
      title="状态监控"
      size="800"
      @close="handleDrawerClose"
      destroy-on-close
    >
      <div class="flex justify-between items-center">
        <el-select
          v-model="selectedService"
          placeholder="请选择服务"
          clearable
          @change="confirmForm"
          @clear="resetForm"
        >
          <el-option
            v-for="service in serviceList"
            :key="service.name"
            :label="service.name"
            :value="service.name"
          >
            <div class="flex items-center">
              <span
                :style="{ backgroundColor: service.isActive ? 'green' : 'red' }"
                class="dot"
              ></span>
              {{ service.name }}
            </div>
          </el-option>
        </el-select>
      </div>
      <el-form v-if="isSelectDisabled" label-position="left">
        <el-row :gutter="5">
          <el-col :span="24" v-for="item in templates" :key="item.name">
            <el-form-item
              v-if="item.type === 'bool' || isComponentVisible"
              :label="item.repr"
            >
              <template v-if="item.type === 'bool'">
                <el-switch
                  v-model="item.value"
                  @change="handleSwitchChange(item.name, item.value)"
                ></el-switch>
              </template>
              <template v-else-if="item.type === 'str' && isComponentVisible">
                <el-input style="width: 300px" v-model="item.value"></el-input>
              </template>
              <template v-else-if="item.type === 'int' && isComponentVisible">
                <el-input-number v-model="item.value"></el-input-number>
              </template>
            </el-form-item>
          </el-col>
          <el-button type="primary" @click="submitForm">确定</el-button>
        </el-row>
      </el-form>
    </el-drawer>
  </div>
</template>

<script setup>
import {
  createMachine,
  deleteMachine,
  deleteMachineByIds,
  updateMachine,
  findMachine,
  getMachineList,
  setMachineService,
} from "@/api/machine";

// 全量引入格式化工具 请按需保留
import {
  getDictFunc,
  formatDate,
  formatBoolean,
  filterDict,
  ReturnArrImg,
  onDownloadFile,
} from "@/utils/format";
import { ElMessage, ElMessageBox } from "element-plus";
import { ref, reactive } from "vue";

defineOptions({
  name: "Machine",
});

// 自动化生成的字典（可能为空）以及字段
const formData = ref({
  name: "",
  description: "",
  ip_addr: "",
  password: "",
  status: false,
  service: "",
});

const drawerVisible = ref(false);
const isSelectDisabled = ref(false);
const isComponentVisible = ref(true);
const selectedService = ref(null);
const serviceList = ref([]);
const currentMachineID = ref(null);

const getServiceList = async (row) => {
  currentMachineID.value = row.ID;
  const res = await findMachine({ ID: row.ID });
  if (res.code === 0) {
    serviceList.value = Object.entries(
      JSON.parse(res.data.remachine.service)
    ).map(([name, status]) => ({
      name,
      isActive: status === 1,
    }));
    if (serviceList.value.length == 0)
      ElMessage.error("您选择的机器当前不支持任何服务，请重新选择！");
    else {
      // console.log(serviceList.value);
      drawerVisible.value = true;
    }
  }
};

function confirmForm() {
  if (typeof selectedService.value == "string") {
    // console.log("confirm成功");
    isSelectDisabled.value = true;
    isComponentVisible.value = true;
    getServiceTemplate(selectedService.value);
  }
}

import { getServiceTemplateList } from "@/api/serviceTemplate";

const templates = ref([]);

const getServiceTemplate = async (serviceName) => {
  const res = await getServiceTemplateList({});
  if (res.code === 0) {
    const serviceTemplateInfo = res.data.list;
    generateTemplates(serviceTemplateInfo, serviceName);
  }
};

function generateTemplates(serviceTemplateInfo, targetService) {
  serviceTemplateInfo.forEach((item) => {
    const { service, template } = item;
    if (service === targetService) {
      templates.value = JSON.parse(template).map((option) => {
        const { default: defaultValue, ...rest } = option;
        let value = defaultValue;

        if (rest.type === "bool") {
          value = Boolean(defaultValue);
        }

        return { ...rest, value };
      });
      // console.log(templates.value);
    }
  });
}

function resetForm() {
  // console.log("reset成功");
  isSelectDisabled.value = false;
  isComponentVisible.value = true;
  selectedService.value = null;
  templates.value = [];
}

function submitForm() {
  // console.log("submit成功");
  ElMessage.success("提交成功");
  const result = {
    machineID: currentMachineID.value,
    services: {
      [selectedService.value]: Object.keys(templates.value).map((index) => ({
        name: templates.value[index].name,
        type: templates.value[index].type,
        value:
          templates.value[index].type === "bool"
            ? templates.value[index].value
              ? "1"
              : "0"
            : templates.value[index].value.toString(),
      })),
    },
  };
  console.log(result);
  updateMachineServiceFunc(result);
}

const updateMachineServiceFunc = async (data) => {
  const res = await setMachineService({ data });
  console.log(res);
  if (res.code === 0) {
    ElMessage.success("更新成功");
  } else {
    ElMessage.error("更新失败");
  }
};

const handleDrawerClose = () => {
  drawerVisible.value = false;
  isSelectDisabled.value = false;
  isComponentVisible.value = true;
  selectedService.value = null;
  currentMachineID.value = null;
  templates.value = [];
  serviceList.value = [];
};

function handleSwitchChange(name, value) {
  // 当 item.value 变为 false 时,隐藏其他表单项
  if (!value && name === "enable") {
    isComponentVisible.value = false;
    console.log(isComponentVisible.value);
  } else if (value && name === "enable") {
    // 当 item.value 变为 true 时,根据 item.type 显示对应的表单项
    isComponentVisible.value = true;
    console.log(isComponentVisible.value);
  }
}

// 验证规则
const rule = reactive({
  name: [
    {
      required: true,
      message: "",
      trigger: ["input", "blur"],
    },
    {
      whitespace: true,
      message: "不能只输入空格",
      trigger: ["input", "blur"],
    },
  ],
  description: [
    {
      required: true,
      message: "",
      trigger: ["input", "blur"],
    },
    {
      whitespace: true,
      message: "不能只输入空格",
      trigger: ["input", "blur"],
    },
  ],
  ip_addr: [
    {
      required: true,
      message: "",
      trigger: ["input", "blur"],
    },
    {
      whitespace: true,
      message: "不能只输入空格",
      trigger: ["input", "blur"],
    },
  ],
  password: [
    {
      required: true,
      message: "",
      trigger: ["input", "blur"],
    },
    {
      whitespace: true,
      message: "不能只输入空格",
      trigger: ["input", "blur"],
    },
  ],
});

const searchRule = reactive({
  createdAt: [
    {
      validator: (rule, value, callback) => {
        if (searchInfo.value.startCreatedAt && !searchInfo.value.endCreatedAt) {
          callback(new Error("请填写结束日期"));
        } else if (
          !searchInfo.value.startCreatedAt &&
          searchInfo.value.endCreatedAt
        ) {
          callback(new Error("请填写开始日期"));
        } else if (
          searchInfo.value.startCreatedAt &&
          searchInfo.value.endCreatedAt &&
          (searchInfo.value.startCreatedAt.getTime() ===
            searchInfo.value.endCreatedAt.getTime() ||
            searchInfo.value.startCreatedAt.getTime() >
              searchInfo.value.endCreatedAt.getTime())
        ) {
          callback(new Error("开始日期应当早于结束日期"));
        } else {
          callback();
        }
      },
      trigger: "change",
    },
  ],
});

const elFormRef = ref();
const elSearchFormRef = ref();

// =========== 表格控制部分 ===========
const page = ref(1);
const total = ref(0);
const pageSize = ref(10);
const tableData = ref([]);
const searchInfo = ref({});

// 重置
const onReset = () => {
  searchInfo.value = {};
  getTableData();
};

// 搜索
const onSubmit = () => {
  elSearchFormRef.value?.validate(async (valid) => {
    if (!valid) return;
    page.value = 1;
    pageSize.value = 10;
    if (searchInfo.value.status === "") {
      searchInfo.value.status = null;
    }
    getTableData();
  });
};

// 分页
const handleSizeChange = (val) => {
  pageSize.value = val;
  getTableData();
};

// 修改页面容量
const handleCurrentChange = (val) => {
  page.value = val;
  getTableData();
};

// 查询
const getTableData = async () => {
  const table = await getMachineList({
    page: page.value,
    pageSize: pageSize.value,
    ...searchInfo.value,
  });
  if (table.code === 0) {
    tableData.value = table.data.list;
    total.value = table.data.total;
    page.value = table.data.page;
    pageSize.value = table.data.pageSize;
  }
};

getTableData();

// ============== 表格控制部分结束 ===============

// 获取需要的字典 可能为空 按需保留
const setOptions = async () => {};

// 获取需要的字典 可能为空 按需保留
setOptions();

// 多选数据
const multipleSelection = ref([]);
// 多选
const handleSelectionChange = (val) => {
  multipleSelection.value = val;
};

// 删除行
const deleteRow = (row) => {
  ElMessageBox.confirm("确定要删除吗?", "提示", {
    confirmButtonText: "确定",
    cancelButtonText: "取消",
    type: "warning",
  }).then(() => {
    deleteMachineFunc(row);
  });
};

// 多选删除
const onDelete = async () => {
  ElMessageBox.confirm("确定要删除吗?", "提示", {
    confirmButtonText: "确定",
    cancelButtonText: "取消",
    type: "warning",
  }).then(async () => {
    const IDs = [];
    if (multipleSelection.value.length === 0) {
      ElMessage({
        type: "warning",
        message: "请选择要删除的数据",
      });
      return;
    }
    multipleSelection.value &&
      multipleSelection.value.map((item) => {
        IDs.push(item.ID);
      });
    const res = await deleteMachineByIds({ IDs });
    if (res.code === 0) {
      ElMessage({
        type: "success",
        message: "删除成功",
      });
      if (tableData.value.length === IDs.length && page.value > 1) {
        page.value--;
      }
      getTableData();
    }
  });
};

// 行为控制标记（弹窗内部需要增还是改）
const type = ref("");

// 更新行
const updateMachineFunc = async (row) => {
  const res = await findMachine({ ID: row.ID });
  type.value = "update";
  if (res.code === 0) {
    formData.value = res.data.remachine;
    dialogFormVisible.value = true;
  }
};

// 删除行
const deleteMachineFunc = async (row) => {
  const res = await deleteMachine({ ID: row.ID });
  if (res.code === 0) {
    ElMessage({
      type: "success",
      message: "删除成功",
    });
    if (tableData.value.length === 1 && page.value > 1) {
      page.value--;
    }
    getTableData();
  }
};

// 弹窗控制标记
const dialogFormVisible = ref(false);

// 查看详情控制标记
const detailShow = ref(false);

// 打开详情弹窗
const openDetailShow = () => {
  detailShow.value = true;
};

// 打开详情
const getDetails = async (row) => {
  // 打开弹窗
  const res = await findMachine({ ID: row.ID });
  if (res.code === 0) {
    formData.value = res.data.remachine;
    openDetailShow();
  }
};

// 关闭详情弹窗
const closeDetailShow = () => {
  detailShow.value = false;
  formData.value = {
    name: "",
    description: "",
    ip_addr: "",
    password: "",
    status: false,
    service: "",
  };
};

// 打开弹窗
const openDialog = () => {
  type.value = "create";
  dialogFormVisible.value = true;
};

// 关闭弹窗
const closeDialog = () => {
  dialogFormVisible.value = false;
  formData.value = {
    name: "",
    description: "",
    ip_addr: "",
    password: "",
    status: false,
    service: "",
  };
};

// 弹窗确定
const enterDialog = async () => {
  elFormRef.value?.validate(async (valid) => {
    if (!valid) return;
    let res;
    switch (type.value) {
      case "create":
        res = await createMachine(formData.value);
        console.log(formData.value);
        console.log(res);
        break;
      case "update":
        res = await updateMachine(formData.value);
        break;
      default:
        res = await createMachine(formData.value);
        break;
    }
    if (res.code === 0) {
      ElMessage({
        type: "success",
        message: "创建/更改成功",
      });
      closeDialog();
      getTableData();
    }
  });
};
</script>

<style>
.dot {
  display: inline-block;
  width: 8px;
  height: 8px;
  border-radius: 50%;
  margin-right: 5px;
}
</style>
