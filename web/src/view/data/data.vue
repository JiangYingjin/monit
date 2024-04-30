<template>

  <div>
    <div class="data-center-box">
      <CenterCard title="设备在线情况">
        <template #action>
          <span
            class="gvaIcon-prompt"
            style="color: #999"
          />
        </template>
        <template #body>
          <ChainRatio />
        </template>
      </CenterCard>
      <CenterCard title="内存使用率"
                  style="grid-column-start: span 3;">
        <template #action>
          <span
              class="gvaIcon-prompt"
              style="color: #999"
          />
        </template>
        <template #body>
<!--          <Order />-->
          <div
              id="chart1"
              style="width: 100%; height: 200px;"
          />
        </template>
      </CenterCard>
<!--      <CenterCard title="回收里程(Km)">-->
<!--        <template #action>-->
<!--          <span-->
<!--            class="gvaIcon-prompt"-->
<!--            style="color: #999"-->
<!--          />-->
<!--        </template>-->
<!--        <template #body>-->
<!--          <ReclaimMileage />-->
<!--        </template>-->
<!--      </CenterCard>-->
<!--      <CenterCard title="设备回收率">-->
<!--        <template #action>-->
<!--          <span-->
<!--            class="gvaIcon-prompt"-->
<!--            style="color: #999"-->
<!--          />-->
<!--        </template>-->
<!--        <template #body>-->
<!--          <RecoveryRate />-->
<!--        </template>-->
<!--      </CenterCard>-->

      <CenterCard
        title="CPU占用率"
        style="grid-column-start: span 4;"
      >
        <template #action>
          <span class="gvaIcon-prompt" style="color: #999" />
        </template>
        <template #body>
          <!-- 循环渲染 initialCharts -->
          <!--    <div-->
          <!--      id="chart"-->
          <!--      style="width: 800px; height: 300px;"-->
          <!--    />-->
          <div
              id="chart"
              style="width: 100%; height: 360px;"
          />
        </template>
      </CenterCard>
