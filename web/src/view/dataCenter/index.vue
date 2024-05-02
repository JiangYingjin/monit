<template>
  <div>
    <div class="data-center-box">
      <CenterCard title="设备在线情况">
        <template #action>
          <span class="gvaIcon-prompt" style="color: #999" />
        </template>
        <template #body>
          <ChainRatio />
        </template>
      </CenterCard>
      <CenterCard title="回收里程(Km)">
        <template #action>
          <span class="gvaIcon-prompt" style="color: #999" />
        </template>
        <template #body>
          <ReclaimMileage />
        </template>
      </CenterCard>
      <CenterCard title="订单完成率">
        <template #action>
          <span class="gvaIcon-prompt" style="color: #999" />
        </template>
        <template #body>
          <Order />
        </template>
      </CenterCard>
      <CenterCard title="设备回收率">
        <template #action>
          <span class="gvaIcon-prompt" style="color: #999" />
        </template>
        <template #body>
          <RecoveryRate />
        </template>
      </CenterCard>



      <CenterCard title="日活跃率" style="grid-column-start: span 3;">
        <template #action>
          <span class="gvaIcon-prompt" style="color: #999" />
        </template>
        <template #body>
          <lineCharts />
        </template>
      </CenterCard>

      <CenterCard title="源接入台数占比" style="grid-auto-columns: 1.5fr">
        <template #action>
          <span class="gvaIcon-prompt" style="color: #999" />
        </template>
        <template #body>
          <part />
        </template>
      </CenterCard>
    </div>
  </div>
</template>

<script setup>
import CenterCard from './dataCenterComponents/centerCard.vue'
import ChainRatio from './dataCenterComponents/chainRatio.vue'
import ReclaimMileage from './dataCenterComponents/ReclaimMileage.vue'
import RecoveryRate from './dataCenterComponents/RecoveryRate.vue'
import Order from './dataCenterComponents/order.vue'
import lineCharts from './dataCenterComponents/lineCharts.vue'
import part from '../data/dataCenterComponents/part.vue'

</script>

<script>
export default {
  name: 'DataCenter',
}
</script>

<style lang="scss" scoped>

.data-center-box{
  width: 100%;
  display: grid;
  grid-template-columns: 1fr 1.5fr 2fr 1.5fr;
  column-gap: 10px;
}
</style>
