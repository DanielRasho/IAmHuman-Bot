<template>
  <section class="section section--create">
    <h1 class="header">Eliminar invitacion</h1>
    <div class="create-input">
      <span class="input-tag">ID invitacion</span>
      <TextInput type="text" v-model="invID" />
    </div>
    <ButtonPrimary text="Delete" :is-disabled="false" @clicked="handleDeleteInvitation" />
  </section>
</template>
<script setup>
import ButtonPrimary from '@/components/ButtonPrimary.vue'
import TextInput from '@/components/TextInput.vue'
import { APISettings } from '@/utils/apiConfig'
import { ref } from 'vue'

const invID = ref('')

async function handleDeleteInvitation() {
  if (invID.value == "") return
  const url = APISettings.baseURL + 'dashboard/?invitation_id=' + invID.value
  console.log(url)
  const token = sessionStorage.getItem('token')
  try {
    const response = await fetch(url, {
      method: 'DELETE',
      headers: {
        Authorization: token,
        'Content-Type': 'application/json'
      }
    })
    let data = await response.json()
    if (!response.ok) throw new Error(data.message)
    alert('Success\n' + data.message)
  } catch (err) {
    alert('Cannot delete invitation\n' + err)
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