<!--      <CenterCard-->
<!--        title="源接入台数占比"-->
<!--        style="grid-auto-columns: 1.5fr"-->
<!--      >-->
<!--        <template #action>-->
<!--          <span class="gvaIcon-prompt" style="color: #999" />-->
<!--        </template>-->
<!--        <template #body>-->
<!--          &lt;!&ndash; 循环渲染 initialCharts &ndash;&gt;-->
<!--          &lt;!&ndash;    <div&ndash;&gt;-->
<!--          &lt;!&ndash;      id="chart"&ndash;&gt;-->
<!--          &lt;!&ndash;      style="width: 800px; height: 300px;"&ndash;&gt;-->
<!--          &lt;!&ndash;    />&ndash;&gt;-->
<!--        </template>-->
<!--      </CenterCard>-->
    </div>
    <br>
    <br>
    <!--    <div class="gva-search-box">-->
    <!--      <el-form-->
    <!--        ref="elSearchFormRef"-->
    <!--        :inline="true"-->
    <!--        :model="searchInfo"-->
    <!--        class="demo-form-inline"-->
    <!--        :rules="searchRule"-->
    <!--        @keyup.enter="onSubmit"-->
    <!--      >-->
    <!--        <el-form-item-->
    <!--          label="创建日期"-->
    <!--          prop="createdAt"-->
    <!--        >-->
    <!--          <template #label>-->
    <!--            <span>-->
    <!--              创建日期-->
    <!--              <el-tooltip content="搜索范围是开始日期（包含）至结束日期（不包含）">-->
    <!--                <el-icon><QuestionFilled /></el-icon>-->
    <!--              </el-tooltip>-->
    <!--            </span>-->
    <!--          </template>-->
    <!--          <el-date-picker-->
    <!--            v-model="searchInfo.startCreatedAt"-->
    <!--            type="datetime"-->
    <!--            placeholder="开始日期"-->
    <!--            :disabled-date="time=> searchInfo.endCreatedAt ? time.getTime() > searchInfo.endCreatedAt.getTime() : false"-->
    <!--          />-->
    <!--          —-->
    <!--          <el-date-picker-->
    <!--            v-model="searchInfo.endCreatedAt"-->
    <!--            type="datetime"-->
    <!--            placeholder="结束日期"-->
    <!--            :disabled-date="time=> searchInfo.startCreatedAt ? time.getTime() < searchInfo.startCreatedAt.getTime() : false"-->
    <!--          />-->
    <!--        </el-form-item>-->

    <!--        <el-form-item>-->
    <!--          <el-button-->
    <!--            type="primary"-->
    <!--            icon="search"-->
    <!--            @click="onSubmit"-->
    <!--          >查询</el-button>-->
    <!--          <el-button-->
    <!--            icon="refresh"-->
    <!--            @click="onReset"-->
    <!--          >重置</el-button>-->
    <!--        </el-form-item>-->
    <!--      </el-form>-->
    <!--    </div>-->
    <!--    <div class="gva-table-box">-->
    <!--      <div class="gva-btn-list">-->
    <!--        <el-button-->
    <!--          type="primary"-->
    <!--          icon="plus"-->
    <!--          @click="openDialog"-->
    <!--        >新增</el-button>-->
    <!--        <el-button-->
    <!--          icon="delete"-->
    <!--          style="margin-left: 10px;"-->
    <!--          :disabled="!multipleSelection.length"-->
    <!--          @click="onDelete"-->
    <!--        >删除</el-button>-->
    <!--      </div>-->
    <!--      <el-table-->
    <!--        ref="multipleTable"-->
    <!--        style="width: 100%"-->
    <!--        tooltip-effect="dark"-->
    <!--        :data="tableData"-->
    <!--        row-key="ID"-->
    <!--        @selection-change="handleSelectionChange"-->
    <!--      >-->
    <!--        <el-table-column-->
    <!--          type="selection"-->
    <!--          width="55"-->
    <!--        />-->

    <!--        <el-table-column-->
    <!--          align="left"-->
    <!--          label="日期"-->
    <!--          width="180"-->
    <!--        >-->
    <!--          <template #default="scope">{{ formatDate(scope.row.CreatedAt) }}</template>-->
    <!--        </el-table-column>-->

    <!--        <el-table-column-->
    <!--          align="left"-->
    <!--          label="数据类型"-->
    <!--          width="120"-->
    <!--        >-->
    <!--          <template #default="scope">-->
    <!--            {{ dataTypeMap[scope.row.dataTypeID.toString()] }}-->
    <!--          </template>-->
    <!--        </el-table-column>-->
    <!--        <el-table-column-->
    <!--          align="left"-->
    <!--          label="值"-->
    <!--          prop="value"-->
    <!--          width="120"-->
    <!--        />-->
    <!--        <el-table-column-->
    <!--          align="left"-->
    <!--          label="机器ID"-->
    <!--          prop="machineID"-->
    <!--          width="120"-->
    <!--        />-->
    <!--        <el-table-column-->
    <!--          align="left"-->
    <!--          label="操作"-->
    <!--          fixed="right"-->
    <!--          min-width="240"-->
    <!--        >-->
    <!--          <template #default="scope">-->
    <!--            <el-button-->
    <!--              type="primary"-->
    <!--              link-->
    <!--              class="table-button"-->
    <!--              @click="getDetails(scope.row)"-->
    <!--            >-->
    <!--              <el-icon style="margin-right: 5px"><InfoFilled /></el-icon>-->
    <!--              查看详情-->
    <!--            </el-button>-->
    <!--            <el-button-->
    <!--              type="primary"-->
    <!--              link-->
    <!--              icon="edit"-->
    <!--              class="table-button"-->
    <!--              @click="updateDataFunc(scope.row)"-->
    <!--            >变更</el-button>-->
    <!--            <el-button-->
    <!--              type="primary"-->
    <!--              link-->
    <!--              icon="delete"-->
    <!--              @click="deleteRow(scope.row)"-->
    <!--            >删除</el-button>-->
    <!--          </template>-->
    <!--        </el-table-column>-->
    <!--      </el-table>-->
    <!--      <div class="gva-pagination">-->
    <!--        <el-pagination-->
    <!--          layout="total, sizes, prev, pager, next, jumper"-->
    <!--          :current-page="page"-->
    <!--          :page-size="pageSize"-->
    <!--          :page-sizes="[10, 30, 50, 100]"-->
    <!--          :total="total"-->
    <!--          @current-change="handleCurrentChange"-->
    <!--          @size-change="handleSizeChange"-->
    <!--        />-->
    <!--      </div>-->
    <!--    </div>-->
    <!--    <el-drawer-->
    <!--      v-model="dialogFormVisible"-->
    <!--      size="800"-->
    <!--      :show-close="false"-->
    <!--      :before-close="closeDialog"-->
    <!--    >-->
    <!--      <template #title>-->
    <!--        <div class="flex justify-between items-center">-->
    <!--          <span class="text-lg">{{ type==='create'?'添加':'修改' }}</span>-->
    <!--          <div>-->
    <!--            <el-button-->
    <!--              type="primary"-->
    <!--              @click="enterDialog"-->
    <!--            >确 定</el-button>-->
    <!--            <el-button @click="closeDialog">取 消</el-button>-->
    <!--          </div>-->
    <!--        </div>-->
    <!--      </template>-->

    <!--      <el-form-->
    <!--        ref="elFormRef"-->
    <!--        :model="formData"-->
    <!--        label-position="top"-->
    <!--        :rules="rule"-->
    <!--        label-width="80px"-->
    <!--      >-->
    <!--        <el-form-item-->
    <!--          label="数据类型:"-->
    <!--          prop="dataTypeID"-->
    <!--        >-->
    <!--          <el-input-->
    <!--            v-model.number="formData.dataTypeID"-->
    <!--            :clearable="true"-->
    <!--            placeholder="请输入数据类型"-->
    <!--          />-->
    <!--        </el-form-item>-->
    <!--        <el-form-item-->
    <!--          label="值:"-->
    <!--          prop="value"-->
    <!--        >-->
    <!--          <el-input-number-->
    <!--            v-model="formData.value"-->
    <!--            style="width:100%"-->
    <!--            :precision="2"-->
    <!--            :clearable="true"-->
    <!--          />-->
    <!--        </el-form-item>-->
    <!--        <el-form-item-->
    <!--          label="机器ID:"-->
    <!--          prop="machineID"-->
    <!--        >-->
    <!--          <el-input-->
    <!--            v-model.number="formData.machineID"-->
    <!--            :clearable="true"-->
    <!--            placeholder="请输入机器ID"-->
    <!--          />-->
    <!--        </el-form-item>-->
    <!--      </el-form>-->
    <!--    </el-drawer>-->

    <!--    <el-drawer-->
    <!--      v-model="detailShow"-->
    <!--      size="800"-->
    <!--      :before-close="closeDetailShow"-->
    <!--      title="查看详情"-->
    <!--      destroy-on-close-->
    <!--    >-->
    <!--      <template #title>-->
    <!--        <div class="flex justify-between items-center">-->
    <!--          <span class="text-lg">查看详情</span>-->
    <!--        </div>-->
    <!--      </template>-->
    <!--      <el-descriptions-->
    <!--        :column="1"-->
    <!--        border-->
    <!--      >-->
    <!--        <el-descriptions-item label="数据类型">-->
    <!--          {{ formData.dataTypeID }}-->
    <!--        </el-descriptions-item>-->
    <!--        <el-descriptions-item label="值">-->
    <!--          {{ formData.value }}-->
    <!--        </el-descriptions-item>-->
    <!--        <el-descriptions-item label="机器ID">-->
    <!--          {{ formData.machineID }}-->
    <!--        </el-descriptions-item>-->
    <!--      </el-descriptions>-->
    <!--    </el-drawer>-->
    <!--    <br>-->

    <div class="filter-box">
      <el-form
        :inline="true"
        :model="filterForm"
        class="filter-form"
        size="small"
        @submit.native.prevent="handleSubmit"
      >
        <el-form-item label="起始时间">
          <el-date-picker
            v-model="filterForm.start_time"
            type="datetime"
            placeholder="选择起始时间"
            style="width: 200px;"
          />
        </el-form-item>
        <el-form-item label="结束时间">
          <el-date-picker
            v-model="filterForm.end_time"
            type="datetime"
            placeholder="选择结束时间"
            style="width: 200px;"
          />
        </el-form-item>
        <el-form-item label="数据类型">
          <el-checkbox
            v-model="allDataTypesChecked"
            @change="handleAllDataTypesChange"
          >全选</el-checkbox>
          <el-select
            v-model="filterForm.dataType"
            multiple
            filterable
            collapse-tags
            placeholder="选择数据类型"
            style="width: 200px;"
            @change="handleDataTypeChange"
          >
            <el-option
              v-for="type in dataTypes"
              :key="type.id"
              :label="type.name"
              :value="type.id"
            />
          </el-select>

        </el-form-item>
        <el-form-item label="机器">
          <el-checkbox
            v-model="allMachinesChecked"
            @change="handleAllMachinesChange"
          >全选</el-checkbox>
          <el-select
            v-model="filterForm.machines"
            multiple
            filterable
            collapse-tags
            placeholder="选择机器"
            style="width: 150px;"
            @change="handleMachinesChange"
          >
            <el-option
              v-for="machine in machines"
              :key="machine.id"
              :label="machine.name"
              :value="machine.id"
            />
          </el-select>
        </el-form-item>
        <el-form-item>
          <el-button
            type="primary"
            native-type="submit"
          >筛选</el-button>
        </el-form-item>
      </el-form>
    </div>
        <div
          v-for="config in chartConfigs"
          :id="config.id"
          :key="config.id"
          style="width: 1200px; height: 500px;"
        />
