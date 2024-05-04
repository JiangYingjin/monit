<template>
  <div>
    <el-row :gutter="5">
      <el-form
        ref="elForm"
        :model="formData"
        size="medium"
        label-width="110px"
        label-position="left"
      >
        <el-col :span="24">
          <el-row :gutter="5">
            <el-col :span="12">
              <el-form-item label="选择命令模板" prop="select">
                <el-cascader
                  v-model="formData.select"
                  :options="selectOptions"
                  :props="selectProps"
                  :disabled="isCascaderDisabled"
                  :style="{ width: '100%' }"
                  placeholder="请选择选择命令模板"
                ></el-cascader>
              </el-form-item>
            </el-col>
            <el-col :span="12" class="button-container">
              <el-form-item size="medium">
                <el-button
                  type="primary"
                  @click="confirmForm"
                  :disabled="isConfirmDisabled"
                  >确定</el-button
                >
                <el-button @click="resetForm" :disabled="isResetDisabled"
                  >重置</el-button
                >
              </el-form-item>
            </el-col>
          </el-row>
        </el-col>
      </el-form>
    </el-row>
    <el-form v-if="isCascaderDisabled" label-position="left">
      <el-row :gutter="5">
        <el-col :span="24" v-for="item in templates" :key="item.name">
          <el-form-item :label="item.repr">
            <template v-if="item.type === 'bool'">
              <el-switch v-model="item.value"></el-switch>
            </template>
            <template v-else-if="item.type === 'str'">
              <el-input style="width: 300px" v-model="item.value"></el-input>
            </template>
            <template v-else-if="item.type === 'int'">
              <el-input-number v-model="item.value"></el-input-number>
            </template>
          </el-form-item>
        </el-col>
        <el-button type="primary" @click="submitForm">确定</el-button>
      </el-row>
    </el-form>
  </div>
</template>

<style>
.button-container {
  display: flex;
  align-items: center;
  justify-content: flex-start;
}
</style>

<script setup>
import { ref, onMounted } from "vue";
import { ElMessage } from "element-plus";

const formData = ref({
  select: [],
});

const selectOptions = ref([]);

const selectProps = {
  multiple: false,
};

const isCascaderDisabled = ref(false);
const isConfirmDisabled = ref(false);
const isResetDisabled = ref(true);

function confirmForm() {
  if (formData.value.select.length == 2) {
    // el-cascader 已选择
    // 执行提交操作
    console.log("confirm成功");
    isCascaderDisabled.value = true;
    isConfirmDisabled.value = true;
    isResetDisabled.value = false;
    getServiceTemplate(formData.value.select[1]);
  } else if (formData.value.select.length == 1) {
    // 选择的机器当前不支持任何服务
    ElMessage.error("您选择的机器当前不支持任何服务，请重新选择！");
  } else {
    // el-cascader 未选择
    ElMessage.error("请选择命令模板！");
  }
}

function resetForm() {
  console.log("reset成功");
  isCascaderDisabled.value = false;
  isConfirmDisabled.value = false;
  isResetDisabled.value = true;
}

function submitForm() {
  console.log("submit成功");
  console.log(templates);
  ElMessage.success("提交成功");
}

import { getMachineList } from "@/api/machine";

const getMachineInfo = async () => {
  const res = await getMachineList({});
  if (res.code === 0) {
    const machineInfo = res.data.list;
    generateOptions(machineInfo);
  }
};

function generateOptions(machineInfo) {
  machineInfo.forEach((item) => {
    const { ID, service } = item;
    const children = [];

    Object.entries(JSON.parse(service)).forEach(([key, value]) => {
      const childOption = {
        value: key,
        label: key,
      };
      children.push(childOption);
    });

    const option = {
      value: ID,
      label: ID.toString(), // 使用 ID 作为选项的 label，您也可以使用其他字段
      children,
    };

    selectOptions.value.push(option);
  });
}

import { getServiceTemplateList } from "@/api/serviceTemplate";

const getServiceTemplate = async (targetService) => {
  const res = await getServiceTemplateList({});
  if (res.code === 0) {
    const serviceTemplateInfo = res.data.list;
    generateTemplates(serviceTemplateInfo, targetService);
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
    }
  });
}

onMounted(() => {
  getMachineInfo();
});
</script>
