<template>
  <div class="body">
    <CreateInvitation />
    <ViewInvitations />
  </div>
</template>
<script setup>
import CreateInvitation from '@/views/Dashboard/CreateInvitation.vue'
import ViewInvitations from '@/views/Dashboard/ViewInvitations.vue'
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
</style>