<!--    <div class="chart-container">-->
<!--      <div class="chart-row" v-for="(row, index) in Math.ceil(chartConfigs.length / 2)" :key="index">-->
<!--        <div v-for="(config, innerIndex) in chartConfigs.slice(index * 2, (index + 1) * 2)" :id="config.id" :key="innerIndex" class="chart-item" style="width: calc(50% - 10px); height: 500px; margin-right: 20px;">-->
<!--          &lt;!&ndash; 这里放置你的曲线图表 &ndash;&gt;-->
<!--          &lt;!&ndash; 每个图表的宽度是50%，减去margin-right的值 &ndash;&gt;-->
<!--        </div>-->
<!--      </div>-->
<!--    </div>-->

    <!--    <div-->
<!--      v-for="config in chartConfigs"-->
<!--      :id="config.id"-->
<!--      :key="config.id"-->
<!--      style="width: 700px;height: 500px"-->
<!--    />-->
    <br>
<!--    <el-radio-group v-model="layout" size="small">-->
<!--      <el-radio label="single">一行一个图表</el-radio>-->
<!--      <el-radio label="double">一行两个图表</el-radio>-->
<!--    </el-radio-group>-->
    <br>
    <br>
<!--    <div class="chart-container">-->
<!--      <el-space>-->
<!--        <div-->
<!--          v-for="(config, index) in chartConfigs"-->
<!--          :key="config.id"-->
<!--          :id="config.id"-->
<!--          class="chart-item"-->
<!--          :style="{ width: getChartWidth() + '%', height: '150px' }"-->
<!--        >-->
<!--          &lt;!&ndash; 这里放置您的曲线图表 &ndash;&gt;-->
<!--        </div>-->
<!--      </el-space>-->
<!--    </div>-->

  </div>

</template>

<script setup>
import {
  createData,
  deleteData,
  deleteDataByIds,
  updateData,
  findData,
  getDataList,
  getData,
} from '@/api/data'
import { getMachineList } from '@/api/machine'
import lineCharts from './dataCenterComponents/lineCharts.vue'

// 全量引入格式化工具 请按需保留
import { getDictFunc, formatDate, formatBoolean, filterDict, ReturnArrImg, onDownloadFile } from '@/utils/format'
import { ElMessage, ElMessageBox } from 'element-plus'
import { ref, reactive } from 'vue'
import { nextTick } from 'vue'

defineOptions({
  name: 'Data'
})
import * as d3 from 'd3' // 使用ES模块导入方式
import { onMounted } from 'vue' // 使用Vue 3提供的onMounted函数
import * as echarts from 'echarts'
import { getDataTypeList } from '@/api/dataType'
import { timestamp } from '@vueuse/core'
import CenterCard from '@/view/dataCenter/dataCenterComponents/centerCard.vue'
import ReclaimMileage from '@/view/dataCenter/dataCenterComponents/ReclaimMileage.vue'
import Order from '@/view/dataCenter/dataCenterComponents/order.vue'
import RecoveryRate from '@/view/dataCenter/dataCenterComponents/RecoveryRate.vue'
import ChainRatio from '@/view/dataCenter/dataCenterComponents/chainRatio.vue'
import part from '../dataCenter/dataCenterComponents/part.vue'
import { FALSE } from 'sass'
import colors from 'tailwindcss/colors' // 导入echarts库
import { computed } from 'vue';
onMounted(() => {
  // getData123()
  // renderChart()
})
const startColor = '#75BFA5' // 起始颜色
const endColor = '#FFA500' // 结束颜色
const colorCount = 0 // 需要的颜色数量

function interpolateColor(startColor, endColor, colorCount) {
  const colorInterpolator = d3.interpolate(startColor, endColor)
  const colorScale = d3.scaleLinear().domain([0, colorCount - 1]).range([0, 1])

  const colors = []
  for (let i = 0; i < colorCount; i++) {
    const t = colorScale(i)
    const interpolatedColor = colorInterpolator(t)
    colors.push(interpolatedColor)
  }

  return colors
}


const getChartsPerRow = computed(() => {
  return layout.value === 'single' ? 1 : 2; // 如果是单行，返回1；如果是两行，返回2
});

// 根据布局类型获取图表容器的宽度
const getChartWidth = computed(() => {
  return 800 / getChartsPerRow.value; // 计算每个图表容器的宽度
});

function renderCharts() {
  chartConfigs.value.forEach(config => {
    // echarts.dispose(document.getElementById(config.id))
    const chart = echarts.init(document.getElementById(config.id), 'light')
    const colors = interpolateColor(startColor, endColor, config.data.length)
    console.log(config.data)
    chart.setOption({
      title: {
        text: config.title,
        left: 'center',
        top: '10px'
      },
      backgroundColor: 'transparent',
      tooltip: {
        trigger: 'axis',
        backgroundColor: 'transparent',
      },
      legend: {
        align: 'left',
        right: '5%',
        top: '0%',
        type: 'plain',
        itemGap: 25,
        itemWidth: 20,
        // icon: 'path://M0 2a2 2 0 0 1 2 -2h14a2 2 0 0 1 2 2v0a2 2 0 0 1 -2 2h-14a2 2 0 0 1 -2 -2z',
        data: config.data.map((series, index) => ({
          name: series.name,
          textStyle: {
            color: colors[index], // 使用相应的颜色
          },
        })),
      },
      // legend: {
      //   data: config.data.map(series => series),
      //   right: 20, // 调整图例到右边的距离
      //   top: 20, // 调整图例到顶部的距离// 根据配置的数据动态生成 legend 数据
      // },
      xAxis: {
        type: 'time',
        // data: globalTimeList,
        boundaryGap: false,
        axisLine: {
          // 坐标轴轴线相关设置。数学上的x轴
          show: false,
          lineStyle: {
            color: '#e1e1e1',
          },
        },
        axisLabel: {
          // 坐标轴刻度标签的相关设置
          textStyle: {
            color: '#92969E',
          },
          // formatter: function(data) {
          //   return data
          // },
        },
        splitLine: {
          show: false,
          lineStyle: {
            color: '#192a44',
          },
        },
        axisTick: {
          show: false,
        },
      },
      yAxis: {
        type: 'value',
        nameTextStyle: {
          color: '#777',
        },
        min: 0,
        splitLine: {
          show: true,
          lineStyle: {
            color: '#e1e1e1',
          },
        },
        axisLine: {
          show: false,
        },
        axisLabel: {
          show: true,
          textStyle: {
            color: '#92969E',
          },
        },
        axisTick: {
          show: false,
        },
        // max: 10 // 设置y轴最大值为100
      },
      // 其他配置项
    })
    // 渲染图表数据
    chart.setOption({
      series: config.data.map((series, index) => ({
        type: 'line',
        name: series.name,
        data: series.data.map(item => [item.time, item.value]),
        symbol: 'circle', // 默认是空心圆（中间是白色的），改成实心圆
        showAllSymbol: true,
        symbolSize: 0,
        smooth: true,
        lineStyle: {
          normal: {
            width: 2,
            color: colors[index], // 线条颜色
          },
        },
        itemStyle: {
          color: colors[index], // 使用相应的颜色
        },
        // areaStyle: {
        //   // 区域填充样式
        //   normal: {
        //     color: new echarts.graphic.LinearGradient(0, 0, 0, 1, [
        //       {
        //         offset: 0,
        //         color: 'rgba(50, 216, 205, .8)'
        //       },
        //       {
        //         offset: 1,
        //         color: 'rgba(255, 255, 255, 0.2)'
        //       }
        //     ], false),
        //     shadowColor: 'rgba(117,191,165,0.52)', // 阴影颜色
        //     shadowBlur: 3,
        //   },
        // },
        areaStyle: {
          normal: {
            color: new echarts.graphic.LinearGradient(0, 0, 0, 1, [
              {
                offset: 0,
                color: colors[index], // 使用相应的颜色
              },
              {
                offset: 1,
                color: 'rgba(255, 255, 255, 0.2)',
              }
            ], false),
            shadowColor: 'rgba(117,191,165,0.52)',
            shadowBlur: 3,
          },
        },
        // 其他配置...
      }))
    })
  })
}

