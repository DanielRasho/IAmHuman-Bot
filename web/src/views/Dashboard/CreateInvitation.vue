<template>
  <section class="section section--create">
    <h1 class="header">Crear invitación</h1>
    <div class="create-input">
      <span class="input-tag">Clase</span>
      <TextInput type="text" />
    </div>
    <div class="create-input">
      <span class="input-tag">Section</span>
      <TextInput type="text" />
    </div>
    <div class="create-input">
      <span class="input-tag">Usos maximos</span>
      <TextInput type="text" />
    </div>
    <div class="create-input">
      <span class="input-tag">Sever ID</span>
      <TextInput type="text" />
    </div>
    <div class="create-input">
      <span class="input-tag">Role ID</span>
      <TextInput type="text" />
    </div>
    <ButtonPrimary text="Create" is-disabled="false" />
  </section>
</template>
<script setup>
import ButtonPrimary from '@/components/ButtonPrimary.vue'
import TextInput from '@/components/TextInput.vue'
</script>
<style scoped>
.header {
  color: var(--white);
  font-weight: bold;
  font-size: 2.2rem;
  margin-bottom: 1rem;
}
.section {
  margin-bottom: 3rem;
}
.create-input {
  margin-bottom: 0.5rem;
}
.input-tag {
  display: inline-block;
  width: 14ch;
}
</style>
