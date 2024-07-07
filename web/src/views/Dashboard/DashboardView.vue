<template>
  <div class="body">
    <CreateInvitation />
    <section class="section section--view">
      <h1 class="header">Invitaciones Existentes</h1>
    </section>
  </div>
</template>
<script setup>
import CreateInvitation from '@/views/Dashboard/CreateInvitation.vue'
</script>
<style scoped>
.body {
  width: 100%;
  padding: 2rem;
}
</style>