function init_charts() {
  const config = initialCharts.value[0]
  // echarts.dispose(document.getElementById(config.id))
  echarts.dispose(document.getElementById('chart'))
  const chart = echarts.init(document.getElementById('chart'), 'light')
  const colors = interpolateColor(startColor, endColor, config.data.length)
  chart.setOption({
    title: {
      text: config.title,
      left: 'center',
      top: '10px'
    },
    backgroundColor: 'transparent',
    tooltip: {
      trigger: 'axis',
      backgroundColor: 'transparent',
    },
    legend: {
      align: 'left',
      right: '5%',
      top: '0%',
      type: 'plain',
      itemGap: 25,
      itemWidth: 20,
      // icon: 'path://M0 2a2 2 0 0 1 2 -2h14a2 2 0 0 1 2 2v0a2 2 0 0 1 -2 2h-14a2 2 0 0 1 -2 -2z',
      data: config.data.map((series, index) => ({
        name: series.name,
        textStyle: {
          color: colors[index], // 使用相应的颜色
        },
      })),
    },
    // legend: {
    //   data: config.data.map(series => series),
    //   right: 20, // 调整图例到右边的距离
    //   top: 20, // 调整图例到顶部的距离// 根据配置的数据动态生成 legend 数据
    // },
    xAxis: {
      type: 'time',
      // data: globalTimeList,
      boundaryGap: false,
      axisLine: {
        // 坐标轴轴线相关设置。数学上的x轴
        show: false,
        lineStyle: {
          color: '#e1e1e1',
        },
      },
      axisLabel: {
        // 坐标轴刻度标签的相关设置
        textStyle: {
          color: '#92969E',
        },
        // formatter: function(data) {
        //   return data
        // },
      },
      splitLine: {
        show: false,
        lineStyle: {
          color: '#192a44',
        },
      },
      axisTick: {
        show: false,
      },
    },
    yAxis: {
      type: 'value',
      nameTextStyle: {
        color: '#777',
      },
      min: 0,
      splitLine: {
        show: true,
        lineStyle: {
          color: '#e1e1e1',
        },
      },
      axisLine: {
        show: false,
      },
      axisLabel: {
        show: true,
        textStyle: {
          color: '#92969E',
        },
      },
      axisTick: {
        show: false,
      },
      // max: 10 // 设置y轴最大值为100
    },
    // 其他配置项
  })
  // 渲染图表数据
  chart.setOption({
    series: config.data.map((series, index) => ({
      type: 'line',
      name: series.name,
      data: series.data.map(item => [item.time, item.value]),
      symbol: 'circle', // 默认是空心圆（中间是白色的），改成实心圆
      showAllSymbol: true,
      symbolSize: 0,
      smooth: true,
      lineStyle: {
        normal: {
          width: 2,
          color: colors[index], // 线条颜色
        },
      },
      itemStyle: {
        color: colors[index], // 使用相应的颜色
      },
      areaStyle: {
        normal: {
          color: new echarts.graphic.LinearGradient(0, 0, 0, 1, [
            {
              offset: 0,
              color: colors[index], // 使用相应的颜色
            },
            {
              offset: 1,
              color: 'rgba(255, 255, 255, 0.2)',
            }
          ], false),
          shadowColor: 'rgba(117,191,165,0.52)',
          shadowBlur: 3,
        },
      },
      // 其他配置...
    }))
  })
}



