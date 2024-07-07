<template>
  <section class="section section--create">
    <h1 class="header">Crear invitación</h1>
    <div class="create-input">
      <span class="input-tag">Clase</span>
      <TextInput type="text" v-model="invClass" />
    </div>
    <div class="create-input">
      <span class="input-tag">Section</span>
      <TextInput type="number" v-model="invSection" />
    </div>
    <div class="create-input">
      <span class="input-tag">Usos maximos</span>
      <TextInput type="number" v-model="invMissingUses" />
    </div>
    <div class="create-input">
      <span class="input-tag">Sever ID</span>
      <TextInput type="text" v-model="invServerID" />
    </div>
    <div class="create-input">
      <span class="input-tag">Role ID</span>
      <TextInput type="text" v-model="invRoleID" />
    </div>
    <ButtonPrimary text="Create" :is-disabled="false" @clicked="handleCreateInvitation" />
  </section>
</template>
<script setup>
import ButtonPrimary from '@/components/ButtonPrimary.vue'
import TextInput from '@/components/TextInput.vue'
import { APISettings } from '@/utils/apiConfig'
import { ref } from 'vue'

const invClass = ref('')
const invSection = ref('')
const invMissingUses = ref('')
const invServerID = ref('')
const invRoleID = ref('')

async function handleCreateInvitation() {
  let section = parseInt(invSection.value)
  let missinUses = parseInt(invMissingUses.value)

  let invitationData = {
    class: invClass.value,
    section: section,
    missing_uses: missinUses,
    server_id: invServerID.value,
    role_id: invRoleID.value
  }
  const url = APISettings.baseURL + 'dashboard/'
  const token = sessionStorage.getItem('token')
  try {
    const response = await fetch(url, {
      method: 'POST',
      headers: {
        Authorization: token
      },
      body: JSON.stringify(invitationData)
    })
    let data = await response.json()
    if (!response.ok) throw new Error(data.message)
    alert('Success\n' + data.message)
  } catch (err) {
    alert('Cannot create invitation\n' + err)
  }
}
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
