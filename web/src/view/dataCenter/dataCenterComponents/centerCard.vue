<template>
  <div class="data-center-card" >
    <div class="data-center-card-head">
      <div class="data-center-card-head-title">{{ title }}</div>
      <slot name="action" />
    </div>
    <div class="data-center-card-body">
      <slot name="body" />
    </div>
  </div>
</template>

<script setup>
defineProps({
  title: String,
})
</script>

<style lang="scss" scoped>
.data-center-card{
  //min-width: 250px;
  flex: auto;
  background-color: #fbfbfd;
  padding: 8px;
  margin-bottom: 12px;
  border-radius: 4px;
  box-shadow: 1px 1px 20px 2px rgba(#eee , .4);
  &-head{
    display: flex;
    align-items: center;
    justify-content: space-between;

    &-title{
      font-size: 14px;
      font-weight: 500;
      padding-left: 10px;
      position: relative;
      &::before{
        content: '';
        position: absolute;
        top: 0;
        left: 0;
        width: 4px;
        height: 100%;
        border-radius: 4px;
        background-color: #1890ff;
      }
    }
  }
  &-body{
    padding-top: 10px;
  }
}
</style>