function init_chart_of_memory() {
  const config = initialChartsOfMemory.value[0]
  console.log(initialCharts.value)
  console.log(initialCharts.value.length)
  console.log(123123)
  // echarts.dispose(document.getElementById(config.id))
  echarts.dispose(document.getElementById('chart1'))
  const chart = echarts.init(document.getElementById('chart1'), 'light', {
    height: '240px'
  })
  const colors = interpolateColor(startColor, endColor, config.data.length)
  chart.setOption({
    title: {
      text: config.title,
      left: 'center',
      top: '10px'
    },
    backgroundColor: 'transparent',
    tooltip: {
      trigger: 'axis',
      backgroundColor: 'transparent',
    },
    legend: {
      align: 'left',
      right: '5%',
      top: '0%',
      type: 'plain',
      itemGap: 25,
      itemWidth: 20,
      // icon: 'path://M0 2a2 2 0 0 1 2 -2h14a2 2 0 0 1 2 2v0a2 2 0 0 1 -2 2h-14a2 2 0 0 1 -2 -2z',
      data: config.data.map((series, index) => ({
        name: series.name,
        textStyle: {
          color: colors[index], // 使用相应的颜色
        },
      })),
    },
    // legend: {
    //   data: config.data.map(series => series),
    //   right: 20, // 调整图例到右边的距离
    //   top: 20, // 调整图例到顶部的距离// 根据配置的数据动态生成 legend 数据
    // },
    xAxis: {
      type: 'time',
      // data: globalTimeList,
      boundaryGap: false,
      axisLine: {
        // 坐标轴轴线相关设置。数学上的x轴
        show: false,
        lineStyle: {
          color: '#e1e1e1',
        },
      },
      axisLabel: {
        // 坐标轴刻度标签的相关设置
        textStyle: {
          color: '#92969E',
        },
        // formatter: function(data) {
        //   return data
        // },
      },
      splitLine: {
        show: false,
        lineStyle: {
          color: '#192a44',
        },
      },
      axisTick: {
        show: false,
      },
    },
    yAxis: {
      type: 'value',
      nameTextStyle: {
        color: '#777',
      },
      min: 0,
      splitLine: {
        show: true,
        lineStyle: {
          color: '#e1e1e1',
        },
      },
      axisLine: {
        show: false,
      },
      axisLabel: {
        show: true,
        textStyle: {
          color: '#92969E',
        },
      },
      axisTick: {
        show: false,
      },
      // max: 10 // 设置y轴最大值为100
    },
    // 其他配置项
  })
  // 渲染图表数据
  chart.setOption({
    series: config.data.map((series, index) => ({
      type: 'line',
      name: series.name,
      data: series.data.map(item => [item.time, item.value]),
      symbol: 'circle', // 默认是空心圆（中间是白色的），改成实心圆
      showAllSymbol: true,
      symbolSize: 0,
      smooth: true,
      lineStyle: {
        normal: {
          width: 2,
          color: colors[index], // 线条颜色
        },
      },
      itemStyle: {
        color: colors[index], // 使用相应的颜色
      },
      areaStyle: {
        normal: {
          color: new echarts.graphic.LinearGradient(0, 0, 0, 1, [
            {
              offset: 0,
              color: colors[index], // 使用相应的颜色
            },
            {
              offset: 1,
              color: 'rgba(255, 255, 255, 0.2)',
            }
          ], false),
          shadowColor: 'rgba(117,191,165,0.52)',
          shadowBlur: 3,
        },
      },
      // 其他配置...
    }))
  })
}

function renderChart() {
  echarts.dispose(document.getElementById('chart'))
  const chart = echarts.init(document.getElementById('chart'), 'light')
  console.log(seriesData)

  // const seriesData1 = [
  //   {
  //     name: 'Series 1',
  //     data: [
  //       { time: '2024-04-16 09:00', value: 10 },
  //       { time: '2024-04-16 10:00', value: 20 },
  //       { time: '2024-04-16 11:00', value: 30 },
  //       // 其他数据...
  //     ]
  //   },
  //   {
  //     name: 'Series 2',
  //     data: [
  //       { time: '2024-04-16 09:00', value: 15 },
  //       { time: '2024-04-16 10:00', value: 25 },
  //       { time: '2024-04-16 11:00', value: 35 },
  //       // 其他数据...
  //     ]
  //   },
  //   // 其他时间序列数据...
  // ]
  const series = seriesData.map(series => ({
    type: 'line',
    name: series.name,
    data: series.data.map(item => [item.time, item.value]),
    // 其他配置...
  }))

  console.log(series)
  console.log()
  chart.setOption({
    // 在这里设置您的图表选项
    title: {
      text: 'CPU占用率', // 标题文本
      left: 'center', // 标题水平居中
      top: '10px' // 标题距离顶部的距离
    },
    // legend: {
    //   data: ['Machine 1', 'Machine 2', 'Machine 3'], // 图例的数据项名称
    //   top: '30px', // 图例距离顶部的距离
    //   left: 'center' // 图例水平居中
    // },
    xAxis: {
      type: 'time',
      // data: globalTimeList,
    },
    yAxis: {
      type: 'value',
      // max: 10 // 设置y轴最大值为100
    },

    series: series
    // data: {
    //   labels: labels,
    //   datasets: [
    // //     {
    //       label: 'Machine 1 CPU Usage',
    //       data: machine1Data,
    //       borderColor: 'rgba(255, 99, 132, 1)',
    //       borderWidth: 1,
    //       fill: false
    //     },
    //     {
    //       label: 'Machine 2 CPU Usage',
    //       data: machine2Data,
    //       borderColor: 'rgba(54, 162, 235, 1)',
    //       borderWidth: 1,
    //       fill: false
    //     },
    //     {
    //       label: 'Machine 3 CPU Usage',
    //       data: machine3Data,
    //       borderColor: 'rgba(75, 192, 192, 1)',
    //       borderWidth: 1,
    //       fill: false
    //     }
    //   ]
    // },
  })

  //
  // // Create chart
  // const ctx = document.getElementById('chart').getContext('2d');
  // const cpuChart = new Chart(ctx, {
  //   type: 'line',
  //   data: {
  //     labels: labels,
  //     datasets: [
  //       {
  //         label: 'Machine 1 CPU Usage',
  //         data: machine1Data,
  //         borderColor: 'rgba(255, 99, 132, 1)',
  //         borderWidth: 1,
  //         fill: false
  //       },
  //       {
  //         label: 'Machine 2 CPU Usage',
  //         data: machine2Data,
  //         borderColor: 'rgba(54, 162, 235, 1)',
  //         borderWidth: 1,
  //         fill: false
  //       },
  //       {
  //         label: 'Machine 3 CPU Usage',
  //         data: machine3Data,
  //         borderColor: 'rgba(75, 192, 192, 1)',
  //         borderWidth: 1,
  //         fill: false
  //       }
  //     ]
  //   },
  //   options: {
  //     scales: {
  //       x: {
  //         type: 'time',
  //         time: {
  //           unit: 'minute'
  //         }
  //       },
  //       y: {
  //         beginAtZero: true,
  //         title: {
  //           display: true,
  //           text: 'CPU Usage (%)'
  //         }
  //       }
  //     }
  //   }
  // })
}

// 自动化生成的字典（可能为空）以及字段
const formData = ref({
  dataTypeID: 0,
  value: 0,
  machineID: 0,
})

// 验证规则
const rule = reactive({
  dataTypeID: [{
    required: true,
    message: '',
    trigger: ['input', 'blur'],
  },
  ],
  value: [{
    required: true,
    message: '',
    trigger: ['input', 'blur'],
  },
  ],
  machineID: [{
    required: true,
    message: '',
    trigger: ['input', 'blur'],
  },
  ],
})

const searchRule = reactive({
  createdAt: [
    { validator: (rule, value, callback) => {
      if (searchInfo.value.startCreatedAt && !searchInfo.value.endCreatedAt) {
        callback(new Error('请填写结束日期'))
      } else if (!searchInfo.value.startCreatedAt && searchInfo.value.endCreatedAt) {
        callback(new Error('请填写开始日期'))
      } else if (searchInfo.value.startCreatedAt && searchInfo.value.endCreatedAt && (searchInfo.value.startCreatedAt.getTime() === searchInfo.value.endCreatedAt.getTime() || searchInfo.value.startCreatedAt.getTime() > searchInfo.value.endCreatedAt.getTime())) {
        callback(new Error('开始日期应当早于结束日期'))
      } else {
        callback()
      }
    }, trigger: 'change' }
  ],
})

