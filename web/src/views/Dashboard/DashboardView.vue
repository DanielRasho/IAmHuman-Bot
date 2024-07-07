<template>
  <div class="body">
    <div class="forms-section">
      <CreateInvitation />
      <DeleteInvitation />
    </div>
    <ViewInvitations />
  </div>
</template>
<script setup>
import CreateInvitation from '@/views/Dashboard/CreateInvitation.vue'
import ViewInvitations from '@/views/Dashboard/ViewInvitations.vue'
import DeleteInvitation from '@/views/Dashboard/DeleteInvitation.vue'
import { onMounted } from 'vue'
import router from '@/router'

onMounted(() => {
  let token = sessionStorage.getItem('token')
  if (token == '') {
    router.replace('login')
  }
})
</script>
<style scoped>
.body {
  width: 100%;
  padding: 2rem;
}
.forms-section{
  display: flex;
  flex-wrap: wrap;
  justify-content: space-around;
}
</style>
