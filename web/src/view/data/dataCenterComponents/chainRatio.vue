<template>
  <div class="chainRatio-box">
    <div
      ref="echart"
      class="chainRatio-box-echarts"
      style="height: 120px"
    />
    <div class="chainRatio-box-values">
      <div class="chainRatio-box-values-item in-line">在线<span>1</span></div>
      <div class="chainRatio-box-values-item out-line">离线<span>3</span></div>
    </div>
  </div>
</template>
<script setup>
import * as echarts from 'echarts'
import { nextTick, onMounted, onUnmounted, ref, shallowRef, watchEffect } from 'vue'

const chart = shallowRef(null)
const echart = ref(null)
const initChart = () => {
  chart.value = echarts.init(echart.value)
  setOptions()
}
const setOptions = () => {
  chart.value.setOption({
    backgroundColor: '#fbfbfb',
    title: {
      text: '4',
      textStyle: {
        color: '#1d1d1f',
        fontSize: 14
      },
      subtext: '总台数',
      subtextStyle: {
        color: '#999',
        fontSize: 13
      },
      itemGap: 20,
      left: 'center',
      top: '40%'
    },
    angleAxis: {
      startAngle: 180,
      max: 360,
      clockwise: true,
      show: false
    },
    radiusAxis: {
      type: 'category',
      show: true,
      axisLabel: {
        show: false
      },
      axisLine: {
        show: false

      },
      axisTick: {
        show: false
      }
    },
    polar: {
      center: ['50%', '50%'],
      radius: '100%',
    },
    series: [
      {
        type: 'bar',
        z: 2,
        data: [50],
        showBackground: true,
        backgroundStyle: {
          color: 'transparent'
        },
        coordinateSystem: 'polar',
        roundCap: false,
        barWidth: 15,
        barGap: '-100%',
        itemStyle: {
          normal: {
            opacity: 1,
            color: new echarts.graphic.LinearGradient(0, 0, 1, 0, [
              {
                offset: 0,
                color: '#DF5341'
              }, {
                offset: 1,
                color: '#DF534D'
              }]),
            shadowBlur: 10,
            shadowColor: '#DF534E',

          }
        }
      },
      {
        type: 'bar',
        z: 1,
        data: [180],
        coordinateSystem: 'polar',
        roundCap: false,
        barWidth: 10,
        barGap: '-100%',
        itemStyle: {
          normal: {
            opacity: 1,
            color: '#5BC2A4'
          }
        }
      },
    ],
  })
}

onMounted(() => {
  nextTick(() => {
    initChart()
    window.addEventListener('resize', () => {
      chart.value?.resize()
    })
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

.chainRatio-box{
  width: 100%;
  height: 120px;
  overflow: hidden;
  position: relative;
  &-echarts{
    width: 100%;
    height: 120px;
    position: absolute;
    left: -20%;
    top: 0;
    bottom: 0;
  }
  &-values{
    position: absolute;
    right: 0;
    top: 0;
    transform: translateY(50%);

    &-item{
      font-size: 13px;
      margin-bottom: 10px;
      color: #777;
      position: relative;
      padding-left: 10px;
      &::before{
        content: '';
        position: absolute;
        width: 8px;
        top: 0;
        left: 0;
        height: 8px;
        border-radius: 50%;
        background-color: var(--color);
        transform: translateY(50%);
      }
      span{
        color: var(--color);
        margin-left: 16px;
      }
    }
  }
}
.in-line{
  --color : #5BC2A4;
}
.out-line{
  --color: #DF534E;
}
</style>