const elFormRef = ref()
const elSearchFormRef = ref()

// =========== 表格控制部分 ===========
const page = ref(1)
const total = ref(0)
const pageSize = ref(10)
const tableData = ref([])
const searchInfo = ref({})
const responseData = ref({})

const machineList = ref([])
const dataTypeList = ref([])

// 重置
const onReset = () => {
  searchInfo.value = {}
  getTableData()
}

// 搜索
const onSubmit = () => {
  elSearchFormRef.value?.validate(async(valid) => {
    if (!valid) return
    page.value = 1
    pageSize.value = 10
    getTableData()
  })
}

// 分页
const handleSizeChange = (val) => {
  pageSize.value = val
  getTableData()
}

// 修改页面容量
const handleCurrentChange = (val) => {
  page.value = val
  getTableData()
}

// 查询
const getTableData = async() => {
  const table = await getDataList({ page: page.value, pageSize: pageSize.value, ...searchInfo.value })
  if (table.code === 0) {
    tableData.value = table.data.list
    total.value = table.data.total
    page.value = table.data.page
    pageSize.value = table.data.pageSize
  }
}

getTableData()

// ============== 表格控制部分结束 ===============

// 获取需要的字典 可能为空 按需保留
const setOptions = async() => {
}

// 获取需要的字典 可能为空 按需保留
setOptions()

// 多选数据
const multipleSelection = ref([])
// 多选
const handleSelectionChange = (val) => {
  multipleSelection.value = val
}

// 删除行
const deleteRow = (row) => {
  ElMessageBox.confirm('确定要删除吗?', '提示', {
    confirmButtonText: '确定',
    cancelButtonText: '取消',
    type: 'warning'
  }).then(() => {
    deleteDataFunc(row)
  })
}

// 多选删除
const onDelete = async() => {
  ElMessageBox.confirm('确定要删除吗?', '提示', {
    confirmButtonText: '确定',
    cancelButtonText: '取消',
    type: 'warning'
  }).then(async() => {
    const IDs = []
    if (multipleSelection.value.length === 0) {
      ElMessage({
        type: 'warning',
        message: '请选择要删除的数据'
      })
      return
    }
    multipleSelection.value &&
        multipleSelection.value.map(item => {
          IDs.push(item.ID)
        })
    const res = await deleteDataByIds({ IDs })
    if (res.code === 0) {
      ElMessage({
        type: 'success',
        message: '删除成功'
      })
      if (tableData.value.length === IDs.length && page.value > 1) {
        page.value--
      }
      getTableData()
    }
  })
}

// 行为控制标记（弹窗内部需要增还是改）
const type = ref('')

// 更新行
const updateDataFunc = async(row) => {
  const res = await findData({ ID: row.ID })
  type.value = 'update'
  if (res.code === 0) {
    formData.value = res.data.redata
    dialogFormVisible.value = true
  }
}

// 删除行
const deleteDataFunc = async(row) => {
  const res = await deleteData({ ID: row.ID })
  if (res.code === 0) {
    ElMessage({
      type: 'success',
      message: '删除成功'
    })
    if (tableData.value.length === 1 && page.value > 1) {
      page.value--
    }
    getTableData()
  }
}

// 弹窗控制标记
const dialogFormVisible = ref(false)

// 查看详情控制标记
const detailShow = ref(false)

// 打开详情弹窗
const openDetailShow = () => {
  detailShow.value = true
}

// 打开详情
const getDetails = async(row) => {
  // 打开弹窗
  const res = await findData({ ID: row.ID })
  if (res.code === 0) {
    formData.value = res.data.redata
    openDetailShow()
  }
}

// 关闭详情弹窗
const closeDetailShow = () => {
  detailShow.value = false
  formData.value = {
    dataTypeID: 0,
    value: 0,
    machineID: 0,
  }
}

// 打开弹窗
const openDialog = () => {
  type.value = 'create'
  dialogFormVisible.value = true
}

// 关闭弹窗
const closeDialog = () => {
  dialogFormVisible.value = false
  formData.value = {
    dataTypeID: 0,
    value: 0,
    machineID: 0,
  }
}
// 弹窗确定
const enterDialog = async() => {
  elFormRef.value?.validate(async(valid) => {
    if (!valid) return
    let res
    switch (type.value) {
      case 'create':
        res = await createData(formData.value)
        break
      case 'update':
        res = await updateData(formData.value)
        break
      default:
        res = await createData(formData.value)
        break
    }
    if (res.code === 0) {
      ElMessage({
        type: 'success',
        message: '创建/更改成功'
      })
      closeDialog()
      getTableData()
    }
  })
}

// 图线筛选框
const filterForm = ref({
  start_time: '', // 使用格式化的开始时间间
  end_time: '', // 使用格式化的当前时间作为结束时
  dataType: '',
  machines: [] // 这里改为数组，以便支持多选
})

// 假设数据类型和机器的选项从后端获取，这里模拟一些示例数据
const dataTypes = ref([])

const machines = ref([])

const allDataTypesChecked = ref(false)
const allMachinesChecked = ref(false)
const handleAllDataTypesChange = (checked) => {
  if (checked) {
    filterForm.value.dataType = dataTypes.value.map(type => type.id)
  } else {
    filterForm.value.dataType = []
  }
}

const handleAllMachinesChange = (checked) => {
  if (checked) {
    filterForm.value.machines = machines.value.map(machine => machine.id)
  } else {
    filterForm.value.machines = []
  }
}
// 监听数据类型选择框的变化
const handleDataTypeChange = (value) => {
  allDataTypesChecked.value = value.length === dataTypes.value.length
}

// 监听机器选择框的变化
const handleMachinesChange = (value) => {
  allMachinesChecked.value = value.length === machines.value.length
}

// 监听表单提交事件
const handleSubmit = () => {
  // 获取筛选条件
  if (filterForm.value.start_time && filterForm.value.end_time) {
    data1.value.start_time = formatDate(filterForm.value.start_time)
    data1.value.end_time = formatDate(filterForm.value.end_time)
  }
  data1.value.machine_ids = filterForm.value.machines
  // data1.value.data_type_id = filterForm.value.dataType
  // for (let i = 0; i < filterForm.value.dataType.length; i++) {
  //   data1.value.data_type_id = filterForm.value.dataType[i]
  // }
  getData123()

  // 根据筛选条件请求数据并更新图表
  // fetchDataAndRenderChart(startTime, endTime, dataType, machines)
}
const initialDataTypes = ref( ['2857619455','301978823']) //cpu-percent,memory-percent,disk-percent //,'3019788237','2963749463'
const data2 = ref({
  machine_ids: [],
  data_type_id: '2857619455',
  start_time: '2024-04-28 00:00:33', // 使用格式化的开始时间间
  end_time: '2024-04-29 00:00:33', // 使用格式化的当前时间作为结束时
  // start_time: '',
  // end_time: '',

})

