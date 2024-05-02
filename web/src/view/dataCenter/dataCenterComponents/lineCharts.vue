<template>
  <div class="lineCharts-box">
    <div
      ref="echart"
      class="lineCharts-box-echarts"
      :style="`width : ${chart?.clientWidth}px`"
    />
  </div>
</template>
<script setup>
import * as echarts from 'echarts'
import { nextTick, onMounted, onUnmounted, ref, watch } from 'vue'

const chart = ref(null)
const echart = ref(null)
const initChart = () => {
  chart.value = echarts.init(echart.value)
  setOptions()
  document.addEventListener('resize', () => {
    chart.value?.resize()
  })
}
const xLabel = ['1月', '2月', '3月', '4月', '5月', '6月', '7月', '8月', '9月', '10月', '11月', '12月']
const goOutSchool = [42871, 40494, 41470, 44968, 43653, 41899, 47615, 43116, 49451, 42149, 48873, 46551]

const setOptions = () => {
  chart.value.setOption({
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
      textStyle: {
        color: '#75BFA5',
      },
      itemGap: 25,
      itemWidth: 20,
      icon: 'path://M0 2a2 2 0 0 1 2 -2h14a2 2 0 0 1 2 2v0a2 2 0 0 1 -2 2h-14a2 2 0 0 1 -2 -2z',
      data: [
        {
          name: '设备数',
        },
        {
          name: '在线目标',
        },
      ],
    },
    grid: {
      top: '15%',
      left: '4%',
      right: '2%',
      bottom: '8%',
      // containLabel: true
    },
    xAxis: [
      {
        type: 'category',
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
          formatter: function(data) {
            return data
          },
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
        data: xLabel,
      },
    ],
    yAxis: [
      {
        name: '单位：笔',
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
          formatter: function(value) {
            if (value !== 0) {
              return `${value / 1000}k`
            }
            return value
          },
        },
        axisTick: {
          show: false,
        },
      },
    ],
    series: [
      {
        name: '在线目标',
        type: 'line',
        showSymbol: false,
        smooth: true,
        markLine: {
          symbol: 'none',
          data: [
            {
              name: '在线目标',
              yAxis: 47000,
              lineStyle: { width: 1.656, color: '#8C9CDA', opacity: 0.8 },
              label: { show: false },
            },
          ],
        },
      },
      {
        name: '设备数',
        type: 'line',
        symbol: 'circle', // 默认是空心圆（中间是白色的），改成实心圆
        showAllSymbol: true,
        symbolSize: 0,
        smooth: true,
        lineStyle: {
          normal: {
            width: 2,
            color: '#75BFA5', // 线条颜色
          },
        },
        areaStyle: {
          // 区域填充样式
          normal: {
            color: new echarts.graphic.LinearGradient(0, 0, 0, 1, [
              {
                offset: 0,
                color: 'rgba(50, 216, 205, .8)'
              },
              {
                offset: 1,
                color: 'rgba(255, 255, 255, 0.2)'
              }
            ], false),
            shadowColor: 'rgba(117,191,165,0.52)', // 阴影颜色
            shadowBlur: 3,
          },
        },
        data: goOutSchool,
      },
    ],
  })
}

onMounted(() => {
  nextTick(() => {
    setTimeout(() => {
      initChart()
    }, 300)
  })
})

onUnmounted(() => {
  if (!chart.value) {
    return
  }
  chart.value.dispose()
  chart.value = null
})
</script>
<style lang="scss" scoped>

.lineCharts-box{
  height: 360px;
  overflow: hidden;
  position: relative;
  &-echarts{
    position: absolute;
    bottom: 0;
    left: 0;
    right: 0;
    z-index: 2;
    width: 100%;
    height: 100%;
  }
}
.in-line{
  --color : #5BC2A4;
}
.out-line{
  --color: #DF534E;
}
</style>
