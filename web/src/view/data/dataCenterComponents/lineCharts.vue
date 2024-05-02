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
import * as d3 from 'd3'

const chart = ref(null)
const echart = ref(null)

props: {
  data: Array // 接收父组件传递的数组
}
const initChart = () => {
  chart.value = echarts.init(echart.value)
  setOptions()
  document.addEventListener('resize', () => {
    chart.value?.resize()
  })
}
const xLabel = ['1月', '2月', '3月', '4月', '5月', '6月', '7月', '8月', '9月', '10月', '11月', '12月']
const goOutSchool = [42871, 40494, 41470, 44968, 43653, 41899, 47615, 43116, 49451, 42149, 48873, 46551]
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
const setOptions = () => {
  const config = data[0]
  const colors = interpolateColor(startColor, endColor, config.data.length)
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
      data: config.data.map((series, index) => ({
        name: series.name,
        textStyle: {
          color: colors[index], // 使用相应的颜色
        },
      })),
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
        type: 'time',
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
    ],
    yAxis: [
      {
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
      },
    ],
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