const data3 = ref({
  machine_ids: [],
  data_type_id: '2857619455',
  start_time: '2024-04-28 00:00:33', // 使用格式化的开始时间间
  end_time: '2024-04-29 00:00:33', // 使用格式化的当前时间作为结束时
  // start_time: '',
  // end_time: '',
})
const datalist = ref ([])
const data1 = ref({
  machine_ids: ['1'],
  data_type_id: '1',
  start_time: '2024-04-28 02:00:33', // 使用格式化的开始时间间
  end_time: '2024-04-30 02:00:43', // 使用格式化的当前时间作为结束时
  // start_time: '',
  // end_time: '',

})

const globalTimeList = []
const globalValuesList = []
const chartConfigs = ref([])
const initialCharts = ref([])
const initialChartsOfMemory = ref([])
const flag1 = ref(false)
const getData123 = async() => {
  // 获取当前时间的Date对象
  // const endTime = new Date()
  //
  // console.log(formatDate(searchInfo.value.startCreatedAt))
  // if (searchInfo.value.startCreatedAt && searchInfo.value.endCreatedAt) {
  //   data1.value.start_time = formatDate(searchInfo.value.startCreatedAt)
  //   data1.value.end_time = formatDate(searchInfo.value.endCreatedAt)
  // }
  // console.log(data1.value.end_time)
  // 打开弹窗
  chartConfigs.value.length = 0
    for (const i in filterForm.value.dataType) {
      console.log(1111)
      console.log(filterForm.value.dataType)
      data1.value.data_type_id = filterForm.value.dataType[i].toString()
      const res = await getData(data1.value)
      if (res.code === 0) {
        // ElMessage({
        //   type: 'success',
        //   message: '成功'
        // })
        responseData.value = res.data
        const time_dic = {}
        const seriesData = []
        for (const key in responseData.value) {
          time_dic[key] = []
          // val_dic[key] = []
          console.log(responseData.value)
          console.log(key)
          responseData.value[key].forEach(item => {
            time_dic[key].push({ time: formatDate(item.CreatedAt), value: item.value })
          })
          console.log(time_dic[key])
          console.log(key)
          seriesData.push({ name: machineMap[key], data: time_dic[key] })
          console.log(seriesData)
          console.log(time_dic[key])
        }
        chartConfigs.value.push({
          id: filterForm.value.dataType[i].toString(),
          title: dataTypeMap[filterForm.value.dataType[i]],
          data: seriesData
        })
        // globalTimeList.length = 0 // 清空时间列表数组
        // globalValuesList.length = 0 // 清空数值列表数组
        // globalTimeList.push(...time)
        // globalValuesList.push(...values)
      }
    }

  console.log(chartConfigs.value)
  nextTick(() => {
    // 在 DOM 更新完成后执行绘制图表的操作
    console.log(initialDataTypes)
    renderCharts()
  })
}

const getInitialData = async() => {
  // 获取当前时间的Date对象
  // const endTime = new Date()
  //
  // console.log(formatDate(searchInfo.value.startCreatedAt))
  // if (searchInfo.value.startCreatedAt && searchInfo.value.endCreatedAt) {
  //   data1.value.start_time = formatDate(searchInfo.value.startCreatedAt)
  //   data1.value.end_time = formatDate(searchInfo.value.endCreatedAt)
  // }
  // console.log(data1.value.end_time)
  // 打开弹窗
  initialCharts.value.length = 0

  const endTime = new Date().toISOString()
  // 设置 startTime 为 endTime 的前一个小时
  const startTime = new Date(new Date(endTime).getTime() - 10 * 60 * 1000)
  const startTimeISOString = startTime.toISOString()
  data2.value.end_time = formatDate(endTime)
  data2.value.start_time = formatDate(startTimeISOString)
  //machine
  const machineIDs = []
  machines.value.forEach(item => {
    machineIDs.push(item.id.toString())
  })
  data2.value.machine_ids = machineIDs
  //time
  // request
  for (const i in initialDataTypes.value) {
    console.log(1111)
    console.log(initialDataTypes.value.dataType)
    data2.value.data_type_id = initialDataTypes.value[i]
    const res = await getData(data2.value)
    if (res.code === 0) {
      // ElMessage({
      //   type: 'success',
      //   message: '成功'
      // })
      responseData.value = res.data
      const time_dic = {}
      const seriesData = []
      for (const key in responseData.value) {
        time_dic[key] = []
        // val_dic[key] = []
        console.log(responseData.value)
        console.log(key)
        responseData.value[key].forEach(item => {
          time_dic[key].push({ time: formatDate(item.CreatedAt), value: item.value })
        })
        console.log(time_dic[key])
        console.log(key)
        seriesData.push({ name: machineMap[key], data: time_dic[key] })
        console.log(seriesData)
        console.log(time_dic[key])
      }
      initialCharts.value.push({
        id: initialDataTypes.value[i],
        title: dataTypeMap[initialDataTypes.value[i]],
        data: seriesData
      })
      // globalTimeList.length = 0 // 清空时间列表数组
      // globalValuesList.length = 0 // 清空数值列表数组
      // globalTimeList.push(...time)
      // globalValuesList.push(...values)
    }
  }

  console.log(chartConfigs.value)
  nextTick(() => {
    // 在 DOM 更新完成后执行绘制图表的操作
    console.log(initialDataTypes)
    // init_charts()
  })
}


