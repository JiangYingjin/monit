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

      <CenterCard
        title="CPU占用率"
        style="grid-column-start: span 4;"
      >
        <template #action>
          <span class="gvaIcon-prompt" style="color: #999" />
        </template>
        <template #body>
          <div
              id="chart"
              style="width: 100%; height: 360px;"
          />
        </template>
      </CenterCard>

    </div>
    <br>
    <br>

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
        <br>
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
            style="width: 200px;"
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
          <el-select v-model="selectedLayout" placeholder="选择布局" style="width: 200px;" >
            <el-option label="一行两个" value="double" />
            <el-option label="一行一个" value="single" />
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

    <el-row :gutter="20">
      <el-col
          v-for="(config, index) in chartConfigs"
          :key="config.id"
          :span="selectedLayout === 'double' ? 12 : 24"
      >
        <div
            :id="config.id"
            style="width: 100%; height: 500px;"
        />
      </el-col>
    </el-row>

    <br>

    <br>
    <br>


  </div>

</template>

<script setup>
import {
  getDataList,
  getData,
} from '@/api/data'
import { getMachineList } from '@/api/machine'
import lineCharts from './dataCenterComponents/lineCharts.vue'

// 全量引入格式化工具 请按需保留
import { getDictFunc, formatDate, formatBoolean, filterDict, ReturnArrImg, onDownloadFile } from '@/utils/format'
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


const selectedLayout = ref(); // 默认值

// 计算属性，用于显示当前布局的描述文本
const currentLayoutText = computed(() => {
  return selectedLayout.value === 'double' ? '一行两个' : '一行一个';
});

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

// =========== 表格控制部分 ===========
const page = ref(1)
const total = ref(0)
const pageSize = ref(10)
const tableData = ref([])
const searchInfo = ref({})
const responseData = ref({})

const machineList = ref([])
const dataTypeList = ref([])

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

const type = ref('')

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
const data_cpu = ref({
  machine_ids: [],
  data_type_id: '2857619455',
  start_time: '', // 使用格式化的开始时间间
  end_time: '', // 使用格式化的当前时间作为结束时
})

const data3 = ref({
  machine_ids: [],
  data_type_id: '2857619455',
  start_time: '2024-04-28 00:00:33', // 使用格式化的开始时间间
  end_time: '2024-04-29 00:00:33', // 使用格式化的当前时间作为结束时
  // start_time: '',
  // end_time: '',
})
const data1 = ref({
  machine_ids: ['1'],
  data_type_id: '1',
  start_time: '2024-04-28 02:00:33', // 使用格式化的开始时间间
  end_time: '2024-04-30 02:00:43', // 使用格式化的当前时间作为结束时
  // start_time: '',
  // end_time: '',

})
const chartConfigs = ref([])
const initialCharts = ref([])
const initialChartsOfMemory = ref([])
const flag1 = ref(false)
const getData123 = async() => {
  chartConfigs.value.length = 0
  for (const i in filterForm.value.dataType) {
    console.log(1111)
    console.log(filterForm.value.dataType)
    data1.value.data_type_id = filterForm.value.dataType[i].toString()
    const res = await getData(data1.value)
    if (res.code === 0) {
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
    }
  }

  console.log(chartConfigs.value)
  nextTick(() => {
    // 在 DOM 更新完成后执行绘制图表的操作
    console.log(initialDataTypes)
    renderCharts()
  })
}

async function getData1234() {
  initialCharts.value.length = 0;
  const machineIDs = machines.value.map(item => item.id.toString());
  data_cpu.value.data_type_id = '2857619455';
  data_cpu.value.machine_ids = machineIDs;
  const endTime = new Date().toISOString();
  // 设置 startTime 为 endTime 的前十分钟
  const startTime = new Date(new Date(endTime).getTime() - 10 * 60 * 1000);
  data_cpu.value.end_time = formatDate(endTime);
  data_cpu.value.start_time = formatDate(startTime.toISOString());
  data3.value.end_time = formatDate(endTime);
  data3.value.start_time = formatDate(startTime.toISOString());

  try {
    const res = await getData(data_cpu.value);
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
      initialCharts.value.push({ id: data_cpu.value.data_type_id.toString(), title: dataTypeMap[data_cpu.value.data_type_id], data: seriesData });
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
  const machineIDs = []
  machines.value.forEach(item => {
    machineIDs.push(item.id.toString())
  })
  data3.value.data_type_id = '3019788237'
  data3.value.machine_ids = machineIDs
  const endTime = new Date().toISOString();
  const startTime = new Date(new Date(endTime).getTime() - 10 * 60 * 1000);
  data3.value.start_time = formatDate(startTime)
  data3.value.end_time = formatDate(endTime)
  console.log(machineIDs)
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
