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
import { nextTick, onMounted, onUnmounted, ref } from 'vue'

const chart = ref(null)
const echart = ref(null)
const initChart = () => {
  chart.value = echarts.init(echart.value)
  setOptions()
  document.addEventListener('resize', () => {
    chart.value?.resize()
  })
}
var data1 = [
  { value: 1220, name: 'Jutze 2D AOI' },
  { value: 1400, name: 'Jutze 3D AOI' },
  { value: 4012, name: 'Jutze Twins' },
  { value: 2488, name: 'Jutze Twins2' },
]
let num = 0
data1.forEach(item => {
  num += item.value
})
var colorList = ['#eae74e', '#5591ef', '#b6c1de', '#81cab8']

const setOptions = () => {
  chart.value.setOption({
    backgroundColor: 'transparent',
    title: {
      text: '总台数',
      top: '23%',
      textAlign: 'center',
      left: '48%',
      textStyle: {
        color: ' #6A7474',
        fontSize: 12,
        fontWeight: '500',
        lineHeight: 12,
      },
      subtext: num,
      subtextStyle: {
        color: '#021C25',
        fontSize: 20,
        fontWeight: '600',
        lineHeight: 20,
      },
    },
    tooltip: {
      trigger: 'item',
      formatter: '{a} <br/>{b}: {c} ({d}%)',
    },
    legend: {
      selectedMode: true, // 取消图例上的点击事件
      type: 'plain',
      orient: 'vertical',
      left: '5%',
      bottom: '10%',
      align: 'left',
      itemGap: 20,
      itemWidth: 12, // 设置宽度
      itemHeight: 12, // 设置高度
      symbolKeepAspect: false,
      textStyle: {
        color: '#000',
        rich: {
          name: {
            verticalAlign: 'right',
            align: 'left',
            width: 70,
            fontSize: 12,
            color: '#666',
          },
          value: {
            align: 'left',
            width: 70,
            fontSize: 12,
            color: '#c1c1c1',
          },
          part1: {
            align: 'left',
            width: 30,
            fontSize: 12,
            color: '#76C0A5',
          },
          part2: {
            align: 'left',
            width: 30,
            fontSize: 12,
            color: '#CF5C50',
          }
        }
      },
      data: data1,
      formatter: function(name) {
        if (data1 && data1.length) {
          for (let i = 0; i < data1.length; i++) {
            if (name === data1[i].name) {
              const value = data1[i].value / num * 100
              const show = value.toFixed(2) > 20 ? `${value.toFixed(2)}% ↑` : `${value.toFixed(2)}% ↓`
              const color = value.toFixed(2) > 20 ? 'part1' : 'part2'
              return (
                '{name| ' + name + '}  ' + '{value| ' + data1[i].value + '}' + '{' + color + '| ' + show + '}'
              )
            }
          }
        }
      }
    },
    series: [
      {
        type: 'pie',
        center: ['50%', '30%'],
        radius: ['60%', '40%'],
        avoidLabelOverlap: false,
        label: {
          show: false,
          position: 'center',
        },
        itemStyle: {
          normal: {
            color: function(params) {
              return colorList[params.dataIndex]
            },
          },
        },
        labelLine: {
          show: true,
        },
        data: data1,
      },
    ],
  })
}

onMounted(() => {
  nextTick(() => {
    initChart()
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
  position: relative;
  &-echarts{
    position: absolute;
    top: 0;
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