// const getData1234 = async() => {
//   initialCharts.value.length = 0
//   if (!flag1.value) {
//     const machineIDs = []
//     machines.value.forEach(item => {
//       machineIDs.push(item.id.toString())
//     })
//     data2.value.data_type_id = '2857619455'
//     data2.value.machine_ids = machineIDs
//     const endTime = new Date().toISOString()
//     // 设置 startTime 为 endTime 的前一个小时
//     const startTime = new Date(new Date(endTime).getTime() - 10 * 60 * 1000)
//     const startTimeISOString = startTime.toISOString()
//     data2.value.end_time = formatDate(endTime)
//     data2.value.start_time = formatDate(startTimeISOString)
//     console.log(data2.value.start_time)
//     data3.value.end_time = formatDate(endTime)
//     data3.value.start_time = formatDate(startTimeISOString)
//     console.log(machineIDs)
//     // data2.value.machine_ids =
//     console.log(initialDataTypes)
//     const res = await getData(data2.value)
//     if (res.code === 0) {
//       // ElMessage({
//       //   type: 'success',
//       //   message: '成功'
//       // })
//       responseData.value = res.data
//       const time_dic = {
//       }
//       const seriesData = []
//       for (const key in responseData.value) {
//         time_dic[key] = []
//         // val_dic[key] = []
//         responseData.value[key].forEach(item => {
//           time_dic[key].push({ time: formatDate(item.CreatedAt), value: item.value })
//         })
//         seriesData.push({ name: machineMap[key], data: time_dic[key] })
//       }
//       initialCharts.value.push({ id: data2.value.data_type_id.toString(), title: dataTypeMap[data2.value.data_type_id], data: seriesData })
//     //   // globalTimeList.length = 0 // 清空时间列表数组
//     //   // globalValuesList.length = 0 // 清空数值列表数组
//     //   // globalTimeList.push(...time)
//     //   // globalValuesList.push(...values)
//     }
//     flag1.value = true
//   }
//   console.log(initialCharts.value)
//   nextTick(() => {
//     // 在 DOM 更新完成后执行绘制图表的操作
//     console.log(1111)
//     init_charts()
//   })
// }

async function getData1234() {
  initialCharts.value.length = 0;
  const machineIDs = machines.value.map(item => item.id.toString());
  data2.value.data_type_id = '2857619455';
  data2.value.machine_ids = machineIDs;
  const endTime = new Date().toISOString();
  // 设置 startTime 为 endTime 的前十分钟
  const startTime = new Date(new Date(endTime).getTime() - 10 * 60 * 1000);
  data2.value.end_time = formatDate(endTime);
  data2.value.start_time = formatDate(startTime.toISOString());
  data3.value.end_time = formatDate(endTime);
  data3.value.start_time = formatDate(startTime.toISOString());

  try {
    const res = await getData(data2.value);
    if (res.code === 0) {
      responseData.value = res.data;
      const time_dic = {};
      const seriesData = [];
      for (const key in responseData.value) {
        time_dic[key] = [];
        responseData.value[key].forEach(item => {
          time_dic[key].push({ time: formatDate(item.CreatedAt), value: item.value });
        });
        seriesData.push({ name: machineMap[key], data: time_dic[key] });
      }
      initialCharts.value.push({ id: data2.value.data_type_id.toString(), title: dataTypeMap[data2.value.data_type_id], data: seriesData });
    }
  } catch (error) {
    console.error('Error fetching data:', error);
  }
  nextTick(() => {
    // 在 DOM 更新完成后执行绘制图表的操作
    console.log(1111)
    init_charts()
  })
}

let flag2 = { value: false }; // 初始设置 flag1 为 false
let refreshIntervalId; // 存储 setInterval 返回的 ID

async function startRefresh() {
  await getData1234(); // 执行一次获取数据和绘制图表的操作
  flag2.value = true; // 将 flag1 设置为 true
  refreshIntervalId = setInterval(getData1234, 15000); // 每隔十秒执行一次 getData1234
}

startRefresh(); // 开始执行定时刷新

// 当需要停止定时刷新时，可以调用 clearInterval(refreshIntervalId);

const getData12345 = async() => {
  initialChartsOfMemory.value.length = 0
  if (true) {
    const machineIDs = []
    machines.value.forEach(item => {
      machineIDs.push(item.id.toString())
    })
    data3.value.data_type_id = '3019788237'
    data3.value.machine_ids = machineIDs
    const endTime = new Date().toISOString();
    // 设置 startTime 为 endTime 的前十分钟
    const startTime = new Date(new Date(endTime).getTime() - 10 * 60 * 1000);
    data3.value.start_time=formatDate(startTime)
    data3.value.end_time=formatDate(endTime)
    console.log(machineIDs)
    // data2.value.machine_ids =
    console.log(initialDataTypes)
    try {
      const res = await getData(data3.value)
      if (res.code === 0) {
        // ElMessage({
        //   type: 'success',
        //   message: '成功'
        // })
        responseData.value = res.data
        const time_dic = {}
        const seriesData = []
        for (const key in responseData.value) {
          time_dic[key] = []
          // val_dic[key] = []
          responseData.value[key].forEach(item => {
            time_dic[key].push({time: formatDate(item.CreatedAt), value: item.value})
          })
          seriesData.push({name: machineMap[key], data: time_dic[key]})
        }
        initialChartsOfMemory.value.push({
          id: data3.value.data_type_id.toString(),
          title: dataTypeMap[data3.value.data_type_id],
          data: seriesData
        })
        //   // globalTimeList.length = 0 // 清空时间列表数组
        //   // globalValuesList.length = 0 // 清空数值列表数组
        //   // globalTimeList.push(...time)
        //   // globalValuesList.push(...values)
      }
      flag1.value = true
      console.log(initialCharts.value)
    } catch (error) {
      console.error('Error fetching data:', error);
    }
    nextTick(() => {
      // 在 DOM 更新完成后执行绘制图表的操作
      console.log(1111)
      init_chart_of_memory()
    })
  }
}
setInterval(getData12345, 15000); // 每隔10秒刷新一次数据
const machineMap = []
const getCurrentMachines = async() => {
  const table = await getMachineList({ page: 0, pageSize: 10000 })
  if (table.code === 0) {
    machineList.value = table.data.list
    machines.value = machineList.value.map(item => ({
      id: item.ID.toString(),
      name: item.name
    }))
    console.log(machines.value)
  }
  // machineIDs.value.length = 0
  machines.value.forEach(item => {
    machineMap[item.id] = item.name
  })
  // getInitialData()
  getData1234()
  getData12345()
}
function sleep(ms) {
  return new Promise(resolve => setTimeout(resolve, ms))
}
getCurrentMachines()
const dataTypeMap = {}
const getCurrentDataTypes = async() => {
  const table = await getDataTypeList({ page: 0, pageSize: 10000 })
  if (table.code === 0) {
    dataTypeList.value = table.data.list
    dataTypes.value = dataTypeList.value.map(item => ({
      id: item.ID.toString(),
      name: item.name
    }))
  }

  dataTypes.value.forEach(item => {
    dataTypeMap[item.id] = item.name
  })

}

getCurrentDataTypes()



// renderChart()
</script>

<script>
export default {
  name: 'DataCenter',
  data() {
    return {
      val: [] // 假设 dataArray 是你要传递的数组变量
    };
  }
}
</script>

<style>
.data-center-box{
  width: 100%;
  display: grid;
  grid-template-columns: 1fr 1.5fr 2fr 1.5fr;
  column-gap: 10px;
}
  /* 样式可根据需要自定义 */

.chart-container {
  display: flex;
}

.chart-item {
  margin-right: 10px; /* 设置曲线之间的间距 */
}

.chart-content {
  width: 100%;
  height: 100%;
}

</style>
